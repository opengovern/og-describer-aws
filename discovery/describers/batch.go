package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/batch/types"

	_ "golang.org/x/tools/go/analysis/passes/ctrlflow"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func BatchComputeEnvironment(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := batch.NewFromConfig(cfg)
	paginator := batch.NewDescribeComputeEnvironmentsPaginator(client, &batch.DescribeComputeEnvironmentsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ComputeEnvironments {
			resource := batchComputeEnvironmentHandle(ctx, v)
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
func batchComputeEnvironmentHandle(ctx context.Context, v types.ComputeEnvironmentDetail) models.Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *v.ComputeEnvironmentArn,
		Name:   *v.ComputeEnvironmentName,
		Description: model.BatchComputeEnvironmentDescription{
			ComputeEnvironment: v,
		},
	}
	return resource
}
func GetBatchComputeEnvironment(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	var values []models.Resource
	ComputeEnvironments := fields["computeEnvironment"]
	client := batch.NewFromConfig(cfg)
	deComputeEnv, err := client.DescribeComputeEnvironments(ctx, &batch.DescribeComputeEnvironmentsInput{
		ComputeEnvironments: []string{ComputeEnvironments},
	})
	if err != nil {
		if isErr(err, "DescribeComputeEnvironmentsNotfound") || isErr(err, "invalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, v := range deComputeEnv.ComputeEnvironments {
		resource := batchComputeEnvironmentHandle(ctx, v)
		values = append(values, resource)
	}
	return values, nil
}

func BatchJob(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := batch.NewFromConfig(cfg)
	paginator := batch.NewDescribeJobQueuesPaginator(client, &batch.DescribeJobQueuesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, jq := range page.JobQueues {
			jobsPaginator := batch.NewListJobsPaginator(client, &batch.ListJobsInput{
				JobQueue: jq.JobQueueName,
			})
			for jobsPaginator.HasMorePages() {
				page, err := jobsPaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}

				for _, v := range page.JobSummaryList {
					resource := batchJobHandle(ctx, v)
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
func batchJobHandle(ctx context.Context, v types.JobSummary) models.Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *v.JobArn,
		Name:   *v.JobName,
		Description: model.BatchJobDescription{
			Job: v,
		},
	}
	return resource
}
func GetBatchJob(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	jobQueues := fields["jobQueues"]
	var values []models.Resource

	client := batch.NewFromConfig(cfg)
	jobqs, err := client.DescribeJobQueues(ctx, &batch.DescribeJobQueuesInput{
		JobQueues: []string{jobQueues},
	})
	if err != nil {
		if isErr(err, "DescribeJobQueuesNotfound") || isErr(err, "invalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, jq := range jobqs.JobQueues {

		listjob, err := client.ListJobs(ctx, &batch.ListJobsInput{
			JobQueue: jq.JobQueueName,
		})
		if err != nil {
			if isErr(err, "ListJobsNotfound") || isErr(err, "invalidParameterValue") {
				return nil, nil
			}
			return nil, err
		}

		for _, v := range listjob.JobSummaryList {
			resource := batchJobHandle(ctx, v)
			values = append(values, resource)
		}

	}
	return values, nil
}

func BatchJobQueue(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := batch.NewFromConfig(cfg)
	paginator := batch.NewDescribeJobQueuesPaginator(client, &batch.DescribeJobQueuesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, jq := range page.JobQueues {
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *jq.JobQueueArn,
				Name:   *jq.JobQueueName,
				Description: model.BatchJobQueueDescription{
					JobQueue: jq,
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
