package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsIamServerCertificate(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_iam_server_certificate",
		Description: "AWS IAM Server Certificate",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListIAMServerCertificate,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "path", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    opengovernance.GetIAMServerCertificate,
			KeyColumns: plugin.AllColumns([]string{"name"}),
		},
		Columns: awsOgRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name that identifies the server certificate.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServerCertificate.ServerCertificateMetadata.ServerCertificateName"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the server certificate.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServerCertificate.ServerCertificateMetadata.Arn"),
			},
			{
				Name:        "server_certificate_id",
				Description: "The stable and unique string identifying the server certificate.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServerCertificate.ServerCertificateMetadata.ServerCertificateId"),
			},
			{
				Name:        "expiration",
				Description: "The date on which the certificate is set to expire.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ServerCertificate.ServerCertificateMetadata.Expiration"),
			},
			{
				Name:        "certificate_body",
				Description: "The contents of the public key certificate.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServerCertificate.CertificateBody"),
			},
			{
				Name:        "certificate_body_length",
				Description: "The contents of the public key certificate.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.BodyLength"),
			},
			{
				Name:        "certificate_chain",
				Description: "The contents of the public key certificate chain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServerCertificate.CertificateChain"),
			},
			{
				Name:        "path",
				Description: "The path to the server certificate.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServerCertificate.ServerCertificateMetadata.Path"),
			},
			{
				Name:        "upload_date",
				Description: "The Amazon Resource Name (ARN) of the account that is designated as the management account for the organization",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ServerCertificate.ServerCertificateMetadata.UploadDate"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached with the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServerCertificate.Tags"),
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServerCertificate.ServerCertificateMetadata.ServerCertificateName"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServerCertificate.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServerCertificate.ServerCertificateMetadata.Arn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS
