package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/fms/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fms"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func FMSPolicy(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := fms.NewFromConfig(cfg)
	paginator := fms.NewListPoliciesPaginator(client, &fms.ListPoliciesInput{})

	var values []models.Resource
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
			resource := models.Resource{
				Region: describeCtx.OGRegion,
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

func GetFMSPolicy(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	policyID := fields["id"]
	client := fms.NewFromConfig(cfg)

	out, err := client.GetPolicy(ctx, &fms.GetPolicyInput{PolicyId: &policyID})
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	tags, err := client.ListTagsForResource(ctx, &fms.ListTagsForResourceInput{
		ResourceArn: out.PolicyArn,
	})
	if err != nil {
		return nil, err
	}
	values = append(values, models.Resource{
		Region: describeCtx.OGRegion,
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
