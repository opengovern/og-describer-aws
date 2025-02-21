package describers

import (
	"context"
	"fmt"

	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dax/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func DAXCluster(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := dax.NewFromConfig(cfg)
	out, err := client.DescribeClusters(ctx, &dax.DescribeClustersInput{})
	if err != nil {
		if strings.Contains(err.Error(), "InvalidParameterValueException") || strings.Contains(err.Error(), "no such host") {
			return nil, nil
		}
		return nil, err
	}

	var values []models.Resource
	for _, cluster := range out.Clusters {
		tags, err := client.ListTags(ctx, &dax.ListTagsInput{
			ResourceName: cluster.ClusterArn,
		})
		if err != nil {
			if strings.Contains(err.Error(), "ClusterNotFoundFault") {
				tags = nil
			} else {
				return nil, err
			}
		}

		resource := dAXClusterHandle(ctx, tags, cluster)
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
func dAXClusterHandle(ctx context.Context, tags *dax.ListTagsOutput, cluster types.Cluster) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		ARN:     *cluster.ClusterArn,
		Name:    *cluster.ClusterName,
		Account: describeCtx.AccountID,
		Description: model.DAXClusterDescription{
			Cluster: cluster,
			Tags:    tags.Tags,
		},
	}
	return resource
}
func GetDAXCluster(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	clusterName := fields["name"]
	var values []models.Resource
	client := dax.NewFromConfig(cfg)

	clusterDescribe, err := client.DescribeClusters(ctx, &dax.DescribeClustersInput{
		ClusterNames: []string{clusterName},
	})
	if err != nil {
		if isErr(err, "DescribeClustersNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, cluster := range clusterDescribe.Clusters {
		tags, err := client.ListTags(ctx, &dax.ListTagsInput{
			ResourceName: cluster.ClusterArn,
		})
		if err != nil {
			if strings.Contains(err.Error(), "ClusterNotFoundFault") {
				tags = nil
			} else {
				return nil, err
			}
		}

		values = append(values, dAXClusterHandle(ctx, tags, cluster))
	}
	return values, nil
}

func DAXParameterGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	client := dax.NewFromConfig(cfg)

	var values []models.Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		parameterGroups, err := client.DescribeParameterGroups(ctx, &dax.DescribeParameterGroupsInput{
			MaxResults: aws.Int32(100),
			NextToken:  prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, parameterGroup := range parameterGroups.ParameterGroups {

			resource := dAXParameterGroupHandle(ctx, parameterGroup)
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}

		return parameterGroups.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
func dAXParameterGroupHandle(ctx context.Context, parameterGroup types.ParameterGroup) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	resource := models.Resource{
		Region:  describeCtx.OGRegion,
		Name:    *parameterGroup.ParameterGroupName,
		Account: describeCtx.AccountID,
		Description: model.DAXParameterGroupDescription{
			ParameterGroup: parameterGroup,
		},
	}
	return resource
}
func GetDAXParameterGroup(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	parameterGroupName := fields["name"]
	var values []models.Resource

	client := dax.NewFromConfig(cfg)
	parameterGroups, err := client.DescribeParameterGroups(ctx, &dax.DescribeParameterGroupsInput{
		ParameterGroupNames: []string{parameterGroupName},
	})
	if err != nil {
		if isErr(err, "DescribeParameterGroupsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, parameterGroup := range parameterGroups.ParameterGroups {
		values = append(values, dAXParameterGroupHandle(ctx, parameterGroup))
	}
	return values, nil
}

func DAXParameter(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)
	//
	client := dax.NewFromConfig(cfg)

	var values []models.Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		parameterGroups, err := client.DescribeParameterGroups(ctx, &dax.DescribeParameterGroupsInput{
			MaxResults: aws.Int32(100),
			NextToken:  prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, parameterGroup := range parameterGroups.ParameterGroups {
			err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
				parameters, err := client.DescribeParameters(ctx, &dax.DescribeParametersInput{
					ParameterGroupName: parameterGroup.ParameterGroupName,
					MaxResults:         aws.Int32(100),
					NextToken:          prevToken,
				})
				if err != nil {
					return nil, err
				}

				for _, parameter := range parameters.Parameters {
					resource := models.Resource{
						Region:  describeCtx.OGRegion,
						Account: describeCtx.AccountID,
						Name:    *parameter.ParameterName,
						Description: model.DAXParameterDescription{
							Parameter:          parameter,
							ParameterGroupName: *parameterGroup.ParameterGroupName,
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

				return parameters.NextToken, nil
			})
			if err != nil {
				return nil, err
			}
		}

		return parameterGroups.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func DAXSubnetGroup(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {

	client := dax.NewFromConfig(cfg)

	var values []models.Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		subnetGroups, err := client.DescribeSubnetGroups(ctx, &dax.DescribeSubnetGroupsInput{
			MaxResults: aws.Int32(100),
			NextToken:  prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, subnetGroup := range subnetGroups.SubnetGroups {

			resource := dAXSubnetGroupHandle(ctx, subnetGroup)
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}
		return subnetGroups.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}
func dAXSubnetGroupHandle(ctx context.Context, subnetGroup types.SubnetGroup) models.Resource {
	describeCtx := model.GetDescribeContext(ctx)
	arn := fmt.Sprintf("arn:%s:dax:%s::subnetgroup:%s", describeCtx.Partition, describeCtx.Region, *subnetGroup.SubnetGroupName)

	resource := models.Resource{
		Account: describeCtx.AccountID,
		Region:  describeCtx.OGRegion,
		Name:    *subnetGroup.SubnetGroupName,
		ARN:     arn,
		Description: model.DAXSubnetGroupDescription{
			SubnetGroup: subnetGroup,
		},
	}
	return resource
}
func GetDAXSubnetGroup(ctx context.Context, cfg aws.Config, fields map[string]string) ([]models.Resource, error) {
	SubnetGroupNames := fields["name"]
	var values []models.Resource
	client := dax.NewFromConfig(cfg)

	subnetGroups, err := client.DescribeSubnetGroups(ctx, &dax.DescribeSubnetGroupsInput{
		SubnetGroupNames: []string{SubnetGroupNames},
	})
	if err != nil {
		if isErr(err, "DescribeSubnetGroupsNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	for _, subnetGroup := range subnetGroups.SubnetGroups {
		values = append(values, dAXSubnetGroupHandle(ctx, subnetGroup))
	}
	return values, nil
}
