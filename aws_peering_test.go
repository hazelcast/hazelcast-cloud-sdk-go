package hazelcastcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAwsPeeringServiceOp_List(t *testing.T) {
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

		if strings.Contains(request.Query, "awsPeerings") {
			fmt.Fprint(w, `{"data":{"response":[{"id":"1"}]}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	peerings, _, _ := NewAwsPeeringService(client).List(context.TODO(), &models.ListAwsPeeringsInput{})

	//then
	assert.Len(t, *peerings, 1)
	assert.Equal(t, (*peerings)[0].Id, "1")
}

func TestAwsPeeringServiceOp_GetProperties(t *testing.T) {
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

		if strings.Contains(request.Query, "awsPeeringProperties") {
			fmt.Fprint(w, `{"data":{"response":{"vpcId":"1","ownerId":"2","region":"us-west-2"}}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	peeringProperties, _, _ := NewAwsPeeringService(client).GetProperties(context.TODO(), &models.GetAwsPeeringPropertiesInput{})

	//then
	assert.Equal(t, (*peeringProperties).VpcId, "1")
	assert.Equal(t, (*peeringProperties).OwnerId, "2")
	assert.Equal(t, (*peeringProperties).Region, "us-west-2")
}

func TestAwsPeeringServiceOp_Delete(t *testing.T) {
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

		if strings.Contains(request.Query, "deleteAwsPeering") {
			fmt.Fprint(w, `{"data":{"response":{"status":"OK"}}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	deleteResult, _, _ := NewAwsPeeringService(client).Delete(context.TODO(), &models.DeleteAwsPeeringInput{})

	//then
	assert.Equal(t, (*deleteResult).Status, "OK")
}

func TestAwsPeeringServiceOp_Accept(t *testing.T) {
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

		if strings.Contains(request.Query, "acceptAwsPeering") {
			fmt.Fprint(w, `{"data":{"response":{"status":"OK"}}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	deleteResult, _, _ := NewAwsPeeringService(client).Accept(context.TODO(), &models.AcceptAwsPeeringInput{})

	//then
	assert.Equal(t, (*deleteResult).Status, "OK")
}