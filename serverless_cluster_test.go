package hazelcastcloud

import (
	"context"
	"encoding/json"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	testClusterId       = "1"
	testCustomerId      = 10
	testClusterName     = "test-cluster"
	testClusterPassword = "hidden"
	testClusterPort     = 30000
	testClusterVersion  = "5.1.1"
	testClusterRegion   = "us-west-2"
	testClusterType     = models.Serverless
)

var createServerlessClusterRequest = models.CreateServerlessClusterInput{
	Name:        testClusterName,
	Region:      testClusterRegion,
	ClusterType: testClusterType,
}

func TestServerlessClusterService_Create(t *testing.T) {
	//given
	server := workingMockServer(t)
	defer server.Close()
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	clusterResponse, _, _ := NewServerlessClusterService(client).Create(context.TODO(), &createServerlessClusterRequest)

	//then
	assert.Equal(t, (*clusterResponse).Id, testClusterId)
	assert.Equal(t, (*clusterResponse).CustomerId, testCustomerId)
	assert.Equal(t, (*clusterResponse).Name, testClusterName)
	assert.Equal(t, (*clusterResponse).Password, testClusterPassword)
	assert.Equal(t, (*clusterResponse).Port, testClusterPort)
	assert.Equal(t, (*clusterResponse).HazelcastVersion, testClusterVersion)
	assert.True(t, (*clusterResponse).IsAutoScalingEnabled)
	assert.True(t, (*clusterResponse).IsHotBackupEnabled)
	assert.True(t, (*clusterResponse).IsHotRestartEnabled)
	assert.False(t, (*clusterResponse).IsIpWhitelistEnabled)
	assert.True(t, (*clusterResponse).IsTlsEnabled)
}

func TestServerlessClusterService_Fail_On_Create(t *testing.T) {
	//given
	server := failingMockServer(t)
	defer server.Close()
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	_, _, createErr := NewServerlessClusterService(client).Create(context.TODO(), &createServerlessClusterRequest)

	//then
	assert.NotNil(t, createErr)
	assert.Contains(t, createErr.Error(), "Internal server error")
}

func workingMockServer(t *testing.T) *httptest.Server {
	serveMux := http.NewServeMux()
	server := httptest.NewServer(serveMux)
	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodPost; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		var request GraphQLQuery
		json.NewDecoder(r.Body).Decode(&request)

		if strings.Contains(request.Query, "createServerlessCluster") {
			responseData, responseFileErr := ioutil.ReadFile("testdata/serverless_cluster_create_response.json")
			assert.NoError(t, responseFileErr)
			w.Write(responseData)
		} else {
			responseData, responseFileErr := ioutil.ReadFile("testdata/access_token_response.json")
			assert.NoError(t, responseFileErr)
			w.Write(responseData)
		}
	})
	return server
}

func failingMockServer(t *testing.T) *httptest.Server {
	serveMux := http.NewServeMux()
	server := httptest.NewServer(serveMux)
	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodPost; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		var request GraphQLQuery
		json.NewDecoder(r.Body).Decode(&request)

		if strings.Contains(request.Query, "apiKey") {
			responseData, responseFileErr := ioutil.ReadFile("testdata/access_token_response.json")
			assert.NoError(t, responseFileErr)
			w.Write(responseData)
		} else {
			responseData, responseFileErr := ioutil.ReadFile("testdata/common_failed_response.json")
			assert.NoError(t, responseFileErr)
			w.Write(responseData)
		}
	})
	return server
}
