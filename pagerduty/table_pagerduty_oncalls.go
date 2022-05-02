package pagerduty

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

//// TABLE DEFINITION

func tablePagerDutyOnCalls(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pagerduty_oncalls",
		Description: "An on-call represents a contiguous unit of time for which a User will be on call for a given Escalation Policy and Escalation Rules.",
		List: &plugin.ListConfig{
			Hydrate: listPagerDutyOnCalls,
		},
		Columns: []*plugin.Column{
			{
				Name:        "escalation_policy",
				Description: "The escalation_policy object.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "user",
				Description: "The user object.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "schedule",
				Description: "The schedule object.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "escalation_level",
				Description: "The escalation level for the on-call.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "start",
				Description: "The start of the on-call. If null, the on-call is a permanent user on-call.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "end",
				Description: "The end of the on-call. If null, the user does not go off-call.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
		},
	}
}

//// LIST FUNCTION

func listPagerDutyOnCalls(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_oncalls.listPagerDutyOnCalls", "connection_error", err)
		return nil, err
	}

	listPage := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		data, err := client.ListOnCallsWithContext(ctx, pagerduty.ListOnCallOptions{})
		return data, err
	}
	listResponse, err := plugin.RetryHydrate(ctx, d, h, listPage, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("pagerduty_oncalls.listPagerDutyOnCalls", "query_error", err)
		return nil, err
	}
	resp := listResponse.(*pagerduty.ListOnCallsResponse)

	for _, oncall := range resp.OnCalls {
		d.StreamListItem(ctx, oncall)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
