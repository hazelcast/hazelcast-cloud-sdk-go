package hazelcastcloud

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestNewClient_Create_A_New_Client_With_Provided_Http_Client(t *testing.T) {
	//given
	rawurl := "https://cloud.hazelcast.com/api/v1"
	parse, _ := url.Parse(rawurl)
	userAgent := "hazelcast-cloud-sdk-go/1.0.0"
	defaultClient := http.DefaultClient

	//when
	client := NewClient(nil)

	//then
	assert.Equal(t, parse, client.BaseURL)
	assert.Equal(t, userAgent, client.UserAgent)
	assert.Equal(t, defaultClient, client.client)
	assert.NotNil(t, client.Auth)
	assert.NotNil(t, client.StarterCluster)
	assert.NotNil(t, client.EnterpriseCluster)
	assert.NotNil(t, client.CloudProvider)
	assert.NotNil(t, client.CloudProvider)
}

func TestNewClient_Create_A_New_Client(t *testing.T) {
	//given
	serveMux := http.NewServeMux()
	server := httptest.NewServer(serveMux)
	defer server.Close()

	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodPost; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
	})

	//when
	client, response, err := New()

	//then
	assert.Nil(t, client)
	assert.Nil(t, response)
	responseErr := err.(*ErrorResponse)
	assert.Equal(t, "You need to provide HZ_CLOUD_API_KEY and HZ_CLOUD_API_SECRET in your environment variables.", responseErr.Message)
	assert.Equal(t, "", responseErr.CorrelationId)
	assert.Nil(t, responseErr.Response)

}

func TestNewFromCredentials(t *testing.T) {
	//given
	serveMux := http.NewServeMux()
	server := httptest.NewServer(serveMux)
	defer server.Close()

	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodPost; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
	})

	//when
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//then
	assert.Equal(t, "token", client.Token)
	server.Close()

}

func TestNewFromCredentials_Return_Nil_Client_When_Login_Response_Is_Not_Valid(t *testing.T) {
	//given
	serveMux := http.NewServeMux()
	server := httptest.NewServer(serveMux)
	defer server.Close()

	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodPost; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"errors":[{"message":"404: Not Found","locations":[{"line":2,"column":3}],"path":["response"],"extensions":{"code":"INTERNAL_SERVER_ERROR","response":{"url":"https://bumblebee.test.hazelcast.cloud/customers/api/login","status":404,"statusText":"Not Found","body":{"application":"customer-service","message":"Invalid api key or api secret","status":404,"error":"Not Found","errorCode":"AccessTokenNotFound","timestamp":1599663026907,"path":"http://bumblebee.test.hazelcast.cloud/customers/api/login","method":"POST"}},"exception":{"stacktrace":["Error: 404: Not Found","    at AuthDatasource.<anonymous> (/usr/src/app/node_modules/apollo-datasource-rest/src/RESTDataSource.ts:137:15)","    at Generator.next (<anonymous>)","    at /usr/src/app/node_modules/apollo-datasource-rest/dist/RESTDataSource.js:8:71","    at new Promise (<anonymous>)","    at __awaiter (/usr/src/app/node_modules/apollo-datasource-rest/dist/RESTDataSource.js:4:12)","    at AuthDatasource.errorFromResponse (/usr/src/app/node_modules/apollo-datasource-rest/dist/RESTDataSource.js:75:16)","    at AuthDatasource.<anonymous> (/usr/src/app/node_modules/apollo-datasource-rest/src/RESTDataSource.ts:102:24)","    at Generator.next (<anonymous>)","    at /usr/src/app/node_modules/apollo-datasource-rest/dist/RESTDataSource.js:8:71","    at new Promise (<anonymous>)"]}}}],"data":{"response":null}}`)
	})

	//when
	client, response, err := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//then
	assert.Nil(t, client)
	assert.NotNil(t, response)
	errorResponse := err.(*ErrorResponse)
	s := err.Error()
	assert.Contains(t, s, "Method:POST")
	assert.Contains(t, s, "Status:200 Message:404: Not Found")
	assert.Equal(t, "404: Not Found", errorResponse.Message)
	assert.Equal(t, "200 OK", errorResponse.Response.Status)
	assert.Equal(t, "", errorResponse.CorrelationId)
	server.Close()
}
