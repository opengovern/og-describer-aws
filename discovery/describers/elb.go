package describers

import (
	"context"
	"fmt"

	"strings"
	"time"

	types3 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func ElasticLoadBalancingV2LoadBalancer(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := elasticloadbalancingv2.NewFromConfig(cfg)
	paginator := elasticloadbalancingv2.NewDescribeLoadBalancersPaginator(client, &elasticloadbalancingv2.DescribeLoadBalancersInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.LoadBalancers {
			resource, err := elasticLoadBalancingV2LoadBalancerHandle(ctx, cfg, v)
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
func elasticLoadBalancingV2LoadBalancerHandle(ctx context.Context, cfg aws.Config, v types.LoadBalancer) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := elasticloadbalancingv2.NewFromConfig(cfg)
	attrs, err := client.DescribeLoadBalancerAttributes(ctx, &elasticloadbalancingv2.DescribeLoadBalancerAttributesInput{
		LoadBalancerArn: v.LoadBalancerArn,
	})
	if err != nil {
		return models.Resource{}, err
	}

	tags, err := client.DescribeTags(ctx, &elasticloadbalancingv2.DescribeTagsInput{
		ResourceArns: []string{*v.LoadBalancerArn},
	})
	if err != nil {
		return models.Resource{}, err
	}

	description := model.ElasticLoadBalancingV2LoadBalancerDescription{
		LoadBalancer: v,
		Attributes:   attrs.Attributes,
	}

	if tags.TagDescriptions != nil && len(tags.TagDescriptions) > 0 {
		description.Tags = tags.TagDescriptions[0].Tags
	}

	resource := models.Resource{
		Region:      describeCtx.OGRegion,
		ARN:         *v.LoadBalancerArn,
		Account:     describeCtx.AccountID,
		Name:        *v.LoadBalancerName,
		Description: description,
	}
	return resource, nil
}
func GetElasticLoadBalancingV2LoadBalancer(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	lbARN := fields["arn"]
	client := elasticloadbalancingv2.NewFromConfig(cfg)
	out, err := client.DescribeLoadBalancers(ctx, &elasticloadbalancingv2.DescribeLoadBalancersInput{
		LoadBalancerArns: []string{lbARN},
	})
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, v := range out.LoadBalancers {
		resource, err := elasticLoadBalancingV2LoadBalancerHandle(ctx, cfg, v)
		if err != nil {
			return nil, err
		}

		values = append(values, resource)
	}
	return values, nil
}

func ElasticLoadBalancingV2Listener(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	lbs, err := ElasticLoadBalancingV2LoadBalancer(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := elasticloadbalancingv2.NewFromConfig(cfg)

	var values []models.Resource
	for _, lb := range lbs {
		arn := lb.Description.(model.ElasticLoadBalancingV2LoadBalancerDescription).LoadBalancer.LoadBalancerArn
		paginator := elasticloadbalancingv2.NewDescribeListenersPaginator(client, &elasticloadbalancingv2.DescribeListenersInput{
			LoadBalancerArn: arn,
		})
		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}

			for _, v := range page.Listeners {
				resource := elasticLoadBalancingV2ListenerHandle(ctx, v)

				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}

	}

	return values, nil
}
func elasticLoadBalancingV2ListenerHandle(ctx context.Context, v types.Listener) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     *v.ListenerArn,
		Name:    nameFromArn(*v.ListenerArn),
		Account: describeCtx.AccountID,
		Description: model.ElasticLoadBalancingV2ListenerDescription{
			Listener: v,
		},
	}
	return resource
}
func GetElasticLoadBalancingV2Listener(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	lbArn := fields["load_balancer_arn"]
	listenerARN := fields["arn"]
	client := elasticloadbalancingv2.NewFromConfig(cfg)
	out, err := client.DescribeListeners(ctx, &elasticloadbalancingv2.DescribeListenersInput{
		ListenerArns:    []string{listenerARN},
		LoadBalancerArn: &lbArn,
	})
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, v := range out.Listeners {
		resource := elasticLoadBalancingV2ListenerHandle(ctx, v)
		values = append(values, resource)
	}
	return values, nil
}

func ElasticLoadBalancingV2ListenerRule(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	listeners, err := ElasticLoadBalancingV2Listener(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := elasticloadbalancingv2.NewFromConfig(cfg)
	var values []models.Resource
	for _, l := range listeners {
		arn := l.Description.(model.ElasticLoadBalancingV2ListenerDescription).Listener.ListenerArn
		err = PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
			output, err := client.DescribeRules(ctx, &elasticloadbalancingv2.DescribeRulesInput{
				ListenerArn: aws.String(*arn),
				Marker:      prevToken,
			})
			if err != nil {
				return nil, err
			}

			for _, v := range output.Rules {
				resource := elasticLoadBalancingV2ListenerRuleHandle(ctx, v)
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
			return output.NextMarker, nil
		})
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}
func elasticLoadBalancingV2ListenerRuleHandle(ctx context.Context, v types.Rule) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     *v.RuleArn,
		Name:    *v.RuleArn,
		Account: describeCtx.AccountID,
		Description: model.ElasticLoadBalancingV2RuleDescription{
			Rule: v,
		},
	}
	return resource
}
func GetElasticLoadBalancingV2ListenerRule(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	arn := fields["arn"]
	var values []models.Resource
	client := elasticloadbalancingv2.NewFromConfig(cfg)

	listeners, err := ElasticLoadBalancingV2Listener(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	for _, l := range listeners {
		if l.ARN != arn {
			continue
		}
		output, err := client.DescribeRules(ctx, &elasticloadbalancingv2.DescribeRulesInput{
			ListenerArn: aws.String(arn),
		})
		if err != nil {
			return nil, err
		}

		for _, v := range output.Rules {
			resource := elasticLoadBalancingV2ListenerRuleHandle(ctx, v)
			values = append(values, resource)
		}
	}
	return values, nil
}

func ElasticLoadBalancingLoadBalancer(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := elasticloadbalancing.NewFromConfig(cfg)
	paginator := elasticloadbalancing.NewDescribeLoadBalancersPaginator(client, &elasticloadbalancing.DescribeLoadBalancersInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.LoadBalancerDescriptions {
			resource, err := elasticLoadBalancingLoadBalancerHandle(ctx, cfg, v)
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
func elasticLoadBalancingLoadBalancerHandle(ctx context.Context, cfg aws.Config, v types3.LoadBalancerDescription) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := elasticloadbalancing.NewFromConfig(cfg)

	attrs, err := client.DescribeLoadBalancerAttributes(ctx, &elasticloadbalancing.DescribeLoadBalancerAttributesInput{
		LoadBalancerName: v.LoadBalancerName,
	})
	if err != nil {
		if isErr(err, "DescribeLoadBalancerAttributesNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	tags, err := client.DescribeTags(ctx, &elasticloadbalancing.DescribeTagsInput{
		LoadBalancerNames: []string{*v.LoadBalancerName},
	})
	if err != nil {
		if isErr(err, "DescribeTagsNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	description := model.ElasticLoadBalancingLoadBalancerDescription{
		LoadBalancer: v,
		Attributes:   attrs.LoadBalancerAttributes,
	}

	if tags.TagDescriptions != nil && len(tags.TagDescriptions) > 0 {
		description.Tags = tags.TagDescriptions[0].Tags
	}

	arn := "arn:" + describeCtx.Partition + ":elasticloadbalancing:" + describeCtx.Region + ":" + describeCtx.AccountID + ":loadbalancer/" + *v.LoadBalancerName
	resource := models.Resource{
		Region:      describeCtx.OGRegion,
		Account:     describeCtx.AccountID,
		ARN:         arn,
		Name:        *v.LoadBalancerName,
		Description: description,
	}
	return resource, nil
}
func GetElasticLoadBalancingLoadBalancer(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	LoadBalancerNames := fields["loadBalancerName"]

	client := elasticloadbalancing.NewFromConfig(cfg)
	out, err := client.DescribeLoadBalancers(ctx, &elasticloadbalancing.DescribeLoadBalancersInput{
		LoadBalancerNames: []string{LoadBalancerNames},
	})
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, loadBalancer := range out.LoadBalancerDescriptions {
		resource, err := elasticLoadBalancingLoadBalancerHandle(ctx, cfg, loadBalancer)
		if err != nil {
			return nil, err
		}
		values = append(values, resource)
	}
	return values, nil
}

func ElasticLoadBalancingV2SslPolicy(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)

	client := elasticloadbalancingv2.NewFromConfig(cfg)

	var values []models.Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.DescribeSSLPolicies(ctx, &elasticloadbalancingv2.DescribeSSLPoliciesInput{
			Marker: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, v := range output.SslPolicies {
			arn := fmt.Sprintf("arn:%s:elbv2:%s:%s:ssl-policy/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, *v.Name)
			resource := models.Resource{
				Region:  describeCtx.OGRegion,
				Name:    *v.Name,
				Account: describeCtx.AccountID,
				ARN:     arn,
				Description: model.ElasticLoadBalancingV2SslPolicyDescription{
					SslPolicy: v,
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

		return output.NextMarker, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func ElasticLoadBalancingV2TargetGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := elasticloadbalancingv2.NewFromConfig(cfg)
	paginator := elasticloadbalancingv2.NewDescribeTargetGroupsPaginator(client, &elasticloadbalancingv2.DescribeTargetGroupsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.TargetGroups {

			resource, err := elasticLoadBalancingV2TargetGroupHandle(ctx, cfg, v)
			if err != nil {
				return nil, err
			}
			emptyResource := models.Resource{}
			if err == nil && resource == emptyResource {
				return nil, nil
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
func elasticLoadBalancingV2TargetGroupHandle(ctx context.Context, cfg aws.Config, v types.TargetGroup) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := elasticloadbalancingv2.NewFromConfig(cfg)
	healthDescriptions, err := client.DescribeTargetHealth(ctx, &elasticloadbalancingv2.DescribeTargetHealthInput{
		TargetGroupArn: v.TargetGroupArn,
	})
	if err != nil {
		if isErr(err, "DescribeTargetHealthNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	tags, err := client.DescribeTags(ctx, &elasticloadbalancingv2.DescribeTagsInput{
		ResourceArns: []string{*v.TargetGroupArn},
	})
	if err != nil {
		if isErr(err, "DescribeTagsNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	var tagsA []types.Tag
	if tags.TagDescriptions != nil && len(tags.TagDescriptions) > 0 {
		tagsA = tags.TagDescriptions[0].Tags
	}

	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     *v.TargetGroupArn,
		Name:    *v.TargetGroupName,
		Account: describeCtx.AccountID,
		Description: model.ElasticLoadBalancingV2TargetGroupDescription{
			TargetGroup: v,
			Health:      healthDescriptions.TargetHealthDescriptions,
			Tags:        tagsA,
		},
	}
	return resource, nil
}
func GetElasticLoadBalancingV2TargetGroup(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	targetGroupArn := fields["targetGroupArn"]
	client := elasticloadbalancingv2.NewFromConfig(cfg)
	resource, err := client.DescribeTargetGroups(ctx, &elasticloadbalancingv2.DescribeTargetGroupsInput{
		TargetGroupArns: []string{targetGroupArn},
	})
	if err != nil {
		if isErr(err, "DescribeTargetGroupsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, v := range resource.TargetGroups {
		resource, err := elasticLoadBalancingV2TargetGroupHandle(ctx, cfg, v)
		if err != nil {
			return nil, err
		}
		emptyResource := models.Resource{}
		if err == nil && resource == emptyResource {
			return nil, nil
		}

		values = append(values, resource)
	}
	return values, nil
}

func ApplicationLoadBalancerMetricRequestCount(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := elasticloadbalancingv2.NewFromConfig(cfg)
	paginator := elasticloadbalancingv2.NewDescribeLoadBalancersPaginator(client, &elasticloadbalancingv2.DescribeLoadBalancersInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, loadBalancer := range page.LoadBalancers {
			if loadBalancer.Type != types.LoadBalancerTypeEnumApplication {
				continue
			}
			arn := strings.SplitN(*loadBalancer.LoadBalancerArn, "/", 2)[1]
			metrics, err := listCloudWatchMetricStatistics(ctx, cfg, "5_MIN", "AWS/ApplicationELB", "RequestCount", "LoadBalancer", arn)
			if err != nil {
				return nil, err
			}
			for _, metric := range metrics {
				resource := models.Resource{
					Region:  describeCtx.OGRegion,
					Account: describeCtx.AccountID,
					ID:      fmt.Sprintf("%s:%s:%s:%s", arn, metric.Timestamp.Format(time.RFC3339), *metric.DimensionName, *metric.DimensionValue),
					Description: model.ApplicationLoadBalancerMetricRequestCountDescription{
						CloudWatchMetricRow: metric,
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
	}

	return values, nil
}

func GetApplicationLoadBalancerMetricRequestCount(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	loadBalancerARN := fields["arn"]
	client := elasticloadbalancingv2.NewFromConfig(cfg)
	out, err := client.DescribeLoadBalancers(ctx, &elasticloadbalancingv2.DescribeLoadBalancersInput{
		LoadBalancerArns: []string{loadBalancerARN},
	})
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, loadBalancer := range out.LoadBalancers {
		if loadBalancer.Type != types.LoadBalancerTypeEnumApplication {
			continue
		}
		arn := strings.SplitN(*loadBalancer.LoadBalancerArn, "/", 2)[1]
		metrics, err := listCloudWatchMetricStatistics(ctx, cfg, "5_MIN", "AWS/ApplicationELB", "RequestCount", "LoadBalancer", arn)
		if err != nil {
			return nil, err
		}
		for _, metric := range metrics {
			values = append(values, models.Resource{
				Region:  describeCtx.OGRegion,
				Account: describeCtx.AccountID,
				ID:      fmt.Sprintf("%s:%s:%s:%s", arn, metric.Timestamp.Format(time.RFC3339), *metric.DimensionName, *metric.DimensionValue),
				Description: model.ApplicationLoadBalancerMetricRequestCountDescription{
					CloudWatchMetricRow: metric,
				},
			})
		}
	}

	return values, nil
}

func ApplicationLoadBalancerMetricRequestCountDaily(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := elasticloadbalancingv2.NewFromConfig(cfg)
	paginator := elasticloadbalancingv2.NewDescribeLoadBalancersPaginator(client, &elasticloadbalancingv2.DescribeLoadBalancersInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, loadBalancer := range page.LoadBalancers {
			if loadBalancer.Type != types.LoadBalancerTypeEnumApplication {
				continue
			}
			arn := strings.SplitN(*loadBalancer.LoadBalancerArn, "/", 2)[1]
			metrics, err := listCloudWatchMetricStatistics(ctx, cfg, "DAILY", "AWS/ApplicationELB", "RequestCount", "LoadBalancer", arn)
			if err != nil {
				return nil, err
			}
			for _, metric := range metrics {
				resource := models.Resource{
					Region:  describeCtx.OGRegion,
					Account: describeCtx.AccountID,
					ID:      fmt.Sprintf("%s:%s:%s:%s", arn, metric.Timestamp.Format(time.RFC3339), *metric.DimensionName, *metric.DimensionValue),
					Description: model.ApplicationLoadBalancerMetricRequestCountDailyDescription{
						CloudWatchMetricRow: metric,
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
	}

	return values, nil
}

func GetApplicationLoadBalancerMetricRequestCountDaily(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	loadBalancerARN := fields["arn"]

	client := elasticloadbalancingv2.NewFromConfig(cfg)
	out, err := client.DescribeLoadBalancers(ctx, &elasticloadbalancingv2.DescribeLoadBalancersInput{LoadBalancerArns: []string{loadBalancerARN}})
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, loadBalancer := range out.LoadBalancers {
		if loadBalancer.Type != types.LoadBalancerTypeEnumApplication {
			continue
		}
		arn := strings.SplitN(*loadBalancer.LoadBalancerArn, "/", 2)[1]
		metrics, err := listCloudWatchMetricStatistics(ctx, cfg, "DAILY", "AWS/ApplicationELB", "RequestCount", "LoadBalancer", arn)
		if err != nil {
			return nil, err
		}
		for _, metric := range metrics {
			values = append(values, models.Resource{
				Account: describeCtx.AccountID,
				Region:  describeCtx.OGRegion,
				ID:      fmt.Sprintf("%s:%s:%s:%s", arn, metric.Timestamp.Format(time.RFC3339), *metric.DimensionName, *metric.DimensionValue),
				Description: model.ApplicationLoadBalancerMetricRequestCountDailyDescription{
					CloudWatchMetricRow: metric,
				},
			})
		}
	}

	return values, nil
}

func NetworkLoadBalancerMetricNetFlowCount(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := elasticloadbalancingv2.NewFromConfig(cfg)
	paginator := elasticloadbalancingv2.NewDescribeLoadBalancersPaginator(client, &elasticloadbalancingv2.DescribeLoadBalancersInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, loadBalancer := range page.LoadBalancers {
			if loadBalancer.Type != types.LoadBalancerTypeEnumNetwork {
				continue
			}
			arn := strings.SplitN(*loadBalancer.LoadBalancerArn, "/", 2)[1]
			metrics, err := listCloudWatchMetricStatistics(ctx, cfg, "5_MIN", "AWS/NetworkELB", "NewFlowCount", "LoadBalancer", arn)
			if err != nil {
				return nil, err
			}
			for _, metric := range metrics {
				resource := models.Resource{
					Account: describeCtx.AccountID,
					Region:  describeCtx.OGRegion,
					ID:      fmt.Sprintf("%s:%s:%s:%s", arn, metric.Timestamp.Format(time.RFC3339), *metric.DimensionName, *metric.DimensionValue),
					Description: model.NetworkLoadBalancerMetricNetFlowCountDescription{
						CloudWatchMetricRow: metric,
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
	}

	return values, nil
}

func NetworkLoadBalancerMetricNetFlowCountDaily(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := elasticloadbalancingv2.NewFromConfig(cfg)
	paginator := elasticloadbalancingv2.NewDescribeLoadBalancersPaginator(client, &elasticloadbalancingv2.DescribeLoadBalancersInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, loadBalancer := range page.LoadBalancers {
			if loadBalancer.Type != types.LoadBalancerTypeEnumNetwork {
				continue
			}
			arn := strings.SplitN(*loadBalancer.LoadBalancerArn, "/", 2)[1]
			metrics, err := listCloudWatchMetricStatistics(ctx, cfg, "DAILY", "AWS/NetworkELB", "NewFlowCount", "LoadBalancer", arn)
			if err != nil {
				return nil, err
			}
			for _, metric := range metrics {
				resource := models.Resource{
					Region:  describeCtx.OGRegion,
					Account: describeCtx.AccountID,
					ID:      fmt.Sprintf("%s:%s:%s:%s", arn, metric.Timestamp.Format(time.RFC3339), *metric.DimensionName, *metric.DimensionValue),
					Description: model.NetworkLoadBalancerMetricNetFlowCountDailyDescription{
						CloudWatchMetricRow: metric,
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
	}

	return values, nil
}
