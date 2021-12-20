package pagerduty

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tablePagerDutyService(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pagerduty_service",
		Description: "A service represents something you monitor (like a web service, email service, or database service).",
		List: &plugin.ListConfig{
			Hydrate: listPagerDutyServices,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getPagerDutyService,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the service.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "An unique identifier of a service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "status",
				Description: "The current state of the service.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "self",
				Description: "The API show URL at which the object is accessible.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The user-provided description of the service.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "auto_resolve_timeout",
				Description: "Time in seconds that an incident is automatically resolved if left open for that long. Value is null if the feature is disabled.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "acknowledgement_timeout",
				Description: "Time in seconds that an incident changes to the Triggered State after being Acknowledged. Value is null if the feature is disabled.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "alert_creation",
				Description: "Whether a service creates only incidents, or both alerts and incidents. A service must create alerts in order to enable incident merging.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "create_at",
				Description: "The date/time when this service was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "html_url",
				Description: "An URL at which the entity is uniquely displayed in the Web app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("HTMLURL"),
			},
			{
				Name:        "last_incident_timestamp",
				Description: "The date/time when the most recent incident was created for this service.",
				Type:        proto.ColumnType_TIMESTAMP,
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
				Name:        "alert_grouping_parameters",
				Description: "Defines how alerts on this service will be automatically grouped into incidents. Note that the alert grouping features are available only on certain plans.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "escalation_policy",
				Description: "Escalation policy associated with the service.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "incident_urgency_rule",
				Description: "A list of incident urgency rules.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "integrations",
				Description: "An array containing integrations that belong to this service. If integrations is passed as an argument, these are full objects - otherwise, these are references.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "scheduled_actions",
				Description: "An array containing scheduled actions for the service.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "support_hours",
				Description: "Defines the service's support hours",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "teams",
				Description: "The set of teams associated with this service.",
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

func listPagerDutyServices(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_service.listPagerDutyServices", "connection_error", err)
		return nil, err
	}

	req := pagerduty.ListServiceOptions{}

	// Additional Filters
	if d.KeyColumnQuals["name"] != nil {
		req.Query = d.KeyColumnQuals["name"].GetStringValue()
	}

	// Retrieve the list of users
	maxResult := uint(100)

	// Reduce the basic request limit down if the user has only requested a small number of rows
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if uint(*limit) < maxResult {
			maxResult = uint(*limit)
		}
	}
	req.APIListObject.Limit = maxResult

	resp, err := client.ListServicesPaginated(ctx, req)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_service.listPagerDutyServices", "query_error", err)
	}

	for _, service := range resp {
		d.StreamListItem(ctx, service)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getPagerDutyService(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_service.getPagerDutyService", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetStringValue()

	// No inputs
	if id == "" {
		return nil, nil
	}

	data, err := client.GetServiceWithContext(ctx, id, &pagerduty.GetServiceOptions{})
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_service.getPagerDutyService", "query_error", err)

		if isNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return *data, nil
}
