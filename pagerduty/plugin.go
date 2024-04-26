package pagerduty

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-pagerduty"

// Plugin creates this (pagerduty) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().Transform(transform.NullIfZeroValue),
		DefaultGetConfig: &plugin.GetConfig{},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		TableMap: map[string]*plugin.Table{
			"pagerduty_escalation_policy":   tablePagerDutyEscalationPolicy(ctx),
			"pagerduty_incident":            tablePagerDutyIncident(ctx),
			"pagerduty_incident_log":        tablePagerDutyIncidentLog(ctx),
			"pagerduty_on_call":             tablePagerDutyOnCall(ctx),
			"pagerduty_priority":            tablePagerDutyPriority(ctx),
			"pagerduty_ruleset":             tablePagerDutyRuleset(ctx),
			"pagerduty_ruleset_rule":        tablePagerDutyRulesetRule(ctx),
			"pagerduty_schedule":            tablePagerDutySchedule(ctx),
			"pagerduty_service":             tablePagerDutyService(ctx),
			"pagerduty_business_service":    tablePagerDutyBusinessService(ctx),
			"pagerduty_service_integration": tablePagerDutyServiceIntegration(ctx),
			"pagerduty_tag":                 tablePagerDutyTag(ctx),
			"pagerduty_team":                tablePagerDutyTeam(ctx),
			"pagerduty_user":                tablePagerDutyUser(ctx),
			"pagerduty_vendor":              tablePagerDutyVendor(ctx),
		},
	}

	return p
}
