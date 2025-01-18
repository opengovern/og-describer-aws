package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func ElastiCacheReplicationGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := elasticache.NewFromConfig(cfg)
	paginator := elasticache.NewDescribeReplicationGroupsPaginator(client, &elasticache.DescribeReplicationGroupsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, item := range page.ReplicationGroups {
			resource := elastiCacheReplicationGroupHandle(ctx, item)
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}
func elastiCacheReplicationGroupHandle(ctx context.Context, item types.ReplicationGroup) models.Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *item.ARN,
		Name:   *item.ARN,
		Description: model.ElastiCacheReplicationGroupDescription{
			ReplicationGroup: item,
		},
	}
	return resource
}
func GetElastiCacheReplicationGroup(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	CacheReplicationId := fields["CacheReplicationId"]
	client := elasticache.NewFromConfig(cfg)
	out, err := client.DescribeReplicationGroups(ctx, &elasticache.DescribeReplicationGroupsInput{
		ReplicationGroupId: &CacheReplicationId,
	})
	if err != nil {
		if isErr(err, "CacheReplicationGroupNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, v := range out.ReplicationGroups {
		values = append(values, elastiCacheReplicationGroupHandle(ctx, v))
	}
	return values, nil
}

func ElastiCacheCluster(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := elasticache.NewFromConfig(cfg)
	paginator := elasticache.NewDescribeCacheClustersPaginator(client, &elasticache.DescribeCacheClustersInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "CacheClusterNotFound") || isErr(err, "InvalidParameterValue") {
				continue
			}
			return nil, err
		}

		for _, cluster := range page.CacheClusters {
			resource, err := elastiCacheClusterHandle(ctx, cluster, client)
			if err != nil {
				return nil, err
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}
func elastiCacheClusterHandle(ctx context.Context, cluster types.CacheCluster, client *elasticache.Client) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	tagsOutput, err := client.ListTagsForResource(ctx, &elasticache.ListTagsForResourceInput{
		ResourceName: cluster.ARN,
	})
	if err != nil {
		if !isErr(err, "CacheClusterNotFound") && !isErr(err, "InvalidParameterValue") {
			return models.Resource{}, err
		} else {
			tagsOutput = &elasticache.ListTagsForResourceOutput{}
		}
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *cluster.ARN,
		Name:   *cluster.ARN,
		Description: model.ElastiCacheClusterDescription{
			Cluster: cluster,
			TagList: tagsOutput.TagList,
		},
	}
	return resource, nil
}
func GetElastiCacheCluster(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	clusterID := fields["id"]
	client := elasticache.NewFromConfig(cfg)
	out, err := client.DescribeCacheClusters(ctx, &elasticache.DescribeCacheClustersInput{
		CacheClusterId: &clusterID,
	})
	if err != nil {
		if isErr(err, "CacheClusterNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, cluster := range out.CacheClusters {
		resource, err := elastiCacheClusterHandle(ctx, cluster, client)
		if err != nil {
			return nil, err
		}
		values = append(values, resource)
	}
	return values, nil
}

func ElastiCacheParameterGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := elasticache.NewFromConfig(cfg)
	paginator := elasticache.NewDescribeCacheParameterGroupsPaginator(client, &elasticache.DescribeCacheParameterGroupsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, cacheParameterGroup := range page.CacheParameterGroups {
			resource := elastiCacheParameterGroupHandle(ctx, cacheParameterGroup)
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}
func elastiCacheParameterGroupHandle(ctx context.Context, cacheParameterGroup types.CacheParameterGroup) models.Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *cacheParameterGroup.ARN,
		Name:   *cacheParameterGroup.CacheParameterGroupName,
		Description: model.ElastiCacheParameterGroupDescription{
			ParameterGroup: cacheParameterGroup,
		},
	}
	return resource
}
func GetElastiCacheParameterGroup(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	cacheParameterGroupName := fields["name"]
	var values []models.Resource
	client := elasticache.NewFromConfig(cfg)

	out, err := client.DescribeCacheParameterGroups(ctx, &elasticache.DescribeCacheParameterGroupsInput{
		CacheParameterGroupName: &cacheParameterGroupName,
	})
	if err != nil {
		return nil, err
	}

	for _, cacheParameterGroup := range out.CacheParameterGroups {
		values = append(values, elastiCacheParameterGroupHandle(ctx, cacheParameterGroup))
	}
	return values, nil
}

func ElastiCacheReservedCacheNode(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := elasticache.NewFromConfig(cfg)
	paginator := elasticache.NewDescribeReservedCacheNodesPaginator(client, &elasticache.DescribeReservedCacheNodesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, reservedCacheNode := range page.ReservedCacheNodes {
			resource := elastiCacheReservedCacheNodeHandle(ctx, reservedCacheNode)
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}
func elastiCacheReservedCacheNodeHandle(ctx context.Context, reservedCacheNode types.ReservedCacheNode) models.Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *reservedCacheNode.ReservationARN,
		ID:     *reservedCacheNode.ReservedCacheNodeId,
		Description: model.ElastiCacheReservedCacheNodeDescription{
			ReservedCacheNode: reservedCacheNode,
		},
	}
	return resource
}
func GetElastiCacheReservedCacheNode(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	id := fields["id"]
	client := elasticache.NewFromConfig(cfg)
	out, err := client.DescribeReservedCacheNodes(ctx, &elasticache.DescribeReservedCacheNodesInput{
		ReservedCacheNodeId: &id,
	})
	if err != nil {
		if isErr(err, "ReservedCacheNodeNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, v := range out.ReservedCacheNodes {
		resource := elastiCacheReservedCacheNodeHandle(ctx, v)
		values = append(values, resource)
	}
	return values, nil
}

func ElastiCacheSubnetGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := elasticache.NewFromConfig(cfg)
	paginator := elasticache.NewDescribeCacheSubnetGroupsPaginator(client, &elasticache.DescribeCacheSubnetGroupsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, cacheSubnetGroup := range page.CacheSubnetGroups {
			resource := elastiCacheSubnetGroupHandle(ctx, cacheSubnetGroup)
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}
func elastiCacheSubnetGroupHandle(ctx context.Context, cacheSubnetGroup types.CacheSubnetGroup) models.Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *cacheSubnetGroup.ARN,
		Name:   *cacheSubnetGroup.CacheSubnetGroupName,
		Description: model.ElastiCacheSubnetGroupDescription{
			SubnetGroup: cacheSubnetGroup,
		},
	}
	return resource
}
func GetElastiCacheSubnetGroup(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	cacheSubnetGroupsName := fields["name"]
	client := elasticache.NewFromConfig(cfg)

	out, err := client.DescribeCacheSubnetGroups(ctx, &elasticache.DescribeCacheSubnetGroupsInput{
		CacheSubnetGroupName: &cacheSubnetGroupsName,
	})
	if err != nil {
		if isErr(err, "DescribeCacheSubnetGroupsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, cacheSubnetGroup := range out.CacheSubnetGroups {

		resource := elastiCacheSubnetGroupHandle(ctx, cacheSubnetGroup)
		values = append(values, resource)

	}
	return values, nil
}
