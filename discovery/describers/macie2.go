package describer

import (
	"context"
	"github.com/opengovern/og-describer-aws/pkg/sdk/models"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/macie2"
	"github.com/opengovern/og-describer-aws/provider/model"
)

func Macie2ClassificationJob(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := macie2.NewFromConfig(cfg)

	var values []models.Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		classificationJobs, err := client.ListClassificationJobs(ctx, &macie2.ListClassificationJobsInput{
			NextToken: prevToken,
		})
		if err != nil {
			if isErr(err, "AccessDeniedException") || strings.Contains(err.Error(), "AccessDeniedException") {
				return
			} else {
				return nil, err
			}
		}

		for _, jobSummary := range classificationJobs.Items {
			resource, err := macie2ClassificationJobHandle(ctx, cfg, *jobSummary.JobId)
			if err != nil {
				if isErr(err, "AccessDeniedException") || strings.Contains(err.Error(), "AccessDeniedException") {
					return nil, nil
				} else {
					return nil, err
				}
			}
			emptyResource := models.Resource{}
			if err == nil && resource == emptyResource {
				continue
			}

			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}

		return classificationJobs.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
func macie2ClassificationJobHandle(ctx context.Context, cfg aws.Config, jobId string) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := macie2.NewFromConfig(cfg)

	job, err := client.DescribeClassificationJob(ctx, &macie2.DescribeClassificationJobInput{
		JobId: &jobId,
	})
	if err != nil {
		if isErr(err, "DescribeClassificationJobNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *job.JobArn,
		Name:   *job.Name,
		ID:     *job.JobId,
		Description: model.Macie2ClassificationJobDescription{
			ClassificationJob: *job,
		},
	}
	return resource, nil
}
func GetMacie2ClassificationJob(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	jobId := fields["jobId"]
	var values []models.Resource

	resource, err := macie2ClassificationJobHandle(ctx, cfg, jobId)
	if err != nil {
		return nil, err
	}
	emptyResource := models.Resource{}
	if err == nil && resource == emptyResource {
		return nil, nil
	}

	values = append(values, resource)
	return values, nil
}
