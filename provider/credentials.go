package provider

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	model "github.com/opengovern/og-describer-aws/pkg/sdk/models"
	"github.com/opengovern/og-describer-aws/provider/configs"
	awsmodel "github.com/opengovern/og-describer-aws/provider/model"
	"github.com/opengovern/og-util/pkg/describe"
	"golang.org/x/net/context"
	"strings"
)

// GenerateAWSConfig creates an AWS configuration using the provided credentials.
// It can assume a role if roleNameToAssume is provided.
func GenerateAWSConfig(awsAccessKeyID string, awsSecretAccessKey string, roleNameToAssume string, externalID string, accountID string) (aws.Config, error) {
	// Step 1: Create base credentials provider
	credsProvider := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
		awsAccessKeyID,
		awsSecretAccessKey,
		"",
	))

	// Step 2: Manually create the AWS Config struct with explicit credentials and region
	cfg := aws.Config{
		Region:      "us-east-2",
		Credentials: credsProvider,
	}

	// Step 3: If a role is specified to assume, perform the AssumeRole operation
	if roleNameToAssume != "" {
		// Construct Role ARN
		roleArn := fmt.Sprintf("arn:aws:iam::%s:role/%s", accountID, roleNameToAssume)

		// Use STS client with the current config
		stsClient := sts.NewFromConfig(cfg)

		// Prepare AssumeRole input
		input := &sts.AssumeRoleInput{
			RoleArn:         aws.String(roleArn),
			RoleSessionName: aws.String("GenerateAWSConfigSession"),
		}
		if externalID != "" {
			input.ExternalId = aws.String(externalID)
		}

		// Perform AssumeRole
		assumeRoleOutput, err := stsClient.AssumeRole(context.TODO(), input)
		if err != nil {
			return aws.Config{}, fmt.Errorf("failed to assume role: %v", err)
		}

		// Update credentials provider with assumed role credentials
		credsProvider = aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
			*assumeRoleOutput.Credentials.AccessKeyId,
			*assumeRoleOutput.Credentials.SecretAccessKey,
			*assumeRoleOutput.Credentials.SessionToken,
		))

		// Update the AWS Config with the new credentials
		cfg.Credentials = credsProvider
	}
	return cfg, nil
}

func GetConfig(config configs.IntegrationCredentials) (*aws.Config, error) {
	cfg, err := GenerateAWSConfig(config.AwsAccessKeyID, config.AwsSecretAccessKey, config.CrossAccountRoleName, config.ExternalID, config.AccountID)
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
	awsMetadata := awsmodel.Metadata{
		Name:         resource.Name,
		AccountID:    job.ProviderID,
		SourceID:     job.IntegrationID,
		Region:       resource.Region,
		ResourceType: strings.ToLower(job.ResourceType),
	}

	awsMetadataBytes, err := json.Marshal(awsMetadata)
	if err != nil {
		return nil, fmt.Errorf("marshal metadata: %v", err.Error())
	}

	err = json.Unmarshal(awsMetadataBytes, &metadata)
	if err != nil {
		return nil, fmt.Errorf("unmarshal metadata: %v", err.Error())
	}
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
