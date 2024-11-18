package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsCostForecastMonthly(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cost_forecast_monthly",
		Description: "AWS Cost Explorer - Cost Forecast (Monthly)",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCostExplorerForcastMonthly,
		},
		Columns: awsOgColumns([]*plugin.Column{
			{
				Name:        "period_start",
				Description: "Start timestamp for this cost metric",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.PeriodStart"),
			},
			{
				Name:        "period_end",
				Description: "End timestamp for this cost metric",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.PeriodEnd"),
			},
			{
				Name:        "mean_value",
				Description: "Average forecasted value",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.MeanValue"),
			},
		},
		),
	}
}
