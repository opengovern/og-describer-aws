package describer

import (
	"context"
	"github.com/opengovern/og-describer-aws/pkg/sdk/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/opengovern/og-describer-aws/provider/model"
)

func SQSQueue(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := sqs.NewFromConfig(cfg)
	paginator := sqs.NewListQueuesPaginator(client, &sqs.ListQueuesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "AWS.SimpleQueueService.NonExistentQueue") {
				return nil, nil
			}
			return nil, err
		}

		for _, url := range page.QueueUrls {
			// url example: http://sqs.us-west-2.amazonaws.com/123456789012/queueName
			// This prevents Implicit memory aliasing in for loop
			url := url

			output, err := client.GetQueueAttributes(ctx, &sqs.GetQueueAttributesInput{
				QueueUrl: &url,
				AttributeNames: []types.QueueAttributeName{
					types.QueueAttributeNameAll,
				},
			})
			if err != nil {
				if isErr(err, "AWS.SimpleQueueService.NonExistentQueue") {
					continue
				}
				return nil, err
			}

			tOutput, err := client.ListQueueTags(ctx, &sqs.ListQueueTagsInput{
				QueueUrl: &url,
			})
			if err != nil {
				if isErr(err, "AWS.SimpleQueueService.NonExistentQueue") {
					tOutput = &sqs.ListQueueTagsOutput{}
				} else {
					return nil, err
				}
			}

			// Add Queue URL since it doesn't exists in the description
			output.Attributes["QueueUrl"] = url

			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    url,
				Name:   nameFromArn(url),
				Description: model.SQSQueueDescription{
					Attributes: output.Attributes,
					Tags:       tOutput.Tags,
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
func sQSQueueHandle(ctx context.Context, url string, queueAttributes *sqs.GetQueueAttributesOutput, tagOutput *sqs.ListQueueTagsOutput) models.Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    url,
		Name:   nameFromArn(url),
		Description: model.SQSQueueDescription{
			Attributes: queueAttributes.Attributes,
			Tags:       tagOutput.Tags,
		},
	}
	return resource
}
func GetSQSQueue(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	QueueName := fields["name"]
	var values []models.Resource
	client := sqs.NewFromConfig(cfg)
	url, err := client.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{
		QueueName: &QueueName,
	})
	if err != nil {
		if isErr(err, "AWS.SimpleQueueService.NonExistentQueue") {
			return nil, nil
		}
		return nil, err
	}

	queueAttributes, err := client.GetQueueAttributes(ctx, &sqs.GetQueueAttributesInput{
		QueueUrl: url.QueueUrl,
	})
	if err != nil {
		if isErr(err, "AWS.SimpleQueueService.NonExistentQueue") {
			return nil, nil
		}
		return nil, err
	}

	tOutput, err := client.ListQueueTags(ctx, &sqs.ListQueueTagsInput{
		QueueUrl: url.QueueUrl,
	})
	if err != nil {
		if isErr(err, "AWS.SimpleQueueService.NonExistentQueue") {
			return nil, nil
		} else {
			return nil, err
		}
	}

	resource := sQSQueueHandle(ctx, *url.QueueUrl, queueAttributes, tOutput)
	values = append(values, resource)
	return values, nil
}
