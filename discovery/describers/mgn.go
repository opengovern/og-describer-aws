package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mgn"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
)

func MGNApplication(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := mgn.NewFromConfig(cfg)
	var values []models.Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		applications, err := client.ListApplications(ctx, &mgn.ListApplicationsInput{
			NextToken: prevToken,
		})
		if err != nil {
			if isErr(err, "UninitializedAccountException") {
				return nil, nil
			}
			return nil, err
		}

		for _, application := range applications.Items {
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *application.Arn,
				Name:   *application.Name,
				ID:     *application.ApplicationID,
				Description: model.MgnApplicationDescription{
					Application: application,
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

		return applications.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
