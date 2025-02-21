package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func RDSDBCluster(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := rds.NewFromConfig(cfg)
	paginator := rds.NewDescribeDBClustersPaginator(client, &rds.DescribeDBClustersInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.DBClusters {
			resource := rDSDBClusterHandle(ctx, client, v)
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
func rDSDBClusterHandle(ctx context.Context, client *rds.Client, v types.DBCluster) models.Resource {
	var actions []types.ResourcePendingMaintenanceActions
	pendingMaintenanceActions, err := client.DescribePendingMaintenanceActions(ctx, &rds.DescribePendingMaintenanceActionsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("db-cluster-id"),
				Values: []string{*v.DBClusterIdentifier},
			},
		},
	})
	if err == nil {
		actions = pendingMaintenanceActions.PendingMaintenanceActions
	}
	describeCtx := model.GetDescribeContext(ctx)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     *v.DBClusterArn,
		Name:    *v.DBClusterIdentifier,
		Account: describeCtx.AccountID,
		Description: model.RDSDBClusterDescription{
			DBCluster:                 v,
			PendingMaintenanceActions: actions,
		},
	}
	return resource
}
func GetRDSDBCluster(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	arn := fields["arn"]
	client := rds.NewFromConfig(cfg)

	out, err := client.DescribeDBClusters(ctx, &rds.DescribeDBClustersInput{
		DBClusterIdentifier: &arn,
	})
	if err != nil {
		if isErr(err, "DescribeDBClustersNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, v := range out.DBClusters {
		resource := rDSDBClusterHandle(ctx, client, v)
		values = append(values, resource)
	}

	return values, nil
}

func RDSDBClusterSnapshot(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := rds.NewFromConfig(cfg)
	paginator := rds.NewDescribeDBClusterSnapshotsPaginator(client, &rds.DescribeDBClusterSnapshotsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.DBClusterSnapshots {
			resource, err := rDSDBClusterSnapshotHandle(ctx, cfg, v)
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
func rDSDBClusterSnapshotHandle(ctx context.Context, cfg aws.Config, v types.DBClusterSnapshot) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := rds.NewFromConfig(cfg)

	attr, err := client.DescribeDBClusterSnapshotAttributes(ctx, &rds.DescribeDBClusterSnapshotAttributesInput{
		DBClusterSnapshotIdentifier: v.DBClusterSnapshotIdentifier,
	})
	if err != nil {
		if isErr(err, "DescribeDBClusterSnapshotAttributesNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	resource := models.Resource{
		Account: describeCtx.AccountID,
		Region:  describeCtx.OGRegion,
		ARN:     *v.DBClusterSnapshotArn,
		Name:    *v.DBClusterSnapshotIdentifier,
		Description: model.RDSDBClusterSnapshotDescription{
			DBClusterSnapshot: v,
			Attributes:        attr.DBClusterSnapshotAttributesResult,
		},
	}
	return resource, nil
}
func GetRDSDBClusterSnapshot(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	DBClusterSnapshotIdentifier := fields["ClusterSnapshotId"]
	SnapshotType := fields["snapshotType"]
	var values []models.Resource

	client := rds.NewFromConfig(cfg)
	clusterSnapshot, err := client.DescribeDBClusterSnapshots(ctx, &rds.DescribeDBClusterSnapshotsInput{
		DBClusterSnapshotIdentifier: &DBClusterSnapshotIdentifier,
		SnapshotType:                &SnapshotType,
	})
	if err != nil {
		if isErr(err, "DescribeDBClusterSnapshotsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, v := range clusterSnapshot.DBClusterSnapshots {
		resource, err := rDSDBClusterSnapshotHandle(ctx, cfg, v)
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

func RDSDBClusterParameterGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := rds.NewFromConfig(cfg)
	paginator := rds.NewDescribeDBClusterParameterGroupsPaginator(client, &rds.DescribeDBClusterParameterGroupsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.DBClusterParameterGroups {
			params, err := client.DescribeDBClusterParameters(ctx, &rds.DescribeDBClusterParametersInput{
				DBClusterParameterGroupName: v.DBClusterParameterGroupName,
			})
			if err != nil {
				return nil, err
			}

			op, err := client.ListTagsForResource(ctx, &rds.ListTagsForResourceInput{
				ResourceName: v.DBClusterParameterGroupArn,
			})
			if err != nil {
				return nil, err
			}

			var tags []types.Tag
			if len(op.TagList) > 0 {
				tags = op.TagList
			}

			resource := models.Resource{
				Account: describeCtx.AccountID,
				Region:  describeCtx.OGRegion,
				ARN:     *v.DBClusterParameterGroupArn,
				Name:    *v.DBClusterParameterGroupName,
				Description: model.RDSDBClusterParameterGroupDescription{
					DBClusterParameterGroup: v,
					Parameters:              params.Parameters,
					Tags:                    tags,
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

func RDSDBInstance(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := rds.NewFromConfig(cfg)
	paginator := rds.NewDescribeDBInstancesPaginator(client, &rds.DescribeDBInstancesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.DBInstances {
			resource, err := rDSDBInstanceHandle(ctx, client, v)
			if err != nil {
				return nil, err
			}
			if resource == nil {
				continue
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

	return values, nil
}
func rDSDBInstanceHandle(ctx context.Context, client *rds.Client, v types.DBInstance) (*models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	pendingMaintenance, err := client.DescribePendingMaintenanceActions(ctx, &rds.DescribePendingMaintenanceActionsInput{
		ResourceIdentifier: v.DBInstanceArn,
	})
	if err != nil {
		if isErr(err, "DescribePendingMaintenanceActionsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, nil
	}
	certificate, err := client.DescribeCertificates(ctx, &rds.DescribeCertificatesInput{
		CertificateIdentifier: v.CACertificateIdentifier,
	})
	if err != nil {
		if isErr(err, "DescribeCertificatesNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, nil
	}

	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     *v.DBInstanceArn,
		Name:    *v.DBInstanceIdentifier,
		Account: describeCtx.AccountID,
		Description: model.RDSDBInstanceDescription{
			DBInstance:         v,
			PendingMaintenance: pendingMaintenance.PendingMaintenanceActions,
			Certificate:        certificate.Certificates,
		},
	}
	return &resource, nil
}
func GetRDSDBInstance(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	dbInstanceId := fields["id"]
	client := rds.NewFromConfig(cfg)

	out, err := client.DescribeDBInstances(ctx, &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: &dbInstanceId,
	})
	if err != nil {
		if isErr(err, "DescribeDBInstancesNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, v := range out.DBInstances {
		resource, err := rDSDBInstanceHandle(ctx, client, v)
		if err != nil {
			return nil, err
		}
		if resource == nil {
			continue
		}

		values = append(values, *resource)
	}
	return values, nil
}

func RDSDBParameterGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := rds.NewFromConfig(cfg)
	paginator := rds.NewDescribeDBParameterGroupsPaginator(client, &rds.DescribeDBParameterGroupsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.DBParameterGroups {
			dbParams, err := client.DescribeDBParameters(ctx, &rds.DescribeDBParametersInput{
				DBParameterGroupName: v.DBParameterGroupName,
			})
			if err != nil {
				return nil, err
			}

			op, err := client.ListTagsForResource(ctx, &rds.ListTagsForResourceInput{
				ResourceName: v.DBParameterGroupArn,
			})
			if err != nil {
				return nil, err
			}

			resource := models.Resource{
				Region:  describeCtx.OGRegion,
				ARN:     *v.DBParameterGroupArn,
				Name:    *v.DBParameterGroupName,
				Account: describeCtx.AccountID,
				Description: model.RDSDBParameterGroupDescription{
					DBParameterGroup: v,
					Parameters:       dbParams.Parameters,
					Tags:             op.TagList,
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

func RDSDBProxy(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := rds.NewFromConfig(cfg)
	paginator := rds.NewDescribeDBProxiesPaginator(client, &rds.DescribeDBProxiesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.DBProxies {
			tags, err := client.ListTagsForResource(ctx, &rds.ListTagsForResourceInput{
				ResourceName: v.DBProxyArn,
			})
			if err != nil {
				tags = &rds.ListTagsForResourceOutput{}
			}

			resource := models.Resource{
				Region:  describeCtx.OGRegion,
				ARN:     *v.DBProxyArn,
				Name:    *v.DBProxyName,
				Account: describeCtx.AccountID,
				Description: model.RDSDBProxyDescription{
					DBProxy: v,
					Tags:    tags.TagList,
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

func RDSDBProxyEndpoint(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := rds.NewFromConfig(cfg)
	paginator := rds.NewDescribeDBProxyEndpointsPaginator(client, &rds.DescribeDBProxyEndpointsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.DBProxyEndpoints {
			resource := models.Resource{
				Region:      describeCtx.OGRegion,
				ARN:         *v.DBProxyEndpointArn,
				Name:        *v.DBProxyEndpointName,
				Account:     describeCtx.AccountID,
				Description: v,
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

func RDSDBProxyTargetGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	proxies, err := RDSDBProxy(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := rds.NewFromConfig(cfg)

	var values []models.Resource
	for _, p := range proxies {
		proxy := p.Description.(types.DBProxy)
		paginator := rds.NewDescribeDBProxyTargetGroupsPaginator(client, &rds.DescribeDBProxyTargetGroupsInput{
			DBProxyName: proxy.DBProxyName,
		})

		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}

			for _, v := range page.TargetGroups {
				resource := models.Resource{
					Region:      describeCtx.OGRegion,
					ARN:         *v.TargetGroupArn,
					Name:        *v.TargetGroupName,
					Account:     describeCtx.AccountID,
					Description: v,
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

func RDSDBSecurityGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := rds.NewFromConfig(cfg)
	paginator := rds.NewDescribeDBSecurityGroupsPaginator(client, &rds.DescribeDBSecurityGroupsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.DBSecurityGroups {
			resource := models.Resource{
				Region:      describeCtx.OGRegion,
				ARN:         *v.DBSecurityGroupArn,
				Name:        *v.DBSecurityGroupName,
				Account:     describeCtx.AccountID,
				Description: v,
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

func RDSDBSubnetGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := rds.NewFromConfig(cfg)
	paginator := rds.NewDescribeDBSubnetGroupsPaginator(client, &rds.DescribeDBSubnetGroupsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.DBSubnetGroups {
			tags, err := client.ListTagsForResource(ctx, &rds.ListTagsForResourceInput{
				ResourceName: v.DBSubnetGroupArn,
			})
			if err != nil {
				return nil, err
			}

			resource := models.Resource{
				Region:  describeCtx.OGRegion,
				ARN:     *v.DBSubnetGroupArn,
				Name:    *v.DBSubnetGroupName,
				Account: describeCtx.AccountID,
				Description: model.RDSDBSubnetGroupDescription{
					DBSubnetGroup: v,
					Tags:          tags,
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

func RDSDBEventSubscription(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := rds.NewFromConfig(cfg)
	paginator := rds.NewDescribeEventSubscriptionsPaginator(client, &rds.DescribeEventSubscriptionsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.EventSubscriptionsList {
			resource := rDSDBEventSubscriptionHandle(ctx, v)
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
func rDSDBEventSubscriptionHandle(ctx context.Context, v types.EventSubscription) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     *v.EventSubscriptionArn,
		Name:    *v.CustSubscriptionId,
		Account: describeCtx.AccountID,
		Description: model.RDSDBEventSubscriptionDescription{
			EventSubscription: v,
		},
	}
	return resource
}
func GetRDSDBEventSubscription(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	subscriptionName := fields["name"]
	var values []models.Resource

	client := rds.NewFromConfig(cfg)
	describes, err := client.DescribeEventSubscriptions(ctx, &rds.DescribeEventSubscriptionsInput{
		SubscriptionName: &subscriptionName,
	})
	if err != nil {
		if isErr(err, "DescribeEventSubscriptionsNotFound") || isErr(err, "invalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, v := range describes.EventSubscriptionsList {
		values = append(values, rDSDBEventSubscriptionHandle(ctx, v))
	}
	return values, nil
}

func RDSGlobalCluster(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := rds.NewFromConfig(cfg)
	paginator := rds.NewDescribeGlobalClustersPaginator(client, &rds.DescribeGlobalClustersInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.GlobalClusters {
			resource := rDSGlobalClusterHandle(ctx, cfg, v)
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
func rDSGlobalClusterHandle(ctx context.Context, cfg aws.Config, v types.GlobalCluster) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	client := rds.NewFromConfig(cfg)

	tags, err := client.ListTagsForResource(ctx, &rds.ListTagsForResourceInput{
		ResourceName: v.GlobalClusterArn,
	})
	if err != nil {
		tags = &rds.ListTagsForResourceOutput{}
	}

	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     *v.GlobalClusterArn,
		Name:    *v.GlobalClusterIdentifier,
		Account: describeCtx.AccountID,
		Description: model.RDSGlobalClusterDescription{
			GlobalCluster: v,
			Tags:          tags.TagList,
		},
	}
	return resource
}
func GetRDSGlobalCluster(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	var values []models.Resource
	arn := fields["arn"]
	client := rds.NewFromConfig(cfg)

	describers, err := client.DescribeGlobalClusters(ctx, &rds.DescribeGlobalClustersInput{
		GlobalClusterIdentifier: &arn,
	})
	if err != nil {
		if isErr(err, "DescribeGlobalClustersNotFound") || isErr(err, "invalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, v := range describers.GlobalClusters {
		values = append(values, rDSGlobalClusterHandle(ctx, cfg, v))
	}
	return values, nil
}

func RDSOptionGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := rds.NewFromConfig(cfg)
	paginator := rds.NewDescribeOptionGroupsPaginator(client, &rds.DescribeOptionGroupsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.OptionGroupsList {
			resource, err := rDSOptionGroupHandle(ctx, cfg, v)
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
func rDSOptionGroupHandle(ctx context.Context, cfg aws.Config, v types.OptionGroup) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := rds.NewFromConfig(cfg)

	tags, err := client.ListTagsForResource(ctx, &rds.ListTagsForResourceInput{
		ResourceName: v.OptionGroupArn,
	})
	if err != nil {
		if isErr(err, "ListTagsForResourceNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     *v.OptionGroupArn,
		Name:    *v.OptionGroupName,
		Account: describeCtx.AccountID,
		Description: model.RDSOptionGroupDescription{
			OptionGroup: v,
			Tags:        tags,
		},
	}
	return resource, nil
}
func GetRDSOptionGroup(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	optionGroupName := fields["name"]
	var values []models.Resource
	client := rds.NewFromConfig(cfg)

	describers, err := client.DescribeOptionGroups(ctx, &rds.DescribeOptionGroupsInput{
		OptionGroupName: &optionGroupName,
	})
	if err != nil {
		if isErr(err, "DescribeOptionGroupsNotFound ") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, v := range describers.OptionGroupsList {
		resource, err := rDSOptionGroupHandle(ctx, cfg, v)
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

func RDSDBSnapshot(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := rds.NewFromConfig(cfg)
	paginator := rds.NewDescribeDBSnapshotsPaginator(client, &rds.DescribeDBSnapshotsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.DBSnapshots {
			resource, err := rDSDBSnapshotHandle(ctx, cfg, v)
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
func rDSDBSnapshotHandle(ctx context.Context, cfg aws.Config, v types.DBSnapshot) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := rds.NewFromConfig(cfg)
	attrs, err := client.DescribeDBSnapshotAttributes(ctx, &rds.DescribeDBSnapshotAttributesInput{
		DBSnapshotIdentifier: v.DBSnapshotIdentifier,
	})
	if err != nil {
		if isErr(err, "DescribeDBSnapshotAttributesNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     *v.DBSnapshotArn,
		Name:    *v.DBSnapshotIdentifier,
		Account: describeCtx.AccountID,
		Description: model.RDSDBSnapshotDescription{
			DBSnapshot:           v,
			DBSnapshotAttributes: attrs.DBSnapshotAttributesResult.DBSnapshotAttributes,
		},
	}

	return resource, nil
}
func GetRDSDBSnapshot(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	dbiResourceId := fields["id"]
	var values []models.Resource
	client := rds.NewFromConfig(cfg)
	describers, err := client.DescribeDBSnapshots(ctx, &rds.DescribeDBSnapshotsInput{
		DbiResourceId: &dbiResourceId,
	})
	if err != nil {
		if isErr(err, "DescribeDBSnapshotsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, v := range describers.DBSnapshots {
		resource, err := rDSDBSnapshotHandle(ctx, cfg, v)
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

func RDSReservedDBInstance(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := rds.NewFromConfig(cfg)
	paginator := rds.NewDescribeReservedDBInstancesPaginator(client, &rds.DescribeReservedDBInstancesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, reservedDBInstance := range page.ReservedDBInstances {
			resource := rDSReservedDBInstanceHandle(ctx, reservedDBInstance)
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
func rDSReservedDBInstanceHandle(ctx context.Context, reservedDBInstance types.ReservedDBInstance) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     *reservedDBInstance.ReservedDBInstanceArn,
		ID:      *reservedDBInstance.ReservedDBInstanceId,
		Account: describeCtx.AccountID,
		Description: model.RDSReservedDBInstanceDescription{
			ReservedDBInstance: reservedDBInstance,
		},
	}
	return resource
}
func GetRDSReservedDBInstance(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	reservedDBInstanceId := fields["id"]
	var values []models.Resource

	client := rds.NewFromConfig(cfg)
	describers, err := client.DescribeReservedDBInstances(ctx, &rds.DescribeReservedDBInstancesInput{
		ReservedDBInstanceId: &reservedDBInstanceId,
	})
	if err != nil {
		if isErr(err, "DescribeReservedDBInstancesNotFound") || isErr(err, "invalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, v := range describers.ReservedDBInstances {
		values = append(values, rDSReservedDBInstanceHandle(ctx, v))
	}
	return values, nil
}

func RDSDBInstanceAutomatedBackup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := rds.NewFromConfig(cfg)
	paginator := rds.NewDescribeDBInstanceAutomatedBackupsPaginator(client, &rds.DescribeDBInstanceAutomatedBackupsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.DBInstanceAutomatedBackups {
			resource := rDSDBInstanceAutomatedBackupHandle(ctx, v)
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
func rDSDBInstanceAutomatedBackupHandle(ctx context.Context, v types.DBInstanceAutomatedBackup) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     *v.DBInstanceArn,
		ID:      *v.DBInstanceIdentifier,
		Account: describeCtx.AccountID,
		Description: model.RDSDBInstanceAutomatedBackupDescription{
			InstanceAutomatedBackup: v,
		},
	}
	return resource
}
func GetRDSDBInstanceAutomatedBackup(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	arn := fields["arn"]
	client := rds.NewFromConfig(cfg)

	out, err := client.DescribeDBInstanceAutomatedBackups(ctx, &rds.DescribeDBInstanceAutomatedBackupsInput{
		DBInstanceAutomatedBackupsArn: &arn,
	})
	if err != nil {
		if isErr(err, "DescribeDBClustersNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, v := range out.DBInstanceAutomatedBackups {
		resource := rDSDBInstanceAutomatedBackupHandle(ctx, v)
		values = append(values, resource)
	}

	return values, nil
}

func RDSDBEngineVersion(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := rds.NewFromConfig(cfg)
	paginator := rds.NewDescribeDBEngineVersionsPaginator(client, &rds.DescribeDBEngineVersionsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		if page.DBEngineVersions == nil {
			continue
		}

		for _, v := range page.DBEngineVersions {
			resource := rDSDBEngineVersionHandle(ctx, v)
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

func rDSDBEngineVersionHandle(ctx context.Context, v types.DBEngineVersion) models.Resource {
	var arn string
	if v.DBEngineVersionArn != nil {
		arn = *v.DBEngineVersionArn
	}
	describeCtx := model.GetDescribeContext(ctx)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     arn,
		Account: describeCtx.AccountID,
		Description: model.RDSDBEngineVersionDescription{
			EngineVersion: v,
		},
	}
	return resource
}

func RDSDBRecommendation(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	logger := GetLoggerFromContext(ctx)

	logger.Info("RDSDBRecommendation start working")

	client := rds.NewFromConfig(cfg)
	paginator := rds.NewDescribeDBRecommendationsPaginator(client, &rds.DescribeDBRecommendationsInput{})

	logger.Info("RDSDBRecommendation start getting pages")
	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		logger.Info("RDSDBRecommendation got page")
		for _, v := range page.DBRecommendations {
			resource := rDSDBRecommendationHandler(ctx, v)
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	logger.Info("RDSDBRecommendation finished")

	return values, nil
}

func rDSDBRecommendationHandler(ctx context.Context, v types.DBRecommendation) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)

	return models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     *v.ResourceArn,
		Account: describeCtx.AccountID,
		Description: model.RDSDBRecommendationDescription{
			DBRecommendation: v,
		},
	}
}
