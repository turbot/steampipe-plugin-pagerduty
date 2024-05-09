package pagerduty

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablePagerDutyPriority(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pagerduty_priority",
		Description: "A priority is a label representing the importance and impact of an incident.",
		List: &plugin.ListConfig{
			Hydrate: listPagerDutyPriorities,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The user-provided short name of the priority.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "The ID of the priority.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "description",
				Description: "The user-provided description of the priority.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "self",
				Description: "The API show URL at which the object is accessible.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "html_url",
				Description: "An URL at which the entity is uniquely displayed in the Web app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("HTMLURL").NullIfZero(),
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

func listPagerDutyPriorities(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_priority.listPagerDutyPriorities", "connection_error", err)
		return nil, err
	}

	listPage := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		data, err := client.ListPrioritiesWithContext(ctx, pagerduty.ListPrioritiesOptions{})
		return data, err
	}
	listResponse, err := plugin.RetryHydrate(ctx, d, h, listPage, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
	if err != nil {
		// If incident priority level is not enabled, API returns 404 Not Found error
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("pagerduty_priority.listPagerDutyPriorities", "query_error", err)
		return nil, err
	}
	resp := listResponse.(*pagerduty.Priorities)

	for _, priority := range resp.Priorities {
		d.StreamListItem(ctx, priority)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
