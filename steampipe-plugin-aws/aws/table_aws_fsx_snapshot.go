package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsFsxSnapshot(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_fsx_snapshot",
		Description: "AWS FSX Snapshot",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("snapshot_id"),
			Hydrate:    opengovernance.GetFSXSnapshot,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListFSXSnapshot,
		},
		Columns: awsOgRegionalColumns([]*plugin.Column{
			{
				Name:        "snapshot_id",
				Description: "The id of the snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.SnapshotId")},
			{
				Name:        "name",
				Description: "The name of the snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.Name")},
			{
				Name:        "resource_arn",
				Description: "The Amazon Resource Name (ARN) of the snapshot",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.ResourceARN")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.SnapshotId")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getFsxSnapshotTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Snapshot.ResourceARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getFsxSnapshotTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.FSXSnapshot).Description.Snapshot.Tags
	return fsxV2TagsToMap(tags)
}
