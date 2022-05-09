package hazelcastcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	testClusterId           = "1"
	testClusterIdInt        = 1
	testCustomerId          = 10
	testClusterName         = "test-cluster"
	testClusterPassword     = "hidden"
	testClusterPort         = 30000
	testClusterVersion      = "5.1.1"
	testClusterRegion       = "us-west-2"
	testClusterType         = models.Serverless
	testClusterRunningState = models.Running
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

func TestServerlessClusterService_List(t *testing.T) {
	//given
	server := workingMockServer(t)
	defer server.Close()

	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	clusterResponses, _, _ := NewServerlessClusterService(client).List(context.TODO())

	//then
	assert.Len(t, *(clusterResponses), 2)
}

func TestServerlessClusterService_Fail_On_List(t *testing.T) {
	//given
	server := failingMockServer(t)
	defer server.Close()

	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	_, _, listErr := NewServerlessClusterService(client).List(context.TODO())

	//then
	assert.NotNil(t, listErr)
	assert.Contains(t, listErr.Error(), "Internal server error")
}

func TestServerlessClusterService_Get(t *testing.T) {
	//given
	server := workingMockServer(t)
	defer server.Close()

	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.GetServerlessClusterInput{ClusterId: testClusterId}

	//when
	clusterResponse, _, _ := NewServerlessClusterService(client).Get(context.TODO(), request)

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
	assert.Equal(t, (*clusterResponse).State, testClusterRunningState)
}

func TestServerlessClusterService_Fail_On_Get(t *testing.T) {
	//given
	server := failingMockServer(t)
	defer server.Close()

	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.GetServerlessClusterInput{ClusterId: testClusterId}

	//when
	_, _, getError := NewServerlessClusterService(client).Get(context.TODO(), request)

	//then
	assert.NotNil(t, getError)
	assert.Contains(t, getError.Error(), "Internal server error")
}

func TestServerlessClusterService_Delete(t *testing.T) {
	//given
	server := workingMockServer(t)
	defer server.Close()

	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.ClusterDeleteInput{ClusterId: testClusterId}

	//when
	clusterResponse, _, _ := NewServerlessClusterService(client).Delete(context.TODO(), request)

	//then
	assert.Equal(t, (*clusterResponse).ClusterId, testClusterIdInt)
}

func TestServerlessClusterService_Fail_On_Delete(t *testing.T) {
	//given
	server := failingMockServer(t)
	defer server.Close()

	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.ClusterDeleteInput{ClusterId: testClusterId}

	//when
	_, _, deleteErr := NewServerlessClusterService(client).Delete(context.TODO(), request)

	//then
	assert.NotNil(t, deleteErr)
	assert.Contains(t, deleteErr.Error(), "Internal server error")
}

func TestServerlessClusterService_Stop(t *testing.T) {
	//given
	server := workingMockServer(t)
	defer server.Close()

	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.ClusterStopInput{ClusterId: testClusterId}

	//when
	clusterResponse, _, _ := NewServerlessClusterService(client).Stop(context.TODO(), request)

	//then
	assert.Equal(t, (*clusterResponse).ClusterId, testClusterIdInt)
}

func TestServerlessClusterService_Fail_On_Stop(t *testing.T) {
	//given
	server := failingMockServer(t)
	defer server.Close()

	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.ClusterStopInput{ClusterId: testClusterId}

	//when
	_, _, deleteErr := NewServerlessClusterService(client).Stop(context.TODO(), request)

	//then
	assert.NotNil(t, deleteErr)
	assert.Contains(t, deleteErr.Error(), "Internal server error")
}

func TestServerlessClusterService_Resume(t *testing.T) {
	//given
	server := workingMockServer(t)
	defer server.Close()

	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.ClusterResumeInput{ClusterId: testClusterId}

	//when
	clusterResponse, _, _ := NewServerlessClusterService(client).Resume(context.TODO(), request)

	//then
	assert.Equal(t, (*clusterResponse).ClusterId, testClusterIdInt)
}

func TestServerlessClusterService_Fail_On_Resume(t *testing.T) {
	//given
	server := failingMockServer(t)
	defer server.Close()

	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.ClusterResumeInput{ClusterId: testClusterId}

	//when
	_, _, deleteErr := NewServerlessClusterService(client).Resume(context.TODO(), request)

	//then
	assert.NotNil(t, deleteErr)
	assert.Contains(t, deleteErr.Error(), "Internal server error")
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

		switch {
		case strings.Contains(request.Query, "createServerlessCluster"):
			responseData, responseFileErr := ioutil.ReadFile("testdata/serverless_cluster_create_response.json")
			assert.NoError(t, responseFileErr)
			w.Write(responseData)
		case strings.Contains(request.Query, "deleteCluster"):
			responseData, responseFileErr := ioutil.ReadFile("testdata/serverless_cluster_delete_response.json")
			assert.NoError(t, responseFileErr)
			w.Write(responseData)
		case strings.Contains(request.Query, "stopCluster"):
			responseData, responseFileErr := ioutil.ReadFile("testdata/serverless_cluster_stop_response.json")
			assert.NoError(t, responseFileErr)
			w.Write(responseData)
		case strings.Contains(request.Query, "resumeCluster"):
			responseData, responseFileErr := ioutil.ReadFile("testdata/serverless_cluster_resume_response.json")
			assert.NoError(t, responseFileErr)
			w.Write(responseData)
		case strings.Contains(request.Query, "clusters"):
			responseData, responseFileErr := ioutil.ReadFile("testdata/serverless_cluster_list_response.json")
			assert.NoError(t, responseFileErr)
			w.Write(responseData)
		case strings.Contains(request.Query, "cluster"):
			responseData, responseFileErr := ioutil.ReadFile("testdata/serverless_cluster_get_response.json")
			assert.NoError(t, responseFileErr)
			w.Write(responseData)
		default:
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

func TestStarterClusterServiceOp_ListUploadedArtifacts(t *testing.T) {
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
		if strings.Contains(request.Query, "customClasses") {
			fmt.Fprint(w, `{"data": {"response": [{"id": 108, "name": "apple-jdk11.jar", "status": "FINISHED"}]}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}
	})

	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.ListUploadedArtifactsInput{}

	//when
	list, _, _ := NewServerlessClusterService(client).ListUploadedArtifacts(context.TODO(), request)

	// then
	assert.NotNil(t, list)
	assert.Len(t, *list, 1)
	assert.Equal(t, (*list)[0].Id, 108)
}

func TestStarterClusterServiceOp_DeleteArtifact(t *testing.T) {
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
		if strings.Contains(request.Query, "deleteCustomClassArtifact") {
			fmt.Fprint(w, `{"data": {"response": {"id": 108, "name": "apple-jdk11.jar", "status": "FINISHED"}}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}
	})

	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.DeleteArtifactInput{
		ClusterId: "109",
	}

	//when
	art, _, _ := NewServerlessClusterService(client).DeleteArtifact(context.TODO(), request)

	// then
	assert.NotNil(t, art)
	assert.Equal(t, (*art).Id, 108)
}
