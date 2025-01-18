package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
)

func AppStreamApplication(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := appstream.NewFromConfig(cfg)

	var values []models.Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.DescribeApplications(ctx, &appstream.DescribeApplicationsInput{
			NextToken: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, item := range output.Applications {
			tags, err := client.ListTagsForResource(ctx, &appstream.ListTagsForResourceInput{
				ResourceArn: item.Arn,
			})
			if err != nil {
				return nil, err
			}
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *item.Arn,
				Name:   *item.Name,
				Description: model.AppStreamApplicationDescription{
					Application: item,
					Tags:        tags.Tags,
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
		return output.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func AppStreamStack(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := appstream.NewFromConfig(cfg)

	var values []models.Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.DescribeStacks(ctx, &appstream.DescribeStacksInput{
			NextToken: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, item := range output.Stacks {

			resource, err := appStreamStackHandle(ctx, cfg, item)
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
		return output.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
func appStreamStackHandle(ctx context.Context, cfg aws.Config, item types.Stack) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := appstream.NewFromConfig(cfg)

	tags, err := client.ListTagsForResource(ctx, &appstream.ListTagsForResourceInput{
		ResourceArn: item.Arn,
	})
	if err != nil {
		if isErr(err, "ListTagsForResourceNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *item.Arn,
		Name:   *item.Name,
		Description: model.AppStreamStackDescription{
			Stack: item,
			Tags:  tags.Tags,
		},
	}
	return resource, nil
}
func GetAppStreamStack(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	name := fields["name"]
	client := appstream.NewFromConfig(cfg)
	out, err := client.DescribeStacks(ctx, &appstream.DescribeStacksInput{
		Names: []string{name},
	})
	if err != nil {
		if isErr(err, "DescribeStacksNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, v := range out.Stacks {

		resource, err := appStreamStackHandle(ctx, cfg, v)
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

func AppStreamFleet(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := appstream.NewFromConfig(cfg)

	var values []models.Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.DescribeFleets(ctx, &appstream.DescribeFleetsInput{
			NextToken: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, item := range output.Fleets {
			resource, err := appStreamFleetHandle(ctx, cfg, item)
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
		return output.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
func appStreamFleetHandle(ctx context.Context, cfg aws.Config, item types.Fleet) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := appstream.NewFromConfig(cfg)
	tags, err := client.ListTagsForResource(ctx, &appstream.ListTagsForResourceInput{
		ResourceArn: item.Arn,
	})
	if err != nil {
		if isErr(err, "ListTagsForResourceNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *item.Arn,
		Name:   *item.Name,
		Description: model.AppStreamFleetDescription{
			Fleet: item,
			Tags:  tags.Tags,
		},
	}

	return resource, nil
}
func GetAppStreamFleet(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	name := fields["name"]
	client := appstream.NewFromConfig(cfg)

	out, err := client.DescribeFleets(ctx, &appstream.DescribeFleetsInput{
		Names: []string{name},
	})
	if err != nil {
		if isErr(err, "DescribeFleetsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, v := range out.Fleets {

		resource, err := appStreamFleetHandle(ctx, cfg, v)
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

func AppStreamImage(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := appstream.NewFromConfig(cfg)

	var values []models.Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.DescribeImages(ctx, &appstream.DescribeImagesInput{
			NextToken: prevToken,
		})
		if err != nil {
			if isErr(err, "AccessDeniedException") {
				return nil, nil
			} else {
				return nil, err
			}
		}

		for _, item := range output.Images {
			resource, err := appStreamImageHandle(ctx, cfg, item)
			emptyResource := models.Resource{}
			if err == nil && resource == emptyResource {
				return nil, nil
			}
			if err != nil {
				if isErr(err, "AccessDeniedException") {
					return nil, nil
				} else {
					return nil, err
				}
			}

			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
		return output.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
func appStreamImageHandle(ctx context.Context, cfg aws.Config, item types.Image) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := appstream.NewFromConfig(cfg)
	tags, err := client.ListTagsForResource(ctx, &appstream.ListTagsForResourceInput{
		ResourceArn: item.Arn,
	})
	if err != nil {
		if isErr(err, "ListTagsForResourceNotFound") || isErr(err, "InvalidParameterValue") || isErr(err, "AccessDeniedException") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *item.Arn,
		Name:   *item.Name,
		Description: model.AppStreamImageDescription{
			Image: item,
			Tags:  tags.Tags,
		},
	}

	return resource, nil
}
func GetAppStreamImage(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	name := fields["name"]
	client := appstream.NewFromConfig(cfg)

	out, err := client.DescribeImages(ctx, &appstream.DescribeImagesInput{
		Names: []string{name},
	})
	if err != nil {
		if isErr(err, "DescribeFleetsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, v := range out.Images {

		resource, err := appStreamImageHandle(ctx, cfg, v)
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
