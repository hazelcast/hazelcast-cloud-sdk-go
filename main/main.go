package main

import (
	"context"
	"fmt"
	hazelcastcloud "github.com/hazelcast/hazelcast-cloud-sdk-go"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
	"google.golang.org/api/compute/v1"
	"os"
	"time"
)

func main() {
	fmt.Println(time.Now().String())
	client, _, clientErr := hazelcastcloud.NewFromCredentials(
		os.Getenv("API_KEY"),
		os.Getenv("API_SECRET"),
		hazelcastcloud.OptionEndpoint("https://optimusprime.test.hazelcast.cloud/api/v1"),
	)
	if clientErr != nil {
		panic(clientErr)
	}
	fmt.Println(time.Now().String())
	peeringProperties, _, propertiesErr := client.GcpPeering.GetProperties(context.Background(), &models.GetGcpPeeringPropertiesInput{
		ClusterId: "53858",
	})
	if propertiesErr != nil {
		panic(propertiesErr)
	}
	fmt.Println(time.Now().String())
	customerPeeringErr := createCustomerVpcPeering("yunus-project", "yunus-vpc", peeringProperties.ProjectId, peeringProperties.NetworkName)
	if customerPeeringErr != nil {
		panic(customerPeeringErr)
	}
	fmt.Println(time.Now().String())
	_, _, acceptErr := client.GcpPeering.Accept(context.Background(), &models.AcceptGcpPeeringInput{
		ClusterId:   53858,
		ProjectId:   "yunus-project",
		NetworkName: "yunus-vpc",
	})
	if acceptErr != nil {
		panic(acceptErr)
	}
	fmt.Println(time.Now().String())
}


func createCustomerVpcPeering(projectIdA string, networkNameA string, projectIdB string, networkNameB string) error {
	computeService, computeServiceErr := compute.NewService(context.Background())
	if computeServiceErr != nil {
		return computeServiceErr
	}
	addPeeringRes, addPeeringErr := computeService.Networks.AddPeering(projectIdA, networkNameA, &compute.NetworksAddPeeringRequest{
		Name:             fmt.Sprintf("from-%s-to-%s", projectIdA, projectIdB),
		PeerNetwork:      "https://www.googleapis.com/compute/v1/projects/" + projectIdB + "/global/networks/" + networkNameB,
		AutoCreateRoutes: true,
	}).Do()
	if addPeeringErr != nil {
		return addPeeringErr
	}
	print(addPeeringRes)
	return nil
}