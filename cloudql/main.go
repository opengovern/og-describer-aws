package main

import (
	"github.com/opengovern/og-describer-aws/cloudql/aws"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: aws.Plugin})
}
