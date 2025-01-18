package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsBatchComputeEnvironment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_batch_compute_environment",
		Description: "AWS Batch ComputeEnvironment",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("compute_environment_name"),
			Hydrate:    opengovernance.GetBatchComputeEnvironment,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListBatchComputeEnvironment,
		},
		Columns: awsOgRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the compute environment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ComputeEnvironment.Uuid")},
			{
				Name:        "compute_environment_name",
				Description: "The name of the compute environment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ComputeEnvironment.ComputeEnvironmentName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the compute environment",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ComputeEnvironment.ComputeEnvironmentArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ComputeEnvironment.ComputeEnvironmentName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ComputeEnvironment.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ComputeEnvironment.ComputeEnvironmentArn").Transform(arnToAkas),
			},
		}),
	}
}
