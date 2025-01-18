package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mwaa"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
)

func MWAAEnvironment(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := mwaa.NewFromConfig(cfg)
	paginator := mwaa.NewListEnvironmentsPaginator(client, &mwaa.ListEnvironmentsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Environments {
			resource, err := mWAAEnvironmentHandle(ctx, cfg, v)
			if err != nil {
				return nil, err
			}
			emptyResource := models.Resource{}
			if err == nil && resource == emptyResource {
				continue
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
func mWAAEnvironmentHandle(ctx context.Context, cfg aws.Config, v string) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := mwaa.NewFromConfig(cfg)
	environment, err := client.GetEnvironment(ctx, &mwaa.GetEnvironmentInput{
		Name: &v,
	})
	if err != nil {
		if isErr(err, "GetEnvironmentNotFound") || isErr(err, "InvalidParameterVaLue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *environment.Environment.Arn,
		Name:   *environment.Environment.Name,
		Description: model.MWAAEnvironmentDescription{
			Environment: *environment.Environment,
		},
	}
	return resource, nil
}
func GetMWAAEnvironment(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	environmentName := fields["name"]
	var values []models.Resource
	resource, err := mWAAEnvironmentHandle(ctx, cfg, environmentName)
	if err != nil {
		return nil, err
	}
	emptyResource := models.Resource{}
	if err == nil && resource == emptyResource {
		return nil, nil
	}
	values = append(values, resource)
	return values, nil
}
