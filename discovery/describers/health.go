package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/health/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/health"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func HealthEvent(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := health.NewFromConfig(cfg)
	paginator := health.NewDescribeEventsPaginator(client, &health.DescribeEventsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, event := range page.Events {
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *event.Arn,
				Description: model.HealthEventDescription{
					Event: event,
				},
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

func HealthAffectedEntity(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := health.NewFromConfig(cfg)
	paginator := health.NewDescribeEventsPaginator(client, &health.DescribeEventsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, event := range page.Events {
			entitiesPaginator := health.NewDescribeAffectedEntitiesPaginator(client, &health.DescribeAffectedEntitiesInput{
				Filter: &types.EntityFilter{
					EventArns: []string{*aws.String(*event.Arn)},
				},
			})
			for entitiesPaginator.HasMorePages() {
				entitiesPage, err := entitiesPaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, entity := range entitiesPage.Entities {
					resource := models.Resource{
						Region: describeCtx.OGRegion,
						ARN:    *event.Arn,
						Description: model.HealthAffectedEntityDescription{
							Entity: entity,
						},
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
		}
	}
	return values, nil
}
