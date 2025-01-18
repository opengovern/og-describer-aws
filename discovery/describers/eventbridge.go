package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
)

func EventBridgeBus(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := eventbridge.NewFromConfig(cfg)

	input := eventbridge.ListEventBusesInput{Limit: aws.Int32(100)}

	var values []models.Resource
	for {
		response, err := client.ListEventBuses(ctx, &input)
		if err != nil {
			if isErr(err, "InvalidParameter") || isErr(err, "ResourceNotFoundException") || isErr(err, "ValidationException") {
				return nil, nil
			}
			return nil, err
		}
		for _, bus := range response.EventBuses {
			tagsOutput, err := client.ListTagsForResource(ctx, &eventbridge.ListTagsForResourceInput{
				ResourceARN: bus.Arn,
			})
			if err != nil {
				if !isErr(err, "InvalidParameter") && !isErr(err, "ResourceNotFoundException") && !isErr(err, "ValidationException") {
					return nil, err
				}
				tagsOutput = &eventbridge.ListTagsForResourceOutput{}
			}

			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *bus.Arn,
				Name:   *bus.Name,
				Description: model.EventBridgeBusDescription{
					Bus:  bus,
					Tags: tagsOutput.Tags,
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
		if response.NextToken == nil {
			break
		}
		input.NextToken = response.NextToken
	}

	return values, nil
}

func EventBridgeRule(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := eventbridge.NewFromConfig(cfg)

	var values []models.Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		listRulesOutput, err := client.ListRules(ctx, &eventbridge.ListRulesInput{
			NextToken: prevToken,
		})
		if err != nil {
			return nil, err
		}
		for _, listRule := range listRulesOutput.Rules {
			rule, err := client.DescribeRule(ctx, &eventbridge.DescribeRuleInput{
				Name: listRule.Name,
			})
			if err != nil {
				return nil, err
			}

			tagsOutput, err := client.ListTagsForResource(ctx, &eventbridge.ListTagsForResourceInput{
				ResourceARN: rule.Arn,
			})
			if err != nil {
				if !isErr(err, "ResourceNotFoundException") && !isErr(err, "ValidationException") {
					return nil, err
				}
				tagsOutput = &eventbridge.ListTagsForResourceOutput{}
			}

			targets, err := client.ListTargetsByRule(ctx, &eventbridge.ListTargetsByRuleInput{
				Rule: listRule.Name,
			})
			if err != nil {
				if !isErr(err, "ResourceNotFoundException") && !isErr(err, "ValidationException") {
					return nil, err
				}
				targets = &eventbridge.ListTargetsByRuleOutput{}
			}

			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *rule.Arn,
				Name:   *rule.Name,
				Description: model.EventBridgeRuleDescription{
					Rule:    *rule,
					Tags:    tagsOutput.Tags,
					Targets: targets.Targets,
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

		return listRulesOutput.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
