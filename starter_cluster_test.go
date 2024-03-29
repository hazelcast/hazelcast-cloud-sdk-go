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

func TestStarterClusterServiceOp_Create(t *testing.T) {
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

		if strings.Contains(request.Query, "createStarterCluster") {
			fmt.Fprint(w, `{"data":{"response":{"id":"123456","customerId":17000,"teamId":null,"name":"test-cluster","password":"e6838c596a0342d4918cf89a8d071023","port":34005,"hazelcastVersion":"3.12.2-4","isAutoScalingEnabled":false,"isHotBackupEnabled":false}}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.CreateStarterClusterInput{}

	//when
	clusterResponse, _, _ := NewStarterClusterService(client).Create(context.TODO(), request)

	//then
	assert.Equal(t, (*clusterResponse).Id, "123456")
	assert.Equal(t, (*clusterResponse).CustomerId, 17000)
	assert.Equal(t, (*clusterResponse).Name, "test-cluster")
	assert.Equal(t, (*clusterResponse).Password, "e6838c596a0342d4918cf89a8d071023")
	assert.Equal(t, (*clusterResponse).Port, 34005)
	assert.Equal(t, (*clusterResponse).HazelcastVersion, "3.12.2-4")
	assert.False(t, (*clusterResponse).IsAutoScalingEnabled)
	assert.False(t, (*clusterResponse).IsHotBackupEnabled)
}

func TestStarterClusterServiceOp_Fail_On_Create(t *testing.T) {
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

		if strings.Contains(request.Query, "createStarterCluster") {
			fmt.Fprint(w, `{"errors":[{"message":"500: Internal server error"}],"data":{"response":null}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.CreateStarterClusterInput{}

	//when
	_, _, createErr := NewStarterClusterService(client).Create(context.TODO(), request)

	//then
	assert.NotNil(t, createErr)
	assert.Contains(t, createErr.Error(), "500: Internal server error")
}

func TestStarterClusterServiceOp_Get(t *testing.T) {
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

		if strings.Contains(request.Query, "cluster") {
			fmt.Fprint(w, `{"data":{"response":{"id":"123456","customerId":17000,"teamId":null,"name":"test-cluster","password":"e6838c596a0342d4918cf89a8d071023","port":34005,"hazelcastVersion":"3.12.2-4","isAutoScalingEnabled":false,"isHotBackupEnabled":false}}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.GetStarterClusterInput{}

	//when
	clusterResponse, _, _ := NewStarterClusterService(client).Get(context.TODO(), request)

	//then
	assert.Equal(t, (*clusterResponse).Id, "123456")
	assert.Equal(t, (*clusterResponse).CustomerId, 17000)
	assert.Equal(t, (*clusterResponse).Name, "test-cluster")
	assert.Equal(t, (*clusterResponse).Password, "e6838c596a0342d4918cf89a8d071023")
	assert.Equal(t, (*clusterResponse).Port, 34005)
	assert.Equal(t, (*clusterResponse).HazelcastVersion, "3.12.2-4")
	assert.False(t, (*clusterResponse).IsAutoScalingEnabled)
	assert.False(t, (*clusterResponse).IsHotBackupEnabled)
}

func TestStarterClusterServiceOp_Fail_On_Get(t *testing.T) {
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

		if strings.Contains(request.Query, "cluster") {
			fmt.Fprint(w, `{"errors":[{"message":"500: Internal server error"}],"data":{"response":null}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.GetStarterClusterInput{}

	//when
	_, _, getErr := NewStarterClusterService(client).Get(context.TODO(), request)

	//then
	assert.NotNil(t, getErr)
	assert.Contains(t, getErr.Error(), "500: Internal server error")
}

func TestStarterClusterServiceOp_List(t *testing.T) {
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

		if strings.Contains(request.Query, "cluster") {
			fmt.Fprint(w, `{"data":{"response":[{"id":"427","name":"demo"},{"id":"429","name":"demo-play2"},{"id":"437","name":"demo-play3"},{"id":"438","name":"demo-sdk"},{"id":"439","name":"mycluster"},{"id":"445","name":"test-cluster"}]}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	clusterResponses, _, _ := NewStarterClusterService(client).List(context.TODO())

	//then
	assert.Len(t, *(clusterResponses), 6)
}

func TestStarterClusterServiceOp_Fail_On_List(t *testing.T) {
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

		if strings.Contains(request.Query, "cluster") {
			fmt.Fprint(w, `{"errors":[{"message":"500: Internal server error"}],"data":{"response":null}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	_, _, listErr := NewStarterClusterService(client).List(context.TODO())

	//then
	assert.NotNil(t, listErr)
	assert.Contains(t, listErr.Error(), "500: Internal server error")
}

func TestStarterClusterServiceOp_Resume(t *testing.T) {
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

		if strings.Contains(request.Query, "resumeCluster") {
			fmt.Fprint(w, `{"data":{"response":{"clusterId":427}}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.ClusterResumeInput{}

	//when
	clusterIdResponse, _, _ := NewStarterClusterService(client).Resume(context.TODO(), request)

	//then
	assert.Equal(t, clusterIdResponse.ClusterId, 427)
}

func TestStarterClusterServiceOp_Fail_On_Resume(t *testing.T) {
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

		if strings.Contains(request.Query, "resumeCluster") {
			fmt.Fprint(w, `{"errors":[{"message":"500: Internal server error"}],"data":{"response":null}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.ClusterResumeInput{}

	//when
	_, _, resumeErr := NewStarterClusterService(client).Resume(context.TODO(), request)

	//then
	assert.NotNil(t, resumeErr)
	assert.Contains(t, resumeErr.Error(), "500: Internal server error")
}

func TestStarterClusterServiceOp_Stop(t *testing.T) {
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

		if strings.Contains(request.Query, "stopCluster") {
			fmt.Fprint(w, `{"data":{"response":{"clusterId":427}}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.ClusterStopInput{}

	//when
	clusterIdResponse, _, _ := NewStarterClusterService(client).Stop(context.TODO(), request)

	//then
	assert.Equal(t, clusterIdResponse.ClusterId, 427)
}

func TestStarterClusterServiceOp_Fail_On_Stop(t *testing.T) {
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

		if strings.Contains(request.Query, "stopCluster") {
			fmt.Fprint(w, `{"errors":[{"message":"500: Internal server error"}],"data":{"response":null}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.ClusterStopInput{}

	//when
	_, _, stopErr := NewStarterClusterService(client).Stop(context.TODO(), request)

	//then
	assert.NotNil(t, stopErr)
	assert.Contains(t, stopErr.Error(), "500: Internal server error")
}

func TestStarterClusterServiceOp_Delete(t *testing.T) {
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

		if strings.Contains(request.Query, "deleteCluster") {
			fmt.Fprint(w, `{"data":{"response":{"clusterId":427}}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.ClusterDeleteInput{}

	//when
	clusterIdResponse, _, _ := NewStarterClusterService(client).Delete(context.TODO(), request)

	//then
	assert.Equal(t, clusterIdResponse.ClusterId, 427)
}

func TestStarterClusterServiceOp_Fail_On_Delete(t *testing.T) {
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

		if strings.Contains(request.Query, "deleteCluster") {
			fmt.Fprint(w, `{"errors":[{"message":"500: Internal server error"}],"data":{"response":null}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.ClusterDeleteInput{}

	//when
	_, _, deleteErr := NewStarterClusterService(client).Delete(context.TODO(), request)

	//then
	assert.NotNil(t, deleteErr)
	assert.Contains(t, deleteErr.Error(), "500: Internal server error")
}

func ExampleStarterClusterService_create() {
	client, _, _ := New()
	create, _, _ := client.StarterCluster.Create(context.Background(), &models.CreateStarterClusterInput{
		Name:                 "example-cluster",
		CloudProvider:        "aws",
		Region:               "us-west-2",
		ClusterType:          models.Free,
		HazelcastVersion:     "3.12",
		TotalMemory:          0.2,
		IsAutoScalingEnabled: false,
		IsHotBackupEnabled:   false,
		IsHotRestartEnabled:  false,
		IsIPWhitelistEnabled: false,
		IsTLSEnabled:         false,
	})
	fmt.Printf("Result: %#v", create)
	//Output:Result: &models.Cluster{Id:"53805", CustomerId:10090, Name:"example-cluster", Password:"2d2a9e5088a94893a1b27fc060efc2e4", Port:31002, HazelcastVersion:"3.12.9", IsAutoScalingEnabled:false, IsHotBackupEnabled:false, IsHotRestartEnabled:false, IsIpWhitelistEnabled:false, IsTlsEnabled:false, ProductType:struct { Name models.ProductTypeName "json:\"name\""; IsFree bool "json:\"isFree\"" }{Name:"Starter", IsFree:true}, State:"PENDING", CreatedAt:"2020-09-08T07:15:13.000Z", StartedAt:"2020-09-08T07:15:13.000Z", StoppedAt:"", Progress:struct { Status string "json:\"status\""; TotalItemCount int "json:\"totalItemCount\""; CompletedItemCount int "json:\"completedItemCount\"" }{Status:"Preparing", TotalItemCount:4, CompletedItemCount:0}, CloudProvider:struct { Name string "json:\"name\""; Region string "json:\"region\""; AvailabilityZones []string "json:\"availabilityZones\"" }{Name:"aws", Region:"us-west-2", AvailabilityZones:[]string{"us-west-2a"}}, DiscoveryTokens:[]models.DiscoveryToken{models.DiscoveryToken{Source:"default", Token:"fd7Zzw2xbxQ692nHqyPUh7Bq9bSLxm2u8tvflzZQ8eqzD1TmUx"}}, Specs:struct { TotalMemory float64 "json:\"totalMemory\""; HeapMemory int "json:\"heapMemory\""; NativeMemory int "json:\"nativeMemory\""; Cpu int "json:\"cpu\""; InstanceType string "json:\"instanceType\""; InstancePerZone int "json:\"instancePerZone\"" }{TotalMemory:0.2, HeapMemory:0, NativeMemory:0, Cpu:0, InstanceType:"", InstancePerZone:0}, Networking:struct { Type string "json:\"type\""; CidrBlock string "json:\"cidrBlock\""; Peering struct { IsEnabled bool "json:\"is_enabled\"" } "json:\"peering\""; PrivateLink struct { Url string "json:\"url\""; State string "json:\"state\"" } "json:\"privateLink\"" }{Type:"", CidrBlock:"", Peering:struct { IsEnabled bool "json:\"is_enabled\"" }{IsEnabled:false}, PrivateLink:struct { Url string "json:\"url\""; State string "json:\"state\"" }{Url:"", State:""}}, DataStructures:models.DataStructureResponse{MapConfigs:[]models.MapConfigResponse{}, JCacheConfigs:[]models.JCacheConfigResponse{}, ReplicatedMapConfigs:[]models.ReplicatedMapConfigResponse{}, QueueConfigs:[]models.QueueConfigResponse{}, SetConfigs:[]models.SetConfigResponse{}, ListConfigs:[]models.ListConfigResponse{}, TopicConfigs:[]models.TopicConfigResponse{}, MultiMapConfigs:[]models.MultiMapConfigResponse{}, RingBufferConfigs:[]models.RingBufferConfigResponse{}, ReliableTopicConfigs:[]models.ReliableTopicConfigResponse{}}}
}

func ExampleStarterClusterService_delete() {
	client, _, _ := New()
	cluster, _, _ := client.StarterCluster.Delete(context.Background(), &models.ClusterDeleteInput{ClusterId: "53805"})
	fmt.Printf("Result: %#v", cluster)
	//Output:Result: &models.ClusterId{ClusterId:53805}
}

func ExampleStarterClusterService_get() {
	client, _, _ := New()
	cluster, _, _ := client.StarterCluster.Get(context.Background(), &models.GetStarterClusterInput{ClusterId: "53805"})
	fmt.Printf("Result: %#v", cluster)
	//Output:Result: &models.Cluster{Id:"53805", CustomerId:10090, Name:"example-cluster", Password:"2d2a9e5088a94893a1b27fc060efc2e4", Port:31002, HazelcastVersion:"3.12.9", IsAutoScalingEnabled:false, IsHotBackupEnabled:false, IsHotRestartEnabled:false, IsIpWhitelistEnabled:false, IsTlsEnabled:false, ProductType:struct { Name models.ProductTypeName "json:\"name\""; IsFree bool "json:\"isFree\"" }{Name:"Starter", IsFree:true}, State:"RUNNING", CreatedAt:"2020-09-08T07:15:13.000Z", StartedAt:"2020-09-08T07:16:37.000Z", StoppedAt:"", Progress:struct { Status string "json:\"status\""; TotalItemCount int "json:\"totalItemCount\""; CompletedItemCount int "json:\"completedItemCount\"" }{Status:"Preparing", TotalItemCount:0, CompletedItemCount:0}, CloudProvider:struct { Name string "json:\"name\""; Region string "json:\"region\""; AvailabilityZones []string "json:\"availabilityZones\"" }{Name:"aws", Region:"us-west-2", AvailabilityZones:[]string{"us-west-2a"}}, DiscoveryTokens:[]models.DiscoveryToken{models.DiscoveryToken{Source:"default", Token:"fd7Zzw2xbxQ692nHqyPUh7Bq9bSLxm2u8tvflzZQ8eqzD1TmUx"}}, Specs:struct { TotalMemory float64 "json:\"totalMemory\""; HeapMemory int "json:\"heapMemory\""; NativeMemory int "json:\"nativeMemory\""; Cpu int "json:\"cpu\""; InstanceType string "json:\"instanceType\""; InstancePerZone int "json:\"instancePerZone\"" }{TotalMemory:0.2, HeapMemory:0, NativeMemory:0, Cpu:0, InstanceType:"", InstancePerZone:0}, Networking:struct { Type string "json:\"type\""; CidrBlock string "json:\"cidrBlock\""; Peering struct { IsEnabled bool "json:\"is_enabled\"" } "json:\"peering\""; PrivateLink struct { Url string "json:\"url\""; State string "json:\"state\"" } "json:\"privateLink\"" }{Type:"", CidrBlock:"", Peering:struct { IsEnabled bool "json:\"is_enabled\"" }{IsEnabled:false}, PrivateLink:struct { Url string "json:\"url\""; State string "json:\"state\"" }{Url:"", State:""}}, DataStructures:models.DataStructureResponse{MapConfigs:[]models.MapConfigResponse{}, JCacheConfigs:[]models.JCacheConfigResponse{}, ReplicatedMapConfigs:[]models.ReplicatedMapConfigResponse{}, QueueConfigs:[]models.QueueConfigResponse{}, SetConfigs:[]models.SetConfigResponse{}, ListConfigs:[]models.ListConfigResponse{}, TopicConfigs:[]models.TopicConfigResponse{}, MultiMapConfigs:[]models.MultiMapConfigResponse{}, RingBufferConfigs:[]models.RingBufferConfigResponse{}, ReliableTopicConfigs:[]models.ReliableTopicConfigResponse{}}}
}

func ExampleStarterClusterService_stop() {
	client, _, _ := New()
	cluster, _, _ := client.StarterCluster.Stop(context.Background(), &models.ClusterStopInput{ClusterId: "53805"})
	fmt.Printf("Result: %#v", cluster)
	//Output:Result: &models.ClusterId{ClusterId:53805}
}

func ExampleStarterClusterService_resume() {
	client, _, _ := New()
	cluster, _, _ := client.StarterCluster.Resume(context.Background(), &models.ClusterResumeInput{ClusterId: "53805"})
	fmt.Printf("Result: %#v", cluster)
	//Output:Result: &models.ClusterId{ClusterId:53805}
}

func ExampleStarterClusterService_list() {
	client, _, _ := New()
	clusters, _, _ := client.StarterCluster.List(context.Background())
	fmt.Printf("Result: %#v", clusters)
	//Output:Result: &[]models.Cluster{models.Cluster{Id:"53805", CustomerId:10090, Name:"example-cluster", Password:"", Port:31002, HazelcastVersion:"3.12.9", IsAutoScalingEnabled:false, IsHotBackupEnabled:false, IsHotRestartEnabled:false, IsIpWhitelistEnabled:false, IsTlsEnabled:false, ProductType:struct { Name models.ProductTypeName "json:\"name\""; IsFree bool "json:\"isFree\"" }{Name:"Starter", IsFree:true}, State:"STOPPED", CreatedAt:"2020-09-08T07:25:40.000Z", StartedAt:"2020-09-08T07:26:48.000Z", StoppedAt:"", Progress:struct { Status string "json:\"status\""; TotalItemCount int "json:\"totalItemCount\""; CompletedItemCount int "json:\"completedItemCount\"" }{Status:"Preparing", TotalItemCount:0, CompletedItemCount:0}, CloudProvider:struct { Name string "json:\"name\""; Region string "json:\"region\""; AvailabilityZones []string "json:\"availabilityZones\"" }{Name:"aws", Region:"us-west-2", AvailabilityZones:[]string{"us-west-2a"}}, DiscoveryTokens:[]models.DiscoveryToken{}, Specs:struct { TotalMemory float64 "json:\"totalMemory\""; HeapMemory int "json:\"heapMemory\""; NativeMemory int "json:\"nativeMemory\""; Cpu int "json:\"cpu\""; InstanceType string "json:\"instanceType\""; InstancePerZone int "json:\"instancePerZone\"" }{TotalMemory:0.2, HeapMemory:0, NativeMemory:0, Cpu:0, InstanceType:"", InstancePerZone:0}, Networking:struct { Type string "json:\"type\""; CidrBlock string "json:\"cidrBlock\""; Peering struct { IsEnabled bool "json:\"is_enabled\"" } "json:\"peering\""; PrivateLink struct { Url string "json:\"url\""; State string "json:\"state\"" } "json:\"privateLink\"" }{Type:"", CidrBlock:"", Peering:struct { IsEnabled bool "json:\"is_enabled\"" }{IsEnabled:false}, PrivateLink:struct { Url string "json:\"url\""; State string "json:\"state\"" }{Url:"", State:""}}, DataStructures:models.DataStructureResponse{MapConfigs:[]models.MapConfigResponse{}, JCacheConfigs:[]models.JCacheConfigResponse{}, ReplicatedMapConfigs:[]models.ReplicatedMapConfigResponse{}, QueueConfigs:[]models.QueueConfigResponse{}, SetConfigs:[]models.SetConfigResponse{}, ListConfigs:[]models.ListConfigResponse{}, TopicConfigs:[]models.TopicConfigResponse{}, MultiMapConfigs:[]models.MultiMapConfigResponse{}, RingBufferConfigs:[]models.RingBufferConfigResponse{}, ReliableTopicConfigs:[]models.ReliableTopicConfigResponse{}}}}
}
