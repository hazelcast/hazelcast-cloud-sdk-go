package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

//This EnterpriseClusterService is used to make operations related with enterprise clusters
type EnterpriseClusterService interface {
	Get(ctx context.Context, input *models.GetEnterpriseClusterInput) (*models.Cluster, *Response, error)
	Create(ctx context.Context, input *models.CreateEnterpriseClusterInput) (*models.Cluster, *Response, error)
	List(ctx context.Context) (*[]models.Cluster, *Response, error)
	Delete(ctx context.Context, input *models.ClusterDeleteInput) (*models.ClusterId, *Response, error)
	ListUploadedArtifacts(ctx context.Context,
		request *models.ListUploadedArtifactsInput) (*[]models.UploadedArtifact, *Response, error)
	UploadArtifact(ctx context.Context,
		request *models.UploadArtifactInput) (*models.UploadedArtifact, *Response, error)
	DeleteArtifact(ctx context.Context,
		request *models.DeleteArtifactInput) (*models.UploadedArtifact, *Response, error)
	DownloadArtifact(ctx context.Context,
		request *models.DownloadArtifactInput) (*models.UploadedArtifactLink, *Response, error)
}

type enterpriseClusterServiceOp struct {
	client *Client
}

func NewEnterpriseClusterService(client *Client) EnterpriseClusterService {
	return &enterpriseClusterServiceOp{client: client}
}

//This function returns detailed configuration of the cluster
func (svc enterpriseClusterServiceOp) Get(ctx context.Context,
	input *models.GetEnterpriseClusterInput) (*models.Cluster, *Response, error) {
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
		return nil, resp, err
	}

	return &cluster, resp, err
}

//This function creates an Enterprise Cluster with a configuration provided in the input
func (svc enterpriseClusterServiceOp) Create(ctx context.Context,
	input *models.CreateEnterpriseClusterInput) (*models.Cluster, *Response, error) {
	var cluster models.Cluster
	var graphqlRequest = models.GraphqlRequest{
		Name:      "createEnterpriseCluster",
		Operation: models.Mutation,
		Input:     input,
		Args:      nil,
		Response:  cluster,
	}
	req, err := svc.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := svc.client.Do(ctx, req, &cluster)
	if err != nil {
		return nil, resp, err
	}

	return &cluster, resp, err
}

//This function lists all non-deleted Enterprise Clusters
func (svc enterpriseClusterServiceOp) List(ctx context.Context) (*[]models.Cluster, *Response, error) {
	var clusterList []models.Cluster
	graphqlRequest := models.GraphqlRequest{
		Name:      "clusters",
		Operation: models.Query,
		Input:     nil,
		Args: models.ClusterListInput{
			ProductType: models.Enterprise,
		},
		Response: clusterList,
	}
	req, err := svc.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := svc.client.Do(ctx, req, &clusterList)
	if err != nil {
		return nil, resp, err
	}

	return &clusterList, resp, err
}

//This function deletes an Enterprise Cluster
func (svc enterpriseClusterServiceOp) Delete(ctx context.Context,
	input *models.ClusterDeleteInput) (*models.ClusterId, *Response, error) {
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
		return nil, resp, err
	}

	return &clusterId, resp, err
}

func (svc enterpriseClusterServiceOp) ListUploadedArtifacts(ctx context.Context,
	request *models.ListUploadedArtifactsInput) (*[]models.UploadedArtifact, *Response, error) {
	var artifact []models.UploadedArtifact
	graphqlRequest := models.GraphqlRequest{
		Name:      "customClasses",
		Operation: models.Query,
		Input:     nil,
		Args:      *request,
		Response:  artifact,
	}
	req, err := svc.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := svc.client.Do(ctx, req, &artifact)
	if err != nil {
		return nil, resp, err
	}

	return &artifact, nil, nil
}

func (svc enterpriseClusterServiceOp) UploadArtifact(ctx context.Context,
	request *models.UploadArtifactInput) (*models.UploadedArtifact, *Response, error) {
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
	req, err := svc.client.NewUploadFileRequest(&graphqlQuery)
	if err != nil {
		return nil, nil, err
	}

	resp, err := svc.client.Do(ctx, req, &artifact)
	if err != nil {
		return nil, resp, err
	}

	return &artifact, nil, nil
}

func (svc enterpriseClusterServiceOp) DeleteArtifact(ctx context.Context,
	request *models.DeleteArtifactInput) (*models.UploadedArtifact, *Response, error) {
	var artifact models.UploadedArtifact
	graphqlQuery := models.GraphqlRequest{
		Name:      "deleteCustomClassArtifact",
		Operation: models.Mutation,
		Input:     nil,
		Args:      *request,
		Response:  artifact,
	}
	req, err := svc.client.NewRequest(&graphqlQuery)
	if err != nil {
		return nil, nil, err
	}

	resp, err := svc.client.Do(ctx, req, &artifact)
	if err != nil {
		return nil, resp, err
	}

	return &artifact, nil, nil
}

func (svc enterpriseClusterServiceOp) DownloadArtifact(ctx context.Context,
	request *models.DownloadArtifactInput) (*models.UploadedArtifactLink, *Response, error) {
	var artifact models.UploadedArtifactLink
	graphqlQuery := models.GraphqlRequest{
		Name:      "downloadCustomClassesArtifact",
		Operation: models.Query,
		Args:      *request,
		Response:  artifact,
	}
	req, err := svc.client.NewRequest(&graphqlQuery)
	if err != nil {
		return nil, nil, err
	}

	resp, err := svc.client.Do(ctx, req, &artifact)
	if err != nil {
		return nil, resp, err
	}

	return &artifact, nil, nil
}
