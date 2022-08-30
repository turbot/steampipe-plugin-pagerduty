package pagerduty

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tablePagerDutySchedule(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pagerduty_schedule",
		Description: "A Schedule determines the time periods that users are On-Call.",
		List: &plugin.ListConfig{
			Hydrate: listPagerDutySchedules,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getPagerDutySchedule,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the schedule.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "An unique identifier of a schedule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "description",
				Description: "The description of the schedule.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "timezone",
				Description: "The time zone of the schedule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("TimeZone").NullIfZero(),
			},
			{
				Name:        "html_url",
				Description: "The API show URL at which the object is accessible.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("HTMLURL").NullIfZero(),
			},
			{
				Name:        "self",
				Description: "The API show URL at which the object is accessible.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "summary",
				Description: "A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "type",
				Description: "The type of object being created.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "escalation_policies",
				Description: "A list of the escalation policies that uses this schedule.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "final_schedule",
				Description: "Specifies the final schedule.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPagerDutySchedule,
			},
			{
				Name:        "override_sub_schedule",
				Description: "Specifies schedule overrides for a given time range.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPagerDutySchedule,
				Transform:   transform.FromField("OverrideSubschedule").NullIfZero(),
			},
			{
				Name:        "schedule_layers",
				Description: "A list of schedule layers.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPagerDutySchedule,
			},
			{
				Name:        "teams",
				Description: "A list of the teams on the schedule.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "users",
				Description: "A list of the users on the schedule.",
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

func listPagerDutySchedules(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_schedule.listPagerDutySchedules", "connection_error", err)
		return nil, err
	}

	req := pagerduty.ListSchedulesOptions{}

	// Additional Filters
	if d.KeyColumnQuals["name"] != nil {
		req.Query = d.KeyColumnQuals["name"].GetStringValue()
	}

	// Retrieve the list of schedules
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
		schedules, err := client.ListSchedulesWithContext(ctx, req)
		return schedules, err
	}
	for {
		listPageResponse, err := plugin.RetryHydrate(ctx, d, h, listPage, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
		if err != nil {
			plugin.Logger(ctx).Error("pagerduty_schedule.listPagerDutySchedules", "query_error", err)
			return nil, err
		}
		listResponse := listPageResponse.(*pagerduty.ListSchedulesResponse)

		for _, schedule := range listResponse.Schedules {
			d.StreamListItem(ctx, schedule)

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

// HYDRATE FUNCTIONS

func getPagerDutySchedule(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_schedule.getPagerDutySchedule", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		id = h.Item.(pagerduty.Schedule).ID
	} else {
		id = d.KeyColumnQuals["id"].GetStringValue()
	}

	// No inputs
	if id == "" {
		return nil, nil
	}

	getDetails := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		data, err := client.GetScheduleWithContext(ctx, id, pagerduty.GetScheduleOptions{})
		return data, err
	}
	getResponse, err := plugin.RetryHydrate(ctx, d, h, getDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_schedule.getPagerDutySchedule", "query_error", err)

		if isNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	getResp := getResponse.(*pagerduty.Schedule)

	return *getResp, nil
}
