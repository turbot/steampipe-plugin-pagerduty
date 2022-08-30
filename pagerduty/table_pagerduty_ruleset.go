package pagerduty

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tablePagerDutyRuleset(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pagerduty_ruleset",
		Description: "Rulesets allow you to route events to an endpoint and create collections of event rules, which define sets of actions to take based on event content.",
		List: &plugin.ListConfig{
			Hydrate: listPagerDutyRulesets,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getPagerDutyRuleset,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the ruleset.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "An unique identifier of a ruleset.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "type",
				Description: "The type of the ruleset. Allowed values are: 'global' and 'default_global'.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creator",
				Description: "A set of information about the user who created the ruleset.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "routing_keys",
				Description: "A list of routing keys for this ruleset.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "team",
				Description: "A set of information about the team that owns the ruleset.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "updater",
				Description: "A set information about the user that has updated the ruleset.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func listPagerDutyRulesets(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_ruleset.listPagerDutyRulesets", "connection_error", err)
		return nil, err
	}

	resp, err := client.ListRulesetsPaginated(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_ruleset.listPagerDutyRulesets", "query_error", err)
		return nil, err
	}

	for _, ruleset := range resp {
		d.StreamListItem(ctx, ruleset)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getPagerDutyRuleset(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_ruleset.getPagerDutyRuleset", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetStringValue()

	// No inputs
	if id == "" {
		return nil, nil
	}

	data, err := client.GetRulesetWithContext(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_ruleset.getPagerDutyRuleset", "query_error", err)

		if isNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return *data, nil
}
