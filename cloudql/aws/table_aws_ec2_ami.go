package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2Ami(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_ami",
		Description: "AWS EC2 AMI",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("image_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidAMIID.NotFound", "InvalidAMIID.Unavailable", "InvalidAMIID.Malformed"}),
			},
			Hydrate: opengovernance.GetEC2AMI,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEC2AMI,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "architecture", Require: plugin.Optional},
				{Name: "description", Require: plugin.Optional},
				{Name: "ena_support", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "hypervisor", Require: plugin.Optional},
				{Name: "image_type", Require: plugin.Optional},
				{Name: "public", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "kernel_id", Require: plugin.Optional},
				{Name: "platform", Require: plugin.Optional},
				{Name: "name", Require: plugin.Optional},
				{Name: "ramdisk_id", Require: plugin.Optional},
				{Name: "root_device_name", Require: plugin.Optional},
				{Name: "root_device_type", Require: plugin.Optional},
				{Name: "state", Require: plugin.Optional},
				{Name: "sriov_net_support", Require: plugin.Optional},
				{Name: "virtualization_type", Require: plugin.Optional},
			},
		},

		Columns: awsOgRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the AMI that was provided during image creation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.Name")},
			{
				Name:        "image_id",
				Description: "The ID of the AMI.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.ImageId")},
			{
				Name:        "state",
				Description: "The current state of the AMI. If the state is available, the image is successfully registered and can be used to launch an instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.State")},
			{
				Name:        "image_type",
				Description: "The type of image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.ImageType")},
			{
				Name:        "image_location",
				Description: "The location of the AMI.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.ImageLocation")},
			{
				Name:        "creation_date",
				Description: "The date and time when the image was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.AMI.CreationDate")},
			{
				Name:        "architecture",
				Description: "The architecture of the image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.Architecture")},
			{
				Name:        "description",
				Description: "The description of the AMI that was provided during image creation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.Description")},
			{
				Name:        "ena_support",
				Description: "Specifies whether enhanced networking with ENA is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.AMI.EnaSupport")},
			{
				Name:        "hypervisor",
				Description: "The hypervisor type of the image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.Hypervisor")},
			{
				Name:        "image_owner_alias",
				Description: "The AWS account alias (for example, amazon, self) or the AWS account ID of the AMI owner.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.ImageOwnerAlias")},
			{
				Name:        "imds_support",
				Description: "If v2.0, it indicates that IMDSv2 is specified in the AMI.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.ImdsSupport")},
			{
				Name:        "kernel_id",
				Description: "The kernel associated with the image, if any. Only applicable for machine images.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.KernelId")},
			{
				Name:        "owner_id",
				Description: "The AWS account ID of the image owner.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.OwnerId")},
			{
				Name:        "platform",
				Description: "This value is set to windows for Windows AMIs; otherwise, it is blank.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.Platform")},
			{
				Name:        "platform_details",
				Description: "The platform details associated with the billing code of the AMI. For more information, see Obtaining Billing Information (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-billing-info.html) in the Amazon Elastic Compute Cloud User Guide.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.PlatformDetails")},
			{
				Name:        "public",
				Description: "Indicates whether the image has public launch permissions. The value is true if this image has public launch permissions or false if it has only implicit and explicit launch permissions.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.AMI.Public")},
			{
				Name:        "ramdisk_id",
				Description: "The RAM disk associated with the image, if any. Only applicable for machine images.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.RamdiskId")},
			{
				Name:        "root_device_name",
				Description: "The device name of the root device volume (for example, /dev/sda1).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.RootDeviceName")},
			{
				Name:        "root_device_type",
				Description: "The type of root device used by the AMI. The AMI can use an EBS volume or an instance store volume.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.RootDeviceType")},
			{
				Name:        "sriov_net_support",
				Description: "Specifies whether enhanced networking with the Intel 82599 Virtual Function interface is enabled.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.SriovNetSupport")},
			{
				Name:        "usage_operation",
				Description: "The operation of the Amazon EC2 instance and the billing code that is associated with the AMI. For the list of UsageOperation codes, see Platform Details and [Usage Operation Billing Codes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-billing-info.html#billing-info) in the Amazon Elastic Compute Cloud User Guide.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.UsageOperation")},
			{
				Name:        "virtualization_type",
				Description: "The type of virtualization of the AMI.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AMI.VirtualizationType")},
			{
				Name:        "block_device_mappings",
				Description: "Any block device mapping entries.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AMI.BlockDeviceMappings")},
			{
				Name:        "product_codes",
				Description: "Any product codes associated with the AMI.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AMI.ProductCodes")},
			{
				Name:        "launch_permissions",
				Description: "The users and groups that have the permissions for creating instances from the AMI.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LaunchPermissions")},
			{
				Name:        "tags_src",
				Description: "A list of tags attached to the AMI.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AMI.Tags")},
			{
				Name:        "is_aws_backup_managed",
				Description: "Is backup managed by AWS",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.From(isAWSBackupManaged)},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getEc2AmiTurbotTitle),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEc2AmiTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEC2AMIAkas)},
		}),
	}
}

// Create Session

//// HYDRATE FUNCTIONS

// create service

// create service

// Get data for turbot defined properties

//// TRANSFORM FUNCTIONS

func getEC2AMIAkas(_ context.Context, d *transform.TransformData) (interface{}, error) {
	image := d.HydrateItem.(opengovernance.EC2AMI).Description.AMI
	metadata := d.HydrateItem.(opengovernance.EC2AMI).Metadata

	akas := []string{"arn:" + metadata.Partition + ":ec2:" + metadata.Region + ":" + metadata.AccountID + ":image/" + *image.ImageId}
	return akas, nil
}

func getEc2AmiTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	image := d.HydrateItem.(opengovernance.EC2AMI).Description.AMI
	return ec2V2TagsToMap(image.Tags)
}

func getEc2AmiTurbotTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.EC2AMI).Description.AMI

	title := data.ImageId
	if data.Name != nil {
		title = data.Name
	}

	return title, nil
}

func isAWSBackupManaged(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.EC2AMI).Description.AMI.Tags
	isAWSManaged := false
	for _, tag := range tags {
		if tag.Key != nil && *tag.Key == "aws:backup:source-resource" {
			isAWSManaged = true
		}
	}

	return isAWSManaged, nil
}
