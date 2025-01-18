package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func CodePipelinePipeline(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := codepipeline.NewFromConfig(cfg)
	paginator := codepipeline.NewListPipelinesPaginator(client, &codepipeline.ListPipelinesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if !isErr(err, "PipelineNotFoundException") {
				return nil, err
			}
			continue
		}

		for _, v := range page.Pipelines {
			pipeline, err := client.GetPipeline(ctx, &codepipeline.GetPipelineInput{
				Name: v.Name,
			})
			if err != nil {
				if !isErr(err, "PipelineNotFoundException") {
					return nil, err
				}
				continue
			}

			resource, err := codePipelinePipelineHandle(ctx, cfg, pipeline)
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
func codePipelinePipelineHandle(ctx context.Context, cfg aws.Config, pipeline *codepipeline.GetPipelineOutput) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := codepipeline.NewFromConfig(cfg)
	tags, err := client.ListTagsForResource(ctx, &codepipeline.ListTagsForResourceInput{
		ResourceArn: pipeline.Metadata.PipelineArn,
	})
	if err != nil {
		if !isErr(err, "InvalidParameter") {
			return models.Resource{}, err
		}
		tags = &codepipeline.ListTagsForResourceOutput{}
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *pipeline.Metadata.PipelineArn,
		Name:   *pipeline.Pipeline.Name,
		Description: model.CodePipelinePipelineDescription{
			Pipeline: *pipeline.Pipeline,
			Metadata: *pipeline.Metadata,
			Tags:     tags.Tags,
		},
	}
	return resource, nil
}
func GetCodePipelinePipeline(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	name := fields["name"]
	var values []models.Resource
	client := codepipeline.NewFromConfig(cfg)

	pipeline, err := client.GetPipeline(ctx, &codepipeline.GetPipelineInput{
		Name: &name,
	})
	if err != nil {
		if isErr(err, "GetPipelineNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	resource, err := codePipelinePipelineHandle(ctx, cfg, pipeline)
	if err != nil {
		return nil, err
	}

	values = append(values, resource)
	return values, nil
}
