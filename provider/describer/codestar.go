package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codestar"
	"github.com/opengovern/og-describer-aws/provider/model"
)

func CodeStarProject(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := codestar.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		projects, err := client.ListProjects(ctx, &codestar.ListProjectsInput{
			MaxResults: aws.Int32(100),
			NextToken:  prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, projectSum := range projects.Projects {
			project, err := client.DescribeProject(ctx, &codestar.DescribeProjectInput{
				Id: projectSum.ProjectId,
			})
			if err != nil {
				return nil, err
			}

			tags, err := client.ListTagsForProject(ctx, &codestar.ListTagsForProjectInput{
				Id: projectSum.ProjectId,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *project.Arn,
				Name:   *project.Id,
				Description: model.CodeStarProjectDescription{
					Project: *project,
					Tags:    tags.Tags,
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

		return projects.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
