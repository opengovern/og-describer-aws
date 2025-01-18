package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/discovery/pkg/es"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsAthenaWorkGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_athena_workgroup",
		Description: "AWS Athena Workgroup",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetAthenaWorkGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAthenaWorkGroup,
		},
		Columns: awsOgRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The workgroup name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WorkGroup.Name"),
			},
			{
				Name:        "description",
				Description: "The workgroup description.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WorkGroup.Description"),
			},
			{
				Name:        "creation_time",
				Description: "The date and time the workgroup was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.WorkGroup.CreationTime"),
			},
			{
				Name:        "state",
				Description: "The state of the workgroup.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WorkGroup.State").Transform(transform.ToString),
			},
			{
				Name:        "additional_configuration",
				Description: "Specifies a user defined JSON string that is passed to the notebook engine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.WorkGroup.Configuration.AdditionalConfiguration"),
			},
			{
				Name:        "bytes_scanned_cutoff_per_query",
				Description: "The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.WorkGroup.Configuration.BytesScannedCutoffPerQuery"),
			},
			{
				Name:        "customer_content_kms_key",
				Description: "Specifies the KMS key that is used to encrypt the user's data stores in Athena.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WorkGroup.Configuration.CustomerContentEncryptionConfiguration.KmsKey"),
			},
			{
				Name:        "enforce_workgroup_configuration",
				Description: "If set to \"true\", the settings for the workgroup override client-side settings.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.WorkGroup.Configuration.EnforceWorkGroupConfiguration"),
			},
			{
				Name:        "effective_engine_version",
				Description: "The engine version on which the query runs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WorkGroup.Configuration.EngineVersion.EffectiveEngineVersion"),
			},
			{
				Name:        "selected_engine_version",
				Description: "The engine version requested by the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WorkGroup.Configuration.EngineVersion.SelectedEngineVersion"),
			},
			{
				Name:        "execution_role",
				Description: "Role used in a notebook session for accessing the user's resources.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WorkGroup.Configuration.ExecutionRole"),
			},
			{
				Name:        "publish_cloudwatch_metrics_enabled",
				Description: "Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.WorkGroup.Configuration.PublishCloudWatchMetricsEnabled"),
			},
			{
				Name:        "requester_pays_enabled",
				Description: "If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.WorkGroup.Configuration.RequesterPaysEnabled"),
			},
			{
				Name:        "s3_acl_option",
				Description: "The Amazon S3 canned ACL that Athena should specify when storing query results.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WorkGroup.Configuration.ResultConfiguration.AclConfiguration.S3AclOption"),
			},
			{
				Name:        "encryption_option",
				Description: "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE_S3), server-side encryption with KMS-managed keys (SSE_KMS), or client-side encryption with KMS-managed keys (CSE_KMS) is used.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WorkGroup.Configuration.ResultConfiguration.EncryptionConfiguration.EncryptionOption"),
			},
			{
				Name:        "result_configuration_kms_key",
				Description: "For SSE_KMS and CSE_KMS, this is the KMS key ARN or ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WorkGroup.Configuration.ResultConfiguration.EncryptionConfiguration.KmsKey"),
			},
			{
				Name:        "expected_bucket_owner",
				Description: "The Amazon Web Services account ID that you expect to be the owner of the Amazon S3 bucket specified by ResultConfiguration$OutputLocation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WorkGroup.Configuration.ResultConfiguration.ExpectedBucketOwner"),
			},
			{
				Name:        "output_location",
				Description: "The location in Amazon S3 where your query results are stored.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.WorkGroup.Configuration.ResultConfiguration.OutputLocation"),
			},
		}),
	}
}

//// LIST FUNCTION

func listAwsAthenaWorkGroups(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	svc, err := AthenaClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("aws_athena_query_execution.listAwsAthenaWorkGroups", "connection_error", err)
		return nil, err
	}
	if svc == nil {
		// Unsupported region, return no data
		return nil, nil
	}

	maxResults := int32(50)

	// Reduce the basic request limit down if the user has only requested a small number of rows
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxResults {
			maxResults = int32(limit)
		}
	}

	listWorkGroupsInput := athena.ListWorkGroupsInput{MaxResults: &maxResults}
	workGroupsPaginator := athena.NewListWorkGroupsPaginator(svc, &listWorkGroupsInput, func(o *athena.ListWorkGroupsPaginatorOptions) {
		o.Limit = maxResults
		o.StopOnDuplicateToken = true
	})

	for workGroupsPaginator.HasMorePages() {
		output, err := workGroupsPaginator.NextPage(ctx)
		if err != nil {
			plugin.Logger(ctx).Error("aws_athena_query_execution.listAwsAthenaWorkGroups", "api_error", err)
			return nil, err
		}

		for _, workGroupSummary := range output.WorkGroups {
			configuration := types.WorkGroupConfiguration{EngineVersion: workGroupSummary.EngineVersion}
			workGroup := types.WorkGroup{
				Name:          workGroupSummary.Name,
				CreationTime:  workGroupSummary.CreationTime,
				Description:   workGroupSummary.Description,
				State:         workGroupSummary.State,
				Configuration: &configuration,
			}

			d.StreamListItem(ctx, workGroup)

			// Context may get cancelled due to manual cancellation or if the limit has been reached
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getAwsAthenaWorkGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	var name string
	if h.Item != nil {
		name = *h.Item.(types.WorkGroup).Name
	} else {
		name = d.EqualsQuals["name"].GetStringValue()
	}

	// Empty input check
	if name == "" {
		return nil, nil
	}

	// Create Session
	svc, err := AthenaClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("aws_athena_query_execution.getAwsAthenaWorkGroup", "connection_error", err)
		return nil, err
	}
	if svc == nil {
		// Unsupported region, return no data
		return nil, nil
	}

	// Build params
	params := &athena.GetWorkGroupInput{WorkGroup: aws.String(name)}

	rowData, err := svc.GetWorkGroup(ctx, params)
	if err != nil {
		plugin.Logger(ctx).Error("aws_athena_query_execution.getAwsAthenaWorkGroup", "api_error", err)
		return nil, err
	}

	return rowData.WorkGroup, nil
}
