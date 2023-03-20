package main

import (
	"github.com/turbot/steampipe-plugin-pagerduty/pagerduty"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: pagerduty.Plugin})
}
