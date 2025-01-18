package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
)

func ServerlessApplicationRepositoryApplication(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := serverlessapplicationrepository.NewFromConfig(cfg)
	paginator := serverlessapplicationrepository.NewListApplicationsPaginator(client, &serverlessapplicationrepository.ListApplicationsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, applicationSummary := range page.Applications {
			resource, err := serverlessApplicationRepositoryApplicationHandle(ctx, cfg, *applicationSummary.ApplicationId)
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
func serverlessApplicationRepositoryApplicationHandle(ctx context.Context, cfg aws.Config, applicationId string) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := serverlessapplicationrepository.NewFromConfig(cfg)

	application, err := client.GetApplication(ctx, &serverlessapplicationrepository.GetApplicationInput{
		ApplicationId: &applicationId,
	})
	if err != nil {
		return models.Resource{}, err
	}

	policy, err := client.GetApplicationPolicy(ctx, &serverlessapplicationrepository.GetApplicationPolicyInput{
		ApplicationId: &applicationId,
	})
	if err != nil {
		policy = &serverlessapplicationrepository.GetApplicationPolicyOutput{}
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *application.ApplicationId,
		Name:   *application.ApplicationId,
		Description: model.ServerlessApplicationRepositoryApplicationDescription{
			Application: *application,
			Statements:  policy.Statements,
		},
	}
	return resource, nil
}
func GetServerlessApplicationRepositoryApplication(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	applicationId := fields["applicationId"]

	var values []models.Resource
	resource, err := serverlessApplicationRepositoryApplicationHandle(ctx, cfg, applicationId)
	if err != nil {
		return nil, err
	}

	values = append(values, resource)
	return values, nil
}
