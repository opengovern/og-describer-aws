package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func GetAccessAnalyzerAnalyzer(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	analyzerName := fields["name"]
	client := accessanalyzer.NewFromConfig(cfg)
	v, err := client.GetAnalyzer(ctx, &accessanalyzer.GetAnalyzerInput{
		AnalyzerName: &analyzerName,
	})
	if err != nil {
		return nil, err
	}

	findings, err := getAnalyzerFindings(ctx, client, v.Analyzer.Arn)
	if err != nil {
		return nil, err
	}

	return []models.Resource{
		{
			Region: describeCtx.OGRegion,
			ARN:    *v.Analyzer.Arn,
			Name:   *v.Analyzer.Name,
			Description: model.AccessAnalyzerAnalyzerDescription{
				Analyzer: *v.Analyzer,
				Findings: findings,
			},
		}}, nil
}

func AccessAnalyzerAnalyzer(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := accessanalyzer.NewFromConfig(cfg)
	paginator := accessanalyzer.NewListAnalyzersPaginator(client, &accessanalyzer.ListAnalyzersInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Analyzers {
			findings, err := getAnalyzerFindings(ctx, client, v.Arn)
			if err != nil {
				return nil, err
			}
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *v.Arn,
				Name:   *v.Name,
				Description: model.AccessAnalyzerAnalyzerDescription{
					Analyzer: v,
					Findings: findings,
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
		}
	}

	return values, nil
}

func getAnalyzerFindings(ctx context.Context, client *accessanalyzer.Client, analyzerArn *string) ([]types.FindingSummary, error) {
	paginator := accessanalyzer.NewListFindingsPaginator(client, &accessanalyzer.ListFindingsInput{
		AnalyzerArn: analyzerArn,
	})

	var findings []types.FindingSummary
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		findings = append(findings, page.Findings...)
	}

	return findings, nil
}

func AccessAnalyzerAnalyzerFinding(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := accessanalyzer.NewFromConfig(cfg)
	paginator := accessanalyzer.NewListAnalyzersPaginator(client, &accessanalyzer.ListAnalyzersInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Analyzers {
			findings, err := getAnalyzerFindings(ctx, client, v.Arn)
			if err != nil {
				return nil, err
			}
			for _, finding := range findings {
				resource := models.Resource{
					Region: describeCtx.OGRegion,
					ID:     *finding.Id,
					Description: model.AccessAnalyzerAnalyzerFindingDescription{
						AnalyzerArn: *v.Arn,
						Finding:     finding,
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
			}
		}
	}

	return values, nil
}
