package pagerduty

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type pagerDutyConfig struct {
	Token *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &pagerDutyConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) pagerDutyConfig {
	if connection == nil || connection.GetConfig() == nil {
		return pagerDutyConfig{}
	}
	config, _ := connection.GetConfig().(pagerDutyConfig)
	return config
}
