package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kafka/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
)

func KafkaCluster(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := kafka.NewFromConfig(cfg)
	paginator := kafka.NewListClustersV2Paginator(client, &kafka.ListClustersV2Input{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, cluster := range page.ClusterInfoList {
			resource, err := kafkaClusterHandle(ctx, cfg, cluster)
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
func kafkaClusterHandle(ctx context.Context, cfg aws.Config, cluster types.Cluster) (models.Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := kafka.NewFromConfig(cfg)

	var configArn string
	var err error
	var operationInfo *types.ClusterOperationInfo
	var configOut *kafka.DescribeConfigurationOutput
	if cluster.ClusterType == types.ClusterTypeProvisioned {
		if cluster.Provisioned.CurrentBrokerSoftwareInfo.ConfigurationArn != nil {
			configArn = *cluster.Provisioned.CurrentBrokerSoftwareInfo.ConfigurationArn
		}
	}

	if len(configArn) >= 1 {
		configOut, err = client.DescribeConfiguration(ctx, &kafka.DescribeConfigurationInput{Arn: &configArn})
		if err != nil {
			if isErr(err, "DescribeConfigurationNotFound") || isErr(err, "InvalidParameterValue") {
				return models.Resource{}, nil
			}
			return models.Resource{}, err
		}
	}

	if cluster.ActiveOperationArn != nil {
		op, err := client.DescribeClusterOperation(ctx, &kafka.DescribeClusterOperationInput{
			ClusterOperationArn: cluster.ActiveOperationArn,
		})
		if err != nil {
			if isErr(err, "DescribeClusterOperationNotFound") || isErr(err, "InvalidParameterValue") {
				return models.Resource{}, nil
			}
			return models.Resource{}, err
		}
		operationInfo = op.ClusterOperationInfo
	}
	resource := models.Resource{
		Region: describeCtx.OGRegion,
		ARN:    *cluster.ClusterArn,
		Name:   *cluster.ClusterName,
		Description: model.KafkaClusterDescription{
			Cluster:              cluster,
			ClusterOperationInfo: operationInfo,
			Configuration:        configOut,
		},
	}
	return resource, nil
}
func GetKafkaCluster(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	clusterName := fields["nameCluster"]
	var values []models.Resource

	client := kafka.NewFromConfig(cfg)
	out, err := client.ListClustersV2(ctx, &kafka.ListClustersV2Input{
		ClusterNameFilter: &clusterName,
	})
	if err != nil {
		if isErr(err, "ListClustersV2NotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, cluster := range out.ClusterInfoList {
		if cluster.ClusterName != &clusterName {
			continue
		}

		resource, err := kafkaClusterHandle(ctx, cfg, cluster)
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
