package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func AccountAlternateContact(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	var values []models.Resource

	contactTypes := []types.AlternateContactType{types.AlternateContactTypeBilling, types.AlternateContactTypeOperations, types.AlternateContactTypeSecurity}
	input := &account.GetAlternateContactInput{
		AccountId: &describeCtx.AccountID,
	}
	for _, contactType := range contactTypes {
		input.AlternateContactType = contactType
		resource, err := accountAlternateContactHandle(ctx, cfg, *input.AccountId, input.AlternateContactType)
		if err != nil {
			return nil, err
		}
		emptyResource := models.Resource{}
		if err == nil && resource == emptyResource {
			continue
		}

		if stream != nil {
			m := *stream
			err := m(resource)
			if err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}
func accountAlternateContactHandle(ctx context.Context, cfg aws.Config, accountId string, contactType types.AlternateContactType) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := account.NewFromConfig(cfg)
	op, err := client.GetAlternateContact(ctx, &account.GetAlternateContactInput{
		AlternateContactType: contactType,
		AccountId:            &accountId,
	})
	if err != nil {
		if isErr(err, "ResourceNotFoundException") {
			op = &account.GetAlternateContactOutput{}
		}
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		Name:   *op.AlternateContact.Name,
		Description: model.AccountAlternateContactDescription{
			AlternateContact: *op.AlternateContact,
			LinkedAccountID:  describeCtx.AccountID,
		},
	}
	return resource, nil
}
func GetAccountAlternateContact(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	accountId := fields["accountId"]
	contactTypes := []types.AlternateContactType{types.AlternateContactTypeBilling, types.AlternateContactTypeOperations, types.AlternateContactTypeSecurity}
	var values []models.Resource
	for _, contactType := range contactTypes {
		resource, err := accountAlternateContactHandle(ctx, cfg, accountId, contactType)
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

func AccountContact(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := account.NewFromConfig(cfg)

	var values []models.Resource

	input := &account.GetContactInformationInput{}
	op, err := client.GetContactInformation(ctx, input)
	if err != nil {
		return nil, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		Name:   *op.ContactInformation.FullName,
		Description: model.AccountContactDescription{
			AlternateContact: *op.ContactInformation,
			LinkedAccountID:  describeCtx.AccountID,
		},
	}
	if stream != nil {
		m := *stream
		err := m(resource)
		if err != nil {
			return nil, err
		}
	} else {
		values = append(values, resource)
	}

	return values, nil
}
