package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func CloudFormationStack(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := cloudformation.NewFromConfig(cfg)
	paginator := cloudformation.NewDescribeStacksPaginator(client, &cloudformation.DescribeStacksInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if !isErr(err, "ValidationError") && !isErr(err, "ResourceNotFoundException") {
				return nil, err
			}
			continue
		}

		for _, v := range page.Stacks {
			resource, err := cloudFormationStackHandle(ctx, cfg, v)
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
func cloudFormationStackHandle(ctx context.Context, cfg aws.Config, v types.Stack) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := cloudformation.NewFromConfig(cfg)

	template, err := client.GetTemplate(ctx, &cloudformation.GetTemplateInput{
		StackName: v.StackName,
	})
	if err != nil {
		if !isErr(err, "ValidationError") && !isErr(err, "ResourceNotFoundException") {
			return models.Resource{}, err
		}
		template = &cloudformation.GetTemplateOutput{}
	}

	stackResources, err := client.DescribeStackResources(ctx, &cloudformation.DescribeStackResourcesInput{
		StackName: v.StackName,
	})
	if err != nil {
		if !isErr(err, "ValidationError") && !isErr(err, "ResourceNotFoundException") {
			return models.Resource{}, err
		}
		stackResources = &cloudformation.DescribeStackResourcesOutput{}
	}

	template.TemplateBody = nil

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *v.StackId,
		Name:   *v.StackName,
		Description: model.CloudFormationStackDescription{
			Stack:          v,
			StackTemplate:  *template,
			StackResources: stackResources.StackResources,
		},
	}
	return resource, nil
}
func GetCloudFormationStack(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	stackName := fields["name"]
	client := cloudformation.NewFromConfig(cfg)
	out, err := client.DescribeStacks(ctx, &cloudformation.DescribeStacksInput{
		StackName: &stackName,
	})
	if err != nil {
		if isErr(err, "DescribeStacksNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}
	var values []models.Resource
	for _, stack := range out.Stacks {
		resource, err := cloudFormationStackHandle(ctx, cfg, stack)
		if err != nil {
			return nil, err
		}
		values = append(values, resource)
	}
	return values, nil
}

func CloudFormationStackSet(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := cloudformation.NewFromConfig(cfg)
	paginator := cloudformation.NewListStackSetsPaginator(client, &cloudformation.ListStackSetsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Summaries {
			resource, err := cloudFormationStackSetHandle(ctx, cfg, *v.StackSetName)
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
func cloudFormationStackSetHandle(ctx context.Context, cfg aws.Config, stackName string) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := cloudformation.NewFromConfig(cfg)
	stackSet, err := client.DescribeStackSet(ctx, &cloudformation.DescribeStackSetInput{
		StackSetName: &stackName,
	})
	if err != nil {
		return models.Resource{}, err
	}

	if stackSet.StackSet.TemplateBody != nil && len(*stackSet.StackSet.TemplateBody) > 5000 {
		v := *stackSet.StackSet.TemplateBody
		stackSet.StackSet.TemplateBody = aws.String(v[:5000])
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *stackSet.StackSet.StackSetARN,
		Name:   *stackSet.StackSet.StackSetName,
		Description: model.CloudFormationStackSetDescription{
			StackSet: *stackSet.StackSet,
		},
	}
	return resource, nil
}
func GetCloudFormationStackSet(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	stackName := fields["name"]

	var values []models.Resource
	resource, err := cloudFormationStackSetHandle(ctx, cfg, stackName)
	if err != nil {
		return nil, err
	}

	values = append(values, resource)
	return values, nil
}

func CloudFormationStackResource(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := cloudformation.NewFromConfig(cfg)
	paginator := cloudformation.NewDescribeStacksPaginator(client, &cloudformation.DescribeStacksInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if !isErr(err, "ValidationError") && !isErr(err, "ResourceNotFoundException") {
				return nil, err
			}
			continue
		}
		for _, v := range page.Stacks {
			resourcesPager := cloudformation.NewListStackResourcesPaginator(client, &cloudformation.ListStackResourcesInput{
				StackName: v.StackName,
			})
			for resourcesPager.HasMorePages() {
				resourcesPage, err := resourcesPager.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, r := range resourcesPage.StackResourceSummaries {
					resource, err := cloudFormationStackResourceHandle(ctx, cfg, *v.StackName, *r.LogicalResourceId)
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
		}
	}

	return values, nil
}
func cloudFormationStackResourceHandle(ctx context.Context, cfg aws.Config, stackName string, resourceId string) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := cloudformation.NewFromConfig(cfg)

	stackResource, err := client.DescribeStackResource(ctx, &cloudformation.DescribeStackResourceInput{
		LogicalResourceId: aws.String(resourceId),
		StackName:         aws.String(stackName),
	})
	if err != nil {
		if !isErr(err, "ValidationError") && !isErr(err, "ResourceNotFoundException") {
			return models.Resource{}, err
		}
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ID:     *stackResource.StackResourceDetail.LogicalResourceId,
		Description: model.CloudFormationStackResourceDescription{
			StackResource: *stackResource.StackResourceDetail,
		},
	}
	return resource, nil
}
func GetCloudFormationStackResource(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	stackName := fields["stack_name"]
	LogicalResourceId := fields["logical_resource_id"]

	resource, err := cloudFormationStackResourceHandle(ctx, cfg, stackName, LogicalResourceId)
	if err != nil {
		return nil, err
	}

	return []models.Resource{resource}, nil
}
