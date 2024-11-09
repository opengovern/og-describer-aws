package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsDynamodbGlobalSecondaryIndex(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_dynamodb_global_secondary_index",
		Description: "AWS DynamoDb GlobalSecondaryIndex",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("index_arn"),
			Hydrate:    opengovernance.GetDynamoDbGlobalSecondaryIndex,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDynamoDbGlobalSecondaryIndex,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "index_name",
				Description: "The name of the global secondary index.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GlobalSecondaryIndex.IndexName")},
			{
				Name:        "index_arn",
				Description: "The Amazon Resource Name (ARN) of the global secondary index",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GlobalSecondaryIndex.IndexArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GlobalSecondaryIndex.IndexName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.GlobalSecondaryIndex.IndexArn").Transform(arnToAkas),
			},
		}),
	}
}
