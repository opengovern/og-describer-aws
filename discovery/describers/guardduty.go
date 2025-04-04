package describers

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func GuardDutyFinding(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	var values []models.Resource

	client := guardduty.NewFromConfig(cfg)

	dpaginator := guardduty.NewListDetectorsPaginator(client, &guardduty.ListDetectorsInput{})
	for dpaginator.HasMorePages() {
		dpage, err := dpaginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, detectorId := range dpage.DetectorIds {
			// This prevents Implicit memory aliasing in for loop
			detectorId := detectorId

			paginator := guardduty.NewListFindingsPaginator(client, &guardduty.ListFindingsInput{
				DetectorId: &detectorId,
			})

			for paginator.HasMorePages() {
				page, err := paginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}

				findings, err := client.GetFindings(ctx, &guardduty.GetFindingsInput{
					DetectorId: &detectorId,
					FindingIds: page.FindingIds,
				})
				if err != nil {
					return nil, err
				}

				for _, item := range findings.Findings {
					resource := models.Resource{
						Region:  describeCtx.OGRegion,
						ARN:     *item.Arn,
						Name:    *item.Id,
						Account: describeCtx.AccountID,
						Description: model.GuardDutyFindingDescription{
							Finding: item,
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

func GuardDutyDetector(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := guardduty.NewFromConfig(cfg)

	var values []models.Resource
	paginator := guardduty.NewListDetectorsPaginator(client, &guardduty.ListDetectorsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, id := range page.DetectorIds {
			id := id
			out, err := client.GetDetector(ctx, &guardduty.GetDetectorInput{
				DetectorId: &id,
			})
			if err != nil {
				return nil, err
			}

			resource := guardDutyDetectorHandle(ctx, out, id)
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
func guardDutyDetectorHandle(ctx context.Context, out *guardduty.GetDetectorOutput, id string) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)

	arn := "arn:" + describeCtx.Partition + ":guardduty:" + describeCtx.Region + ":" + describeCtx.AccountID + ":detector/" + id
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     arn,
		Account: describeCtx.AccountID,
		Name:    id,
		Description: model.GuardDutyDetectorDescription{
			DetectorId: id,
			Detector:   out,
		},
	}
	return resource
}
func GetGuardDutyDetector(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	detectorId := fields["id"]
	client := guardduty.NewFromConfig(cfg)
	var values []models.Resource

	out, err := client.GetDetector(ctx, &guardduty.GetDetectorInput{
		DetectorId: &detectorId,
	})
	if err != nil {
		if isErr(err, "GuardDutyDetectorNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	resource := guardDutyDetectorHandle(ctx, out, detectorId)
	values = append(values, resource)
	return values, nil
}

func GuardDutyFilter(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := guardduty.NewFromConfig(cfg)
	paginator := guardduty.NewListDetectorsPaginator(client, &guardduty.ListDetectorsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, detectorId := range page.DetectorIds {
			filterPaginator := guardduty.NewListFiltersPaginator(client, &guardduty.ListFiltersInput{
				DetectorId: &detectorId,
			})
			for filterPaginator.HasMorePages() {
				filterPage, err := filterPaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, filter := range filterPage.FilterNames {
					filterOutput, err := client.GetFilter(ctx, &guardduty.GetFilterInput{
						DetectorId: &detectorId,
						FilterName: &filter,
					})
					if err != nil {
						if isErr(err, "GetFilterNotFound") || isErr(err, "InvalidParameterValue") {
							return nil, nil
						}
						return nil, err
					}

					resource := guardDutyFilterHandle(ctx, filterOutput, detectorId, filter)
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
func guardDutyFilterHandle(ctx context.Context, filterOutput *guardduty.GetFilterOutput, detectorId string, filter string) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)

	arn := fmt.Sprintf("arn:%s:guardduty:%s:%s:detector/%s/filter/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, detectorId, filter)
	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    arn,
		Name:   *filterOutput.Name,
		Description: model.GuardDutyFilterDescription{
			Filter:     *filterOutput,
			DetectorId: detectorId,
		},
	}
	return resource
}
func GetGuardDutyFilter(ctx context.Context, cfg aws.Config, field map[string]string) ([]models.Resource, error) {
	detectorId := field["id"]

	client := guardduty.NewFromConfig(cfg)
	out, err := client.ListFilters(ctx, &guardduty.ListFiltersInput{
		DetectorId: &detectorId,
	})
	if err != nil {
		if isErr(err, "ListFiltersNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, filter := range out.FilterNames {
		filterOutput, err := client.GetFilter(ctx, &guardduty.GetFilterInput{
			DetectorId: &detectorId,
			FilterName: &filter,
		})
		if err != nil {
			if isErr(err, "GetFilterNotFound") || isErr(err, "InvalidParameterValue") {
				return nil, nil
			}
			return nil, err
		}

		resource := guardDutyFilterHandle(ctx, filterOutput, detectorId, filter)
		values = append(values, resource)
	}
	return values, nil
}

func GuardDutyIPSet(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := guardduty.NewFromConfig(cfg)
	paginator := guardduty.NewListDetectorsPaginator(client, &guardduty.ListDetectorsInput{})
	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, detectorId := range page.DetectorIds {
			ipSetsPaginator := guardduty.NewListIPSetsPaginator(client, &guardduty.ListIPSetsInput{
				DetectorId: &detectorId,
			})

			for ipSetsPaginator.HasMorePages() {
				ipSetPage, err := ipSetsPaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, ipSetId := range ipSetPage.IpSetIds {
					ipSetOutput, err := client.GetIPSet(ctx, &guardduty.GetIPSetInput{
						DetectorId: &detectorId,
						IpSetId:    &ipSetId,
					})
					if err != nil {
						return nil, err
					}

					resource := guardDutyIPSetHandle(ctx, ipSetOutput, ipSetId, detectorId)
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
func guardDutyIPSetHandle(ctx context.Context, ipSetOutput *guardduty.GetIPSetOutput, ipSetId string, detectorId string) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)

	arn := fmt.Sprintf("arn:%s:guardduty:%s:%s:detector/%s/ipset/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, detectorId, ipSetId)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     arn,
		Account: describeCtx.AccountID,
		Name:    *ipSetOutput.Name,
		Description: model.GuardDutyIPSetDescription{
			IPSet:      *ipSetOutput,
			DetectorId: detectorId,
			IPSetId:    ipSetId,
		},
	}
	return resource
}
func GetGuardDutyIPSet(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	detectorId := fields["id"]

	client := guardduty.NewFromConfig(cfg)
	out, err := client.ListIPSets(ctx, &guardduty.ListIPSetsInput{
		DetectorId: &detectorId,
	})
	if err != nil {
		if isErr(err, "ListIPSetsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, ipSetId := range out.IpSetIds {
		ipSetOutput, err := client.GetIPSet(ctx, &guardduty.GetIPSetInput{
			DetectorId: &detectorId,
			IpSetId:    &ipSetId,
		})
		if err != nil {
			if isErr(err, "GetIPSetNotFound") || isErr(err, "InvalidParameterValue") {
				return nil, nil
			}
			return nil, err
		}

		values = append(values, guardDutyIPSetHandle(ctx, ipSetOutput, ipSetId, detectorId))
	}
	return values, nil
}

func GuardDutyMember(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := guardduty.NewFromConfig(cfg)
	paginator := guardduty.NewListDetectorsPaginator(client, &guardduty.ListDetectorsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, detectorId := range page.DetectorIds {
			membersPaginator := guardduty.NewListMembersPaginator(client, &guardduty.ListMembersInput{
				DetectorId: &detectorId,
			})

			for membersPaginator.HasMorePages() {
				membersPage, err := membersPaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}

				for _, member := range membersPage.Members {
					resource := guardDutyMemberHandle(ctx, member)
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
func guardDutyMemberHandle(ctx context.Context, member types.Member) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		Name:    *member.AccountId,
		Account: describeCtx.AccountID,
		Description: model.GuardDutyMemberDescription{
			Member: member,
		},
	}
	return resource
}
func GetGuardDutyMember(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	memberId := fields["id"]
	client := guardduty.NewFromConfig(cfg)

	members, err := client.ListMembers(ctx, &guardduty.ListMembersInput{
		DetectorId: &memberId,
	})
	if err != nil {
		if isErr(err, "ListMembersNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, member := range members.Members {
		resource := guardDutyMemberHandle(ctx, member)
		values = append(values, resource)
	}
	return values, nil
}

func GuardDutyPublishingDestination(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := guardduty.NewFromConfig(cfg)
	paginator := guardduty.NewListDetectorsPaginator(client, &guardduty.ListDetectorsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, detectorId := range page.DetectorIds {
			publishingDestinationsPaginator := guardduty.NewListPublishingDestinationsPaginator(client, &guardduty.ListPublishingDestinationsInput{
				DetectorId: &detectorId,
			})

			for publishingDestinationsPaginator.HasMorePages() {
				publishingDestinationsPage, err := publishingDestinationsPaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, destination := range publishingDestinationsPage.Destinations {

					destinationOutput, err := client.DescribePublishingDestination(ctx, &guardduty.DescribePublishingDestinationInput{
						DestinationId: destination.DestinationId,
						DetectorId:    &detectorId,
					})
					if err != nil {
						return nil, err
					}

					resource := guardDutyPublishingDestinationHandle(ctx, detectorId, destinationOutput, destination)
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
func guardDutyPublishingDestinationHandle(ctx context.Context, detectorId string, destinationOutput *guardduty.DescribePublishingDestinationOutput, destination types.Destination) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)

	arn := fmt.Sprintf("arn:%s:guardduty:%s:%s:detector/%s/publishingDestination/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, detectorId, *destination.DestinationId)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     arn,
		ID:      *destinationOutput.DestinationId,
		Account: describeCtx.AccountID,
		Description: model.GuardDutyPublishingDestinationDescription{
			PublishingDestination: *destinationOutput,
			DetectorId:            detectorId,
		},
	}
	return resource
}
func GetGuardDutyPublishingDestination(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	detectorId := fields["id"]

	client := guardduty.NewFromConfig(cfg)
	publishingDestinations, err := client.ListPublishingDestinations(ctx, &guardduty.ListPublishingDestinationsInput{
		DetectorId: &detectorId,
	})
	if err != nil {
		if isErr(err, "ListPublishingDestinationsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, destination := range publishingDestinations.Destinations {
		destinationOutput, err := client.DescribePublishingDestination(ctx, &guardduty.DescribePublishingDestinationInput{
			DestinationId: destination.DestinationId,
			DetectorId:    &detectorId,
		})
		if err != nil {
			if isErr(err, "DescribePublishingDestinationNotFound") || isErr(err, "InvalidParameterValue") {
				return nil, nil
			}
			return nil, err
		}

		resource := guardDutyPublishingDestinationHandle(ctx, detectorId, destinationOutput, destination)
		values = append(values, resource)
	}
	return values, nil
}

func GuardDutyThreatIntelSet(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := guardduty.NewFromConfig(cfg)
	paginator := guardduty.NewListDetectorsPaginator(client, &guardduty.ListDetectorsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, detectorId := range page.DetectorIds {
			threatIntelSetsPaginator := guardduty.NewListThreatIntelSetsPaginator(client, &guardduty.ListThreatIntelSetsInput{
				DetectorId: &detectorId,
			})

			for threatIntelSetsPaginator.HasMorePages() {
				threatIntelSetsPage, err := threatIntelSetsPaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, threatIntelSetId := range threatIntelSetsPage.ThreatIntelSetIds {

					threatIntelSetOutput, err := client.GetThreatIntelSet(ctx, &guardduty.GetThreatIntelSetInput{
						DetectorId:       &detectorId,
						ThreatIntelSetId: &threatIntelSetId,
					})
					if err != nil {
						return nil, err
					}

					resource := guardDutyThreatIntelSetHandle(ctx, threatIntelSetOutput, detectorId, threatIntelSetId)
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
func guardDutyThreatIntelSetHandle(ctx context.Context, threatIntelSetOutput *guardduty.GetThreatIntelSetOutput, detectorId string, threatIntelSetId string) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	arn := fmt.Sprintf("arn:%s:guardduty:%s:%s:detector/%s/threatintelset/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, detectorId, threatIntelSetId)

	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     arn,
		Name:    *threatIntelSetOutput.Name,
		Account: describeCtx.AccountID,
		Description: model.GuardDutyThreatIntelSetDescription{
			ThreatIntelSet:   *threatIntelSetOutput,
			DetectorId:       detectorId,
			ThreatIntelSetID: threatIntelSetId,
		},
	}
	return resource
}
func GetGuardDutyThreatIntelSet(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	detectorId := fields["id"]

	client := guardduty.NewFromConfig(cfg)
	threatIntelSets, err := client.ListThreatIntelSets(ctx, &guardduty.ListThreatIntelSetsInput{
		DetectorId: &detectorId,
	})
	if err != nil {
		if isErr(err, "ListThreatIntelSetsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, threatIntelSetId := range threatIntelSets.ThreatIntelSetIds {
		threatIntelSetOutput, err := client.GetThreatIntelSet(ctx, &guardduty.GetThreatIntelSetInput{
			DetectorId:       &detectorId,
			ThreatIntelSetId: &threatIntelSetId,
		})
		if err != nil {
			if isErr(err, "GetThreatIntelSetNotFound") || isErr(err, "InvalidParameterValue") {
				return nil, nil
			}
			return nil, err
		}

		resource := guardDutyThreatIntelSetHandle(ctx, threatIntelSetOutput, detectorId, threatIntelSetId)
		values = append(values, resource)
	}
	return values, nil
}
