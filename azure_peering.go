package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

type AzurePeeringService interface {
	List(ctx context.Context, input *models.ListAzurePeeringsInput) (*[]models.AzurePeering, *Response, error)
	Delete(ctx context.Context, input *models.DeleteAzurePeeringInput) (*models.Result, *Response, error)
	GetProperties(ctx context.Context, input *models.GetAzurePeeringPropertiesInput) (*models.AzurePeeringProperties, *Response, error)
}

type azurePeeringServiceOp struct {
	client *Client
}

func NewAzurePeeringService(client *Client) AzurePeeringService {
	return &azurePeeringServiceOp{client: client}
}

//This function returns needed properties to initialize Azure vNet Peering.
func (p azurePeeringServiceOp) GetProperties(ctx context.Context, input *models.GetAzurePeeringPropertiesInput) (*models.AzurePeeringProperties, *Response, error) {
	var peeringProperties models.AzurePeeringProperties
	graphqlRequest := models.GraphqlRequest{
		Name:      "azurePeeringProperties",
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

//This function returns the list of Azure vNet Peerings for the cluster you specified in Input.
func (p azurePeeringServiceOp) List(ctx context.Context, input *models.ListAzurePeeringsInput) (*[]models.AzurePeering, *Response, error) {
	var peeringList []models.AzurePeering
	graphqlRequest := models.GraphqlRequest{
		Name:      "azurePeerings",
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

//This function deletes Azure vNet Peering for the peering you specified its ID in Input.
func (p azurePeeringServiceOp) Delete(ctx context.Context, input *models.DeleteAzurePeeringInput) (*models.Result, *Response, error) {
	var peeringResult models.Result
	graphqlRequest := models.GraphqlRequest{
		Name:      "deleteAzurePeering",
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
