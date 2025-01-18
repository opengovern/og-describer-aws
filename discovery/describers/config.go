package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func ConfigConfigurationRecorder(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := configservice.NewFromConfig(cfg)
	out, err := client.DescribeConfigurationRecorders(ctx, &configservice.DescribeConfigurationRecordersInput{})
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, item := range out.ConfigurationRecorders {
		status, err := client.DescribeConfigurationRecorderStatus(ctx, &configservice.DescribeConfigurationRecorderStatusInput{
			ConfigurationRecorderNames: []string{*item.Name},
		})
		if err != nil {
			return nil, err
		}

		arn := "arn:" + describeCtx.Partition + ":config:" + describeCtx.Region + ":" + describeCtx.AccountID + ":config-recorder" + "/" + *item.Name
		resource := models.Resource{
			Region: describeCtx.OGRegion,
			ARN:    arn,
			Name:   *item.Name,
			Description: model.ConfigConfigurationRecorderDescription{
				ConfigurationRecorder:        item,
				ConfigurationRecordersStatus: status.ConfigurationRecordersStatus[0],
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

	return values, nil
}

func ConfigAggregateAuthorization(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := configservice.NewFromConfig(cfg)
	paginator := configservice.NewDescribeAggregationAuthorizationsPaginator(client, &configservice.DescribeAggregationAuthorizationsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, item := range page.AggregationAuthorizations {
			tags, err := client.ListTagsForResource(ctx, &configservice.ListTagsForResourceInput{
				ResourceArn: item.AggregationAuthorizationArn,
			})
			if err != nil {
				return nil, err
			}

			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *item.AggregationAuthorizationArn,
				ID:     *item.AuthorizedAccountId,
				Description: model.ConfigAggregationAuthorizationDescription{
					AggregationAuthorization: item,
					Tags:                     tags.Tags,
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

func ConfigConformancePack(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := configservice.NewFromConfig(cfg)
	paginator := configservice.NewDescribeConformancePacksPaginator(client, &configservice.DescribeConformancePacksInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, item := range page.ConformancePackDetails {
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *item.ConformancePackArn,
				ID:     *item.ConformancePackId,
				Name:   *item.ConformancePackName,
				Description: model.ConfigConformancePackDescription{
					ConformancePack: item,
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

func ConfigRule(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := configservice.NewFromConfig(cfg)
	paginator := configservice.NewDescribeConfigRulesPaginator(client, &configservice.DescribeConfigRulesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		configRuleList := make([]string, 0, len(page.ConfigRules))
		for _, item := range page.ConfigRules {
			configRuleList = append(configRuleList, *item.ConfigRuleName)
		}
		complianceList, err := client.DescribeComplianceByConfigRule(ctx, &configservice.DescribeComplianceByConfigRuleInput{
			ConfigRuleNames: configRuleList,
		})
		if err != nil {
			return nil, err
		}

		complianceMap := make(map[string]types.ComplianceByConfigRule)
		for _, item := range complianceList.ComplianceByConfigRules {
			complianceMap[*item.ConfigRuleName] = item
		}

		for _, item := range page.ConfigRules {
			tags, err := client.ListTagsForResource(ctx, &configservice.ListTagsForResourceInput{
				ResourceArn: item.ConfigRuleArn,
			})
			if err != nil {
				return nil, err
			}

			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *item.ConfigRuleArn,
				ID:     *item.ConfigRuleId,
				Name:   *item.ConfigRuleName,
				Description: model.ConfigRuleDescription{
					Rule:       item,
					Compliance: complianceMap[*item.ConfigRuleName],
					Tags:       tags.Tags,
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

func ConfigRetentionConfiguration(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := configservice.NewFromConfig(cfg)
	paginator := configservice.NewDescribeRetentionConfigurationsPaginator(client, &configservice.DescribeRetentionConfigurationsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, item := range page.RetentionConfigurations {
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				Name:   *item.Name,
				Description: model.ConfigRetentionConfigurationDescription{
					RetentionConfiguration: item,
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
