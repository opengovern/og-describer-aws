package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/drs"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func DRSSourceServer(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := drs.NewFromConfig(cfg)
	paginator := drs.NewDescribeSourceServersPaginator(client, &drs.DescribeSourceServersInput{
		MaxResults: aws.Int32(100),
	})

	var values []models.Resource
	pageNo := 0
	for paginator.HasMorePages() && pageNo < 5 {
		pageNo++
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if !isErr(err, "UninitializedAccountException") {
				return nil, err
			}
			continue
		}

		for _, v := range page.Items {
			launchConfiguration, err := client.GetLaunchConfiguration(ctx, &drs.GetLaunchConfigurationInput{
				SourceServerID: v.SourceServerID,
			})
			if err != nil {
				return nil, err
			}

			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *v.Arn,
				Name:   *v.SourceServerID,
				Description: model.DRSSourceServerDescription{
					SourceServer:        v,
					LaunchConfiguration: *launchConfiguration,
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

func DRSRecoveryInstance(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := drs.NewFromConfig(cfg)
	paginator := drs.NewDescribeRecoveryInstancesPaginator(client, &drs.DescribeRecoveryInstancesInput{
		MaxResults: aws.Int32(100),
	})

	var values []models.Resource
	pageNo := 0
	for paginator.HasMorePages() && pageNo < 5 {
		pageNo++
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if !isErr(err, "UninitializedAccountException") {
				return nil, err
			}
			continue
		}

		for _, v := range page.Items {
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *v.Arn,
				Name:   *v.RecoveryInstanceID,
				Description: model.DRSRecoveryInstanceDescription{
					RecoveryInstance: v,
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

func DRSJob(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := drs.NewFromConfig(cfg)
	paginator := drs.NewDescribeJobsPaginator(client, &drs.DescribeJobsInput{
		MaxResults: aws.Int32(100),
	})

	var values []models.Resource
	pageNo := 0
	for paginator.HasMorePages() && pageNo < 5 {
		pageNo++
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if !isErr(err, "UninitializedAccountException") {
				return nil, err
			}
			continue
		}

		for _, v := range page.Items {
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *v.Arn,
				ID:     *v.JobID,
				Description: model.DRSJobDescription{
					Job: v,
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

func DRSRecoverySnapshot(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := drs.NewFromConfig(cfg)
	paginator := drs.NewDescribeSourceServersPaginator(client, &drs.DescribeSourceServersInput{
		MaxResults: aws.Int32(100),
	})

	var values []models.Resource
	sourceServerPageNo := 0
	for paginator.HasMorePages() && sourceServerPageNo < 5 {
		sourceServerPageNo++
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if !isErr(err, "UninitializedAccountException") {
				return nil, err
			}
			continue
		}

		for _, sourceServer := range page.Items {
			recoverySnapshotPaginator := drs.NewDescribeRecoverySnapshotsPaginator(client, &drs.DescribeRecoverySnapshotsInput{
				MaxResults:     aws.Int32(100),
				SourceServerID: sourceServer.SourceServerID,
			})
			recoverySnapshotPageNo := 0
			for recoverySnapshotPaginator.HasMorePages() && recoverySnapshotPageNo < 5 {
				recoverySnapshotPageNo++

				recoverySnapshotPage, err := recoverySnapshotPaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}

				for _, recoverySnapshot := range recoverySnapshotPage.Items {
					resource := models.Resource{
						Region: describeCtx.OGRegion,
						ID:     *recoverySnapshot.SnapshotID,
						Description: model.DRSRecoverySnapshotDescription{
							RecoverySnapshot: recoverySnapshot,
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
