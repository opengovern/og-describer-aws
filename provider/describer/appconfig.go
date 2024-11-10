package describer

import (
	"context"
	"fmt"
	"github.com/opengovern/og-describer-aws/pkg/sdk/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appconfig"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/opengovern/og-describer-aws/provider/model"
)

func AppConfigApplication(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := appconfig.NewFromConfig(cfg)
	paginator := appconfig.NewListApplicationsPaginator(client, &appconfig.ListApplicationsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, application := range page.Items {
			resource, err := appConfigApplicationHandle(ctx, cfg, application)
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
func appConfigApplicationHandle(ctx context.Context, cfg aws.Config, application types.Application) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := appconfig.NewFromConfig(cfg)
	arn := fmt.Sprintf("arn:%s:appconfig:%s:%s:application/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, *application.Id)

	tags, err := client.ListTagsForResource(ctx, &appconfig.ListTagsForResourceInput{
		ResourceArn: aws.String(arn),
	})
	if err != nil {
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ID:     *application.Id,
		Name:   *application.Name,
		ARN:    arn,
		Description: model.AppConfigApplicationDescription{
			Application: application,
			Tags:        tags.Tags,
		},
	}
	return resource, nil
}
func GetAppConfigApplication(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	var values []models.Resource
	applicationId := fields["id"]
	client := appconfig.NewFromConfig(cfg)
	applications, err := client.ListApplications(ctx, &appconfig.ListApplicationsInput{})
	if err != nil {
		if isErr(err, "ListApplicationsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, application := range applications.Items {
		if application.Id != &applicationId {
			continue
		}
		resource, err := appConfigApplicationHandle(ctx, cfg, application)
		if err != nil {
			return nil, err
		}
		values = append(values, resource)
	}
	return values, nil
}
