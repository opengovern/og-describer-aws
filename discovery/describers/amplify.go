package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/amplify/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
)

func AmplifyApp(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := amplify.NewFromConfig(cfg)

	var values []models.Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.ListApps(ctx, &amplify.ListAppsInput{
			MaxResults: 100,
			NextToken:  prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, item := range output.Apps {
			resource := amplifyAppHandle(ctx, item)
			if stream != nil {
				m := *stream
				err := m(resource)
				if err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
		return output.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
func amplifyAppHandle(ctx context.Context, item types.App) models.Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := models.Resource{
		Region: describeCtx.OGRegion,
		Name:   *item.Name,
		ARN:    *item.AppArn,
		ID:     *item.AppId,
		Description: model.AmplifyAppDescription{
			App: item,
		},
	}
	return resource
}
func GetAmplifyApp(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	appId := fields["appId"]
	client := amplify.NewFromConfig(cfg)

	out, err := client.ListApps(ctx, &amplify.ListAppsInput{})
	if err != nil {
		if isErr(err, "ListAppsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, app := range out.Apps {
		if *app.AppId != appId {
			continue
		}
		resource := amplifyAppHandle(ctx, app)
		values = append(values, resource)
	}

	return values, nil
}
