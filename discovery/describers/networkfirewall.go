package describer

import (
	"context"
	"github.com/opengovern/og-describer-aws/pkg/sdk/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/opengovern/og-describer-aws/provider/model"
)

func NetworkFirewallFirewall(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := networkfirewall.NewFromConfig(cfg)
	paginator := networkfirewall.NewListFirewallsPaginator(client, &networkfirewall.ListFirewallsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Firewalls {
			resource, err := NetworkFirewallFirewallHandle(ctx, cfg, v)
			if err != nil {
				return nil, err
			}
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
	}

	return values, nil
}
func NetworkFirewallFirewallHandle(ctx context.Context, cfg aws.Config, v types.FirewallMetadata) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := networkfirewall.NewFromConfig(cfg)
	firewall, err := client.DescribeFirewall(ctx, &networkfirewall.DescribeFirewallInput{
		FirewallName: v.FirewallName,
		FirewallArn:  v.FirewallArn,
	})
	if err != nil {
		return models.Resource{}, err
	}

	firewallLogging, err := client.DescribeLoggingConfiguration(ctx, &networkfirewall.DescribeLoggingConfigurationInput{
		FirewallArn: firewall.Firewall.FirewallArn,
	})
	if err != nil {
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *v.FirewallArn,
		Name:   *v.FirewallName,
		Description: model.NetworkFirewallFirewallDescription{
			Firewall:             *firewall.Firewall,
			FirewallStatus:       *firewall.FirewallStatus,
			LoggingConfiguration: firewallLogging,
		},
	}
	return resource, nil
}
func GetNetworkFirewallFirewall(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	var values []models.Resource
	firewallArn := fields["firewallArn"]
	client := networkfirewall.NewFromConfig(cfg)

	listFirewalls, err := client.ListFirewalls(ctx, &networkfirewall.ListFirewallsInput{})
	if err != nil {
		if isErr(err, "ListFirewallsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, v := range listFirewalls.Firewalls {
		if *v.FirewallArn != firewallArn {
			continue
		}
		resource, err := NetworkFirewallFirewallHandle(ctx, cfg, v)
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

func NetworkFirewallPolicy(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := networkfirewall.NewFromConfig(cfg)
	paginator := networkfirewall.NewListFirewallPoliciesPaginator(client, &networkfirewall.ListFirewallPoliciesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.FirewallPolicies {
			if v.Arn == nil {
				continue
			}

			data, err := client.DescribeFirewallPolicy(ctx, &networkfirewall.DescribeFirewallPolicyInput{
				FirewallPolicyArn:  v.Arn,
				FirewallPolicyName: v.Name,
			})
			if err != nil {
				return nil, err
			}

			var name string
			if v.Name != nil {
				name = *v.Name
			} else {
				name = *v.Arn
			}
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *v.Arn,
				Name:   name,
				Description: model.NetworkFirewallFirewallPolicyDescription{
					FirewallPolicy:         data.FirewallPolicy,
					FirewallPolicyResponse: data.FirewallPolicyResponse,
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

func NetworkFirewallRuleGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := networkfirewall.NewFromConfig(cfg)
	paginator := networkfirewall.NewListRuleGroupsPaginator(client, &networkfirewall.ListRuleGroupsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.RuleGroups {
			if v.Arn == nil {
				continue
			}

			data, err := client.DescribeRuleGroup(ctx, &networkfirewall.DescribeRuleGroupInput{
				RuleGroupArn:  v.Arn,
				RuleGroupName: v.Name,
			})
			if err != nil {
				return nil, err
			}

			var name string
			if v.Name != nil {
				name = *v.Name
			} else {
				name = *v.Arn
			}
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *v.Arn,
				Name:   name,
				Description: model.NetworkFirewallRuleGroupDescription{
					RuleGroup:         data.RuleGroup,
					RuleGroupResponse: data.RuleGroupResponse,
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
