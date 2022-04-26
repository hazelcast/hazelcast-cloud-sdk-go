package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

// ServerlessClusterService is used to interact with serverless clusters.
type ServerlessClusterService struct {
	client *Client
}

// NewServerlessClusterService returns a new instance of the service.
func NewServerlessClusterService(client *Client) ServerlessClusterService {
	return ServerlessClusterService{client}
}

// Create creates a serverless cluster according to configuration provided in the request.
func (svc ServerlessClusterService) Create(ctx context.Context, input *models.CreateServerlessClusterInput) (*models.Cluster, *Response, error) {
	var cluster models.Cluster
	var graphqlRequest = models.GraphqlRequest{
		Operation: models.Mutation,
		Name:      "createServerlessCluster",
		Input:     *input,
		Args:      nil,
		Response:  cluster,
	}
	req, err := svc.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := svc.client.Do(ctx, req, &cluster)
	if err != nil {
		return nil, nil, err
	}

	return &cluster, resp, nil
}

func (svc ServerlessClusterService) List(ctx context.Context) (*[]models.Cluster, *Response, error) {
	var clusterList []models.Cluster
	graphqlRequest := models.GraphqlRequest{
		Name:      "clusters",
		Operation: models.Query,
		Input:     nil,
		Args: models.ClusterListInput{
			ProductType: models.Starter,
		},
		Response: clusterList,
	}
	req, err := svc.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := svc.client.Do(ctx, req, &clusterList)
	if err != nil {
		return nil, nil, err
	}

	return &clusterList, resp, err
}

func (svc ServerlessClusterService) Get(ctx context.Context, input *models.GetServerlessClusterInput) (*models.Cluster, *Response, error) {
	var cluster models.Cluster
	var graphqlRequest = models.GraphqlRequest{
		Name:      "cluster",
		Operation: models.Query,
		Input:     nil,
		Args:      *input,
		Response:  cluster,
	}
	req, err := svc.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := svc.client.Do(ctx, req, &cluster)
	if err != nil {
		return nil, nil, err
	}

	return &cluster, resp, err
}
