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

func TestAzurePeeringServiceOp_List(t *testing.T) {
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

		if strings.Contains(request.Query, "azurePeering") {
			fmt.Fprint(w, `{"data":{"response":[{"id":"1"}]}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	peerings, _, _ := NewAzurePeeringService(client).List(context.TODO(), &models.ListAzurePeeringsInput{})

	//then
	assert.Len(t, *peerings, 1)
	assert.Equal(t, (*peerings)[0].Id, "1")
}

func TestAzurePeeringServiceOp_Fail_On_List(t *testing.T) {
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

		if strings.Contains(request.Query, "azurePeering") {
			fmt.Fprint(w, `{"errors":[{"message":"500: Internal server error"}],"data":{"response":null}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	_, _, listErr := NewAzurePeeringService(client).List(context.TODO(), &models.ListAzurePeeringsInput{})

	//then
	assert.NotNil(t, listErr)
	assert.Contains(t, listErr.Error(), "500: Internal server error")
}

func TestAzurePeeringServiceOp_GetProperties(t *testing.T) {
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

		if strings.Contains(request.Query, "azurePeeringProperties") {
			fmt.Fprint(w, `{"data":{"response":{"appRegistrationId":"1","appRegistrationKey":"2","tenantId":"3","vnetName":"4","subscriptionId":"5","resourceGroupName":"6"}}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	peeringProperties, _, _ := NewAzurePeeringService(client).GetProperties(context.TODO(), &models.GetAzurePeeringPropertiesInput{})

	//then
	assert.Equal(t, (*peeringProperties).AppRegistrationId, "1")
	assert.Equal(t, (*peeringProperties).AppRegistrationKey, "2")
	assert.Equal(t, (*peeringProperties).TenantId, "3")
	assert.Equal(t, (*peeringProperties).VnetName, "4")
	assert.Equal(t, (*peeringProperties).SubscriptionId, "5")
	assert.Equal(t, (*peeringProperties).ResourceGroupName, "6")
}

func TestAzurePeeringServiceOp_Fail_On_GetProperties(t *testing.T) {
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

		if strings.Contains(request.Query, "azurePeeringProperties") {
			fmt.Fprint(w, `{"errors":[{"message":"500: Internal server error"}],"data":{"response":null}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	_, _, acceptErr := NewAzurePeeringService(client).GetProperties(context.TODO(), &models.GetAzurePeeringPropertiesInput{})

	//then
	assert.NotNil(t, acceptErr)
	assert.Contains(t, acceptErr.Error(), "500: Internal server error")
}

func TestAzurePeeringServiceOp_Delete(t *testing.T) {
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

		if strings.Contains(request.Query, "deleteAzurePeering") {
			fmt.Fprint(w, `{"data":{"response":{"status":"OK"}}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	deleteResult, _, _ := NewAzurePeeringService(client).Delete(context.TODO(), &models.DeleteAzurePeeringInput{})

	//then
	assert.Equal(t, (*deleteResult).Status, "OK")
}

func TestAzurePeeringServiceOp_Fail_On_Delete(t *testing.T) {
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

		if strings.Contains(request.Query, "deleteAzurePeering") {
			fmt.Fprint(w, `{"errors":[{"message":"500: Internal server error"}],"data":{"response":null}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	_, _, deleteErr := NewAzurePeeringService(client).Delete(context.TODO(), &models.DeleteAzurePeeringInput{})

	//then
	assert.NotNil(t, deleteErr)
	assert.Contains(t, deleteErr.Error(), "500: Internal server error")
}
