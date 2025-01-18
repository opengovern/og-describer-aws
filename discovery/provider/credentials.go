package provider

import (
	"encoding/json"
	model "github.com/opengovern/og-describer-aws/discovery/pkg/models"
	"github.com/opengovern/og-util/pkg/describe"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"fmt"
	"golang.org/x/net/context"

)

// AccountCredentialsFromMap TODO: converts a map to a configs.IntegrationCredentials.
func AccountCredentialsFromMap(m map[string]any) (model.IntegrationCredentials, error) {
	mj, err := json.Marshal(m)
	if err != nil {
		return model.IntegrationCredentials{}, err
	}

	var c model.IntegrationCredentials
	err = json.Unmarshal(mj, &c)
	if err != nil {
		return model.IntegrationCredentials{}, err
	}

	return c, nil
}

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

func GetConfig(config model.IntegrationCredentials) (*aws.Config, error) {
	cfg, err := GenerateAWSConfig(config.AwsAccessKeyID, config.AwsSecretAccessKey, config.CrossAccountRoleName, config.ExternalID, config.AccountID)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
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
func GetAdditionalParameters(job describe.DescribeJob) (map[string]string, error) {
	additionalParameters := make(map[string]string)

	if _, ok := job.IntegrationLabels["OrganizationName"]; ok {
		additionalParameters["OrganizationName"] = job.IntegrationLabels["OrganizationName"]
	}

	return additionalParameters, nil
}
