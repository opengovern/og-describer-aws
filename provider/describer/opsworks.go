package describer

import (
	"context"
	"github.com/opengovern/og-describer-aws/pkg/sdk/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	"github.com/opengovern/og-describer-aws/provider/model"
)

func OpsWorksCMServer(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := opsworkscm.NewFromConfig(cfg)
	paginator := opsworkscm.NewDescribeServersPaginator(client, &opsworkscm.DescribeServersInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Servers {
			resource, err := opsWorksCMServerHandle(ctx, cfg, v)
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
func opsWorksCMServerHandle(ctx context.Context, cfg aws.Config, v types.Server) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := opsworkscm.NewFromConfig(cfg)

	tags, err := client.ListTagsForResource(ctx, &opsworkscm.ListTagsForResourceInput{
		ResourceArn: v.ServerArn,
	})
	if err != nil {
		if isErr(err, "ListTagsForResourceNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *v.ServerArn,
		Name:   *v.ServerName,
		Description: model.OpsWorksCMServerDescription{
			Server: v,
			Tags:   tags.Tags,
		},
	}
	return resource, nil
}
func GetOpsWorksCMServer(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	serverName := fields["name"]
	client := opsworkscm.NewFromConfig(cfg)

	server, err := client.DescribeServers(ctx, &opsworkscm.DescribeServersInput{
		ServerName: &serverName,
	})
	if err != nil {
	}

	var values []models.Resource
	for _, v := range server.Servers {

		resource, err := opsWorksCMServerHandle(ctx, cfg, v)
		emptyResource := models.Resource{}
		if err == nil && resource == emptyResource {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}

		values = append(values, resource)
	}
	return values, nil
}
