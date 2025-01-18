package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2TransitGatewayRouteTable(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_transit_gateway_route_table",
		Description: "AWS EC2 Transit Gateway Route Table",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("transit_gateway_route_table_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidRouteTableID.NotFound", "InvalidRouteTableId.Unavailable", "InvalidRouteTableId.Malformed"}),
			},
			Hydrate: opengovernance.GetEC2TransitGatewayRouteTable,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEC2TransitGatewayRouteTable,
		},

		Columns: awsOgRegionalColumns([]*plugin.Column{
			{
				Name:        "transit_gateway_route_table_id",
				Description: "The ID of the transit gateway route table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGatewayRouteTable.TransitGatewayRouteTableId")},
			{
				Name:        "transit_gateway_id",
				Description: "The ID of the transit gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGatewayRouteTable.TransitGatewayId")},
			{
				Name:        "state",
				Description: "The state of the transit gateway route table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TransitGatewayRouteTable.State")},
			{
				Name:        "creation_time",
				Description: "The creation time of transit gateway route table.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.TransitGatewayRouteTable.CreationTime")},
			{
				Name:        "default_association_route_table",
				Description: "Indicates whether this is the default association route table for the transit gateway.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.TransitGatewayRouteTable.DefaultAssociationRouteTable")},
			{
				Name:        "default_propagation_route_table",
				Description: "Indicates whether this is the default propagation route table for the transit gateway.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.TransitGatewayRouteTable.DefaultPropagationRouteTable")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TransitGatewayRouteTable.Tags")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEc2TransitGatewayRouteTableTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getEc2TransitGatewayRouteTableTurbotTitle),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getEc2TransitGatewayRouteTableTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.EC2TransitGatewayRouteTable).Description.TransitGatewayRouteTable
	return ec2V2TagsToMap(data.Tags)
}

func getEc2TransitGatewayRouteTableTurbotTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.EC2TransitGatewayRouteTable).Description.TransitGatewayRouteTable
	title := data.TransitGatewayRouteTableId
	if data.Tags != nil {
		for _, i := range data.Tags {
			if *i.Key == "Name" {
				title = i.Value
			}
		}
	}
	return title, nil
}
