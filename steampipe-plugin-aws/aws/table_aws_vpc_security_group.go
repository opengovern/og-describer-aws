package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsVpcSecurityGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_security_group",
		Description: "AWS VPC Security Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("group_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidGroupId.Malformed", "InvalidGroupId.NotFound", "InvalidGroup.NotFound"}),
			},
			Hydrate: opengovernance.GetEC2SecurityGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEC2SecurityGroup,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "description", Require: plugin.Optional},
				{Name: "group_name", Require: plugin.Optional},
				{Name: "owner_id", Require: plugin.Optional},
				{Name: "vpc_id", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "group_name",
				Description: "The friendly name that identifies the security group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SecurityGroup.GroupName"),
			},
			{
				Name:        "group_id",
				Description: "Contains the unique ID to identify a security group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SecurityGroup.GroupId"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the security group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN"),
			},
			{
				Name:        "description",
				Description: "A description of the security group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SecurityGroup.Description"),
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC for the security group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SecurityGroup.VpcId"),
			},
			{
				Name:        "owner_id",
				Description: "Contains the AWS account ID of the owner of the security group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SecurityGroup.OwnerId"),
			},
			{
				Name:        "ip_permissions",
				Description: "A list of inbound rules associated with the security group",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SecurityGroup.IpPermissions"),
			},
			{
				Name:        "ip_permissions_egress",
				Description: "A list of outbound rules associated with the security group",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SecurityGroup.IpPermissionsEgress"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags that are attached to the security group",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SecurityGroup.Tags"),
			},

			// Standard columns for all tables
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getVpcSecurityGroupTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SecurityGroup.GroupName"),
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

func getVpcSecurityGroupTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	securityGroup := d.HydrateItem.(opengovernance.EC2SecurityGroup).Description.SecurityGroup

	// Get the resource tags
	if securityGroup.Tags != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range securityGroup.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
