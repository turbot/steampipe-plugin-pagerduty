package pagerduty

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type pagerDutyConfig struct {
	Token *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &pagerDutyConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) pagerDutyConfig {
	if connection == nil || connection.Config == nil {
		return pagerDutyConfig{}
	}
	config, _ := connection.Config.(pagerDutyConfig)
	return config
}
