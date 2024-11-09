package aws

import (
	"context"

	"github.com/aws/aws-sdk-go/service/elasticache"
	opengovernance "github.com/opengovern/og-describer-aws/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// // TABLE DEFINITION
func tableAwsElasticacheRedisMetricCacheHitsHourly(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_elasticache_redis_metric_cache_hits_hourly",
		Description: "AWS Elasticache Redis CacheHits metric (Hourly)",
		List: &plugin.ListConfig{
			ParentHydrate: opengovernance.ListElastiCacheCluster,
			Hydrate:       listElastiCacheMetricCacheHitsHourly,
		},

		Columns: awsKaytuRegionalColumns(cwMetricColumns(
			[]*plugin.Column{
				{
					Name:        "cache_cluster_id",
					Description: "The cache cluster id.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("DimensionValue"),
				},
			})),
	}
}

func listElastiCacheMetricCacheHitsHourly(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	cacheClusterConfiguration := h.Item.(*elasticache.CacheCluster)
	return listCWMetricStatistics(ctx, d, "Hourly", "AWS/ElastiCache", "CacheHits", "CacheClusterId", *cacheClusterConfiguration.CacheClusterId)
}
