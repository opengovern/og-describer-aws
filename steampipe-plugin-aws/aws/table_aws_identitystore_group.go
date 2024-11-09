package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsIdentityStoreGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_identitystore_group",
		Description: "AWS Identity Store Group",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListIdentityStoreGroup,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "identity_store_id",
				Description: "The globally unique identifier for the identity store.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.IdentityStoreId")},
			{
				Name:        "name",
				Description: "Contains the group's display name value.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.DisplayName")},
			{
				Name:        "id",
				Description: "The identifier for a group in the identity store.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.GroupId"),
			},
			{
				Name:        "external_ids",
				Description: "The identifier for a group in the identity store.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Group.ExternalIds"),
			},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.DisplayName")},

			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Group.GroupId").Transform(arnToAkas),
			},
		}),
	}
}
