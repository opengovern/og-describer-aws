package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2SslPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_ssl_policy",
		Description: "AWS EC2 SSL Policy",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "region"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"SSLPolicyNotFound"}),
			},
			Hydrate: opengovernance.GetElasticLoadBalancingV2SslPolicy,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListElasticLoadBalancingV2SslPolicy,
		},

		Columns: awsOgRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SslPolicy.Name")},
			{
				Name:        "ciphers",
				Description: "A list of ciphers.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SslPolicy.Ciphers")},
			{
				Name:        "ssl_protocols",
				Description: "A list of protocols.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SslPolicy.SslProtocols")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SslPolicy.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
