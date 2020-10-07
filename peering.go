package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

type PeeringService interface {
	List(ctx context.Context, input *models.PeeringListInput) (*[]models.Peering, *Response, error)
	CreateGcpVpcNetworkPeering(ctx context.Context, input *models.CreateGcpVpcNetworkPeeringInput) (*models.Result, *Response, error)
}

type peeringServiceOp struct {
	client *Client
}

func NewPeeringService(client *Client) PeeringService {
	return &peeringServiceOp{client: client}
}

func (p peeringServiceOp) List(ctx context.Context, input *models.PeeringListInput) (*[]models.Peering, *Response, error) {
	var peeringList []models.Peering
	graphqlRequest := models.GraphqlRequest{
		Name:      "peerings",
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

func (p peeringServiceOp) CreateGcpVpcNetworkPeering(ctx context.Context, input *models.CreateGcpVpcNetworkPeeringInput) (*models.Result, *Response, error) {
	var peeringResult models.Result
	graphqlRequest := models.GraphqlRequest{
		Name:      "createGcpVpcNetworkPeering",
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
