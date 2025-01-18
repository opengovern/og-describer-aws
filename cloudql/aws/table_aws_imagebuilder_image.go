package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsImageBuilderImage(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_imagebuilder_image",
		Description: "AWS ImageBuilder Image",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetImageBuilderImage,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListImageBuilderImage,
		},
		Columns: awsOgRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the image",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Arn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Image.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Image.Arn").Transform(arnToAkas),
			},
		}),
	}
}
