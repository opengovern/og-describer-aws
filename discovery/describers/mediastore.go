package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/mediastore/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediastore"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
)

func MediaStoreContainer(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := mediastore.NewFromConfig(cfg)

	var values []models.Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		containers, err := client.ListContainers(ctx, &mediastore.ListContainersInput{
			NextToken: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, container := range containers.Containers {
			resource := mediaStoreContainerHandle(ctx, cfg, container)

			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}

		return containers.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
func mediaStoreContainerHandle(ctx context.Context, cfg aws.Config, container types.Container) models.Resource {
	describeCtx := GetDescribeContext(ctx)
	client := mediastore.NewFromConfig(cfg)

	policy, err := client.GetContainerPolicy(ctx, &mediastore.GetContainerPolicyInput{
		ContainerName: container.Name,
	})
	if err != nil {
		policy = nil
	}

	tags, err := client.ListTagsForResource(ctx, &mediastore.ListTagsForResourceInput{
		Resource: container.ARN,
	})
	if err != nil {
		tags = &mediastore.ListTagsForResourceOutput{}
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *container.ARN,
		Name:   *container.Name,
		Description: model.MediaStoreContainerDescription{
			Container: container,
			Policy:    policy,
			Tags:      tags.Tags,
		},
	}

	return resource
}
func GetMediaStoreContainer(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	containerName := fields["name"]
	var values []models.Resource
	client := mediastore.NewFromConfig(cfg)
	out, err := client.DescribeContainer(ctx, &mediastore.DescribeContainerInput{
		ContainerName: &containerName,
	})
	if err != nil {
		return nil, err
	}

	values = append(values, mediaStoreContainerHandle(ctx, cfg, *out.Container))
	return values, nil
}
