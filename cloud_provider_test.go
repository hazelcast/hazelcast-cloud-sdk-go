package hazelcastcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCloudProviderServiceOp_List(t *testing.T) {
	//given
	serveMux := http.NewServeMux()
	server := httptest.NewServer(serveMux)
	defer server.Close()

	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodPost; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		var request GraphQLQuery
		json.NewDecoder(r.Body).Decode(&request)

		if strings.Contains(request.Query, "cloudProviders") {
			fmt.Fprint(w, `{"data":{"response":[{"name":"aws"},{"name":"azure"},{"name":"google"}]}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	cloudProviders, _, _ := NewCloudProviderService(client).List(context.TODO())

	//then
	assert.Len(t, *cloudProviders, 3)
	assert.Equal(t, (*cloudProviders)[0].Name, "aws")
	assert.Equal(t, (*cloudProviders)[1].Name, "azure")
	assert.Equal(t, (*cloudProviders)[2].Name, "google")
}

func ExampleCloudProviderService_list() {
	client, _, _ := New()
	cloudProviders, _, _ := client.CloudProvider.List(context.Background())
	fmt.Printf("Result: %#v", cloudProviders)
	//Output:Result: &[]models.CloudProvider{models.CloudProvider{Name:"aws", IsEnabledForStarter:true, IsEnabledForEnterprise:true}, models.CloudProvider{Name:"azure", IsEnabledForStarter:false, IsEnabledForEnterprise:true}, models.CloudProvider{Name:"google", IsEnabledForStarter:false, IsEnabledForEnterprise:false}}
}
