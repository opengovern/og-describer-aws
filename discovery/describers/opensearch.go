package describers

import (
	"context"
	"math"

	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opensearch"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
)

func OpenSearchDomain(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := opensearch.NewFromConfig(cfg)

	domainNamesOutput, err := client.ListDomainNames(ctx, &opensearch.ListDomainNamesInput{})
	if err != nil {
		return nil, err
	}
	domainNames := make([]string, 0, len(domainNamesOutput.DomainNames))
	for _, domainName := range domainNamesOutput.DomainNames {
		domainNames = append(domainNames, *domainName.DomainName)
	}

	var values []models.Resource
	// OpenSearch API only allows 5 domains per request
	for i := 0; i < len(domainNames); i = i + 5 {
		domains, err := client.DescribeDomains(ctx, &opensearch.DescribeDomainsInput{
			DomainNames: domainNames[i:int(math.Min(float64(i+5), float64(len(domainNames))))],
		})
		if err != nil {
			return nil, err
		}

		for _, domain := range domains.DomainStatusList {
			resource, err := openSearchDomainHandle(ctx, cfg, domain)
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
func openSearchDomainHandle(ctx context.Context, cfg aws.Config, domain types.DomainStatus) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := opensearch.NewFromConfig(cfg)

	tags, err := client.ListTags(ctx, &opensearch.ListTagsInput{
		ARN: domain.ARN,
	})
	if err != nil {
		if isErr(err, "ListTagsNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *domain.ARN,
		Name:   *domain.DomainName,
		Description: model.OpenSearchDomainDescription{
			Domain: domain,
			Tags:   tags.TagList,
		},
	}
	return resource, nil
}
func GetOpenSearchDomain(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	client := opensearch.NewFromConfig(cfg)
	var values []models.Resource

	domainName := fields["domainName"]
	domain, err := client.DescribeDomain(ctx, &opensearch.DescribeDomainInput{
		DomainName: &domainName,
	})
	if err != nil {
		return nil, err
	}

	resource, err := openSearchDomainHandle(ctx, cfg, *domain.DomainStatus)
	emptyResource := models.Resource{}
	if err == nil && resource == emptyResource {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	values = append(values, resource)
	return values, nil
}
