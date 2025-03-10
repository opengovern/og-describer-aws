package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/discovery/pkg/es"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsS3MultiRegionAccessPoint(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_s3_multi_region_access_point",
		Description: "AWS S3 Multi Region Access Point",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListS3MultiRegionAccessPoint,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameter", "InvalidRequest"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{Name: "account_id", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "account_id"}),
			Hydrate:    opengovernance.GetS3MultiRegionAccessPoint,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchMultiRegionAccessPoint", "InvalidParameter", "InvalidRequest"}),
			},
		},
		Columns: awsOgColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the Multi-Region Access Point.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Report.Name"),
			},
			{
				Name:        "alias",
				Description: "The alias for the Multi-Region Access Point.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Report.Alias"),
			},
			{
				Name:        "created_at",
				Description: "When the Multi-Region Access Point create request was received.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Report.CreatedAt"),
			},
			{
				Name:        "status",
				Description: "The current status of the Multi-Region Access Point. CREATING and DELETING are temporary states that exist while the request is propagating and being completed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Report.Status"),
			},
			{
				Name:        "public_access_block",
				Description: "The PublicAccessBlock configuration that you want to apply to this Amazon S3 account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Report.PublicAccessBlock"),
			},
			{
				Name:        "regions",
				Description: "A collection of the Regions and buckets associated with the Multi-Region Access Point.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Report.Regions"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Report.Name"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

func listS3MultiRegionAccessPoints(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Get account details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("aws_s3_multi_region_access_point.listS3MultiRegionAccessPoints", "common_data_error", err)
		return nil, err
	}
	commonColumnData := commonData.(*awsCommonColumnData)

	// Create Session
	svc, err := S3ControlMultiRegionAccessClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("aws_s3_multi_region_access_point.listS3MultiRegionAccessPoints", "client_error", err)
		return nil, err
	}
	if svc == nil {
		return nil, nil
	}

	ownerAccountId := d.EqualsQuals["account_id"].GetStringValue()
	if ownerAccountId != "" && ownerAccountId != commonColumnData.AccountId {
		return nil, nil
	}

	maxItems := int32(100)

	// If the requested number of items is less than the paging max limit
	// set the limit to that instead
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxItems {
			maxItems = limit
		}
	}

	input := &s3control.ListMultiRegionAccessPointsInput{
		AccountId:  aws.String(commonColumnData.AccountId),
		MaxResults: maxItems,
	}

	paginator := s3control.NewListMultiRegionAccessPointsPaginator(svc, input, func(o *s3control.ListMultiRegionAccessPointsPaginatorOptions) {
		o.Limit = maxItems
		o.StopOnDuplicateToken = true
	})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			plugin.Logger(ctx).Error("aws_s3_multi_region_access_point.listS3MultiRegionAccessPoints", "api_error", err)
			return nil, err
		}

		for _, accessPoint := range output.AccessPoints {
			d.StreamListItem(ctx, accessPoint)

			// Context may get cancelled due to manual cancellation or if the limit has been reached
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getS3MultiRegionAccessPoint(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// Get account details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("aws_s3_multi_region_access_point.getS3MultiRegionAccessPoint", "common_data_error", err)
		return nil, err
	}
	commonColumnData := commonData.(*awsCommonColumnData)

	// Create Session
	svc, err := S3ControlMultiRegionAccessClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("aws_s3_multi_region_access_point.getS3MultiRegionAccessPoint", "client_error", err)
		return nil, err
	}
	if svc == nil {
		return nil, nil
	}

	var name, ownerAccountId string
	if h.Item != nil {
		multiRegionAccessPoint := h.Item.(types.MultiRegionAccessPointReport)
		name = *multiRegionAccessPoint.Name
		ownerAccountId = commonColumnData.AccountId
	} else {
		name = d.EqualsQuals["name"].GetStringValue()
		ownerAccountId = d.EqualsQuals["account_id"].GetStringValue()
	}

	// Build params
	params := &s3control.GetMultiRegionAccessPointInput{
		Name:      aws.String(name),
		AccountId: aws.String(ownerAccountId),
	}

	// execute list call
	item, err := svc.GetMultiRegionAccessPoint(ctx, params)
	if err != nil {
		plugin.Logger(ctx).Error("aws_s3_multi_region_access_point.getS3MultiRegionAccessPoint", "api_error", err)
		return nil, err
	}

	return item.AccessPoint, nil
}

func multiRegionAccessPointName(item interface{}) string {
	switch item := item.(type) {
	case types.MultiRegionAccessPointReport:
		return *item.Alias
	case *types.MultiRegionAccessPointReport:
		return *item.Name
	}
	return ""
}
