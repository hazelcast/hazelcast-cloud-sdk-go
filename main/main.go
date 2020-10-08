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
	"time"
)

func main() {

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
					ID: &customerVnet,
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

func GcpPeering() {
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
