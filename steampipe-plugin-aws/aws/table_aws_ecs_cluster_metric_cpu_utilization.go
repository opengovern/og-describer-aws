package aws

import (
	"context"
	"strings"

	opengovernance "github.com/opengovern/og-describer-aws/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEcsClusterMetricCpuUtilization(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ecs_cluster_metric_cpu_utilization",
		Description: "AWS ECS Cluster Cloudwatch Metrics - CPU Utilization",
		List: &plugin.ListConfig{
			ParentHydrate: opengovernance.ListECSCluster,
			Hydrate:       listEcsClusterMetricCpuUtilization,
		},

		Columns: awsKaytuRegionalColumns(cwMetricColumns(
			[]*plugin.Column{
				{
					Name:        "cluster_name",
					Description: "A user-generated string that you use to identify your cluster.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("DimensionValue"),
				},
			})),
	}
}

func listEcsClusterMetricCpuUtilization(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data := h.Item.(opengovernance.ECSCluster).Description.Cluster
	clusterName := strings.Split(*data.ClusterArn, "/")[1]
	return listCWMetricStatistics(ctx, d, "5_MIN", "AWS/ECS", "CPUUtilization", "ClusterName", clusterName)
}
