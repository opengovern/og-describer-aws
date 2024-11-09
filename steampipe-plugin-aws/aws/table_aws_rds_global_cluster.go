package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsRDSGlobalCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_rds_global_cluster",
		Description: "AWS RDS GlobalCluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("global_cluster_identifier"),
			Hydrate:    opengovernance.GetRDSGlobalCluster,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRDSGlobalCluster,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "global_cluster_identifier",
				Description: "The id of the global cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GlobalCluster.GlobalClusterIdentifier")},
			{
				Name:        "global_cluster_arn",
				Description: "The Amazon Resource Name (ARN) of the global cluster",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GlobalCluster.GlobalClusterArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GlobalCluster.GlobalClusterIdentifier")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getRDSGlobalClusterTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.GlobalCluster.GlobalClusterArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getRDSGlobalClusterTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.RDSGlobalCluster).Description.Tags
	return rdsV2TagsToMap(tags)
}
