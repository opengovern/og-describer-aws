package describer

import (
	"context"
	"github.com/opengovern/og-describer-aws/pkg/sdk/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	"github.com/opengovern/og-describer-aws/provider/model"
)

func TimestreamDatabase(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := timestreamwrite.NewFromConfig(cfg)
	paginator := timestreamwrite.NewListDatabasesPaginator(client, &timestreamwrite.ListDatabasesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Databases {
			tags, err := client.ListTagsForResource(ctx, &timestreamwrite.ListTagsForResourceInput{ResourceARN: v.Arn})
			if err != nil {
				return nil, err
			}

			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *v.Arn,
				Name:   *v.DatabaseName,
				Description: model.TimestreamDatabaseDescription{
					Database: v,
					Tags:     tags,
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
