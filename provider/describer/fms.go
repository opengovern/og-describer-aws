package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/fms/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fms"
	"github.com/opengovern/og-describer-aws/provider/model"
)

func FMSPolicy(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := fms.NewFromConfig(cfg)
	paginator := fms.NewListPoliciesPaginator(client, &fms.ListPoliciesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.PolicyList {
			tags, err := client.ListTagsForResource(ctx, &fms.ListTagsForResourceInput{
				ResourceArn: v.PolicyArn,
			})
			if err != nil {
				return nil, err
			}
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.PolicyArn,
				Name:   *v.PolicyName,
				Description: model.FMSPolicyDescription{
					Policy: v,
					Tags:   tags.TagList,
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

func GetFMSPolicy(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	policyID := fields["id"]
	client := fms.NewFromConfig(cfg)

	out, err := client.GetPolicy(ctx, &fms.GetPolicyInput{PolicyId: &policyID})
	if err != nil {
		return nil, err
	}

	var values []Resource
	tags, err := client.ListTagsForResource(ctx, &fms.ListTagsForResourceInput{
		ResourceArn: out.PolicyArn,
	})
	if err != nil {
		return nil, err
	}
	values = append(values, Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    *out.PolicyArn,
		Name:   *out.Policy.PolicyName,
		Description: model.FMSPolicyDescription{
			Policy: types.PolicySummary{
				DeleteUnusedFMManagedResources: out.Policy.DeleteUnusedFMManagedResources,
				PolicyArn:                      out.PolicyArn,
				PolicyId:                       out.Policy.PolicyId,
				PolicyName:                     out.Policy.PolicyName,
				RemediationEnabled:             out.Policy.RemediationEnabled,
				ResourceType:                   out.Policy.ResourceType,
				//SecurityServiceType:            out.Policy.SecurityServiceType, //TODO-Saleh
			},
			Tags: tags.TagList,
		},
	})

	return values, nil
}
