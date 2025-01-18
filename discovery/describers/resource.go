package describers

import (
	"strings"
)

type Resource struct {
	// ARN uniquely identifies an AWS resource across regions, accounts and types.
	ARN string
	// ID doesn't uniquely identifies a resource. It will be used to create a
	// unique identifier by concating PARTITION|REGION|ACCOUNT|TYPE|ID
	ID          string
	Description interface{}

	Name      string
	Account   string
	Region    string
	Partition string
	Type      string
}

func CompositeID(list ...string) string {
	normList := make([]string, 0, len(list))
	for _, v := range list {
		v = strings.TrimSpace(v)
		v = strings.ToLower(v)
		if v == "" {
			continue
		}

		normList = append(normList, v)

	}
	return strings.Join(normList, "|")
}

// nameFromArn returns the name from ARN.
// It **assumes** that the name of the resource is always the last element in the arn.
func nameFromArn(arn string) string {
	parts := strings.Split(arn, ":")
	name := parts[len(parts)-1]

	parts = strings.Split(name, "/")
	name = parts[len(parts)-1]

	return name
}
