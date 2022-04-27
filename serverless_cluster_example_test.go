package hazelcastcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

func ExampleServerlessClusterService_create() {
	server := exampleTestMockServer()
	defer server.Close()
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	cluster, _, _ := client.ServerlessCluster.Create(context.Background(), &createServerlessClusterRequest)
	fmt.Printf("Result: %#v", cluster)
	//Output:Result: &models.Cluster{Id:"1", CustomerId:10, Name:"test-cluster", ReleaseName:"pr-1", Password:"hidden", Port:30000, HazelcastVersion:"5.1.1", IsAutoScalingEnabled:true, IsHotBackupEnabled:true, IsHotRestartEnabled:true, IsIpWhitelistEnabled:false, IsTlsEnabled:true, ProductType:struct { Name models.ProductTypeName "json:\"name\""; IsFree bool "json:\"isFree\"" }{Name:"STARTER", IsFree:false}, ClusterType:struct { Name models.ClusterType "json:\"name\"" }{Name:"SERVERLESS"}, State:"PENDING", CreatedAt:"2022-04-21T11:49:02.000Z", StartedAt:"2022-04-21T11:49:02.000Z", StoppedAt:"", Progress:struct { Status string "json:\"status\""; TotalItemCount int "json:\"totalItemCount\""; CompletedItemCount int "json:\"completedItemCount\"" }{Status:"Preparing", TotalItemCount:4, CompletedItemCount:0}, CloudProvider:struct { Name string "json:\"name\""; Region string "json:\"region\""; AvailabilityZones []string "json:\"availabilityZones\"" }{Name:"aws", Region:"us-west-2", AvailabilityZones:[]string{"us-west-2a"}}, DiscoveryTokens:[]models.DiscoveryToken{models.DiscoveryToken{Source:"default", Token:"hidden"}}, Specs:struct { TotalMemory float64 "json:\"totalMemory\""; HeapMemory int "json:\"heapMemory\""; NativeMemory int "json:\"nativeMemory\""; Cpu int "json:\"cpu\""; InstanceType string "json:\"instanceType\""; InstancePerZone int "json:\"instancePerZone\"" }{TotalMemory:10, HeapMemory:0, NativeMemory:0, Cpu:0, InstanceType:"", InstancePerZone:0}, Networking:struct { Type string "json:\"type\""; CidrBlock string "json:\"cidrBlock\""; Peering struct { IsEnabled bool "json:\"is_enabled\"" } "json:\"peering\""; PrivateLink struct { Url string "json:\"url\""; State string "json:\"state\"" } "json:\"privateLink\"" }{Type:"", CidrBlock:"", Peering:struct { IsEnabled bool "json:\"is_enabled\"" }{IsEnabled:false}, PrivateLink:struct { Url string "json:\"url\""; State string "json:\"state\"" }{Url:"", State:""}}, DataStructures:models.DataStructureResponse{MapConfigs:[]models.MapConfigResponse{}, JCacheConfigs:[]models.JCacheConfigResponse{}, ReplicatedMapConfigs:[]models.ReplicatedMapConfigResponse{}, QueueConfigs:[]models.QueueConfigResponse{}, SetConfigs:[]models.SetConfigResponse{}, ListConfigs:[]models.ListConfigResponse{}, TopicConfigs:[]models.TopicConfigResponse{}, MultiMapConfigs:[]models.MultiMapConfigResponse{}, RingBufferConfigs:[]models.RingBufferConfigResponse{}, ReliableTopicConfigs:[]models.ReliableTopicConfigResponse{}}}
}

func ExampleServerlessClusterService_list() {
	server := exampleTestMockServer()
	defer server.Close()
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	clusters, _, _ := client.ServerlessCluster.List(context.Background())
	fmt.Printf("Result: %#v", clusters)
	//Output:Result: &[]models.Cluster{models.Cluster{Id:"1", CustomerId:10, Name:"test-cluster1", ReleaseName:"pr-1", Password:"hidden", Port:30000, HazelcastVersion:"5.1.1", IsAutoScalingEnabled:true, IsHotBackupEnabled:true, IsHotRestartEnabled:true, IsIpWhitelistEnabled:false, IsTlsEnabled:true, ProductType:struct { Name models.ProductTypeName "json:\"name\""; IsFree bool "json:\"isFree\"" }{Name:"STARTER", IsFree:false}, ClusterType:struct { Name models.ClusterType "json:\"name\"" }{Name:"SERVERLESS"}, State:"RUNNING", CreatedAt:"2022-04-21T11:49:02.000Z", StartedAt:"2022-04-21T11:49:02.000Z", StoppedAt:"", Progress:struct { Status string "json:\"status\""; TotalItemCount int "json:\"totalItemCount\""; CompletedItemCount int "json:\"completedItemCount\"" }{Status:"Running", TotalItemCount:4, CompletedItemCount:4}, CloudProvider:struct { Name string "json:\"name\""; Region string "json:\"region\""; AvailabilityZones []string "json:\"availabilityZones\"" }{Name:"aws", Region:"us-west-2", AvailabilityZones:[]string{"us-west-2a"}}, DiscoveryTokens:[]models.DiscoveryToken{models.DiscoveryToken{Source:"default", Token:"hidden"}}, Specs:struct { TotalMemory float64 "json:\"totalMemory\""; HeapMemory int "json:\"heapMemory\""; NativeMemory int "json:\"nativeMemory\""; Cpu int "json:\"cpu\""; InstanceType string "json:\"instanceType\""; InstancePerZone int "json:\"instancePerZone\"" }{TotalMemory:10, HeapMemory:0, NativeMemory:0, Cpu:0, InstanceType:"", InstancePerZone:0}, Networking:struct { Type string "json:\"type\""; CidrBlock string "json:\"cidrBlock\""; Peering struct { IsEnabled bool "json:\"is_enabled\"" } "json:\"peering\""; PrivateLink struct { Url string "json:\"url\""; State string "json:\"state\"" } "json:\"privateLink\"" }{Type:"", CidrBlock:"", Peering:struct { IsEnabled bool "json:\"is_enabled\"" }{IsEnabled:false}, PrivateLink:struct { Url string "json:\"url\""; State string "json:\"state\"" }{Url:"", State:""}}, DataStructures:models.DataStructureResponse{MapConfigs:[]models.MapConfigResponse{}, JCacheConfigs:[]models.JCacheConfigResponse{}, ReplicatedMapConfigs:[]models.ReplicatedMapConfigResponse{}, QueueConfigs:[]models.QueueConfigResponse{}, SetConfigs:[]models.SetConfigResponse{}, ListConfigs:[]models.ListConfigResponse{}, TopicConfigs:[]models.TopicConfigResponse{}, MultiMapConfigs:[]models.MultiMapConfigResponse{}, RingBufferConfigs:[]models.RingBufferConfigResponse{}, ReliableTopicConfigs:[]models.ReliableTopicConfigResponse{}}}, models.Cluster{Id:"2", CustomerId:10, Name:"test-cluster2", ReleaseName:"pr-2", Password:"hidden", Port:30000, HazelcastVersion:"5.1.1", IsAutoScalingEnabled:true, IsHotBackupEnabled:true, IsHotRestartEnabled:true, IsIpWhitelistEnabled:false, IsTlsEnabled:true, ProductType:struct { Name models.ProductTypeName "json:\"name\""; IsFree bool "json:\"isFree\"" }{Name:"STARTER", IsFree:false}, ClusterType:struct { Name models.ClusterType "json:\"name\"" }{Name:"SERVERLESS"}, State:"PENDING", CreatedAt:"2022-04-26T13:05:06.000Z", StartedAt:"2022-04-26T13:05:06.000Z", StoppedAt:"", Progress:struct { Status string "json:\"status\""; TotalItemCount int "json:\"totalItemCount\""; CompletedItemCount int "json:\"completedItemCount\"" }{Status:"Preparing", TotalItemCount:4, CompletedItemCount:0}, CloudProvider:struct { Name string "json:\"name\""; Region string "json:\"region\""; AvailabilityZones []string "json:\"availabilityZones\"" }{Name:"aws", Region:"us-west-2", AvailabilityZones:[]string{"us-west-2a"}}, DiscoveryTokens:[]models.DiscoveryToken{models.DiscoveryToken{Source:"default", Token:"hidden"}}, Specs:struct { TotalMemory float64 "json:\"totalMemory\""; HeapMemory int "json:\"heapMemory\""; NativeMemory int "json:\"nativeMemory\""; Cpu int "json:\"cpu\""; InstanceType string "json:\"instanceType\""; InstancePerZone int "json:\"instancePerZone\"" }{TotalMemory:10, HeapMemory:0, NativeMemory:0, Cpu:0, InstanceType:"", InstancePerZone:0}, Networking:struct { Type string "json:\"type\""; CidrBlock string "json:\"cidrBlock\""; Peering struct { IsEnabled bool "json:\"is_enabled\"" } "json:\"peering\""; PrivateLink struct { Url string "json:\"url\""; State string "json:\"state\"" } "json:\"privateLink\"" }{Type:"", CidrBlock:"", Peering:struct { IsEnabled bool "json:\"is_enabled\"" }{IsEnabled:false}, PrivateLink:struct { Url string "json:\"url\""; State string "json:\"state\"" }{Url:"", State:""}}, DataStructures:models.DataStructureResponse{MapConfigs:[]models.MapConfigResponse{}, JCacheConfigs:[]models.JCacheConfigResponse{}, ReplicatedMapConfigs:[]models.ReplicatedMapConfigResponse{}, QueueConfigs:[]models.QueueConfigResponse{}, SetConfigs:[]models.SetConfigResponse{}, ListConfigs:[]models.ListConfigResponse{}, TopicConfigs:[]models.TopicConfigResponse{}, MultiMapConfigs:[]models.MultiMapConfigResponse{}, RingBufferConfigs:[]models.RingBufferConfigResponse{}, ReliableTopicConfigs:[]models.ReliableTopicConfigResponse{}}}}
}

func ExampleServerlessClusterService_get() {
	server := exampleTestMockServer()
	defer server.Close()
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	cluster, _, _ := client.ServerlessCluster.Get(context.Background(), &models.GetServerlessClusterInput{ClusterId: testClusterId})
	fmt.Printf("Result: %#v", cluster)
	//Output:Result: &models.Cluster{Id:"1", CustomerId:10, Name:"test-cluster", ReleaseName:"pr-1", Password:"hidden", Port:30000, HazelcastVersion:"5.1.1", IsAutoScalingEnabled:true, IsHotBackupEnabled:true, IsHotRestartEnabled:true, IsIpWhitelistEnabled:false, IsTlsEnabled:true, ProductType:struct { Name models.ProductTypeName "json:\"name\""; IsFree bool "json:\"isFree\"" }{Name:"STARTER", IsFree:false}, ClusterType:struct { Name models.ClusterType "json:\"name\"" }{Name:"SERVERLESS"}, State:"RUNNING", CreatedAt:"2022-04-21T11:49:02.000Z", StartedAt:"2022-04-21T11:49:02.000Z", StoppedAt:"", Progress:struct { Status string "json:\"status\""; TotalItemCount int "json:\"totalItemCount\""; CompletedItemCount int "json:\"completedItemCount\"" }{Status:"Running", TotalItemCount:4, CompletedItemCount:4}, CloudProvider:struct { Name string "json:\"name\""; Region string "json:\"region\""; AvailabilityZones []string "json:\"availabilityZones\"" }{Name:"aws", Region:"us-west-2", AvailabilityZones:[]string{"us-west-2a"}}, DiscoveryTokens:[]models.DiscoveryToken{models.DiscoveryToken{Source:"default", Token:"hidden"}}, Specs:struct { TotalMemory float64 "json:\"totalMemory\""; HeapMemory int "json:\"heapMemory\""; NativeMemory int "json:\"nativeMemory\""; Cpu int "json:\"cpu\""; InstanceType string "json:\"instanceType\""; InstancePerZone int "json:\"instancePerZone\"" }{TotalMemory:10, HeapMemory:0, NativeMemory:0, Cpu:0, InstanceType:"", InstancePerZone:0}, Networking:struct { Type string "json:\"type\""; CidrBlock string "json:\"cidrBlock\""; Peering struct { IsEnabled bool "json:\"is_enabled\"" } "json:\"peering\""; PrivateLink struct { Url string "json:\"url\""; State string "json:\"state\"" } "json:\"privateLink\"" }{Type:"", CidrBlock:"", Peering:struct { IsEnabled bool "json:\"is_enabled\"" }{IsEnabled:false}, PrivateLink:struct { Url string "json:\"url\""; State string "json:\"state\"" }{Url:"", State:""}}, DataStructures:models.DataStructureResponse{MapConfigs:[]models.MapConfigResponse{}, JCacheConfigs:[]models.JCacheConfigResponse{}, ReplicatedMapConfigs:[]models.ReplicatedMapConfigResponse{}, QueueConfigs:[]models.QueueConfigResponse{}, SetConfigs:[]models.SetConfigResponse{}, ListConfigs:[]models.ListConfigResponse{}, TopicConfigs:[]models.TopicConfigResponse{}, MultiMapConfigs:[]models.MultiMapConfigResponse{}, RingBufferConfigs:[]models.RingBufferConfigResponse{}, ReliableTopicConfigs:[]models.ReliableTopicConfigResponse{}}}
}

func ExampleServerlessClusterService_delete() {
	server := exampleTestMockServer()
	defer server.Close()
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	cluster, _, _ := client.ServerlessCluster.Delete(context.Background(), &models.ClusterDeleteInput{ClusterId: testClusterId})
	fmt.Printf("Result: %#v", cluster)
	//Output:Result: &models.ClusterId{ClusterId:1}
}

func ExampleServerlessClusterService_stop() {
	server := exampleTestMockServer()
	defer server.Close()
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	cluster, _, _ := client.ServerlessCluster.Stop(context.Background(), &models.ClusterStopInput{ClusterId: testClusterId})
	fmt.Printf("Result: %#v", cluster)
	//Output:Result: &models.ClusterId{ClusterId:1}
}

func ExampleServerlessClusterService_resume() {
	server := exampleTestMockServer()
	defer server.Close()
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	cluster, _, _ := client.ServerlessCluster.Resume(context.Background(), &models.ClusterResumeInput{ClusterId: testClusterId})
	fmt.Printf("Result: %#v", cluster)
	//Output:Result: &models.ClusterId{ClusterId:1}
}

func exampleTestMockServer() *httptest.Server {
	serveMux := http.NewServeMux()
	server := httptest.NewServer(serveMux)
	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodPost; m != r.Method {
			fmt.Printf("Request method = %v, expected %v", r.Method, m)
		}
		var request GraphQLQuery
		json.NewDecoder(r.Body).Decode(&request)

		switch {
		case strings.Contains(request.Query, "createServerlessCluster"):
			responseData, _ := ioutil.ReadFile("testdata/serverless_cluster_create_response.json")
			w.Write(responseData)
		case strings.Contains(request.Query, "deleteCluster"):
			responseData, _ := ioutil.ReadFile("testdata/serverless_cluster_delete_response.json")
			w.Write(responseData)
		case strings.Contains(request.Query, "stopCluster"):
			responseData, _ := ioutil.ReadFile("testdata/serverless_cluster_stop_response.json")
			w.Write(responseData)
		case strings.Contains(request.Query, "resumeCluster"):
			responseData, _ := ioutil.ReadFile("testdata/serverless_cluster_resume_response.json")
			w.Write(responseData)
		case strings.Contains(request.Query, "clusters"):
			responseData, _ := ioutil.ReadFile("testdata/serverless_cluster_list_response.json")
			w.Write(responseData)
		case strings.Contains(request.Query, "cluster"):
			responseData, _ := ioutil.ReadFile("testdata/serverless_cluster_get_response.json")
			w.Write(responseData)
		default:
			responseData, _ := ioutil.ReadFile("testdata/access_token_response.json")
			w.Write(responseData)
		}
	})
	return server
}
