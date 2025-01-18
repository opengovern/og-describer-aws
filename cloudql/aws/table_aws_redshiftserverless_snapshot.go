package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsRedshiftServerlessSnapshot(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_redshiftserverless_snapshot",
		Description: "AWS RedshiftServerless Snapshot",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("snapshot_name"),
			Hydrate:    opengovernance.GetRedshiftServerlessSnapshot,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRedshiftServerlessSnapshot,
		},
		Columns: awsOgRegionalColumns([]*plugin.Column{
			{
				Name:        "snapshot_name",
				Description: "The name of the snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.SnapshotName")},
			{
				Name:        "snapshot_arn",
				Description: "The Amazon Resource Name (ARN) of the snapshot",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.SnapshotArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.SnapshotName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getRedshiftServerlessSnapshotTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Snapshot.SnapshotArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getRedshiftServerlessSnapshotTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.RedshiftServerlessSnapshot).Description.Tags
	return redshiftServerlessV2TagsToMap(tags)
}
