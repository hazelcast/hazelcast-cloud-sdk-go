package main

import (
	"context"
	hazelcastcloud "github.com/hazelcast/hazelcast-cloud-sdk-go"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
	"os"
)

func main() {
	client, response, err := hazelcastcloud.NewFromCredentials(os.Getenv("API_KEY"), os.Getenv("API_SECRET"), hazelcastcloud.OptionEndpoint("http://localhost:4000"))
	list, response, err := client.GcpPeering.List(context.Background(), &models.ListGcpPeeringInput{ClusterId: "53851"})
	asd, response, err := client.GcpPeering.Delete(context.Background(), &models.DeleteGcpPeeringInput{Id: (*list)[0].Id})
	print(response, err, list, asd)
}
