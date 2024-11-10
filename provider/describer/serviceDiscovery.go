package describer

import (
	"context"
	"github.com/opengovern/og-describer-aws/pkg/sdk/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/opengovern/og-describer-aws/provider/model"
)

func ServiceDiscoveryService(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := servicediscovery.NewFromConfig(cfg)

	paginator := servicediscovery.NewListServicesPaginator(client, &servicediscovery.ListServicesInput{})
	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, item := range page.Services {
			tag, err := client.ListTagsForResource(ctx, &servicediscovery.ListTagsForResourceInput{
				ResourceARN: item.Arn,
			})
			if err != nil {
				return nil, err
			}

			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ID:     *item.Id,
				Description: model.ServiceDiscoveryServiceDescription{
					Service: item,
					Tags:    tag.Tags,
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

func ServiceDiscoveryNamespace(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := servicediscovery.NewFromConfig(cfg)

	paginator := servicediscovery.NewListNamespacesPaginator(client, &servicediscovery.ListNamespacesInput{})
	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Namespaces {
			tag, err := client.ListTagsForResource(ctx, &servicediscovery.ListTagsForResourceInput{
				ResourceARN: v.Arn,
			})
			if err != nil {
				return nil, err
			}

			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ID:     *v.Id,
				Name:   *v.Name,
				Description: model.ServiceDiscoveryNamespaceDescription{
					Namespace: v,
					Tags:      tag.Tags,
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
func ServiceDiscoveryInstance(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := servicediscovery.NewFromConfig(cfg)

	paginator := servicediscovery.NewListServicesPaginator(client, &servicediscovery.ListServicesInput{})
	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Services {
			resources, err := getServiceDiscoveryInstances(ctx, cfg, v.Id)
			if err != nil {
				return nil, err
			}
			for _, resource := range resources {
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
	return values, nil
}

func getServiceDiscoveryInstances(ctx context.Context, cfg aws.Config, id *string) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := servicediscovery.NewFromConfig(cfg)

	paginator := servicediscovery.NewListInstancesPaginator(client, &servicediscovery.ListInstancesInput{ServiceId: id})
	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Instances {
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ID:     *v.Id,
				Name:   *v.Id,
				Description: model.ServiceDiscoveryInstanceDescription{
					Instance: v,
				},
			}
			values = append(values, resource)
		}
	}
	return values, nil
}
