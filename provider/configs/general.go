package configs

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "aws"                                    // example: aws, azure
	IntegrationName      = integration.Type("INTEGRATION_NAME")     // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describer-aws" // example: github.com/opengovern/og-describer-aws
)
