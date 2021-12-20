package pagerduty

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type pagerDutyConfig struct {
	ApiUrlOverride *string `cty:"api_url_override"`
	Token          *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
	"api_url_override": {
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
