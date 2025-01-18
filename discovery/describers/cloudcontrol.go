package describer

import (
	"context"
	"github.com/opengovern/og-describer-aws/pkg/sdk/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/opengovern/og-describer-aws/provider/model"
)

func CloudControlResource(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := cloudcontrol.NewFromConfig(cfg)

	paginator := cloudcontrol.NewListResourcesPaginator(client, &cloudcontrol.ListResourcesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ResourceDescriptions {
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ID:     *v.Identifier,
				Description: model.CloudControlResourceDescription{
					Resource: v,
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
