package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsCostByLinkedAccountDaily(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cost_by_account_daily",
		Description: "AWS Cost Explorer - Cost by Linked Account (Daily)",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCostExplorerByAccountDaily,
		},
		Columns: awsOgColumns(
			ogCostExplorerColumns([]*plugin.Column{
				{
					Name:        "linked_account_id",
					Description: "The AWS Account ID.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("Description.Dimension1"),
				},
			}),
		),
	}
}
