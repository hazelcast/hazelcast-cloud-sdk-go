package models

//Type of Peering list Input
type ListAwsPeeringsInput struct {
	ClusterId string `json:"clusterId"`
}

//Type of AcceptAwsPeeringInput input to initiate peering connection from Hazelcast to your Project
type AcceptAwsPeeringInput struct {
	ClusterId           string                            `json:"clusterId"`
	VpcId               string                            `json:"vpcId"`
	VpcCidr             string                            `json:"vpcCidr"`
	PeeringConnectionId string                            `json:"peeringConnectionId"`
	Subnets             []AcceptAwsVpcPeeringInputSubnets `json:"subnets"`
}

//Type of AcceptAwsVpcPeeringInputSubnets input to initiate peering connection subnets field.
type AcceptAwsVpcPeeringInputSubnets struct {
	SubnetId   string `json:"subnetId"`
	SubnetCidr string `json:"subnetCidr"`
}

//Type of DeleteAwsPeeringInput Input
type DeleteAwsPeeringInput struct {
	Id string `json:"id"`
}

//Type of AwsPeering list object
type AwsPeering struct {
	Id         string `json:"id"`
	VpcId      string `json:"vpcId"`
	VpcCidr    string `json:"vpcCidr"`
	SubnetId   string `json:"subnetId"`
	SubnetCidr string `json:"subnetCidr"`
}

//Type of AwsPeeringPropertiesInput to get properties
type GetAwsPeeringPropertiesInput struct {
	ClusterId string `json:"clusterId"`
}

//Type of AwsPeeringProperties to collect needed information for AWS VPC Peering Connection
type AwsPeeringProperties struct {
	VpcId   string `json:"vpcId"`
	VpcCidr string `json:"vpcCidr"`
	OwnerId string `json:"ownerId"`
	Region  string `json:"region"`
}
