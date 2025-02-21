package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func SecretsManagerSecret(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := secretsmanager.NewFromConfig(cfg)
	paginator := secretsmanager.NewListSecretsPaginator(client, &secretsmanager.ListSecretsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, item := range page.SecretList {
			resource, err := secretsManagerSecretHandle(ctx, cfg, item.ARN, item.Name)
			if err != nil {
				return nil, err
			}

			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}
func secretsManagerSecretHandle(ctx context.Context, cfg aws.Config, Arn *string, Name *string) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := secretsmanager.NewFromConfig(cfg)
	out, err := client.DescribeSecret(ctx, &secretsmanager.DescribeSecretInput{
		SecretId: Arn,
	})
	if err != nil {
		return models.Resource{}, err
	}

	policy, err := client.GetResourcePolicy(ctx, &secretsmanager.GetResourcePolicyInput{
		SecretId: Arn,
	})
	if err != nil {
		return models.Resource{}, err
	}

	resource := models.Resource{
		Account: describeCtx.AccountID,
		Region:  describeCtx.OGRegion,
		ARN:     *Arn,
		Name:    *Name,
		Description: model.SecretsManagerSecretDescription{
			Secret:         out,
			ResourcePolicy: policy.ResourcePolicy,
		},
	}
	return resource, nil
}
func GetSecretsManagerSecret(ctx context.Context, cfg aws.Config, field map[string]string) ([]models.Resource, error) {
	secretId := field["id"]
	var values []models.Resource

	client := secretsmanager.NewFromConfig(cfg)
	secretValue, err := client.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: &secretId,
	})
	if err != nil {
		return nil, err
	}

	resource, err := secretsManagerSecretHandle(ctx, cfg, secretValue.ARN, secretValue.Name)
	if err != nil {
		return nil, err
	}

	values = append(values, resource)
	return values, nil
}
