package main

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/authorization/mgmt/authorization"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/graphrbac/graphrbac"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2017-09-01/network"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/google/uuid"
	hazelcastcloud "github.com/hazelcast/hazelcast-cloud-sdk-go"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
	"google.golang.org/api/compute/v1"
	"os"
	"strconv"
	"time"
)

func main() {
	//GcpPeering(53878, "yunus-project","yunus-vpc")
	AwsPeeringTest()
}

func AwsPeeringTest() {
	client, _, clientErr := hazelcastcloud.NewFromCredentials(
		os.Getenv("API_KEY"),
		os.Getenv("API_SECRET"),
		hazelcastcloud.OptionEndpoint("https://optimusprime.test.hazelcast.cloud/api/v1"),
	)
	if clientErr != nil {
		panic(clientErr)
	}
		peeringProperties, _, propertiesErr := client.AwsPeering.GetProperties(context.Background(), &models.GetAwsPeeringPropertiesInput{
		ClusterId: "53879",
	})

	fmt.Println(peeringProperties, propertiesErr)
}

func AzurePeering() {
	client, _, clientErr := hazelcastcloud.NewFromCredentials(
		os.Getenv("API_KEY"),
		os.Getenv("API_SECRET"),
		hazelcastcloud.OptionEndpoint("https://optimusprime.test.hazelcast.cloud/api/v1"),
	)
	if clientErr != nil {
		panic(clientErr)
	}
	peeringProperties, _, propertiesErr := client.AzurePeering.GetProperties(context.Background(), &models.GetAzurePeeringPropertiesInput{
		ClusterId: "53858",
	})
	if propertiesErr != nil {
		panic(propertiesErr)
	}

	true := true
	false := false
	customerVnet := "<customerVnet>"
	customerSubscriptionId := "<customerSubscription>"
	customerResourceGroup := "<customerSubscription>"
	customerTenantId := "<customerTenantId>"
	hazelcastVnetId := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/VirtualNetworks/%s", peeringProperties.SubscriptionId, peeringProperties.ResourceGroupName, peeringProperties.VnetName)
	customerVnetId := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/VirtualNetworks/%s", customerSubscriptionId, customerResourceGroup, customerVnet)

	customerAuth, customerAuthErr := auth.NewAuthorizerFromEnvironment()
	if customerAuthErr != nil {
		panic(customerAuthErr)
	}
	env, envErr := azure.EnvironmentFromName("AzurePublicCloud")
	if envErr != nil {
		panic(envErr)
	}
	oauthConfig, oauthConfigErr := adal.NewOAuthConfig(env.ActiveDirectoryEndpoint, peeringProperties.TenantId)
	if oauthConfigErr != nil {
		panic(oauthConfigErr)
	}
	token, tokenErr := adal.NewServicePrincipalToken(*oauthConfig,
		peeringProperties.AppRegistrationId, peeringProperties.AppRegistrationKey, env.ResourceManagerEndpoint)
	if tokenErr != nil {
		panic(tokenErr)
	}
	hazelcastAuth := autorest.NewBearerAuthorizer(token)
	customerVnetPeeringClient := network.NewVirtualNetworkPeeringsClient(customerSubscriptionId)
	customerVnetPeeringClient.Authorizer = customerAuth
	hazelcastVnetPeeringClient := network.NewVirtualNetworkPeeringsClient(peeringProperties.SubscriptionId)
	hazelcastVnetPeeringClient.Authorizer = hazelcastAuth
	customerServicePrincipalClient := graphrbac.NewServicePrincipalsClient(customerTenantId)
	customerServicePrincipalClient.Authorizer = customerAuth
	customerRoleAssignmentClient := authorization.NewRoleAssignmentsClient(customerSubscriptionId)
	customerRoleAssignmentClient.Authorizer = customerAuth

	_, createSpErr := customerServicePrincipalClient.Create(context.Background(), graphrbac.ServicePrincipalCreateParameters{
		AppID: &peeringProperties.AppRegistrationId,
	})
	if createSpErr != nil {
		panic(createSpErr)
	}

	roleDefinitionId := "4d97b98b-1d4f-4787-a291-c67834d212e7"
	_, roleAssignmentErr := customerRoleAssignmentClient.Create(context.Background(), hazelcastVnetId, "hazelcast-cloud", authorization.RoleAssignmentCreateParameters{
		Properties: &authorization.RoleAssignmentProperties{
			RoleDefinitionID: &roleDefinitionId,
			PrincipalID:      &peeringProperties.AppRegistrationId,
		},
	})
	if roleAssignmentErr != nil {
		panic(roleAssignmentErr)
	}

	customerPeerList, peerListErr := customerVnetPeeringClient.List(context.Background(), customerResourceGroup, customerVnet)
	if peerListErr != nil {
		panic(peerListErr)
	}
	for _, v := range customerPeerList.Values() {
		_, _ = customerVnetPeeringClient.Delete(context.Background(), customerResourceGroup, customerVnet, *v.Name)
	}

	hazelcastPeerList, peerListErr := hazelcastVnetPeeringClient.List(context.Background(), peeringProperties.ResourceGroupName, peeringProperties.VnetName)
	if peerListErr != nil {
		panic(peerListErr)
	}
	for _, v := range hazelcastPeerList.Values() {
		_, _ = hazelcastVnetPeeringClient.Delete(context.Background(), peeringProperties.ResourceGroupName, peeringProperties.VnetName, *v.Name)
	}

	hzPeerName := uuid.New().String()
	_, hazelcastPeerErr := hazelcastVnetPeeringClient.CreateOrUpdate(context.Background(),
		peeringProperties.ResourceGroupName, peeringProperties.VnetName, "hz-to-customer", network.VirtualNetworkPeering{
			VirtualNetworkPeeringPropertiesFormat: &network.VirtualNetworkPeeringPropertiesFormat{
				AllowVirtualNetworkAccess: &true,
				AllowForwardedTraffic:     &true,
				AllowGatewayTransit:       &false,
				UseRemoteGateways:         &false,
				RemoteVirtualNetwork: &network.SubResource{
					ID: &customerVnetId,
				},
			},
			Name: &hzPeerName,
		})
	if hazelcastPeerErr != nil {
		panic(hazelcastPeerErr)
	}

	customerPeerName := uuid.New().String()
	_, customerPeerErr := customerVnetPeeringClient.CreateOrUpdate(context.Background(),
		customerResourceGroup, customerVnet, "customer-to-hz", network.VirtualNetworkPeering{
			VirtualNetworkPeeringPropertiesFormat: &network.VirtualNetworkPeeringPropertiesFormat{
				AllowVirtualNetworkAccess: &true,
				AllowForwardedTraffic:     &true,
				AllowGatewayTransit:       &false,
				UseRemoteGateways:         &false,
				RemoteVirtualNetwork: &network.SubResource{
					ID: &hazelcastVnetId,
				},
			},
			Name: &customerPeerName,
		})
	if customerPeerErr != nil {
		panic(customerPeerErr)
	}

	print("Done")

}

func GcpPeering(clusterId int, customerProject string, customerNetwork string) {
	fmt.Println(time.Now().String())
	client, _, clientErr := hazelcastcloud.NewFromCredentials(
		os.Getenv("API_KEY"),
		os.Getenv("API_SECRET"),
		hazelcastcloud.OptionEndpoint("https://optimusprime.test.hazelcast.cloud/api/v1"),
	)
	if clientErr != nil {
		panic(clientErr)
	}

	peeringProperties, _, propertiesErr := client.GcpPeering.GetProperties(context.Background(), &models.GetGcpPeeringPropertiesInput{
		ClusterId: strconv.Itoa(clusterId),
	})
	if propertiesErr != nil {
		panic(propertiesErr)
	}
	computeService, computeServiceErr := compute.NewService(context.Background())
	if computeServiceErr != nil {
		panic(computeServiceErr)
	}

	_, addPeeringErr := computeService.Networks.AddPeering(customerProject, customerNetwork, &compute.NetworksAddPeeringRequest{
		Name:             fmt.Sprintf("%s-%s", peeringProperties.ProjectId, peeringProperties.NetworkName),
		PeerNetwork:      "https://www.googleapis.com/compute/v1/projects/" + peeringProperties.ProjectId + "/global/networks/" + peeringProperties.NetworkName,
		AutoCreateRoutes: true,
	}).Do()

	if addPeeringErr != nil {
		panic(addPeeringErr)
	}

	_, _, acceptErr := client.GcpPeering.Accept(context.Background(), &models.AcceptGcpPeeringInput{
		ClusterId:   clusterId,
		ProjectId:   customerProject,
		NetworkName: customerNetwork,
	})
	if acceptErr != nil {
		panic(acceptErr)
	}

	list, _, _ := client.GcpPeering.List(context.Background(), &models.ListGcpPeeringsInput{ClusterId: strconv.Itoa(clusterId)})
	asd, _, propertiesErr := client.GcpPeering.Delete(context.Background(), &models.DeleteGcpPeeringInput{Id: (*list)[0].Id})


	fmt.Println(list,asd, clientErr)
	fmt.Println(time.Now().String())
}