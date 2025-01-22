package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func GlacierVault(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)

	client := glacier.NewFromConfig(cfg)
	paginator := glacier.NewListVaultsPaginator(client, &glacier.ListVaultsInput{
		AccountId: &describeCtx.AccountID,
	})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if !isErr(err, "ResourceNotFoundException") && !isErr(err, "InvalidParameter") {
				return nil, err
			}
			continue
		}

		for _, vault := range page.VaultList {
			accessPolicy, err := client.GetVaultAccessPolicy(ctx, &glacier.GetVaultAccessPolicyInput{
				AccountId: &describeCtx.AccountID,
				VaultName: vault.VaultName,
			})
			if err != nil {
				if !isErr(err, "ResourceNotFoundException") && !isErr(err, "InvalidParameter") {
					return nil, err
				}
				accessPolicy = &glacier.GetVaultAccessPolicyOutput{
					Policy: &types.VaultAccessPolicy{},
				}
			}

			lockPolicy, err := client.GetVaultLock(ctx, &glacier.GetVaultLockInput{
				AccountId: &describeCtx.AccountID,
				VaultName: vault.VaultName,
			})
			if err != nil {
				if !isErr(err, "ResourceNotFoundException") && !isErr(err, "InvalidParameter") {
					return nil, err
				}
				lockPolicy = &glacier.GetVaultLockOutput{}
			}

			tags, err := client.ListTagsForVault(ctx, &glacier.ListTagsForVaultInput{
				AccountId: &describeCtx.AccountID,
				VaultName: vault.VaultName,
			})
			if err != nil {
				if !isErr(err, "ResourceNotFoundException") && !isErr(err, "InvalidParameter") {
					return nil, err
				}
				tags = &glacier.ListTagsForVaultOutput{}
			}

			var vaultNotificationConfig types.VaultNotificationConfig
			notifications, err := client.GetVaultNotifications(ctx, &glacier.GetVaultNotificationsInput{
				AccountId: &describeCtx.AccountID,
				VaultName: vault.VaultName,
			})
			if err != nil {
				if !isErr(err, "ResourceNotFoundException") && !isErr(err, "InvalidParameter") {
					return nil, err
				}
			} else if notifications != nil && notifications.VaultNotificationConfig != nil {
				vaultNotificationConfig = *notifications.VaultNotificationConfig
			}

			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *vault.VaultARN,
				Name:   *vault.VaultName,
				Description: model.GlacierVaultDescription{
					Vault:        vault,
					AccessPolicy: *accessPolicy.Policy,
					LockPolicy: types.VaultLockPolicy{
						Policy: lockPolicy.Policy,
					},
					VaultNotificationConfig: vaultNotificationConfig,
					Tags:                    tags.Tags,
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

func GetGlacierVault(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	vaultName := fields["name"]

	client := glacier.NewFromConfig(cfg)
	vault, err := client.DescribeVault(ctx, &glacier.DescribeVaultInput{
		AccountId: &describeCtx.AccountID,
		VaultName: &vaultName,
	})
	if err != nil {
		if !isErr(err, "ResourceNotFoundException") && !isErr(err, "InvalidParameter") {
			return nil, err
		}
		return nil, nil
	}

	var values []models.Resource
	accessPolicy, err := client.GetVaultAccessPolicy(ctx, &glacier.GetVaultAccessPolicyInput{
		AccountId: &describeCtx.AccountID,
		VaultName: vault.VaultName,
	})
	if err != nil {
		if !isErr(err, "ResourceNotFoundException") && !isErr(err, "InvalidParameter") {
			return nil, err
		}
		accessPolicy = &glacier.GetVaultAccessPolicyOutput{
			Policy: &types.VaultAccessPolicy{},
		}
	}

	lockPolicy, err := client.GetVaultLock(ctx, &glacier.GetVaultLockInput{
		AccountId: &describeCtx.AccountID,
		VaultName: vault.VaultName,
	})
	if err != nil {
		if !isErr(err, "ResourceNotFoundException") && !isErr(err, "InvalidParameter") {
			return nil, err
		}
		lockPolicy = &glacier.GetVaultLockOutput{}
	}

	tags, err := client.ListTagsForVault(ctx, &glacier.ListTagsForVaultInput{
		AccountId: &describeCtx.AccountID,
		VaultName: vault.VaultName,
	})
	if err != nil {
		if !isErr(err, "ResourceNotFoundException") && !isErr(err, "InvalidParameter") {
			return nil, err
		}
		tags = &glacier.ListTagsForVaultOutput{}
	}

	values = append(values, models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *vault.VaultARN,
		Name:   *vault.VaultName,
		Description: model.GlacierVaultDescription{
			Vault: types.DescribeVaultOutput{
				CreationDate:      vault.CreationDate,
				LastInventoryDate: vault.LastInventoryDate,
				NumberOfArchives:  vault.NumberOfArchives,
				SizeInBytes:       vault.SizeInBytes,
				VaultARN:          vault.VaultARN,
				VaultName:         vault.VaultName,
			},
			AccessPolicy: *accessPolicy.Policy,
			LockPolicy: types.VaultLockPolicy{
				Policy: lockPolicy.Policy,
			},
			Tags: tags.Tags,
		},
	})

	return values, nil
}
