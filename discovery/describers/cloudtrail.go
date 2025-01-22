package describers

import (
	"context"

	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func CloudTrailTrail(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := cloudtrail.NewFromConfig(cfg)
	paginator := cloudtrail.NewListTrailsPaginator(client, &cloudtrail.ListTrailsInput{})

	//stsClient := sts.NewFromConfig(cfg)
	//identity, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	//if err != nil {
	//	if isErr(err, "GetCallerIdentityNotFound") || isErr(err, "InvalidParameterValue") {
	//		return nil, nil
	//	}
	//	return nil, err
	//}

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		var trails []string
		for _, trail := range page.Trails {
			// Ignore trails that don't belong this region (Based on steampipe)
			if !strings.EqualFold(*trail.HomeRegion, cfg.Region) {
				continue
			}

			if trail.TrailARN != nil {
				// Ignore trails that don't belong to this account (Based on steampipe)
				//if aws.ToString(identity.Account) != arnToAccountId(*trail.TrailARN) {
				//	continue
				//}

				trails = append(trails, *trail.TrailARN)
			} else if trail.Name != nil {
				trails = append(trails, *trail.Name)
			} else {
				continue
			}
		}

		output, err := client.DescribeTrails(ctx, &cloudtrail.DescribeTrailsInput{
			IncludeShadowTrails: aws.Bool(false),
			TrailNameList:       trails,
		})
		if err != nil {
			return nil, err
		}

		for _, v := range output.TrailList {
			resource, err := cloudTrailTrailHandle(ctx, cfg, v)
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
func cloudTrailTrailHandle(ctx context.Context, cfg aws.Config, v types.Trail) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := cloudtrail.NewFromConfig(cfg)
	statusOutput, err := client.GetTrailStatus(ctx, &cloudtrail.GetTrailStatusInput{
		Name: v.TrailARN,
	})
	if err != nil {
		if isErr(err, "GetTrailStatusNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	selectorOutput, err := client.GetEventSelectors(ctx, &cloudtrail.GetEventSelectorsInput{
		TrailName: v.TrailARN,
	})
	if err != nil {
		if isErr(err, "GetEventSelectorsNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	tagsOutput, err := client.ListTags(ctx, &cloudtrail.ListTagsInput{
		ResourceIdList: []string{*v.TrailARN},
	})
	if err != nil {
		if isErr(err, "ListTagsNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	var tags []types.Tag
	if len(tagsOutput.ResourceTagList) > 0 {
		tags = tagsOutput.ResourceTagList[0].TagsList
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *v.TrailARN,
		Name:   *v.Name,
		Description: model.CloudTrailTrailDescription{
			Trail:                  v,
			TrailStatus:            *statusOutput,
			EventSelectors:         selectorOutput.EventSelectors,
			AdvancedEventSelectors: selectorOutput.AdvancedEventSelectors,
			Tags:                   tags,
		},
	}
	return resource, nil
}
func GetCloudTrailTrail(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	name := fields["name"]
	client := cloudtrail.NewFromConfig(cfg)

	out, err := client.GetTrail(ctx, &cloudtrail.GetTrailInput{
		Name: &name,
	})
	if err != nil {
		if isErr(err, "GetTrailNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}
	trail := out.Trail

	stsClient := sts.NewFromConfig(cfg)
	identity, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		if isErr(err, "GetCallerIdentityNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var trails []string
	// Ignore trails that don't belong this region (Based on steampipe)
	if !strings.EqualFold(*trail.HomeRegion, cfg.Region) {
		return nil, nil
	}
	if trail.TrailARN != nil {
		// Ignore trails that don't belong to this account (Based on steampipe)
		if aws.ToString(identity.Account) != arnToAccountId(*trail.TrailARN) {
			return nil, nil
		}

		trails = append(trails, *trail.TrailARN)
	} else if trail.Name != nil {
		trails = append(trails, *trail.Name)
	}

	output, err := client.DescribeTrails(ctx, &cloudtrail.DescribeTrailsInput{
		TrailNameList: trails,
	})
	if err != nil {
		if isErr(err, "DescribeTrailsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, v := range output.TrailList {
		resource, err := cloudTrailTrailHandle(ctx, cfg, v)
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

func CloudTrailChannel(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := cloudtrail.NewFromConfig(cfg)
	paginator := cloudtrail.NewListChannelsPaginator(client, &cloudtrail.ListChannelsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, channel := range page.Channels {
			output, err := client.GetChannel(ctx, &cloudtrail.GetChannelInput{
				Channel: channel.ChannelArn,
			})
			if err != nil {
				return nil, err
			}

			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *channel.ChannelArn,
				Name:   *channel.Name,
				Description: model.CloudTrailChannelDescription{
					Channel: *output,
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

func CloudTrailEventDataStore(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := cloudtrail.NewFromConfig(cfg)
	paginator := cloudtrail.NewListEventDataStoresPaginator(client, &cloudtrail.ListEventDataStoresInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, eventDataStore := range page.EventDataStores {
			resource, err := cloudTrailEventDataStoreHandle(ctx, cfg, eventDataStore)
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
func cloudTrailEventDataStoreHandle(ctx context.Context, cfg aws.Config, eventDataStore types.EventDataStore) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := cloudtrail.NewFromConfig(cfg)

	output, err := client.GetEventDataStore(ctx, &cloudtrail.GetEventDataStoreInput{
		EventDataStore: eventDataStore.EventDataStoreArn,
	})
	if err != nil {
		if isErr(err, "GetEventDataStoreNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *eventDataStore.EventDataStoreArn,
		Name:   *eventDataStore.Name,
		Description: model.CloudTrailEventDataStoreDescription{
			EventDataStore: *output,
		},
	}
	return resource, nil
}
func GetCloudTrailEventDataStore(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	eventDataStoreName := fields["name"]
	client := cloudtrail.NewFromConfig(cfg)

	out, err := client.ListEventDataStores(ctx, &cloudtrail.ListEventDataStoresInput{})
	if err != nil {
		if isErr(err, "ListEventDataStoresNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, evenDataStore := range out.EventDataStores {
		if !strings.EqualFold(eventDataStoreName, *evenDataStore.Name) {
			continue
		}

		resource, err := cloudTrailEventDataStoreHandle(ctx, cfg, evenDataStore)
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

func CloudTrailImport(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := cloudtrail.NewFromConfig(cfg)
	paginator := cloudtrail.NewListImportsPaginator(client, &cloudtrail.ListImportsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "UnsupportedOperationException") {
				return nil, nil
			}
			return nil, err
		}

		for _, trailImport := range page.Imports {
			output, err := client.GetImport(ctx, &cloudtrail.GetImportInput{
				ImportId: trailImport.ImportId,
			})
			if err != nil {
				if isErr(err, "UnsupportedOperationException") {
					continue
				}
				return nil, err
			}

			resource := models.Resource{
				Region: describeCtx.OGRegion,
				Name:   *trailImport.ImportId,
				Description: model.CloudTrailImportDescription{
					Import: *output,
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

func CloudTrailQuery(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := cloudtrail.NewFromConfig(cfg)
	paginator := cloudtrail.NewListEventDataStoresPaginator(client, &cloudtrail.ListEventDataStoresInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, eventDataStore := range page.EventDataStores {
			queryPaginator := cloudtrail.NewListQueriesPaginator(client, &cloudtrail.ListQueriesInput{
				EventDataStore: eventDataStore.EventDataStoreArn,
			})
			for queryPaginator.HasMorePages() {
				page, err := queryPaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}

				for _, query := range page.Queries {
					output, err := client.DescribeQuery(ctx, &cloudtrail.DescribeQueryInput{
						QueryId: query.QueryId,
					})
					if err != nil {
						return nil, err
					}

					resource := models.Resource{
						Region: describeCtx.OGRegion,
						Name:   *query.QueryId,
						Description: model.CloudTrailQueryDescription{
							Query:             *output,
							EventDataStoreARN: *eventDataStore.EventDataStoreArn,
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

func CloudTrailTrailEvent(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := cloudwatchlogs.NewFromConfig(cfg)

	logGroups, err := CloudWatchLogsLogGroup(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, logGroup := range logGroups {
		paginator := cloudwatchlogs.NewFilterLogEventsPaginator(client, &cloudwatchlogs.FilterLogEventsInput{
			LogGroupName: logGroup.Description.(model.CloudWatchLogsLogGroupDescription).LogGroup.LogGroupName,
			Limit:        aws.Int32(100),
		})

		// To avoid throttling, don't fetching everything. Only the first 5 pages!
		pageNo := 0
		for paginator.HasMorePages() && pageNo < 5 {
			pageNo++
			page, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}

			for _, event := range page.Events {
				resource := models.Resource{
					Region: describeCtx.OGRegion,
					ID:     *event.EventId,
					Description: model.CloudTrailTrailEventDescription{
						TrailEvent:   event,
						LogGroupName: logGroup.Name,
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

func arnToAccountId(arn string) string {
	if arn != "" {
		return strings.Split(arn, ":")[4]
	}

	return ""
}
