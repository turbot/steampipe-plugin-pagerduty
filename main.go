package main

import (
	"github.com/turbot/steampipe-plugin-pagerduty/pagerduty"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: pagerduty.Plugin})
}
