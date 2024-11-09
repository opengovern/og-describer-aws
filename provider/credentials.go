package provider

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	model "github.com/opengovern/og-describer-aws/pkg/sdk/models"
	"github.com/opengovern/og-describer-aws/provider/configs"
	"github.com/opengovern/og-util/pkg/describe"
	"golang.org/x/net/context"
)

// GenerateAWSConfig creates an AWS configuration using the provided credentials provider.
func GenerateAWSConfig(credsProvider aws.CredentialsProvider) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credsProvider),
	)
	if err != nil {
		return aws.Config{}, fmt.Errorf("failed to load configuration: %v", err)
	}
	return cfg, nil
}

func GetConfig(config configs.IntegrationCredentials) (*aws.Config, error) {
	credsProvider := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
		config.AwsAccessKeyID,
		config.AwsSecretAccessKey,
		"",
	))

	cfg, err := GenerateAWSConfig(credsProvider)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// AccountCredentialsFromMap TODO: converts a map to an configs.IntegrationCredentials.
func AccountCredentialsFromMap(m map[string]any) (configs.IntegrationCredentials, error) {
	mj, err := json.Marshal(m)
	if err != nil {
		return configs.IntegrationCredentials{}, err
	}

	var c configs.IntegrationCredentials
	err = json.Unmarshal(mj, &c)
	if err != nil {
		return configs.IntegrationCredentials{}, err
	}

	return c, nil
}

// GetResourceMetadata TODO: Get metadata as a map to add to the resources
func GetResourceMetadata(job describe.DescribeJob, resource model.Resource) (map[string]string, error) {
	metadata := make(map[string]string)

	return metadata, nil
}

// AdjustResource TODO: Do any needed adjustment on resource object before storing
func AdjustResource(job describe.DescribeJob, resource *model.Resource) error {
	return nil
}

// GetAdditionalParameters TODO: pass additional parameters needed in describer wrappers in /provider/describer_wrapper.go
func GetAdditionalParameters(job describe.DescribeJob) (map[string]any, error) {
	additionalParameters := make(map[string]any)

	additionalParameters["accountID"] = job.ProviderID

	return additionalParameters, nil
}
