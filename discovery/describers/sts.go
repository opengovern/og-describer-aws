package describers

import (
	"context"

	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func STSAccount(ctx context.Context, cfg aws.Config) (string, error) {
	svc := sts.NewFromConfig(cfg)

	acc, err := svc.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return "", err
	}

	return *acc.Account, nil
}

func STSCallerIdentity(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := sts.NewFromConfig(cfg)
	var values []models.Resource
	ci, err := client.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return nil, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ID:     *ci.UserId,
		ARN:    *ci.Arn,
		Description: model.STSCallerIdentityDescription{
			UsrId:   *ci.UserId,
			Account: *ci.Account,
			Arn:     *ci.Arn,
		},
	}

	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	} else {
		values = append(values, resource)
	}

	return values, nil
}
