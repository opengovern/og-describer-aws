package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func IdentityStoreGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := identitystore.NewFromConfig(cfg)
	ssoadminClient := ssoadmin.NewFromConfig(cfg)
	paginator := ssoadmin.NewListInstancesPaginator(ssoadminClient, &ssoadmin.ListInstancesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Instances {
			paginator2 := identitystore.NewListGroupsPaginator(client, &identitystore.ListGroupsInput{IdentityStoreId: v.IdentityStoreId})
			for paginator2.HasMorePages() {
				page2, err2 := paginator2.NextPage(ctx)
				if err2 != nil {
					return nil, err2
				}

				for _, group := range page2.Groups {
					resource := models.Resource{
						Region:  describeCtx.OGRegion,
						ID:      *group.GroupId,
						Account: describeCtx.AccountID,
						Name:    *group.DisplayName,
						Description: model.IdentityStoreGroupDescription{
							Group: group,
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

func IdentityStoreUser(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := identitystore.NewFromConfig(cfg)
	ssoadminClient := ssoadmin.NewFromConfig(cfg)
	paginator := ssoadmin.NewListInstancesPaginator(ssoadminClient, &ssoadmin.ListInstancesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, i := range page.Instances {
			paginator2 := identitystore.NewListUsersPaginator(client, &identitystore.ListUsersInput{IdentityStoreId: i.IdentityStoreId})
			page2, err2 := paginator2.NextPage(ctx)
			if err2 != nil {
				return nil, err2
			}
			for _, user := range page2.Users {
				var primaryEmail *string
				for _, e := range user.Emails {
					if e.Primary {
						primaryEmail = e.Value
					}
				}
				resource := models.Resource{
					Region:  describeCtx.OGRegion,
					ID:      *user.UserId,
					Account: describeCtx.AccountID,
					Name:    *user.UserName,
					Description: model.IdentityStoreUserDescription{
						User:         user,
						PrimaryEmail: primaryEmail,
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

	return values, nil
}

func IdentityStoreGroupMembership(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := identitystore.NewFromConfig(cfg)

	ssoadminClient := ssoadmin.NewFromConfig(cfg)
	paginator := ssoadmin.NewListInstancesPaginator(ssoadminClient, &ssoadmin.ListInstancesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, i := range page.Instances {
			paginator2 := identitystore.NewListGroupsPaginator(client, &identitystore.ListGroupsInput{IdentityStoreId: i.IdentityStoreId})
			page2, err2 := paginator2.NextPage(ctx)
			if err2 != nil {
				return nil, err2
			}

			for _, group := range page2.Groups {
				membershipPaginator := identitystore.NewListGroupMembershipsPaginator(client, &identitystore.ListGroupMembershipsInput{
					GroupId:         group.GroupId,
					IdentityStoreId: group.IdentityStoreId,
				})
				for membershipPaginator.HasMorePages() {
					membershipPage, err := membershipPaginator.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					for _, membership := range membershipPage.GroupMemberships {
						resource := models.Resource{
							Region:  describeCtx.OGRegion,
							ID:      *membership.MembershipId,
							Account: describeCtx.AccountID,
							Description: model.IdentityStoreGroupMembershipDescription{
								GroupId:         membership.GroupId,
								IdentityStoreId: membership.IdentityStoreId,
								MembershipId:    membership.MembershipId,
								MemberId:        membership.MemberId,
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
	}

	return values, nil
}
