package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsBackupRegionSetting(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_backup_region_settings",
		Description: "AWS Backup Region Settings",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListBackupRegionSetting,
		},
		Columns: awsOgRegionalColumns([]*plugin.Column{
			{
				Name:        "resource_type_management_preference",
				Description: "Resource Type Management Preference.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourceTypeManagementPreference"),
			},
			{
				Name:        "resource_type_opt_in_preference",
				Description: "Resource Type Opt In Preference.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourceTypeOptInPreference"),
			},
			// Steampipe standard columns
			{
				Name:        "region",
				Description: "The AWS Region in which the settings are for.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Region"),
			},
		}),
	}
}
