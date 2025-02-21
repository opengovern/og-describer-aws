package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func RamPrincipalAssociation(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := ram.NewFromConfig(cfg)

	var values []models.Resource
	paginator := ram.NewGetResourceShareAssociationsPaginator(client, &ram.GetResourceShareAssociationsInput{AssociationType: types.ResourceShareAssociationTypePrincipal})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, association := range page.ResourceShareAssociations {
			resource, err := ramPrincipalAssociationHandle(ctx, cfg, association, *association.ResourceShareArn)
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
func GetRamPrincipalAssociation(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	resourceShareArn := fields["ResourceShareArn"]
	client := ram.NewFromConfig(cfg)

	associations, err := client.GetResourceShareAssociations(ctx, &ram.GetResourceShareAssociationsInput{
		ResourceShareArns: []string{resourceShareArn},
	})
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, association := range associations.ResourceShareAssociations {
		resource, err := ramPrincipalAssociationHandle(ctx, cfg, association, resourceShareArn)
		if err != nil {
			return nil, err
		}
		values = append(values, resource)
	}
	return values, nil
}
func ramPrincipalAssociationHandle(ctx context.Context, cfg aws.Config, association types.ResourceShareAssociation, resourceShareArn string) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := ram.NewFromConfig(cfg)

	permissionPaginator := ram.NewListResourceSharePermissionsPaginator(client, &ram.ListResourceSharePermissionsInput{
		ResourceShareArn: &resourceShareArn,
	})

	var permissions []types.ResourceSharePermissionSummary
	for permissionPaginator.HasMorePages() {
		permissionPage, err := permissionPaginator.NextPage(ctx)
		if err != nil {
			return models.Resource{}, err
		}
		permissions = append(permissions, permissionPage.Permissions...)
	}

	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		Name:    *association.ResourceShareName,
		ARN:     *association.ResourceShareArn,
		Account: describeCtx.AccountID,
		Description: model.RamPrincipalAssociationDescription{
			PrincipalAssociation:    association,
			ResourceSharePermission: permissions,
		},
	}
	return resource, nil
}

func RamResourceAssociation(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := ram.NewFromConfig(cfg)

	var values []models.Resource
	paginator := ram.NewGetResourceShareAssociationsPaginator(client, &ram.GetResourceShareAssociationsInput{AssociationType: types.ResourceShareAssociationTypeResource})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, association := range page.ResourceShareAssociations {
			resource, err := ramResourceAssociationHandle(ctx, cfg, association, *association.ResourceShareArn)
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
func ramResourceAssociationHandle(ctx context.Context, cfg aws.Config, association types.ResourceShareAssociation, resourceShareArn string) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := ram.NewFromConfig(cfg)
	permissionPaginator := ram.NewListResourceSharePermissionsPaginator(client, &ram.ListResourceSharePermissionsInput{
		ResourceShareArn: &resourceShareArn,
	})
	var permissions []types.ResourceSharePermissionSummary
	for permissionPaginator.HasMorePages() {
		permissionPage, err := permissionPaginator.NextPage(ctx)
		if err != nil {
			return models.Resource{}, err
		}
		permissions = append(permissions, permissionPage.Permissions...)
	}

	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		Name:    *association.ResourceShareName,
		ARN:     *association.ResourceShareArn,
		Account: describeCtx.AccountID,
		Description: model.RamResourceAssociationDescription{
			ResourceAssociation:     association,
			ResourceSharePermission: permissions,
		},
	}
	return resource, nil
}
func GetRamResourceAssociation(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	resourceShareArn := fields["resourceShareArn"]
	client := ram.NewFromConfig(cfg)

	associations, err := client.GetResourceShareAssociations(ctx, &ram.GetResourceShareAssociationsInput{
		ResourceShareArns: []string{resourceShareArn},
	})
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, association := range associations.ResourceShareAssociations {

		resource, err := ramResourceAssociationHandle(ctx, cfg, association, resourceShareArn)
		if err != nil {
			return nil, err
		}

		values = append(values, resource)
	}
	return values, nil
}
