package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

type AwsPeeringService interface {
	List(ctx context.Context, input *models.ListAwsPeeringsInput) (*[]models.AwsPeering, *Response, error)
	Accept(ctx context.Context, input *models.AcceptAwsPeeringInput) (*models.Result, *Response, error)
	Delete(ctx context.Context, input *models.DeleteAwsPeeringInput) (*models.Result, *Response, error)
	GetProperties(ctx context.Context, input *models.GetAwsPeeringPropertiesInput) (*models.AwsPeeringProperties, *Response, error)
}

type awsPeeringServiceOp struct {
	client *Client
}

func NewAwsPeeringService(client *Client) AwsPeeringService {
	return &awsPeeringServiceOp{client: client}
}

//This function returns needed properties to initialize Amazon Web Services VPC Peering.
func (p awsPeeringServiceOp) GetProperties(ctx context.Context, input *models.GetAwsPeeringPropertiesInput) (*models.AwsPeeringProperties, *Response, error) {
	var peeringProperties models.AwsPeeringProperties
	graphqlRequest := models.GraphqlRequest{
		Name:      "awsPeeringProperties",
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

//This function returns the list of Amazon Web Services VPC Peerings for the cluster you specified in Input.
func (p awsPeeringServiceOp) List(ctx context.Context, input *models.ListAwsPeeringsInput) (*[]models.AwsPeering, *Response, error) {
	var peeringList []models.AwsPeering
	graphqlRequest := models.GraphqlRequest{
		Name:      "awsPeerings",
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

//This function accept Amazon Web Services VPC Peering that comes from the specified parameters you provided in Input.
func (p awsPeeringServiceOp) Accept(ctx context.Context, input *models.AcceptAwsPeeringInput) (*models.Result, *Response, error) {
	var peeringResult models.Result
	graphqlRequest := models.GraphqlRequest{
		Name:      "acceptAwsPeering",
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

//This function deletes Amazon Web Services VPC Peering for the peering you specified its ID in Input.
func (p awsPeeringServiceOp) Delete(ctx context.Context, input *models.DeleteAwsPeeringInput) (*models.Result, *Response, error) {
	var peeringResult models.Result
	graphqlRequest := models.GraphqlRequest{
		Name:      "deleteAwsPeering",
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