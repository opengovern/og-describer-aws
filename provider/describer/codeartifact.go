package describer

import (
	"context"
	"github.com/opengovern/og-describer-aws/pkg/sdk/models"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/opengovern/og-describer-aws/provider/model"
)

func CodeArtifactRepository(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := codeartifact.NewFromConfig(cfg)
	paginator := codeartifact.NewListRepositoriesPaginator(client, &codeartifact.ListRepositoriesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Repositories {
			resource, err := codeArtifactRepositoryHandle(ctx, cfg, v)
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
func codeArtifactRepositoryHandle(ctx context.Context, cfg aws.Config, v types.RepositorySummary) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := codeartifact.NewFromConfig(cfg)

	tags, err := client.ListTagsForResource(ctx, &codeartifact.ListTagsForResourceInput{
		ResourceArn: v.Arn,
	})
	if err != nil {
		if isErr(err, "ListTagsForResourceNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}
	var policy types.ResourcePolicy
	policyOutput, err := client.GetRepositoryPermissionsPolicy(ctx, &codeartifact.GetRepositoryPermissionsPolicyInput{
		Domain:      v.DomainName,
		Repository:  v.Name,
		DomainOwner: v.DomainOwner,
	})
	if err != nil {
		if strings.Contains(err.Error(), "ResourceNotFoundException") {
			policy = types.ResourcePolicy{}
		} else {
			if isErr(err, "GetRepositoryPermissionsPolicyNotFound") || isErr(err, "InvalidParameterValue") {
				return models.Resource{}, nil
			}
			return models.Resource{}, err
		}
	} else {
		policy = *policyOutput.Policy
	}
	description, err := client.DescribeRepository(ctx, &codeartifact.DescribeRepositoryInput{
		Domain:      v.Name,
		DomainOwner: v.DomainOwner,
		Repository:  v.Name,
	})
	if err != nil {
		if isErr(err, "DescribeRepositoryNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	resultData := []string{}

	for _, item := range description.Repository.ExternalConnections {

		// Build the params
		params := &codeartifact.GetRepositoryEndpointInput{
			Repository:  description.Repository.Name,
			Domain:      description.Repository.DomainName,
			DomainOwner: description.Repository.DomainOwner,
			Format:      item.PackageFormat,
		}

		if err != nil {
			return models.Resource{}, err
		}

		data, err := client.GetRepositoryEndpoint(ctx, params)

		if err != nil {
			return models.Resource{}, err
		}
		resultData = append(resultData, *data.RepositoryEndpoint)
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *v.Arn,
		Name:   *v.Name,
		Description: model.CodeArtifactRepositoryDescription{
			Repository:  v,
			Policy:      policy,
			Description: *description.Repository,
			Endpoints:   resultData,
			Tags:        tags.Tags,
		},
	}
	return resource, nil
}
func GetCodeArtifactRepository(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	repository := fields["repository"]
	client := codeartifact.NewFromConfig(cfg)

	out, err := client.DescribeRepository(ctx, &codeartifact.DescribeRepositoryInput{
		Repository: &repository,
	})
	if err != nil {
		if isErr(err, "DescribeRepositoryNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	Repo := types.RepositorySummary{
		Arn:                  out.Repository.Arn,
		AdministratorAccount: out.Repository.AdministratorAccount,
		Description:          out.Repository.Description,
		DomainName:           out.Repository.DomainName,
		DomainOwner:          out.Repository.DomainOwner,
		Name:                 out.Repository.Name,
	}

	var values []models.Resource

	resource, err := codeArtifactRepositoryHandle(ctx, cfg, Repo)
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

func CodeArtifactDomain(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := codeartifact.NewFromConfig(cfg)
	paginator := codeartifact.NewListDomainsPaginator(client, &codeartifact.ListDomainsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Domains {
			resource, err := CodeArtifactDomainHandle(ctx, cfg, v)
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
func CodeArtifactDomainHandle(ctx context.Context, cfg aws.Config, v types.DomainSummary) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := codeartifact.NewFromConfig(cfg)
	tags, err := client.ListTagsForResource(ctx, &codeartifact.ListTagsForResourceInput{
		ResourceArn: v.Arn,
	})
	if err != nil {
		if isErr(err, "ListTagsForResourceNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	domain, err := client.DescribeDomain(ctx, &codeartifact.DescribeDomainInput{
		Domain:      v.Name,
		DomainOwner: v.Owner,
	})
	if err != nil {
		if isErr(err, "DescribeDomainNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	policy, err := client.GetDomainPermissionsPolicy(ctx, &codeartifact.GetDomainPermissionsPolicyInput{
		Domain:      v.Name,
		DomainOwner: v.Owner,
	})
	if err != nil {
		if isErr(err, "ResourceNotFoundException") {
			policy = &codeartifact.GetDomainPermissionsPolicyOutput{}
		} else {
			return models.Resource{}, err
		}
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *v.Arn,
		Description: model.CodeArtifactDomainDescription{
			Domain: *domain.Domain,
			Policy: *policy.Policy,
			Tags:   tags.Tags,
		},
	}
	if v.Name != nil {
		resource.Name = *v.Name
	}

	return resource, nil
}
func GetCodeArtifactDomain(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	domainName := fields["domainName"]
	client := codeartifact.NewFromConfig(cfg)
	domains, err := client.ListDomains(ctx, &codeartifact.ListDomainsInput{})
	if err != nil {
		if isErr(err, "ListDomainsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, v := range domains.Domains {
		if *v.Name != domainName {
			continue
		}

		resource, err := CodeArtifactDomainHandle(ctx, cfg, v)
		if err != nil {
			return nil, err
		}
		emptyResource := models.Resource{}
		if err == nil && resource == emptyResource {
			return nil, nil
		}

		values = append(values, resource)
	}

	return values, nil
}
