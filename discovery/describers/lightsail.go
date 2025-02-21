package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func LightsailInstance(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := lightsail.NewFromConfig(cfg)

	var values []models.Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		instances, err := client.GetInstances(ctx, &lightsail.GetInstancesInput{
			PageToken: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, instance := range instances.Instances {
			resource := lightsailInstanceHandle(ctx, instance)

			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}

		return instances.NextPageToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
func lightsailInstanceHandle(ctx context.Context, instance types.Instance) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     *instance.Arn,
		Account: describeCtx.AccountID,
		Name:    *instance.Name,
		Description: model.LightsailInstanceDescription{
			Instance: instance,
		},
	}
	return resource
}
func GetLightsailInstance(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	instanceName := fields["name"]
	var values []models.Resource

	client := lightsail.NewFromConfig(cfg)
	instance, err := client.GetInstance(ctx, &lightsail.GetInstanceInput{
		InstanceName: &instanceName,
	})
	if err != nil {
		if isErr(err, "GetInstanceNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	values = append(values, lightsailInstanceHandle(ctx, *instance.Instance))
	return values, nil
}
