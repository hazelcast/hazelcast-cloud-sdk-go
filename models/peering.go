package models

//Type of Peering list Input
type PeeringListInput struct {
	ClusterId string `json:"clusterId"`
}

//Type of CreateGcpVpcNetworkPeeringInput input to initiate peering connection from Hazelcast to your Project
type CreateGcpVpcNetworkPeeringInput struct {
	ClusterId   int    `json:"clusterId"`
	ProjectId   string `json:"projectId"`
	NetworkName string `json:"networkName"`
}

//Type of Peering lists object
type Peering struct {
	Id                  int    `json:"id"`
	ClusterId           int    `json:"clusterId"`
	VpcId               string `json:"vpcId"`
	SubnetId            string `json:"subnetId"`
	SubnetCidr          string `json:"subnetCidr"`
	PeeringConnectionId string    `json:"peeringConnectionId"`
}
