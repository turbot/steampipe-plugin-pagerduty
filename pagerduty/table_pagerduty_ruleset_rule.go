package pagerduty

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

//// TABLE DEFINITION

func tablePagerDutyRulesetRule(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pagerduty_ruleset_rule",
		Description: "An event rule allows you to set actions that should be taken on events that meet your designated rule criteria.",
		List: &plugin.ListConfig{
			ParentHydrate: listPagerDutyRulesets,
			Hydrate:       listPagerDutyRulesetRules,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getPagerDutyRulesetRule,
			KeyColumns: plugin.AllColumns([]string{"ruleset_id", "id"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The ID of the event rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "ruleset_id",
				Description: "The ID of the ruleset.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RulesetID"),
			},
			{
				Name:        "disabled",
				Description: "Indicates whether the Event Rule is disabled and would therefore not be evaluated.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "position",
				Description: "Position/index of the Event Rule in the Ruleset. Starting from position 0 (the first rule), rules are evaluated one-by-one until a matching rule is found.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Position"),
			},
			{
				Name:        "catch_all",
				Description: "Indicates whether the Event Rule is the last Event Rule of the Ruleset that serves as a catch-all. It has limited functionality compared to other rules and always matches.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "self",
				Description: "The API show URL at which the object is accessible.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "actions",
				Description: "A set of actions that defines when an event matches this rule, the actions that will be taken to change the resulting alert and incident.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "conditions",
				Description: "A set of information defined the conditions resulting alert and incident.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "time_frame",
				Description: "Time-based conditions for limiting when the rule is active.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
		},
	}
}

type rulesetRuleInfo = struct {
	pagerduty.RulesetRule
	RulesetID string
}

//// LIST FUNCTION

func listPagerDutyRulesetRules(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Ruleset's information
	rulesetData := h.Item.(*pagerduty.Ruleset)

	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_ruleset_rule.listPagerDutyRulesetRules", "connection_error", err)
		return nil, err
	}

	resp, err := client.ListRulesetRulesPaginated(ctx, rulesetData.ID)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_ruleset_rule.listPagerDutyRulesetRules", "query_error", err)
		return nil, err
	}

	for _, rules := range resp {
		d.StreamListItem(ctx, rulesetRuleInfo{*rules, rulesetData.ID})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getPagerDutyRulesetRule(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_ruleset_rule.getPagerDutyRulesetRule", "connection_error", err)
		return nil, err
	}
	rulesetID := d.KeyColumnQuals["ruleset_id"].GetStringValue()
	ruleID := d.KeyColumnQuals["id"].GetStringValue()

	// No inputs
	if rulesetID == "" || ruleID == "" {
		return nil, nil
	}

	data, err := client.GetRulesetRuleWithContext(ctx, rulesetID, ruleID)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_ruleset_rule.getPagerDutyRulesetRule", "query_error", err)

		if isNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return rulesetRuleInfo{*data, rulesetID}, nil
}
