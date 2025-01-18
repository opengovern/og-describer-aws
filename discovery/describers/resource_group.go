package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
)

func ResourceGroups(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := resourcegroups.NewFromConfig(cfg)
	paginator := resourcegroups.NewListGroupsPaginator(client, &resourcegroups.ListGroupsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.GroupIdentifiers {
			resources, err := client.ListGroupResources(ctx, &resourcegroups.ListGroupResourcesInput{Group: v.GroupArn})
			if err != nil {
				return nil, err
			}

			tags, err := client.GetTags(ctx, &resourcegroups.GetTagsInput{Arn: v.GroupArn})
			if err != nil {
				return nil, err
			}

			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *v.GroupArn,
				Name:   *v.GroupName,
				Description: model.ResourceGroupsGroupDescription{
					GroupIdentifier: v,
					Resources:       resources.Resources,
					Tags:            tags,
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
