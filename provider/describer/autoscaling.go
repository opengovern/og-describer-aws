package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/opengovern/og-describer-aws/provider/model"
)

func AutoScalingAutoScalingGroup(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := autoscaling.NewFromConfig(cfg)
	paginator := autoscaling.NewDescribeAutoScalingGroupsPaginator(client, &autoscaling.DescribeAutoScalingGroupsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.AutoScalingGroups {
			var desc model.AutoScalingGroupDescription

			sg := v // Do this to avoid the pointer being replaces by the for loop

			desc.AutoScalingGroup = &sg
			desc.Policies, err = getAutoScalingPolicies(ctx, cfg, v.AutoScalingGroupName)
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region:      describeCtx.OGRegion,
				ARN:         *v.AutoScalingGroupARN,
				Name:        *v.AutoScalingGroupName,
				Description: desc,
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

func GetAutoScalingAutoScalingGroup(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	autoScalingGroupName := fields["name"]
	client := autoscaling.NewFromConfig(cfg)

	out, err := client.DescribeAutoScalingGroups(ctx, &autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: []string{autoScalingGroupName},
	})
	if err != nil {
		return nil, err
	}

	var values []Resource

	for _, v := range out.AutoScalingGroups {
		var desc model.AutoScalingGroupDescription

		sg := v // Do this to avoid the pointer being replaces by the for loop

		desc.AutoScalingGroup = &sg
		desc.Policies, err = getAutoScalingPolicies(ctx, cfg, v.AutoScalingGroupName)
		if err != nil {
			return nil, err
		}

		values = append(values, Resource{
			Region:      describeCtx.OGRegion,
			ARN:         *v.AutoScalingGroupARN,
			Name:        *v.AutoScalingGroupName,
			Description: desc,
		})
	}

	return values, nil
}

func getAutoScalingPolicies(ctx context.Context, cfg aws.Config, asgName *string) ([]types.ScalingPolicy, error) {
	client := autoscaling.NewFromConfig(cfg)
	paginator := autoscaling.NewDescribePoliciesPaginator(client, &autoscaling.DescribePoliciesInput{
		AutoScalingGroupName: asgName,
	})

	var values []types.ScalingPolicy
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		values = append(values, page.ScalingPolicies...)
	}

	return values, nil
}

func AutoScalingLaunchConfiguration(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := autoscaling.NewFromConfig(cfg)
	paginator := autoscaling.NewDescribeLaunchConfigurationsPaginator(client, &autoscaling.DescribeLaunchConfigurationsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.LaunchConfigurations {
			resource := Resource{
				Region: describeCtx.OGRegion,
				ARN:    *v.LaunchConfigurationARN,
				Name:   *v.LaunchConfigurationName,
				Description: model.AutoScalingLaunchConfigurationDescription{
					LaunchConfiguration: v,
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

func GetAutoScalingLaunchConfiguration(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	launchConfigurationName := fields["name"]
	client := autoscaling.NewFromConfig(cfg)
	out, err := client.DescribeLaunchConfigurations(ctx, &autoscaling.DescribeLaunchConfigurationsInput{
		LaunchConfigurationNames: []string{launchConfigurationName},
	})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range out.LaunchConfigurations {
		values = append(values, Resource{
			Region: describeCtx.OGRegion,
			ARN:    *v.LaunchConfigurationARN,
			Name:   *v.LaunchConfigurationName,
			Description: model.AutoScalingLaunchConfigurationDescription{
				LaunchConfiguration: v,
			},
		})
	}

	return values, nil
}

func AutoScalingLifecycleHook(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	groups, err := AutoScalingAutoScalingGroup(ctx, cfg, nil)
	if groups != nil {
		return nil, err
	}

	client := autoscaling.NewFromConfig(cfg)

	var values []Resource
	for _, g := range groups {
		group := g.Description.(types.AutoScalingGroup)
		output, err := client.DescribeLifecycleHooks(ctx, &autoscaling.DescribeLifecycleHooksInput{
			AutoScalingGroupName: group.AutoScalingGroupName,
		})
		if err != nil {
			return nil, err
		}

		for _, v := range output.LifecycleHooks {
			resource := Resource{
				Region:      describeCtx.OGRegion,
				ID:          CompositeID(*v.AutoScalingGroupName, *v.LifecycleHookName),
				Name:        *v.AutoScalingGroupName,
				Description: v,
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

func AutoScalingScalingPolicy(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := autoscaling.NewFromConfig(cfg)
	paginator := autoscaling.NewDescribePoliciesPaginator(client, &autoscaling.DescribePoliciesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ScalingPolicies {
			resource := Resource{
				Region:      describeCtx.OGRegion,
				ARN:         *v.PolicyARN,
				Name:        *v.PolicyName,
				Description: v,
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

func AutoScalingScheduledAction(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := autoscaling.NewFromConfig(cfg)
	paginator := autoscaling.NewDescribeScheduledActionsPaginator(client, &autoscaling.DescribeScheduledActionsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ScheduledUpdateGroupActions {
			resource := Resource{
				Region:      describeCtx.OGRegion,
				ARN:         *v.ScheduledActionARN,
				Name:        *v.ScheduledActionName,
				Description: v,
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

func AutoScalingWarmPool(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	groups, err := AutoScalingAutoScalingGroup(ctx, cfg, nil)
	if groups != nil {
		return nil, err
	}

	client := autoscaling.NewFromConfig(cfg)

	var values []Resource
	for _, g := range groups {
		group := g.Description.(types.AutoScalingGroup)

		err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
			output, err := client.DescribeWarmPool(ctx, &autoscaling.DescribeWarmPoolInput{
				AutoScalingGroupName: group.AutoScalingGroupName,
				NextToken:            prevToken,
			})
			if err != nil {
				return nil, err
			}

			for _, v := range output.Instances {
				resource := Resource{
					Region:      describeCtx.OGRegion,
					ID:          CompositeID(*group.AutoScalingGroupName, *v.InstanceId), // TODO
					Name:        *v.LaunchConfigurationName,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
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

	}

	return values, nil
}
