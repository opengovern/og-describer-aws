package describers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/simspaceweaver"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/opengovern/og-describer-aws/discovery/pkg/models"
	model "github.com/opengovern/og-describer-aws/discovery/provider"
)

func SimSpaceWeaverSimulation(ctx context.Context, cfg aws.Config, stream *models.StreamSender) ([]models.Resource, error) {
	describeCtx := model.GetDescribeContext(ctx)

	client := simspaceweaver.NewFromConfig(cfg)
	paginator := simspaceweaver.NewListSimulationsPaginator(client, &simspaceweaver.ListSimulationsInput{})

	var values []models.Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "AccessDeniedException") {
				continue
			}
			return nil, err
		}

		for _, v := range page.Simulations {
			data, err := client.DescribeSimulation(ctx, &simspaceweaver.DescribeSimulationInput{
				Simulation: v.Name,
			})
			if err != nil {
				return nil, err
			}

			var name string
			if v.Name != nil {
				name = *v.Name
			}

			tags, err := client.ListTagsForResource(ctx, &simspaceweaver.ListTagsForResourceInput{
				ResourceArn: v.Arn,
			})
			if err != nil {
				return nil, err
			}

			resource := models.Resource{
				Account: describeCtx.AccountID,
				Region:  describeCtx.OGRegion,
				ARN:     *v.Arn,
				Name:    name,
				Description: model.SimSpaceWeaverSimulationDescription{
					Simulation:     v,
					SimulationItem: data,
					Tags:           tags.Tags,
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
