package aws

import (
	"context"

	"github.com/aws/aws-sdk-go/service/elasticache"
	opengovernance "github.com/opengovern/og-describer-aws/SDK/generated"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// // TABLE DEFINITION
func tableAwsElasticacheRedisMetricGetTypeCmdsHourly(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_elasticache_redis_metric_get_type_cmds_hourly",
		Description: "AWS Elasticache Redis GetTypeCmds metric(Hourly)",
		List: &plugin.ListConfig{
			ParentHydrate: opengovernance.ListElastiCacheCluster,
			Hydrate:       listElastiCacheMetricGetTypeCmdsHourly,
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

func listElastiCacheMetricGetTypeCmdsHourly(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	cacheClusterConfiguration := h.Item.(*elasticache.CacheCluster)
	return listCWMetricStatistics(ctx, d, "Hourly", "AWS/ElastiCache", "GetTypeCmds", "CacheClusterId", *cacheClusterConfiguration.CacheClusterId)
}
