package models

//Type of ListAzurePeeringsInput to list connections
type ListAzurePeeringsInput struct {
	ClusterId string `json:"clusterId"`
}

//Type of DeleteAzurePeeringInput Input
type DeleteAzurePeeringInput struct {
	Id string `json:"id"`
}

//Type of AzurePeering list object
type AzurePeering struct {
	Id string `json:"id"`
}

//Type of GetAzurePeeringPropertiesInput to get properties
type GetAzurePeeringPropertiesInput struct {
	ClusterId string `json:"clusterId"`
}

//Type of AzurePeeringProperties to collect needed information for Azure vNet Peering Connection
type AzurePeeringProperties struct {
	AppRegistrationId  string `json:"appRegistrationId"`
	AppRegistrationKey string `json:"appRegistrationKey"`
	TenantId           string `json:"tenantId"`
	VnetName           string `json:"vnetName"`
	SubscriptionId     string `json:"subscriptionId"`
	ResourceGroupName  string `json:"resourceGroupName"`
}
