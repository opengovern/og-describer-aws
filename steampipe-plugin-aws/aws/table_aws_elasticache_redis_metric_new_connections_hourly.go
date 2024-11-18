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
func tableAwsElasticacheRedisMetricNewConnectionsHourly(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_elasticache_redis_metric_new_connections_hourly",
		Description: "AWS Elasticache Redis NewConnections metric (Hourly)",
		List: &plugin.ListConfig{
			ParentHydrate: opengovernance.ListElastiCacheCluster,
			Hydrate:       listElastiCacheMetricNewConnectionsHourly,
		},

		Columns: awsOgRegionalColumns(cwMetricColumns(
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

func listElastiCacheMetricNewConnectionsHourly(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	cacheClusterConfiguration := h.Item.(*elasticache.CacheCluster)
	return listCWMetricStatistics(ctx, d, "Hourly", "AWS/ElastiCache", "NewConnections", "CacheClusterId", *cacheClusterConfiguration.CacheClusterId)
}
