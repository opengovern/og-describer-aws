package describers

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation"
	lakeformationTypes "github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func GlueCatalogDatabase(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := glue.NewFromConfig(cfg)
	paginator := glue.NewGetDatabasesPaginator(client, &glue.GetDatabasesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, database := range page.DatabaseList {
			resource := glueCatalogDatabaseHandle(ctx, database)
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
func glueCatalogDatabaseHandle(ctx context.Context, database types.Database) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	arn := fmt.Sprintf("arn:aws:glue:%s:%s:database/%s", describeCtx.Region, describeCtx.AccountID, *database.Name)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		Name:    *database.Name,
		Account: describeCtx.AccountID,
		ARN:     arn,
		Description: model.GlueCatalogDatabaseDescription{
			Database: database,
		},
	}
	return resource
}
func GetGlueCatalogDatabase(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	name := fields["name"]
	client := glue.NewFromConfig(cfg)
	var values []models.Resource

	out, err := client.GetDatabase(ctx, &glue.GetDatabaseInput{
		Name: &name,
	})
	if err != nil {
		if isErr(err, "GetDatabaseNotFound") || isErr(err, "invalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}
	database := out.Database

	values = append(values, glueCatalogDatabaseHandle(ctx, *database))
	return values, nil
}

func GlueCatalogTable(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := glue.NewFromConfig(cfg)
	paginator := glue.NewGetDatabasesPaginator(client, &glue.GetDatabasesInput{})

	lakeformationClient := lakeformation.NewFromConfig(cfg)

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, database := range page.DatabaseList {
			tablePaginator := glue.NewGetTablesPaginator(client, &glue.GetTablesInput{
				DatabaseName: database.Name,
				CatalogId:    database.CatalogId,
			})
			for tablePaginator.HasMorePages() {
				tablePage, err := tablePaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, table := range tablePage.TableList {
					resource, err := glueCatalogTableHandle(ctx, lakeformationClient, table, *database.Name)
					if err != nil {
						return nil, err
					}

					if stream != nil {
						if err := (*stream)(*resource); err != nil {
							return nil, err
						}
					} else {
						values = append(values, *resource)
					}
				}
			}
		}
	}

	return values, nil
}
func glueCatalogTableHandle(ctx context.Context, client *lakeformation.Client, table types.Table, databaseName string) (*models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	arn := fmt.Sprintf("arn:aws:glue:%s:%s:table/%s/%s", describeCtx.Region, describeCtx.AccountID, databaseName, *table.Name)

	if table.ViewOriginalText != nil && len(*table.ViewOriginalText) > 5000 {
		v := *table.ViewOriginalText
		table.ViewOriginalText = aws.String(v[:5000])
	}

	if table.ViewExpandedText != nil && len(*table.ViewExpandedText) > 5000 {
		v := *table.ViewExpandedText
		table.ViewExpandedText = aws.String(v[:5000])
	}

	params := &lakeformation.GetResourceLFTagsInput{
		CatalogId:          table.CatalogId,
		ShowAssignedLFTags: aws.Bool(true),
		Resource: &lakeformationTypes.Resource{
			Table: &lakeformationTypes.TableResource{
				CatalogId:    table.CatalogId,
				DatabaseName: table.DatabaseName,
				Name:         table.Name,
			},
		},
	}

	var lfTableTags []lakeformationTypes.LFTagPair
	lftags, err := client.GetResourceLFTags(ctx, params)
	if err != nil {
		return nil, err
	}
	if lftags != nil && len(lftags.LFTagsOnTable) > 0 {
		lfTableTags = lftags.LFTagsOnTable
	}

	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		Name:    *table.Name,
		Account: describeCtx.AccountID,
		ARN:     arn,
		Description: model.GlueCatalogTableDescription{
			Table:  table,
			LfTags: lfTableTags,
		},
	}
	return &resource, nil
}
func GetGlueCatalogTable(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	name := fields["databaseName"]
	catalogId := fields["catalogId"]
	client := glue.NewFromConfig(cfg)
	lakeformationClient := lakeformation.NewFromConfig(cfg)

	tables, err := client.GetTables(ctx, &glue.GetTablesInput{
		DatabaseName: &name,
		CatalogId:    &catalogId,
	})
	if err != nil {
		if isErr(err, "GetTablesNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, table := range tables.TableList {
		value, err := glueCatalogTableHandle(ctx, lakeformationClient, table, name)
		if err != nil {
			return nil, err
		}
		values = append(values, *value)
	}
	return values, nil
}

func GlueConnection(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := glue.NewFromConfig(cfg)
	paginator := glue.NewGetConnectionsPaginator(client, &glue.GetConnectionsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, connection := range page.ConnectionList {
			resource := glueConnectionHandle(ctx, connection)
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
func glueConnectionHandle(ctx context.Context, connection types.Connection) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	arn := fmt.Sprintf("arn:aws:glue:%s:%s:connection/%s", describeCtx.Region, describeCtx.AccountID, *connection.Name)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		Name:    *connection.Name,
		ARN:     arn,
		Account: describeCtx.AccountID,
		Description: model.GlueConnectionDescription{
			Connection: connection,
		},
	}
	return resource
}
func GetGlueConnection(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	name := fields["name"]

	client := glue.NewFromConfig(cfg)

	out, err := client.GetConnection(ctx, &glue.GetConnectionInput{
		Name: &name,
	})
	if err != nil {
		if isErr(err, "GetConnectionNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	connection := out.Connection

	values = append(values, glueConnectionHandle(ctx, *connection))
	return values, nil
}

func GlueCrawler(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := glue.NewFromConfig(cfg)
	paginator := glue.NewGetCrawlersPaginator(client, &glue.GetCrawlersInput{})
	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, crawler := range page.Crawlers {
			resource := glueCrawlerHandle(ctx, crawler)
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
func glueCrawlerHandle(ctx context.Context, crawler types.Crawler) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	arn := fmt.Sprintf("arn:aws:glue:%s:%s:crawler/%s", describeCtx.Region, describeCtx.AccountID, *crawler.Name)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		Account: describeCtx.AccountID,
		Name:    *crawler.Name,
		ARN:     arn,
		Description: model.GlueCrawlerDescription{
			Crawler: crawler,
		},
	}
	return resource
}
func GetGlueCrawler(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	name := fields["name"]

	client := glue.NewFromConfig(cfg)

	out, err := client.GetCrawler(ctx, &glue.GetCrawlerInput{
		Name: &name,
	})
	if err != nil {
		if isErr(err, "GetCrawlerNotfound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}
	crawler := out.Crawler

	var values []models.Resource
	values = append(values, glueCrawlerHandle(ctx, *crawler))
	return values, nil
}

func GlueDataCatalogEncryptionSettings(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := glue.NewFromConfig(cfg)

	settings, err := client.GetDataCatalogEncryptionSettings(ctx, &glue.GetDataCatalogEncryptionSettingsInput{})
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	resource := glueDataCatalogEncryptionSettingsHandle(ctx, settings)
	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	} else {
		values = append(values, resource)
	}

	return values, nil
}
func glueDataCatalogEncryptionSettingsHandle(ctx context.Context, settings *glue.GetDataCatalogEncryptionSettingsOutput) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		Account: describeCtx.AccountID,
		Description: model.GlueDataCatalogEncryptionSettingsDescription{
			DataCatalogEncryptionSettings: *settings.DataCatalogEncryptionSettings,
		},
	}
	return resource
}
func GetGlueDataCatalogEncryptionSettings(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	catalogId := fields["catalogId"]

	client := glue.NewFromConfig(cfg)
	settings, err := client.GetDataCatalogEncryptionSettings(ctx, &glue.GetDataCatalogEncryptionSettingsInput{
		CatalogId: &catalogId,
	})
	if err != nil {
		if isErr(err, "GetDataCatalogEncryptionSettingsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	values = append(values, glueDataCatalogEncryptionSettingsHandle(ctx, settings))
	return values, nil
}

func GlueDataQualityRuleset(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := glue.NewFromConfig(cfg)
	paginator := glue.NewListDataQualityRulesetsPaginator(client, &glue.ListDataQualityRulesetsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, listRuleset := range page.Rulesets {
			ruleset, err := client.GetDataQualityRuleset(ctx, &glue.GetDataQualityRulesetInput{
				Name: listRuleset.Name,
			})
			if err != nil {
				return nil, err
			}

			resource := models.Resource{
				Region:  describeCtx.OGRegion,
				Name:    *listRuleset.Name,
				Account: describeCtx.AccountID,
				Description: model.GlueDataQualityRulesetDescription{
					DataQualityRuleset: *ruleset,
					RulesetRuleCount:   listRuleset.RuleCount,
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

func GlueDevEndpoint(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := glue.NewFromConfig(cfg)
	paginator := glue.NewGetDevEndpointsPaginator(client, &glue.GetDevEndpointsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, devEndpoint := range page.DevEndpoints {
			resource := glueDevEndpointHandle(ctx, devEndpoint)
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
func glueDevEndpointHandle(ctx context.Context, devEndpoint types.DevEndpoint) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	arn := fmt.Sprintf("arn:aws:glue:%s:%s:devEndpoint/%s", describeCtx.Region, describeCtx.AccountID, *devEndpoint.EndpointName)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		Name:    *devEndpoint.EndpointName,
		ARN:     arn,
		Account: describeCtx.AccountID,
		Description: model.GlueDevEndpointDescription{
			DevEndpoint: devEndpoint,
		},
	}
	return resource
}
func GetGlueDevEndpoint(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	endpointName := fields["name"]
	client := glue.NewFromConfig(cfg)

	out, err := client.GetDevEndpoint(ctx, &glue.GetDevEndpointInput{
		EndpointName: &endpointName,
	})
	if err != nil {
		if isErr(err, "GetDevEndpointNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	values = append(values, glueDevEndpointHandle(ctx, *out.DevEndpoint))
	return values, nil
}

func GlueJob(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := glue.NewFromConfig(cfg)
	paginator := glue.NewGetJobsPaginator(client, &glue.GetJobsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, job := range page.Jobs {
			resource := glueJobHandle(ctx, cfg, job)
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
func glueJobHandle(ctx context.Context, cfg aws.Config, job types.Job) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	client := glue.NewFromConfig(cfg)

	arn := fmt.Sprintf("arn:aws:glue:%s:%s:job/%s", describeCtx.Region, describeCtx.AccountID, *job.Name)

	bookmark, err := client.GetJobBookmark(ctx, &glue.GetJobBookmarkInput{
		JobName: job.Name,
	})
	if err != nil {
		bookmark = &glue.GetJobBookmarkOutput{}
	}

	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		Name:    *job.Name,
		Account: describeCtx.AccountID,
		ARN:     arn,
		Description: model.GlueJobDescription{
			Job:      job,
			Bookmark: bookmark.JobBookmarkEntry,
		},
	}
	return resource
}
func GetGlueJob(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	jobName := fields["name"]

	client := glue.NewFromConfig(cfg)

	var values []models.Resource
	out, err := client.GetJob(ctx, &glue.GetJobInput{
		JobName: &jobName,
	})
	if err != nil {
		if isErr(err, "GetJobNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	values = append(values, glueJobHandle(ctx, cfg, *out.Job))
	return values, nil
}

func GlueSecurityConfiguration(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := glue.NewFromConfig(cfg)
	paginator := glue.NewGetSecurityConfigurationsPaginator(client, &glue.GetSecurityConfigurationsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, securityConfiguration := range page.SecurityConfigurations {
			resource := glueSecurityConfigurationHandle(ctx, securityConfiguration)
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
func glueSecurityConfigurationHandle(ctx context.Context, securityConfiguration types.SecurityConfiguration) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	arn := fmt.Sprintf("arn:aws:glue:%s:%s:security-configuration/%s", describeCtx.Region, describeCtx.AccountID, *securityConfiguration.Name)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		Account: describeCtx.AccountID,
		Name:    *securityConfiguration.Name,
		ARN:     arn,
		Description: model.GlueSecurityConfigurationDescription{
			SecurityConfiguration: securityConfiguration,
		},
	}
	return resource
}
func GetGlueSecurityConfiguration(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	name := fields["name"]
	client := glue.NewFromConfig(cfg)

	out, err := client.GetSecurityConfiguration(ctx, &glue.GetSecurityConfigurationInput{
		Name: &name,
	})
	if err != nil {
		if isErr(err, "GetSecurityConfigurationNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	values = append(values, glueSecurityConfigurationHandle(ctx, *out.SecurityConfiguration))
	return values, nil
}
