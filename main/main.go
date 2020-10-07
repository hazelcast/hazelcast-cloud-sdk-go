package main

import (
	"context"
	hazelcastcloud "github.com/hazelcast/hazelcast-cloud-sdk-go"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

func main() {
	client, response, err := hazelcastcloud.NewFromCredentials("joYCidIG4mlZM2fFLxBxrXcYA",
		"STJyBnPfrwbHQFezK7i5lltxk3vKnpgz3FokNMmR2L2tTlmGRHYAiEEgPf1k",hazelcastcloud.OptionEndpoint("http://localhost:4000"))
	list, response, err := client.Peering.List(context.Background(), &models.PeeringListInput{ClusterId: "53851"})
	create, response, err := client.Peering.CreateGcpVpcNetworkPeering(context.Background(), &models.CreateGcpVpcNetworkPeeringInput{
		ClusterId:   53851,
		ProjectId:   "test-asas2342423d",
		NetworkName: "test-gfdgdg",
		},
	)
	print(create ,list, response, err)
}