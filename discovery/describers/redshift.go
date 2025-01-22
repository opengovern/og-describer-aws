package describers

import (
	"context"
	"errors"
	"fmt"

	"strings"

	types2 "github.com/aws/aws-sdk-go-v2/service/redshiftserverless/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless"
	"github.com/aws/smithy-go"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func RedshiftCluster(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := redshift.NewFromConfig(cfg)
	paginator := redshift.NewDescribeClustersPaginator(client, &redshift.DescribeClustersInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Clusters {
			resource, err := redshiftClusterHandle(ctx, cfg, v)
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
func redshiftClusterHandle(ctx context.Context, cfg aws.Config, v types.Cluster) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := redshift.NewFromConfig(cfg)
	logStatus, err := client.DescribeLoggingStatus(ctx, &redshift.DescribeLoggingStatusInput{
		ClusterIdentifier: v.ClusterIdentifier,
	})
	if err != nil {
		if isErr(err, "DescribeLoggingStatusNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	sactions, err := client.DescribeScheduledActions(ctx, &redshift.DescribeScheduledActionsInput{
		Filters: []types.ScheduledActionFilter{
			{
				Name:   types.ScheduledActionFilterNameClusterIdentifier,
				Values: []string{*v.ClusterIdentifier},
			},
		},
	})
	if err != nil {
		if isErr(err, "DescribeScheduledActionsNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *v.ClusterNamespaceArn,
		Name:   *v.ClusterIdentifier,
		Description: model.RedshiftClusterDescription{
			Cluster:          v,
			LoggingStatus:    logStatus,
			ScheduledActions: sactions.ScheduledActions,
		},
	}
	return resource, nil
}
func GetRedshiftCluster(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	tagKey := fields["tagKey"]
	tagValue := fields["tagValue"]
	client := redshift.NewFromConfig(cfg)

	out, err := client.DescribeClusters(ctx, &redshift.DescribeClustersInput{
		TagKeys:   []string{tagKey},
		TagValues: []string{tagValue},
	})
	if err != nil {
		if isErr(err, "DescribeClustersNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, cluster := range out.Clusters {
		resource, err := redshiftClusterHandle(ctx, cfg, cluster)
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

func RedshiftEventSubscription(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := redshift.NewFromConfig(cfg)
	paginator := redshift.NewDescribeEventSubscriptionsPaginator(client, &redshift.DescribeEventSubscriptionsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.EventSubscriptionsList {
			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ID:     *v.CustSubscriptionId,
				Name:   *v.CustSubscriptionId,
				Description: model.RedshiftEventSubscriptionDescription{
					EventSubscription: v,
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

func RedshiftClusterParameterGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := redshift.NewFromConfig(cfg)
	paginator := redshift.NewDescribeClusterParameterGroupsPaginator(client, &redshift.DescribeClusterParameterGroupsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ParameterGroups {
			resource, err := redshiftClusterParameterGroupHandle(ctx, cfg, v)
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
func redshiftClusterParameterGroupHandle(ctx context.Context, cfg aws.Config, v types.ClusterParameterGroup) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := redshift.NewFromConfig(cfg)

	param, err := client.DescribeClusterParameters(ctx, &redshift.DescribeClusterParametersInput{
		ParameterGroupName: v.ParameterGroupName,
	})
	if err != nil {
		if isErr(err, "DescribeClusterParametersNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	arn := "arn:" + describeCtx.Partition + ":redshift:" + describeCtx.Region + ":" + describeCtx.AccountID + ":parametergroup"
	if strings.HasPrefix(*v.ParameterGroupName, ":") {
		arn = arn + *v.ParameterGroupName
	} else {
		arn = arn + ":" + *v.ParameterGroupName
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    arn,
		Name:   *v.ParameterGroupName,
		Description: model.RedshiftClusterParameterGroupDescription{
			ClusterParameterGroup: v,
			Parameters:            param.Parameters,
		},
	}
	return resource, nil
}
func GetRedshiftClusterParameterGroup(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	ParameterGroupName := fields["name"]
	client := redshift.NewFromConfig(cfg)

	out, err := client.DescribeClusterParameterGroups(ctx, &redshift.DescribeClusterParameterGroupsInput{
		ParameterGroupName: &ParameterGroupName,
	})
	if err != nil {
		if isErr(err, "DescribeClusterParameterGroupsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, v := range out.ParameterGroups {

		resource, err := redshiftClusterParameterGroupHandle(ctx, cfg, v)
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

func RedshiftClusterSecurityGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := redshift.NewFromConfig(cfg)
	paginator := redshift.NewDescribeClusterSecurityGroupsPaginator(client, &redshift.DescribeClusterSecurityGroupsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			var ae smithy.APIError
			if errors.As(err, &ae) && (ae.ErrorMessage() == "VPC-by-Default customers cannot use cluster security groups") {
				return nil, nil
			}

			return nil, err
		}

		for _, v := range page.ClusterSecurityGroups {
			resource := models.Resource{
				Region:      describeCtx.OGRegion,
				ID:          *v.ClusterSecurityGroupName,
				Name:        *v.ClusterSecurityGroupName,
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

func RedshiftClusterSubnetGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := redshift.NewFromConfig(cfg)
	paginator := redshift.NewDescribeClusterSubnetGroupsPaginator(client, &redshift.DescribeClusterSubnetGroupsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ClusterSubnetGroups {
			resource := models.Resource{
				Region:      describeCtx.OGRegion,
				ID:          *v.ClusterSubnetGroupName,
				Name:        *v.ClusterSubnetGroupName,
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

func RedshiftSnapshot(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := redshift.NewFromConfig(cfg)
	paginator := redshift.NewDescribeClusterSnapshotsPaginator(client, &redshift.DescribeClusterSnapshotsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "ClusterSnapshotNotFound") {
				continue
			}
			return nil, err
		}

		for _, v := range page.Snapshots {
			resource := redshiftSnapshotHandle(ctx, v)
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
func redshiftSnapshotHandle(ctx context.Context, v types.Snapshot) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	arn := fmt.Sprintf("arn:%s:redshift:%s:%s:snapshot:%s/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, *v.ClusterIdentifier, *v.SnapshotIdentifier)
	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    arn,
		Name:   *v.SnapshotIdentifier,
		Description: model.RedshiftSnapshotDescription{
			Snapshot: v,
		},
	}
	return resource
}
func GetRedshiftSnapshot(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	clusterIdentifier := fields["id"]

	client := redshift.NewFromConfig(cfg)

	out, err := client.DescribeClusterSnapshots(ctx, &redshift.DescribeClusterSnapshotsInput{
		ClusterIdentifier: &clusterIdentifier,
	})
	if err != nil {
		if isErr(err, "ClusterSnapshotNotFound") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, v := range out.Snapshots {
		values = append(values, redshiftSnapshotHandle(ctx, v))
	}

	return values, nil
}

func RedshiftServerlessNamespace(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := redshiftserverless.NewFromConfig(cfg)
	paginator := redshiftserverless.NewListNamespacesPaginator(client, &redshiftserverless.ListNamespacesInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Namespaces {
			resource, err := redshiftServerlessNamespaceHandle(ctx, cfg, v)
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
func redshiftServerlessNamespaceHandle(ctx context.Context, cfg aws.Config, v types2.Namespace) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := redshiftserverless.NewFromConfig(cfg)
	tags, err := client.ListTagsForResource(ctx, &redshiftserverless.ListTagsForResourceInput{
		ResourceArn: v.NamespaceArn,
	})
	if err != nil {
		if isErr(err, "ListTagsForResourceNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *v.NamespaceArn,
		Name:   *v.NamespaceName,
		Description: model.RedshiftServerlessNamespaceDescription{
			Namespace: v,
			Tags:      tags.Tags,
		},
	}
	return resource, nil
}
func GetRedshiftServerlessNamespace(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	namespaceName := fields["name"]
	var values []models.Resource
	client := redshiftserverless.NewFromConfig(cfg)

	namespaces, err := client.GetNamespace(ctx, &redshiftserverless.GetNamespaceInput{
		NamespaceName: &namespaceName,
	})
	if err != nil {
		if isErr(err, "GetNamespaceNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	resource, err := redshiftServerlessNamespaceHandle(ctx, cfg, *namespaces.Namespace)
	emptyResource := models.Resource{}
	if err == nil && resource == emptyResource {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	values = append(values, resource)
	return values, nil
}

func RedshiftServerlessSnapshot(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := redshiftserverless.NewFromConfig(cfg)
	paginator := redshiftserverless.NewListSnapshotsPaginator(client, &redshiftserverless.ListSnapshotsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Snapshots {
			resource, err := redshiftServerlessSnapshotHandle(ctx, cfg, v)
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
func redshiftServerlessSnapshotHandle(ctx context.Context, cfg aws.Config, v types2.Snapshot) (models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := redshiftserverless.NewFromConfig(cfg)
	tags, err := client.ListTagsForResource(ctx, &redshiftserverless.ListTagsForResourceInput{
		ResourceArn: v.NamespaceArn,
	})
	if err != nil {
		if isErr(err, "ListTagsForResourceNotFound") || isErr(err, "InvalidParameterValue") {
			return models.Resource{}, nil
		}
		return models.Resource{}, err
	}

	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *v.NamespaceArn,
		Name:   *v.NamespaceName,
		Description: model.RedshiftServerlessSnapshotDescription{
			Snapshot: v,
			Tags:     tags.Tags,
		},
	}
	return resource, nil
}
func GetRedshiftServerlessSnapshot(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	snapshot := fields["name"]
	var values []models.Resource
	client := redshiftserverless.NewFromConfig(cfg)

	out, err := client.GetSnapshot(ctx, &redshiftserverless.GetSnapshotInput{
		SnapshotName: &snapshot,
	})
	if err != nil {
		if isErr(err, "GetSnapshotNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	resource, err := redshiftServerlessSnapshotHandle(ctx, cfg, *out.Snapshot)
	emptyResource := models.Resource{}
	if err == nil && resource == emptyResource {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	values = append(values, resource)
	return values, nil
}

func RedshiftServerlessWorkgroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	client := redshiftserverless.NewFromConfig(cfg)
	paginator := redshiftserverless.NewListWorkgroupsPaginator(client, &redshiftserverless.ListWorkgroupsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Workgroups {
			tags, err := client.ListTagsForResource(ctx, &redshiftserverless.ListTagsForResourceInput{
				ResourceArn: v.WorkgroupArn,
			})
			if err != nil {
				return nil, err
			}

			resource := models.Resource{
				Region: describeCtx.OGRegion,
				ARN:    *v.WorkgroupArn,
				Name:   *v.WorkgroupName,
				Description: model.RedshiftServerlessWorkgroupDescription{
					Workgroup: v,
					Tags:      tags.Tags,
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

func RedshiftSubnetGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := redshift.NewFromConfig(cfg)
	paginator := redshift.NewDescribeClusterSubnetGroupsPaginator(client, &redshift.DescribeClusterSubnetGroupsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, clusterSubnetGroup := range page.ClusterSubnetGroups {
			resource := redshiftSubnetGroupHandle(ctx, clusterSubnetGroup)
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
func redshiftSubnetGroupHandle(ctx context.Context, clusterSubnetGroup types.ClusterSubnetGroup) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	arn := fmt.Sprintf("arn:%s:redshift:%s:%s:subnetgroup:%s", describeCtx.Partition, describeCtx.OGRegion, describeCtx.AccountID, *clusterSubnetGroup.ClusterSubnetGroupName)
	resource := models.Resource{
		Region: describeCtx.OGRegion,
		Name:   *clusterSubnetGroup.ClusterSubnetGroupName,
		ARN:    arn,
		Description: model.RedshiftSubnetGroupDescription{
			ClusterSubnetGroup: clusterSubnetGroup,
		},
	}
	return resource
}
func GetRedshiftSubnetGroup(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	ClusterSubnetGroupName := fields["name"]
	client := redshift.NewFromConfig(cfg)

	clusterSubnets, err := client.DescribeClusterSubnetGroups(ctx, &redshift.DescribeClusterSubnetGroupsInput{
		ClusterSubnetGroupName: &ClusterSubnetGroupName,
	})
	if err != nil {
		if isErr(err, "DescribeClusterSubnetGroupsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, clusterSubnetGroup := range clusterSubnets.ClusterSubnetGroups {
		values = append(values, redshiftSubnetGroupHandle(ctx, clusterSubnetGroup))
	}
	return values, nil
}
