package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsSsoAdminAccountAssignment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ssoadmin_account_assignment",
		Description: "AWS SSO Account Assignment",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSSOAdminAccountAssignment,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "target_account_id",
				Description: "The identifier of the AWS account from which to list the assignments.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccountAssignment.AccountId"),
			},
			{
				Name:        "instance_arn",
				Description: "The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.InstanceArn"),
			},
			{
				Name:        "permission_set_arn",
				Description: "The ARN of the permission set from which to list assignments.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccountAssignment.PermissionSetArn"),
			},
			{
				Name:        "principal_type",
				Description: "The entity type for which the assignment will be created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccountAssignment.PrincipalType"),
			},
			{
				Name:        "principal_id",
				Description: "An identifier for an object in IAM Identity Center, such as a user or group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccountAssignment.PrincipalId"),
			},
		}),
	}
}
