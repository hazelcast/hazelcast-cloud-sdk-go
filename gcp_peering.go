package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

type GcpPeeringService interface {
	List(ctx context.Context, input *models.ListGcpPeeringInput) (*[]models.GcpPeering, *Response, error)
	Accept(ctx context.Context, input *models.AcceptGcpPeeringInput) (*models.Result, *Response, error)
	Delete(ctx context.Context, input *models.DeleteGcpPeeringInput) (*models.Result, *Response, error)
}

type gcpPeeringServiceOp struct {
	client *Client
}

func NewGcpPeeringService(client *Client) GcpPeeringService {
	return &gcpPeeringServiceOp{client: client}
}

func (p gcpPeeringServiceOp) List(ctx context.Context, input *models.ListGcpPeeringInput) (*[]models.GcpPeering, *Response, error) {
	var peeringList []models.GcpPeering
	graphqlRequest := models.GraphqlRequest{
		Name:      "listGcpPeerings",
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
