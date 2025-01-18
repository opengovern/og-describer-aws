package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func AthenaWrokgroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := athena.NewFromConfig(cfg)
	pager := athena.NewListWorkGroupsPaginator(client, &athena.ListWorkGroupsInput{})
	var values []models.Resource
	for pager.HasMorePages() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, item := range page.WorkGroups {
			resource, err := authenaWorkgroupHandle(ctx, cfg, item)
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

func authenaWorkgroupHandle(ctx context.Context, cfg aws.Config, item types.WorkGroupSummary) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := athena.NewFromConfig(cfg)
	output, err := client.GetWorkGroup(ctx, &athena.GetWorkGroupInput{
		WorkGroup: item.Name,
	})
	if err != nil {
		if isErr(err, "GetWorkGroupNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}
	resource := models.Resource{
		Region: describeCtx.OGRegion,
		Name:   *output.WorkGroup.Name,
		Description: model.AthenaWorkGroupDescription{
			WorkGroup: output.WorkGroup,
		},
	}

	return resource, nil
}

func AthenaQueryExecution(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := athena.NewFromConfig(cfg)
	pager := athena.NewListQueryExecutionsPaginator(client, &athena.ListQueryExecutionsInput{})
	var values []models.Resource
	for pager.HasMorePages() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, item := range page.QueryExecutionIds {
			resource, err := authenaQueryExecutionHandle(ctx, cfg, item)
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

func authenaQueryExecutionHandle(ctx context.Context, cfg aws.Config, id string) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := athena.NewFromConfig(cfg)
	output, err := client.GetQueryExecution(ctx, &athena.GetQueryExecutionInput{
		QueryExecutionId: aws.String(id),
	})
	if err != nil {
		if isErr(err, "GetQueryExecutionNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}
	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ID:     *output.QueryExecution.QueryExecutionId,
		Name:   *output.QueryExecution.Query,
		Description: model.AthenaQueryExecutionDescription{
			QueryExecution: output.QueryExecution,
		},
	}

	return resource, nil
}
