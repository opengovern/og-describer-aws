package describers

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func ServiceQuotasService(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := servicequotas.NewFromConfig(cfg)

	servicesPaginator := servicequotas.NewListServicesPaginator(client, &servicequotas.ListServicesInput{})

	var values []models.Resource
	for servicesPaginator.HasMorePages() {
		servicesPage, err := servicesPaginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, service := range servicesPage.Services {
			arn := fmt.Sprintf("arn:%s:servicequotas:%s:%s:%s", describeCtx.Partition, describeCtx.OGRegion, describeCtx.AccountID, *service.ServiceCode)

			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    arn,
				Name:   *service.ServiceName,
				ID:     *service.ServiceCode,
				Description: model.ServiceQuotasServiceDescription{
					Service: service,
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

func ServiceQuotasDefaultServiceQuota(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := servicequotas.NewFromConfig(cfg)

	var values []models.Resource
	servicesPaginator := servicequotas.NewListServicesPaginator(client, &servicequotas.ListServicesInput{})
	for servicesPaginator.HasMorePages() {
		servicesPage, err := servicesPaginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, service := range servicesPage.Services {
			paginator := servicequotas.NewListAWSDefaultServiceQuotasPaginator(client, &servicequotas.ListAWSDefaultServiceQuotasInput{
				ServiceCode: service.ServiceCode,
			})
			for paginator.HasMorePages() {
				page, err := paginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}

				for _, quota := range page.Quotas {
					resource := models.Resource{
						Region: describeCtx.OGRegion,
						ARN:    *quota.QuotaArn + "--default",
						Name:   *quota.QuotaName,
						ID:     *quota.QuotaCode,
						Description: model.ServiceQuotasDefaultServiceQuotaDescription{
							DefaultServiceQuota: quota,
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
		}
	}
	return values, nil
}

func ServiceQuotasServiceQuota(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := servicequotas.NewFromConfig(cfg)

	var values []models.Resource
	servicesPaginator := servicequotas.NewListServicesPaginator(client, &servicequotas.ListServicesInput{})
	for servicesPaginator.HasMorePages() {
		servicesPage, err := servicesPaginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, service := range servicesPage.Services {
			paginator := servicequotas.NewListServiceQuotasPaginator(client, &servicequotas.ListServiceQuotasInput{
				ServiceCode: service.ServiceCode,
			})
			for paginator.HasMorePages() {
				page, err := paginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}

				for _, quota := range page.Quotas {
					tags, err := client.ListTagsForResource(ctx, &servicequotas.ListTagsForResourceInput{
						ResourceARN: quota.QuotaArn,
					})
					if err != nil {
						tags = &servicequotas.ListTagsForResourceOutput{}
					}
					resource := models.Resource{
						Region: describeCtx.OGRegion,
						ARN:    *quota.QuotaArn,
						Name:   *quota.QuotaName,
						ID:     *quota.QuotaCode,
						Description: model.ServiceQuotasServiceQuotaDescription{
							ServiceQuota: quota,
							Tags:         tags.Tags,
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
		}
	}
	return values, nil
}

func ServiceQuotasServiceQuotaChangeRequest(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := servicequotas.NewFromConfig(cfg)

	var values []models.Resource

	paginator := servicequotas.NewListRequestedServiceQuotaChangeHistoryPaginator(client, &servicequotas.ListRequestedServiceQuotaChangeHistoryInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, requestedQuota := range page.RequestedQuotas {
			tags, err := client.ListTagsForResource(ctx, &servicequotas.ListTagsForResourceInput{
				ResourceARN: requestedQuota.QuotaArn,
			})
			if err != nil {
				tags = &servicequotas.ListTagsForResourceOutput{}
			}

			arn := fmt.Sprintf("arn:aws:servicequotas:%s:%s:changeRequest/%s", describeCtx.OGRegion, describeCtx.AccountID, *requestedQuota.Id)
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    arn,
				ID:     *requestedQuota.Id,
				Description: model.ServiceQuotasServiceQuotaChangeRequestDescription{
					ServiceQuotaChangeRequest: requestedQuota,
					Tags:                      tags.Tags,
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
