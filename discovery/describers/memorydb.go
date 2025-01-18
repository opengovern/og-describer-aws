package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/memorydb/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/memorydb"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
)

func MemoryDbCluster(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := memorydb.NewFromConfig(cfg)

	var values []models.Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		clusters, err := client.DescribeClusters(ctx, &memorydb.DescribeClustersInput{
			NextToken: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, cluster := range clusters.Clusters {
			resource, err := memoryDbClusterHandle(ctx, cfg, cluster)
			if err != nil {
				return nil, err
			}
			emptyResource := models.Resource{}
			if err == nil && resource == emptyResource {
				return nil, nil
			}

			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}

		return clusters.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
func memoryDbClusterHandle(ctx context.Context, cfg aws.Config, cluster types.Cluster) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := memorydb.NewFromConfig(cfg)
	tags, err := client.ListTags(ctx, &memorydb.ListTagsInput{
		ResourceArn: cluster.ARN,
	})
	if err != nil {
		if isErr(err, "ListTagsNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *cluster.ARN,
		Name:   *cluster.Name,
		Description: model.MemoryDbClusterDescription{
			Cluster: cluster,
			Tags:    tags.TagList,
		},
	}
	return resource, nil
}
func GetMemoryDbCluster(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	clusterName := fields["clusterName"]
	client := memorydb.NewFromConfig(cfg)

	describers, err := client.DescribeClusters(ctx, &memorydb.DescribeClustersInput{
		ClusterName: &clusterName,
	})
	if err != nil {
		if isErr(err, "DescribeClustersNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, cluster := range describers.Clusters {
		resource, err := memoryDbClusterHandle(ctx, cfg, cluster)
		if err != nil {
			return nil, err
		}
		emptyResource := models.Resource{}
		if err == nil && resource == emptyResource {
			return nil, nil
		}
		values = append(values, resource)
	}
	return values, nil
}
