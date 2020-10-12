package models

//Type of Peering list Input
type ListAwsPeeringsInput struct {
	ClusterId string `json:"clusterId"`
}

//Type of AcceptAwsPeeringInput input to initiate peering connection from Hazelcast to your Project
type AcceptAwsPeeringInput struct {
	ClusterId   int    `json:"clusterId"`
	ProjectId   string `json:"projectId"`
	NetworkName string `json:"networkName"`
}

//Type of DeleteAwsPeeringInput Input
type DeleteAwsPeeringInput struct {
	Id string `json:"id"`
}

//Type of AwsPeering list object
type AwsPeering struct {
	Id          string `json:"id"`
}

//Type of AwsPeeringPropertiesInput to get properties
type GetAwsPeeringPropertiesInput struct {
	ClusterId string `json:"clusterId"`
}

//Type of AwsPeeringProperties to collect needed information for AWS VPC Peering Connection
type AwsPeeringProperties struct {
	VpcId   string `json:"vpcId"`
	OwnerId string `json:"ownerId"`
	Region string `json:"region"`
}
