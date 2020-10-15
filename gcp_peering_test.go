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

func TestGcpPeeringServiceOp_List(t *testing.T) {
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

		if strings.Contains(request.Query, "gcpPeerings") {
			fmt.Fprint(w, `{"data":{"response":[{"id":"1","projectId":"2","networkName":"test"}]}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	peerings, _, _ := NewGcpPeeringService(client).List(context.TODO(), &models.ListGcpPeeringsInput{})

	//then
	assert.Len(t, *peerings, 1)
	assert.Equal(t, (*peerings)[0].Id, "1")
	assert.Equal(t, (*peerings)[0].ProjectId, "2")
	assert.Equal(t, (*peerings)[0].NetworkName, "test")
}

func TestGcpPeeringServiceOp_GetProperties(t *testing.T) {
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

		if strings.Contains(request.Query, "gcpPeeringProperties") {
			fmt.Fprint(w, `{"data":{"response":{"projectId":"1","networkName":"test"}}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	peeringProperties, _, _ := NewGcpPeeringService(client).GetProperties(context.TODO(), &models.GetGcpPeeringPropertiesInput{})

	//then
	assert.Equal(t, (*peeringProperties).ProjectId, "1")
	assert.Equal(t, (*peeringProperties).NetworkName, "test")
}

func TestGcpPeeringServiceOp_Delete(t *testing.T) {
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

		if strings.Contains(request.Query, "deleteGcpPeering") {
			fmt.Fprint(w, `{"data":{"response":{"status":"OK"}}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	deleteResult, _, _ := NewGcpPeeringService(client).Delete(context.TODO(), &models.DeleteGcpPeeringInput{})

	//then
	assert.Equal(t, (*deleteResult).Status, "OK")
}

func TestGcpPeeringServiceOp_Accept(t *testing.T) {
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

		if strings.Contains(request.Query, "acceptGcpPeering") {
			fmt.Fprint(w, `{"data":{"response":{"status":"OK"}}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	deleteResult, _, _ := NewGcpPeeringService(client).Accept(context.TODO(), &models.AcceptGcpPeeringInput{})

	//then
	assert.Equal(t, (*deleteResult).Status, "OK")
}