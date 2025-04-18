package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEmrInstanceGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_emr_instance_group",
		Description: "AWS EMR Instance Group",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEMRInstanceGroup,
		},

		Columns: awsOgRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the instance group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceGroup.Name")},
			{
				Name:        "id",
				Description: "The identifier of the instance group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceGroup.Id")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the instance group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN"),
			},
			{
				Name:        "cluster_id",
				Description: "The unique identifier for the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClusterID")},
			{
				Name:        "instance_group_type",
				Description: "The type of the instance group. Valid values are MASTER, CORE or TASK.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceGroup.InstanceGroupType")},
			{
				Name:        "instance_type",
				Description: "The EC2 instance type for all instances in the instance group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceGroup.InstanceType")},
			{
				Name:        "state",
				Description: "The current state of the instance group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceGroup.Status.State")},
			{
				Name:        "bid_price",
				Description: "The maximum price you are willing to pay for Spot Instances. If specified, indicates that the instance group uses Spot Instances.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceGroup.BidPrice")},
			{
				Name:        "configurations_version",
				Description: "The version number of the requested configuration specification for this instance group.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.InstanceGroup.ConfigurationsVersion")},
			{
				Name:        "ebs_optimized",
				Description: "Indicates whether the instance group is EBS-optimized, or not.  An Amazon EBS-optimized instance uses an optimized configuration stack and provides additional, dedicated capacity for Amazon EBS I/O.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.InstanceGroup.EbsOptimized")},
			{
				Name:        "last_successfully_applied_configurations_version",
				Description: "The version number of a configuration specification that was successfully applied for an instance group last time.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.InstanceGroup.LastSuccessfullyAppliedConfigurationsVersion")},
			{
				Name:        "market",
				Description: "The marketplace to provision instances for this group. Valid values are ON_DEMAND or SPOT.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceGroup.Market")},
			{
				Name:        "requested_instance_count",
				Description: "The target number of instances for the instance group.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.InstanceGroup.RequestedInstanceCount")},
			{
				Name:        "running_instance_count",
				Description: "The number of instances currently running in this instance group.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.InstanceGroup.RunningInstanceCount")},
			{
				Name:        "autoscaling_policy",
				Description: "An automatic scaling policy for a core instance group or task instance group in an Amazon EMR cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceGroup.AutoScalingPolicy")},
			{
				Name:        "configurations",
				Description: "A list of configurations supplied for an EMR cluster instance group. Only availbale for Amazon EMR releases 4.x or later.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceGroup.Configurations")},
			{
				Name:        "ebs_block_devices",
				Description: "The EBS block devices that are mapped to this instance group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceGroup.EbsBlockDevices")},
			{
				Name:        "last_successfully_applied_configurations",
				Description: "A list of configurations that were successfully applied for an instance group last time.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceGroup.LastSuccessfullyAppliedConfigurations")},
			{
				Name:        "shrink_policy",
				Description: "Policy for customizing shrink operations.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceGroup.ShrinkPolicy")},
			{
				Name:        "state_change_reason",
				Description: "The status change reason details for the instance group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceGroup.Status.StateChangeReason")},
			{
				Name:        "status_timeline",
				Description: "The timeline of the instance group status over time.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.InstanceGroup.Status.Timeline")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(EmrInstanceGroupTitle),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func EmrInstanceGroupTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.EMRInstanceGroup).Description.InstanceGroup

	if data.Name == nil || *data.Name == "" {
		return data.Id, nil
	}
	return data.Name, nil
}
