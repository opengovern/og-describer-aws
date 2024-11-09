package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/synthetics"
)

func SyntheticsCanary(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := synthetics.NewFromConfig(cfg)
	paginator := synthetics.NewDescribeCanariesPaginator(client, &synthetics.DescribeCanariesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Canaries {
			resource := Resource{
				Region:      describeCtx.OGRegion,
				ID:          *v.Id,
				Name:        *v.Name,
				Description: v,
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
