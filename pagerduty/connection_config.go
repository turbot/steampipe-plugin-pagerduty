package pagerduty

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type pagerDutyConfig struct {
	Token *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
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
