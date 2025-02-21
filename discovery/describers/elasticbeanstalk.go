package describers

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func ElasticBeanstalkEnvironment(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := elasticbeanstalk.NewFromConfig(cfg)

	out, err := client.DescribeEnvironments(ctx, &elasticbeanstalk.DescribeEnvironmentsInput{})
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, item := range out.Environments {

		resource, err := elasticBeanstalkEnvironmentHandle(ctx, cfg, item)
		emptyResource := models.Resource{}
		if err == nil && resource == emptyResource {
			return nil, nil
		}
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

	return values, nil
}
func elasticBeanstalkEnvironmentHandle(ctx context.Context, cfg aws.Config, item types.EnvironmentDescription) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)

	client := elasticbeanstalk.NewFromConfig(cfg)

	tags, err := client.ListTagsForResource(ctx, &elasticbeanstalk.ListTagsForResourceInput{
		ResourceArn: item.EnvironmentArn,
	})
	if err != nil {
		if isErr(err, "ListTagsForResourceNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	managedActions, err := client.DescribeEnvironmentManagedActions(ctx, &elasticbeanstalk.DescribeEnvironmentManagedActionsInput{
		EnvironmentId:   item.EnvironmentId,
		EnvironmentName: item.EnvironmentName,
	})
	if err != nil {
		if isErr(err, "DescribeEnvironmentManagedActionsNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	params := &elasticbeanstalk.DescribeConfigurationSettingsInput{
		ApplicationName: item.ApplicationName,
		EnvironmentName: item.EnvironmentName,
	}

	var configurationSettingsDescription []types.ConfigurationSettingsDescription
	configurationSettings, err := client.DescribeConfigurationSettings(ctx, params)
	if err != nil {
		if isErr(err, "DescribeConfigurationSettingsNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}
	if configurationSettings != nil && len(configurationSettings.ConfigurationSettings) > 0 {
		configurationSettingsDescription = configurationSettings.ConfigurationSettings
	}

	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     *item.EnvironmentArn,
		Name:    *item.EnvironmentName,
		Account: describeCtx.AccountID,
		Description: model.ElasticBeanstalkEnvironmentDescription{
			EnvironmentDescription: item,
			ManagedAction:          managedActions.ManagedActions,
			Tags:                   tags.ResourceTags,
			ConfigurationSetting:   configurationSettingsDescription,
		},
	}
	return resource, nil
}
func GetElasticBeanstalkEnvironment(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	client := elasticbeanstalk.NewFromConfig(cfg)

	environmentName := fields["name"]

	out, err := client.DescribeEnvironments(ctx, &elasticbeanstalk.DescribeEnvironmentsInput{
		EnvironmentNames: []string{environmentName},
	})
	if err != nil {
		if isErr(err, "DescribeEnvironmentsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, item := range out.Environments {

		resource, err := elasticBeanstalkEnvironmentHandle(ctx, cfg, item)
		emptyResource := models.Resource{}
		if err == nil && resource == emptyResource {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}

		values = append(values, resource)
	}
	return values, nil
}

func ElasticBeanstalkApplication(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := elasticbeanstalk.NewFromConfig(cfg)
	out, err := client.DescribeApplications(ctx, &elasticbeanstalk.DescribeApplicationsInput{})
	if err != nil {
		if !isErr(err, "ResourceNotFoundException") && !isErr(err, "InsufficientPrivilegesException") && !strings.Contains(err.Error(), "Access Denied") {
			return nil, err
		}
		return nil, nil
	}

	var values []models.Resource
	for _, item := range out.Applications {
		resource, err := elasticBeanstalkApplicationHandle(ctx, cfg, item)
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

	return values, nil
}
func elasticBeanstalkApplicationHandle(ctx context.Context, cfg aws.Config, item types.ApplicationDescription) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)

	client := elasticbeanstalk.NewFromConfig(cfg)

	tags, err := client.ListTagsForResource(ctx, &elasticbeanstalk.ListTagsForResourceInput{
		ResourceArn: item.ApplicationArn,
	})
	if err != nil {
		if !isErr(err, "ResourceNotFoundException") && !isErr(err, "InsufficientPrivilegesException") {
			return models.Resource{}, err
		}
		tags = &elasticbeanstalk.ListTagsForResourceOutput{}
	}

	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     *item.ApplicationArn,
		Name:    *item.ApplicationName,
		Account: describeCtx.AccountID,
		Description: model.ElasticBeanstalkApplicationDescription{
			Application: item,
			Tags:        tags.ResourceTags,
		},
	}

	return resource, nil
}
func GetElasticBeanstalkApplication(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	applicationName := fields["name"]

	client := elasticbeanstalk.NewFromConfig(cfg)

	out, err := client.DescribeApplications(ctx, &elasticbeanstalk.DescribeApplicationsInput{
		ApplicationNames: []string{applicationName},
	})
	if err != nil {
		if isErr(err, "DescribeApplicationsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, v := range out.Applications {
		resource, err := elasticBeanstalkApplicationHandle(ctx, cfg, v)
		if err != nil {
			return nil, err
		}

		values = append(values, resource)
	}
	return values, nil
}

func ElasticBeanstalkPlatform(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := elasticbeanstalk.NewFromConfig(cfg)
	paginator := elasticbeanstalk.NewListPlatformVersionsPaginator(client, &elasticbeanstalk.ListPlatformVersionsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, item := range page.PlatformSummaryList {
			platform, err := client.DescribePlatformVersion(ctx, &elasticbeanstalk.DescribePlatformVersionInput{
				PlatformArn: item.PlatformArn,
			})
			if err != nil {
				return nil, err
			}

			resource := models.Resource{
				Region:  describeCtx.OGRegion,
				ARN:     *platform.PlatformDescription.PlatformArn,
				Name:    *platform.PlatformDescription.PlatformName,
				Account: describeCtx.AccountID,
				Description: model.ElasticBeanstalkPlatformDescription{
					Platform: *platform.PlatformDescription,
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

func ElasticBeanstalkApplicationVersion(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := elasticbeanstalk.NewFromConfig(cfg)

	pagesLeft := true
	out, err := client.DescribeApplications(ctx, &elasticbeanstalk.DescribeApplicationsInput{})
	if err != nil {
		if !isErr(err, "ResourceNotFoundException") && !isErr(err, "InsufficientPrivilegesException") && !strings.Contains(err.Error(), "Access Denied") {
			return nil, err
		}
		return nil, nil
	}

	var values []models.Resource
	for _, item := range out.Applications {
		params := &elasticbeanstalk.DescribeApplicationVersionsInput{
			ApplicationName: item.ApplicationName,
		}

		if pagesLeft {
			applicationVersions, err := client.DescribeApplicationVersions(ctx, params)
			if err != nil {
				return nil, err
			}

			for _, v := range applicationVersions.ApplicationVersions {
				resource, err := elasticBeanstalkApplicationVersionHandle(ctx, cfg, v)
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
			if applicationVersions.NextToken != nil {
				pagesLeft = true
				params.NextToken = applicationVersions.NextToken
			} else {
				pagesLeft = false
			}
		}
	}
	return values, nil
}

func elasticBeanstalkApplicationVersionHandle(ctx context.Context, cfg aws.Config, v types.ApplicationVersionDescription) (models.Resource, error) {
	client := elasticbeanstalk.NewFromConfig(cfg)
	describeCtx := model.GetDescribeContext(ctx)

	tags, err := client.ListTagsForResource(ctx, &elasticbeanstalk.ListTagsForResourceInput{
		ResourceArn: v.ApplicationVersionArn,
	})
	if err != nil {
		if !isErr(err, "ResourceNotFoundException") && !isErr(err, "InsufficientPrivilegesException") {
			return models.Resource{}, err
		}
		tags = &elasticbeanstalk.ListTagsForResourceOutput{}
	}

	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     *v.ApplicationVersionArn,
		Name:    *v.ApplicationName,
		Account: describeCtx.AccountID,
		Description: model.ElasticBeanstalkApplicationVersionDescription{
			ApplicationVersion: v,
			Tags:               tags.ResourceTags,
		},
	}
	return resource, nil
}
