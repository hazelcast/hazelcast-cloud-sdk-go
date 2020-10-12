package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

type GcpPeeringService interface {
	List(ctx context.Context, input *models.ListGcpPeeringsInput) (*[]models.GcpPeering, *Response, error)
	Accept(ctx context.Context, input *models.AcceptGcpPeeringInput) (*models.Result, *Response, error)
	Delete(ctx context.Context, input *models.DeleteGcpPeeringInput) (*models.Result, *Response, error)
	GetProperties(ctx context.Context, input *models.GetGcpPeeringPropertiesInput) (*models.GcpPeeringProperties, *Response, error)
}

type gcpPeeringServiceOp struct {
	client *Client
}

func NewGcpPeeringService(client *Client) GcpPeeringService {
	return &gcpPeeringServiceOp{client: client}
}

//This function returns needed properties to initialize Google Cloud Platform VPC Peering
func (p gcpPeeringServiceOp) GetProperties(ctx context.Context, input *models.GetGcpPeeringPropertiesInput) (*models.GcpPeeringProperties, *Response, error) {
	var peeringProperties models.GcpPeeringProperties
	graphqlRequest := models.GraphqlRequest{
		Name:      "gcpPeeringProperties",
		Operation: models.Query,
		Input:     nil,
		Args:      *input,
		Response:  peeringProperties,
	}
	req, err := p.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := p.client.Do(ctx, req, &peeringProperties)
	if err != nil {
		return nil, resp, err
	}

	return &peeringProperties, resp, err
}

//This function returns the list of Google Cloud Platform VPC Peerings for the cluster you specified in Input.
func (p gcpPeeringServiceOp) List(ctx context.Context, input *models.ListGcpPeeringsInput) (*[]models.GcpPeering, *Response, error) {
	var peeringList []models.GcpPeering
	graphqlRequest := models.GraphqlRequest{
		Name:      "gcpPeerings",
		Operation: models.Query,
		Input:     nil,
		Args:      *input,
		Response:  peeringList,
	}
	req, err := p.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := p.client.Do(ctx, req, &peeringList)
	if err != nil {
		return nil, resp, err
	}

	return &peeringList, resp, err
}

//This function accept Google Cloud Platform VPC Peering that comes from the specified parameters you provided in Input.
func (p gcpPeeringServiceOp) Accept(ctx context.Context, input *models.AcceptGcpPeeringInput) (*models.Result, *Response, error) {
	var peeringResult models.Result
	graphqlRequest := models.GraphqlRequest{
		Name:      "acceptGcpPeering",
		Operation: models.Mutation,
		Input:     *input,
		Args:      nil,
		Response:  peeringResult,
	}
	req, err := p.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := p.client.Do(ctx, req, &peeringResult)
	if err != nil {
		return nil, resp, err
	}

	return &peeringResult, resp, err
}

//This function deletes Google Cloud Platform VPC Peering for the peering you specified its ID in Input
func (p gcpPeeringServiceOp) Delete(ctx context.Context, input *models.DeleteGcpPeeringInput) (*models.Result, *Response, error) {
	var peeringResult models.Result
	graphqlRequest := models.GraphqlRequest{
		Name:      "deleteGcpPeering",
		Operation: models.Mutation,
		Input:     nil,
		Args:      *input,
		Response:  peeringResult,
	}
	req, err := p.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := p.client.Do(ctx, req, &peeringResult)
	if err != nil {
		return nil, resp, err
	}

	return &peeringResult, resp, err
}