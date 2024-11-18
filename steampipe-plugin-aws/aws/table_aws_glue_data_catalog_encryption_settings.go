package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAwsGlueDataCatalogEncryptionSettings(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_glue_data_catalog_encryption_settings",
		Description: "AWS Glue Data Catalog Encryption Settings",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListGlueDataCatalogEncryptionSettings,
		},

		Columns: awsOgRegionalColumns([]*plugin.Column{
			{
				Name:        "encryption_at_rest",
				Description: "A list of public keys to be used by the DataCatalogEncryptionSettingss for authentication.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataCatalogEncryptionSettings.EncryptionAtRest")},
			{
				Name:        "connection_password_encryption",
				Description: "A list of security group identifiers used in this DataCatalogEncryptionSettings.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataCatalogEncryptionSettings.ConnectionPasswordEncryption")},
		}),
	}
}
