package describers

import (
	"context"
	_ "database/sql/driver"

	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/shield"
)

func ShieldProtectionGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := shield.NewFromConfig(cfg)
	paginator := shield.NewListProtectionGroupsPaginator(client, &shield.ListProtectionGroupsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if !isErr(err, "ResourceNotFoundException") {
				return nil, err
			}
			continue
		}

		for _, v := range page.ProtectionGroups {
			resource, err := shieldProtectionGroupHandle(ctx, cfg, v)
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
	}

	return values, nil
}
func shieldProtectionGroupHandle(ctx context.Context, cfg aws.Config, v types.ProtectionGroup) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := shield.NewFromConfig(cfg)

	tags, err := client.ListTagsForResource(ctx, &shield.ListTagsForResourceInput{
		ResourceARN: v.ProtectionGroupArn,
	})
	if err != nil {
		if isErr(err, "ListTagsForResourceNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *v.ProtectionGroupArn,
		Name:   *v.ProtectionGroupId,
		Description: model.ShieldProtectionGroupDescription{
			ProtectionGroup: v,
			Tags:            tags.Tags,
		},
	}
	return resource, nil
}
func GetShieldProtectionGroup(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	protectionGroupId := fields["id"]
	client := shield.NewFromConfig(cfg)
	var values []models.Resource

	out, err := client.DescribeProtectionGroup(ctx, &shield.DescribeProtectionGroupInput{
		ProtectionGroupId: &protectionGroupId,
	})
	if err != nil {
		if isErr(err, "DescribeProtectionGroupNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	resource, err := shieldProtectionGroupHandle(ctx, cfg, *out.ProtectionGroup)
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
