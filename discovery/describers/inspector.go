package describers

import (
	"context"

	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func InspectorAssessmentRun(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := inspector.NewFromConfig(cfg)
	paginator := inspector.NewListAssessmentRunsPaginator(client, &inspector.ListAssessmentRunsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		if len(page.AssessmentRunArns) == 0 {
			continue
		}

		assessmentRuns, err := client.DescribeAssessmentRuns(ctx, &inspector.DescribeAssessmentRunsInput{
			AssessmentRunArns: page.AssessmentRunArns,
		})
		if err != nil {
			return nil, err
		}

		for _, assessmentRun := range assessmentRuns.AssessmentRuns {
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				Name:   *assessmentRun.Name,
				ARN:    *assessmentRun.Arn,
				Description: model.InspectorAssessmentRunDescription{
					AssessmentRun: assessmentRun,
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

func InspectorAssessmentTarget(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := inspector.NewFromConfig(cfg)
	paginator := inspector.NewListAssessmentTargetsPaginator(client, &inspector.ListAssessmentTargetsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		if len(page.AssessmentTargetArns) == 0 {
			continue
		}

		assessmentTargets, err := client.DescribeAssessmentTargets(ctx, &inspector.DescribeAssessmentTargetsInput{
			AssessmentTargetArns: page.AssessmentTargetArns,
		})
		if err != nil {
			return nil, err
		}

		for _, assessmentTarget := range assessmentTargets.AssessmentTargets {
			resource := inspectorAssessmentTargetHandle(ctx, assessmentTarget)
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
func inspectorAssessmentTargetHandle(ctx context.Context, assessmentTarget types.AssessmentTarget) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	resource := models.Resource{
		Region: describeCtx.OGRegion,
		Name:   *assessmentTarget.Name,
		ARN:    *assessmentTarget.Arn,
		Description: model.InspectorAssessmentTargetDescription{
			AssessmentTarget: assessmentTarget,
		},
	}
	return resource
}
func GetInspectorAssessmentTarget(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	AssessmentTargetArn := fields["arn"]
	var values []models.Resource
	client := inspector.NewFromConfig(cfg)

	describeAssessments, err := client.DescribeAssessmentTargets(ctx, &inspector.DescribeAssessmentTargetsInput{
		AssessmentTargetArns: []string{AssessmentTargetArn},
	})
	if err != nil {
		if isErr(err, "DescribeAssessmentTargetsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, assessmentTarget := range describeAssessments.AssessmentTargets {
		resource := inspectorAssessmentTargetHandle(ctx, assessmentTarget)
		values = append(values, resource)
	}
	return values, nil
}

func InspectorAssessmentTemplate(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := inspector.NewFromConfig(cfg)
	paginator := inspector.NewListAssessmentTemplatesPaginator(client, &inspector.ListAssessmentTemplatesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		if len(page.AssessmentTemplateArns) == 0 {
			continue
		}

		assessmentTemplates, err := client.DescribeAssessmentTemplates(ctx, &inspector.DescribeAssessmentTemplatesInput{
			AssessmentTemplateArns: page.AssessmentTemplateArns,
		})
		if err != nil {
			return nil, err
		}

		for _, assessmentTemplate := range assessmentTemplates.AssessmentTemplates {
			tags, err := client.ListTagsForResource(ctx, &inspector.ListTagsForResourceInput{
				ResourceArn: assessmentTemplate.Arn,
			})
			if err != nil {
				return nil, err
			}

			eventSubscriptions, err := client.ListEventSubscriptions(ctx, &inspector.ListEventSubscriptionsInput{
				ResourceArn: assessmentTemplate.Arn,
			})
			if err != nil {
				return nil, err
			}

			resource := inspectorAssessmentTemplateHandle(ctx, assessmentTemplate, eventSubscriptions, tags)
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
func inspectorAssessmentTemplateHandle(ctx context.Context, assessmentTemplate types.AssessmentTemplate, eventSubscriptions *inspector.ListEventSubscriptionsOutput, tags *inspector.ListTagsForResourceOutput) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	resource := models.Resource{
		Region: describeCtx.OGRegion,
		Name:   *assessmentTemplate.Name,
		ARN:    *assessmentTemplate.Arn,
		Description: model.InspectorAssessmentTemplateDescription{
			AssessmentTemplate: assessmentTemplate,
			EventSubscriptions: eventSubscriptions.Subscriptions,
			Tags:               tags.Tags,
		},
	}
	return resource
}
func GetInspectorAssessmentTemplate(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	arn := fields["arn"]
	client := inspector.NewFromConfig(cfg)

	var values []models.Resource
	assessmentTemplates, err := client.DescribeAssessmentTemplates(ctx, &inspector.DescribeAssessmentTemplatesInput{
		AssessmentTemplateArns: []string{arn},
	})
	if err != nil {
		return nil, err
	}

	for _, assessmentTemplate := range assessmentTemplates.AssessmentTemplates {
		tags, err := client.ListTagsForResource(ctx, &inspector.ListTagsForResourceInput{
			ResourceArn: assessmentTemplate.Arn,
		})
		if err != nil {
			return nil, err
		}

		eventSubscriptions, err := client.ListEventSubscriptions(ctx, &inspector.ListEventSubscriptionsInput{
			ResourceArn: assessmentTemplate.Arn,
		})
		if err != nil {
			return nil, err
		}

		resource := inspectorAssessmentTemplateHandle(ctx, assessmentTemplate, eventSubscriptions, tags)
		values = append(values, resource)
	}
	return values, nil
}

func InspectorExclusion(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := inspector.NewFromConfig(cfg)
	paginator := inspector.NewListAssessmentRunsPaginator(client, &inspector.ListAssessmentRunsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, assessmentRun := range page.AssessmentRunArns {
			exclusionsPaginator := inspector.NewListExclusionsPaginator(client, &inspector.ListExclusionsInput{
				AssessmentRunArn: &assessmentRun,
			})

			for exclusionsPaginator.HasMorePages() {
				exclusionPage, err := exclusionsPaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}

				exclusions, err := client.DescribeExclusions(ctx, &inspector.DescribeExclusionsInput{
					ExclusionArns: exclusionPage.ExclusionArns,
				})
				if err != nil {
					return nil, err
				}

				for _, exclusion := range exclusions.Exclusions {
					resource := models.Resource{
						Region: describeCtx.OGRegion,
						Name:   *exclusion.Title,
						ARN:    *exclusion.Arn,
						Description: model.InspectorExclusionDescription{
							Exclusion: exclusion,
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

func InspectorFinding(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := inspector.NewFromConfig(cfg)
	paginator := inspector.NewListFindingsPaginator(client, &inspector.ListFindingsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		if len(page.FindingArns) == 0 {
			continue
		}

		findings, err := client.DescribeFindings(ctx, &inspector.DescribeFindingsInput{
			FindingArns: page.FindingArns,
		})
		if err != nil {
			return nil, err
		}

		for _, finding := range findings.Findings {
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				Name:   *finding.Title,
				ID:     *finding.Id,
				ARN:    *finding.Arn,
				Description: model.InspectorFindingDescription{
					Finding:     finding,
					FailedItems: findings.FailedItems,
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

func Inspector2Coverage(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := inspector2.NewFromConfig(cfg)
	paginator := inspector2.NewListCoveragePaginator(client, &inspector2.ListCoverageInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, coveredResource := range page.CoveredResources {
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ID:     *coveredResource.ResourceId,
				Description: model.Inspector2CoverageDescription{
					CoveredResource: coveredResource,
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

func Inspector2CoverageStatistic(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := inspector2.NewFromConfig(cfg)
	paginator := inspector2.NewListCoverageStatisticsPaginator(client, &inspector2.ListCoverageStatisticsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		resource := models.Resource{
			Region: describeCtx.OGRegion,
			Description: model.Inspector2CoverageStatisticDescription{
				TotalCounts: page.TotalCounts,
				Counts:      page.CountsByGroup,
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

	return values, nil
}

func Inspector2CoverageMember(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := inspector2.NewFromConfig(cfg)
	associated, err := Inspector2CoverageMemberHelper(ctx, cfg, client, true)
	if err != nil {
		return nil, err
	}
	notAssociated, err := Inspector2CoverageMemberHelper(ctx, cfg, client, false)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	values = append(values, associated...)
	for _, resource := range notAssociated {
		if !ContainsResource(resource, values) {
			values = append(values, resource)
		}
	}
	return values, nil
}

func ContainsResource(val models.Resource, values []models.Resource) bool {
	for _, v := range values {
		if reflect.DeepEqual(val, v) {
			return true
		}
	}
	return false
}

func Inspector2CoverageMemberHelper(ctx context.Context, cfg aws.Config, client *inspector2.Client, onlyAssociated bool) ([]models.Resource, error) {
	input := &inspector2.ListMembersInput{
		OnlyAssociated: &onlyAssociated,
	}
	paginator := inspector2.NewListMembersPaginator(client, input)
	describeCtx := model.GetDescribeContext(ctx)

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, member := range page.Members {
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				Description: model.Inspector2MemberDescription{
					Member:         member,
					OnlyAssociated: onlyAssociated,
				},
			}

			values = append(values, resource)
		}

	}
	return values, nil
}

func Inspector2Finding(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := inspector2.NewFromConfig(cfg)
	paginator := inspector2.NewListFindingsPaginator(client, &inspector2.ListFindingsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		for _, finding := range page.Findings {
			if err != nil {
				return nil, err
			}
			for _, v := range finding.Resources {
				resource := models.Resource{
					Region: describeCtx.OGRegion,
					Description: model.Inspector2FindingDescription{
						Finding:  finding,
						Resource: v,
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
