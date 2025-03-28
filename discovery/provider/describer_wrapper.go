package provider

import (
	"fmt"

	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/smithy-go"
	model "github.com/opengovern/og-describer-aws/discovery/pkg/models"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"golang.org/x/net/context"
)

type describeContextKey string

var (
	key            describeContextKey = "describe_ctx"
	triggerTypeKey string             = "trigger_type"
)

type DescribeContext struct {
	AccountID string
	Region    string
	OGRegion  string
	Partition string
}

func GetDescribeContext(ctx context.Context) DescribeContext {
	describe, ok := ctx.Value(key).(DescribeContext)
	if !ok {
		panic("context key not found")
	}
	return describe
}

func GetTriggerTypeFromContext(ctx context.Context) enums.DescribeTriggerType {
	tt, ok := ctx.Value(triggerTypeKey).(enums.DescribeTriggerType)
	if !ok {
		return ""
	}
	return tt
}

func WithTriggerType(ctx context.Context, tt enums.DescribeTriggerType) context.Context {
	return context.WithValue(ctx, triggerTypeKey, tt)
}
func WithDescribeContext(ctx context.Context, describeCtx DescribeContext) context.Context {
	return context.WithValue(ctx, key, describeCtx)
}
func ParallelDescribeRegionalSingleResource(describe func(context.Context, aws.Config, map[string]string) ([]model.Resource, error)) model.ResourceDescriber {
	type result struct {
		region    string
		resources []model.Resource
		err       error
		errorCode string
	}
	return func(ctx context.Context, config model.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalData map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		cfg, err := GetConfig(config)
		if err != nil {
			return nil, err
		}

		regionsObjs, err := getAllRegions(ctx, *cfg, false)
		if err != nil {
			return nil, err
		}
		var regions []string
		for _, region := range regionsObjs {
			regions = append(regions, *region.RegionName)
		}
		input := make(chan result, len(regions))
		for _, region := range regions {
			go func(r string) {
				defer func() {
					if err := recover(); err != nil {
						//stack := debug.Stack()
						//input <- result{region: r, resources: nil, err: fmt.Errorf("paniced: %v\n%s", err, string(stack))}
						input <- result{region: r, resources: nil, err: fmt.Errorf("paniced: %v", err)}
					}
				}()
				// Make a shallow copy and override the default region
				rCfg := cfg.Copy()
				rCfg.Region = r

				partition, _ := PartitionOf(r)
				ctx = WithDescribeContext(ctx, DescribeContext{
					AccountID: additionalData["accountID"],
					Region:    region,
					OGRegion:  region,
					Partition: partition,
				})
				ctx = WithTriggerType(ctx, triggerType)
				resources, err := describe(ctx, rCfg, map[string]string{})
				errCode := ""
				if err != nil {
					var ae smithy.APIError
					if errors.As(err, &ae) {
						errCode = ae.ErrorCode()
					}
				}
				input <- result{region: r, resources: resources, err: err, errorCode: errCode}
			}(region)
		}

		var output []model.Resource
		for range regions {
			resp := <-input
			if resp.err != nil {
				return nil, resp.err
			}

			if resp.resources == nil {
				resp.resources = []model.Resource{}
			}

			partition, _ := PartitionOf(resp.region)
			for i := range resp.resources {
				resp.resources[i].Account = additionalData["accountID"]
				resp.resources[i].Partition = partition
				resp.resources[i].Region = resp.region
			}

			output = append(output, resp.resources...)
		}

		return output, nil
	}
}

func SequentialDescribeRegional(describe func(context.Context, aws.Config, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, config model.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalData map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		cfg, err := GetConfig(config)
		if err != nil {
			return nil, err
		}

		regionsObjs, err := getAllRegions(ctx, *cfg, false)
		if err != nil {
			return nil, err
		}
		var regions []string
		for _, region := range regionsObjs {
			regions = append(regions, *region.RegionName)
		}
		var result []model.Resource
		for _, region := range regions {
			// Make a shallow copy and override the default region
			rCfg := cfg.Copy()
			rCfg.Region = region

			partition, _ := PartitionOf(region)
			ctx = WithDescribeContext(ctx, DescribeContext{
				AccountID: additionalData["accountID"],
				Region:    region,
				OGRegion:  region,
				Partition: partition,
			})
			ctx = WithTriggerType(ctx, triggerType)
			resources, err := describe(ctx, rCfg, stream)
			if err != nil {
				return nil, err
			}

			if resources == nil {
				resources = []model.Resource{}
			}

			for i := range resources {
				resources[i].Account = additionalData["accountID"]
			}

			result = append(result, resources...)
		}
		return result, nil
	}
}

// Parallel describe the resources across the reigons. Failure in one regions won't affect
// other regions.
func ParallelDescribeRegional(describe func(context.Context, aws.Config, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	type result struct {
		region    string
		resources []model.Resource
		err       error
		errorCode string
	}
	return func(ctx context.Context, config model.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalData map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		fmt.Println("ParallelDescribeRegional")
		cfg, err := GetConfig(config)
		if err != nil {
			return nil, err
		}

		regionsObjs, err := getAllRegions(ctx, *cfg, false)
		if err != nil {
			return nil, err
		}
		var regions []string
		for _, region := range regionsObjs {
			regions = append(regions, *region.RegionName)
		}
		input := make(chan result, len(regions))
		for _, region := range regions {
			go func(r string) {
				defer func() {
					if err := recover(); err != nil {
						//stack := debug.Stack()
						//input <- result{region: r, resources: nil, err: fmt.Errorf("paniced: %v\n%s", err, string(stack))}
						input <- result{region: r, resources: nil, err: fmt.Errorf("paniced: %v", err)}
					}
				}()
				// Make a shallow copy and override the default region
				rCfg := cfg.Copy()
				rCfg.Region = r

				partition, _ := PartitionOf(r)
				describeCtx := DescribeContext{
					AccountID: additionalData["accountID"],
					Region:    r,
					OGRegion:  r,
					Partition: partition,
				}

				fmt.Println("ParallelDescribeRegional for region", r, rCfg.Region, describeCtx.Region)
				ctx = WithDescribeContext(ctx, describeCtx)
				ctx = WithTriggerType(ctx, triggerType)
				fmt.Println("running describe")
				resources, err := describe(ctx, rCfg, stream)
				fmt.Println("describe finished", err)
				errCode := ""
				if err != nil {
					var ae smithy.APIError
					if errors.As(err, &ae) {
						errCode = ae.ErrorCode()
					}
				}
				input <- result{region: r, resources: resources, err: err, errorCode: errCode}
			}(region)
		}

		var output []model.Resource
		for range regions {
			fmt.Println("ParallelDescribeRegional waiting for result")
			resp := <-input
			fmt.Println("ParallelDescribeRegional got a result")
			if resp.err != nil {
				return nil, resp.err
			}

			if resp.resources == nil {
				resp.resources = []model.Resource{}
			}

			partition, _ := PartitionOf(resp.region)
			for i := range resp.resources {
				resp.resources[i].Account = additionalData["accountID"]
				resp.resources[i].Partition = partition
				resp.resources[i].Region = resp.region
			}

			output = append(output, resp.resources...)
		}
		fmt.Println("ParallelDescribeRegional finished")

		return output, nil
	}
}

// Sequentially describe the resources. If anyone of the regions fails, it will move on to the next region.
func SequentialDescribeGlobal(describe func(context.Context, aws.Config, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, config model.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalData map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		var output []model.Resource
		cfg, err := GetConfig(config)
		if err != nil {
			return nil, err
		}

		regionsObjs, err := getAllRegions(ctx, *cfg, false)
		if err != nil {
			return nil, err
		}
		var regions []string
		for _, region := range regionsObjs {
			regions = append(regions, *region.RegionName)
		}
		for _, region := range regions {
			// Make a shallow copy and override the default region
			rCfg := cfg.Copy()
			rCfg.Region = region

			partition, _ := PartitionOf(region)
			ctx = WithDescribeContext(ctx, DescribeContext{
				AccountID: additionalData["accountID"],
				Region:    region,
				OGRegion:  "global",
				Partition: partition,
			})
			ctx = WithTriggerType(ctx, triggerType)
			resources, err := describe(ctx, rCfg, stream)
			if err != nil {
				return nil, err
			}

			if resources == nil {
				resources = []model.Resource{}
			}

			for i := range resources {
				resources[i].Account = additionalData["accountID"]
				resources[i].Region = "global"
			}

			output = append(output, resources...)

			break
		}

		m := map[string]interface{}{}
		var newV []model.Resource
		for _, v := range output {
			if _, ok := m[v.UniqueID()]; ok {
				continue
			}

			m[v.UniqueID()] = struct{}{}
			newV = append(newV, v)
		}

		return newV, nil
	}
}
