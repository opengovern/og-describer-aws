package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func FirehoseDeliveryStream(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := firehose.NewFromConfig(cfg)

	var values []models.Resource

	err := PaginateRetrieveAll(func(prevToken *string) (lastName *string, err error) {
		deliveryStreams, err := client.ListDeliveryStreams(ctx, &firehose.ListDeliveryStreamsInput{
			ExclusiveStartDeliveryStreamName: prevToken,
		})
		if err != nil {
			return nil, err
		}
		for _, deliveryStreamName := range deliveryStreams.DeliveryStreamNames {
			lastName = &deliveryStreamName

			resource, err := FirehoseDeliveryStreamHandle(ctx, cfg, *lastName)
			if err != nil {
				return nil, err
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
		if deliveryStreams.HasMoreDeliveryStreams == nil || !*deliveryStreams.HasMoreDeliveryStreams {
			return nil, nil
		}
		return lastName, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
func FirehoseDeliveryStreamHandle(ctx context.Context, cfg aws.Config, deliveryStreamName string) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := firehose.NewFromConfig(cfg)

	deliveryStream, err := client.DescribeDeliveryStream(ctx, &firehose.DescribeDeliveryStreamInput{
		DeliveryStreamName: &deliveryStreamName,
	})
	if err != nil {
		if isErr(err, "DescribeDeliveryStreamNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	tags, err := client.ListTagsForDeliveryStream(ctx, &firehose.ListTagsForDeliveryStreamInput{
		DeliveryStreamName: &deliveryStreamName,
	})
	if err != nil {
		if isErr(err, "ListTagsForDeliveryStreamNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *deliveryStream.DeliveryStreamDescription.DeliveryStreamARN,
		Name:   *deliveryStream.DeliveryStreamDescription.DeliveryStreamName,
		Description: model.FirehoseDeliveryStreamDescription{
			DeliveryStream: *deliveryStream.DeliveryStreamDescription,
			Tags:           tags.Tags,
		},
	}
	return resource, nil
}
func GetFirehoseDeliveryStream(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	deliveryStreamName := fields["name"]
	var values []models.Resource

	resource, err := FirehoseDeliveryStreamHandle(ctx, cfg, deliveryStreamName)
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
