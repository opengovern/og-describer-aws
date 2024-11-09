package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsAppstreamApplication(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_appstream_application",
		Description: "AWS AppStream Application",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetAppStreamApplication,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAppStreamApplication,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the application",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.Arn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Application.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Application.Arn").Transform(arnToAkas),
			},
		}),
	}
}
