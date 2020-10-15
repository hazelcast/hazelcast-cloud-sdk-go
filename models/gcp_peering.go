package models

//Type of Peering list Input
type ListGcpPeeringsInput struct {
	ClusterId string `json:"clusterId"`
}

//Type of AcceptGcpPeeringInput input to initiate peering connection from Hazelcast to your Project
type AcceptGcpPeeringInput struct {
	ClusterId   string `json:"clusterId"`
	ProjectId   string `json:"projectId"`
	NetworkName string `json:"networkName"`
}

//Type of DeleteGcpPeeringInput Input
type DeleteGcpPeeringInput struct {
	Id string `json:"id"`
}

//Type of GcpPeering list object
type GcpPeering struct {
	Id          string `json:"id"`
	ProjectId   string `json:"projectId"`
	NetworkName string `json:"networkName"`
}

//Type of GcpPeeringPropertiesInput to get properties
type GetGcpPeeringPropertiesInput struct {
	ClusterId string `json:"clusterId"`
}

//Type of GcpPeeringProperties to collect needed information for GCP VPC Peering Connection
type GcpPeeringProperties struct {
	ProjectId   string `json:"projectId"`
	NetworkName string `json:"networkName"`
}
