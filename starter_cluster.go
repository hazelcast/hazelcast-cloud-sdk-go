package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

//This StarterClusterService is used to interact with Starter Clusters.
type StarterClusterService interface {
	Get(ctx context.Context, request *models.GetStarterClusterInput) (*models.Cluster, *Response, error)
	Create(ctx context.Context, request *models.CreateStarterClusterInput) (*models.Cluster, *Response, error)
	List(ctx context.Context) (*[]models.Cluster, *Response, error)
	Resume(ctx context.Context, request *models.ClusterResumeInput) (*models.ClusterId, *Response, error)
	Stop(ctx context.Context, request *models.ClusterStopInput) (*models.ClusterId, *Response, error)
	Delete(ctx context.Context, request *models.ClusterDeleteInput) (*models.ClusterId, *Response, error)
	ListUploadedArtifacts(ctx context.Context, request *models.ListUploadedArtifactsInput) (*[]models.UploadedArtifact, *Response, error)
	UploadArtifact(ctx context.Context, request *models.UploadArtifactInput) (*models.UploadedArtifact, *Response, error)
	DeleteArtifact(ctx context.Context, request *models.DeleteArtifactInput) (*models.UploadedArtifact, *Response, error)
}

type starterClusterServiceOp struct {
	client *Client
}

func NewStarterClusterService(client *Client) StarterClusterService {
	return &starterClusterServiceOp{client: client}
}

//This function gets detailed configuration of started cluster according to cluster id that provided on the request
func (c starterClusterServiceOp) Get(ctx context.Context, input *models.GetStarterClusterInput) (*models.Cluster, *Response, error) {
	var cluster models.Cluster
	var graphqlRequest = models.GraphqlRequest{
		Name:      "cluster",
		Operation: models.Query,
		Input:     nil,
		Args:      *input,
		Response:  cluster,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &cluster)
	if err != nil {
		return nil, resp, err
	}

	return &cluster, resp, err
}

//This function creates started cluster according to configuration provided in the request
func (c starterClusterServiceOp) Create(ctx context.Context, input *models.CreateStarterClusterInput) (*models.Cluster, *Response, error) {
	var cluster models.Cluster
	var graphqlRequest = models.GraphqlRequest{
		Name:      "createStarterCluster",
		Operation: models.Mutation,
		Input:     *input,
		Args:      nil,
		Response:  cluster,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &cluster)
	if err != nil {
		return nil, resp, err
	}

	return &cluster, resp, err
}

//This function list all non-deleted Starter Cluster
func (c starterClusterServiceOp) List(ctx context.Context) (*[]models.Cluster, *Response, error) {
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
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &clusterList)
	if err != nil {
		return nil, resp, err
	}

	return &clusterList, resp, err
}

//This function resume a stopped Starter Cluster
func (c starterClusterServiceOp) Resume(ctx context.Context, input *models.ClusterResumeInput) (*models.ClusterId, *Response, error) {
	var clusterId models.ClusterId
	graphqlRequest := models.GraphqlRequest{
		Name:      "resumeCluster",
		Operation: models.Mutation,
		Input:     nil,
		Args:      *input,
		Response:  clusterId,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &clusterId)
	if err != nil {
		return nil, nil, err
	}

	return &clusterId, resp, err
}

//This function stops a running Starter Cluster
func (c starterClusterServiceOp) Stop(ctx context.Context, input *models.ClusterStopInput) (*models.ClusterId, *Response, error) {
	var clusterId models.ClusterId
	graphqlRequest := models.GraphqlRequest{
		Name:      "stopCluster",
		Operation: models.Mutation,
		Input:     nil,
		Args:      *input,
		Response:  clusterId,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &clusterId)
	if err != nil {
		return nil, resp, err
	}

	return &clusterId, resp, err
}

//This function deletes a starter Starter Cluster
func (c starterClusterServiceOp) Delete(ctx context.Context, input *models.ClusterDeleteInput) (*models.ClusterId, *Response, error) {
	var clusterId models.ClusterId
	graphqlRequest := models.GraphqlRequest{
		Name:      "deleteCluster",
		Operation: models.Mutation,
		Input:     nil,
		Args:      *input,
		Response:  clusterId,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &clusterId)
	if err != nil {
		return nil, resp, err
	}

	return &clusterId, resp, err
}

func (c starterClusterServiceOp) ListUploadedArtifacts(ctx context.Context, request *models.ListUploadedArtifactsInput) (*[]models.UploadedArtifact, *Response, error) {
	var artifact []models.UploadedArtifact
	graphqlRequest := models.GraphqlRequest{
		Name:      "customClasses",
		Operation: models.Query,
		Input:     nil,
		Args:      *request,
		Response:  artifact,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &artifact)
	if err != nil {
		return nil, resp, err
	}

	return &artifact, nil, nil
}

func (c starterClusterServiceOp) UploadArtifact(ctx context.Context, request *models.UploadArtifactInput) (*models.UploadedArtifact, *Response, error) {
	var artifact models.UploadedArtifact
	graphqlQuery := models.GraphqlRequest{
		Name:      "uploadCustomClassArtifact",
		Operation: models.Mutation,
		Input:     nil,
		Args: models.UploadArtifactArgs{
			ClusterId: request.ClusterId,
		},
		Response: artifact,
		UploadFile: models.UploadFile{
			FileName: request.FileName,
			Content:  request.Content,
		},
	}
	req, err := c.client.NewUploadFileRequest(&graphqlQuery)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &artifact)
	if err != nil {
		return nil, resp, err
	}

	return &artifact, nil, nil
}

func (c starterClusterServiceOp) DeleteArtifact(ctx context.Context, request *models.DeleteArtifactInput) (*models.UploadedArtifact, *Response, error) {
	var artifact models.UploadedArtifact
	graphqlQuery := models.GraphqlRequest{
		Name:      "deleteCustomClassArtifact",
		Operation: models.Mutation,
		Input:     nil,
		Args:      *request,
		Response:  artifact,
	}
	req, err := c.client.NewRequest(&graphqlQuery)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &artifact)
	if err != nil {
		return nil, resp, err
	}

	return &artifact, nil, nil
}
