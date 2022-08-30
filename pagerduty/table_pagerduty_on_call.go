package pagerduty

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tablePagerDutyOnCall(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pagerduty_on_call",
		Description: "An on-call represents a contiguous unit of time for which a User will be on call for a given Escalation Policy and Escalation Rules.",
		List: &plugin.ListConfig{
			Hydrate: listPagerDutyOnCalls,
		},
		Columns: []*plugin.Column{
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
			{
				Name:        "escalation_policy",
				Description: "The escalation_policy object.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "schedule",
				Description: "The schedule object.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "user_on_call",
				Description: "The user object.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("User"),
			},
		},
	}
}

//// LIST FUNCTION

func listPagerDutyOnCalls(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_on_call.listPagerDutyOnCalls", "connection_error", err)
		return nil, err
	}

	req := pagerduty.ListOnCallOptions{}

	// Retrieve the list of on calls
	maxResult := uint(100)

	// Reduce the basic request limit down if the user has only requested a small number of rows
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if uint(*limit) < maxResult {
			maxResult = uint(*limit)
		}
	}
	req.APIListObject.Limit = maxResult

	listPage := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		data, err := client.ListOnCallsWithContext(ctx, req)
		return data, err
	}
	for {
		listPageResponse, err := plugin.RetryHydrate(ctx, d, h, listPage, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
		if err != nil {
			if isNotFoundError(err) {
				return nil, nil
			}
			plugin.Logger(ctx).Error("pagerduty_on_call.listPagerDutyOnCalls", "query_error", err)
			return nil, err
		}
		listResponse := listPageResponse.(*pagerduty.ListOnCallsResponse)

		for _, oncall := range listResponse.OnCalls {
			d.StreamListItem(ctx, oncall)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if !listResponse.APIListObject.More {
			break
		}
		req.APIListObject.Offset = listResponse.APIListObject.Offset + listResponse.APIListObject.Limit
	}

	return nil, nil
}
