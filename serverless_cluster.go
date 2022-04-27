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

// List provides the ability to get a list of serverless clusters.
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

// Get retrieves a serverless cluster by its ID.
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

// Delete deletes a serverless cluster by its ID.
func (svc ServerlessClusterService) Delete(ctx context.Context, input *models.ClusterDeleteInput) (*models.ClusterId, *Response, error) {
	var clusterId models.ClusterId
	graphqlRequest := models.GraphqlRequest{
		Name:      "deleteCluster",
		Operation: models.Mutation,
		Input:     nil,
		Args:      *input,
		Response:  clusterId,
	}
	req, err := svc.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := svc.client.Do(ctx, req, &clusterId)
	if err != nil {
		return nil, nil, err
	}

	return &clusterId, resp, err
}

// Stop provides the ability to stop a serverless cluster by its ID.
func (svc ServerlessClusterService) Stop(ctx context.Context, input *models.ClusterStopInput) (*models.ClusterId, *Response, error) {
	var clusterId models.ClusterId
	graphqlRequest := models.GraphqlRequest{
		Name:      "stopCluster",
		Operation: models.Mutation,
		Input:     nil,
		Args:      *input,
		Response:  clusterId,
	}
	req, err := svc.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := svc.client.Do(ctx, req, &clusterId)
	if err != nil {
		return nil, nil, err
	}

	return &clusterId, resp, err
}

// Resume provides the ability to resume a serverless cluster by its ID.
func (svc ServerlessClusterService) Resume(ctx context.Context, input *models.ClusterResumeInput) (*models.ClusterId, *Response, error) {
	var clusterId models.ClusterId
	graphqlRequest := models.GraphqlRequest{
		Name:      "resumeCluster",
		Operation: models.Mutation,
		Input:     nil,
		Args:      *input,
		Response:  clusterId,
	}
	req, err := svc.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := svc.client.Do(ctx, req, &clusterId)
	if err != nil {
		return nil, nil, err
	}

	return &clusterId, resp, err
}
