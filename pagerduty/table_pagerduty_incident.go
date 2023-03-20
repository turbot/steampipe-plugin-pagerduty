package pagerduty

import (
	"context"
	"time"

	"github.com/PagerDuty/go-pagerduty"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablePagerDutyIncident(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pagerduty_incident",
		Description: "An incident represents a problem or an issue that needs to be addressed and resolved.",
		List: &plugin.ListConfig{
			Hydrate: listPagerDutyIncidents,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "created_at",
					Require:   plugin.Optional,
					Operators: []string{">", ">=", "=", "<", "<="},
				},
				{
					Name:    "incident_key",
					Require: plugin.Optional,
				},
				{
					Name:    "status",
					Require: plugin.Optional,
				},
				{
					Name:    "urgency",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getPagerDutyIncident,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "An unique identifier of the incident.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "incident_number",
				Description: "The number of the incident. This is unique across your account.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "status",
				Description: "The current status of the incident.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "urgency",
				Description: "The current urgency of the incident.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "summary",
				Description: "A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "The date/time the incident was first triggered.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "description",
				Description: "The description of the incident.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "html_url",
				Description: "The API show URL at which the object is accessible.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("HTMLURL").NullIfZero(),
			},
			{
				Name:        "incident_key",
				Description: "The incident's de-duplication key.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_mergeable",
				Description: "Indicates whether the incident's alerts can be merged with another incident, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "last_status_change_at",
				Description: "The time at which the status of the incident last changed.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "self",
				Description: "The API show URL at which the object is accessible.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "type",
				Description: "The type of object being created.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "acknowledgements",
				Description: "A list of all acknowledgements for this incident. This list will be empty if the 'Incident.status' is resolved or triggered.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "alert_counts",
				Description: "Describes the count of triggered and resolved alerts.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "assignments",
				Description: "A list of all assignments for this incident. This list will be empty if the 'Incident.status' is resolved.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "body",
				Description: "Describes the additional incident details.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "conference_bridge",
				Description: "Specifies the contact information that allows responders to easily connect and collaborate during major incident response.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "escalation_policy",
				Description: "Specifies the escalation policy assigned to this incident.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "first_trigger_log_entry",
				Description: "Specifies the first log entry when the incident was triggered.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "last_status_change_by",
				Description: "The agent (user, service or integration) that created or modified the incident log entry.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "pending_actions",
				Description: "A list of pending_actions on the incident. A pending_action object contains a type of action which can be escalate, unacknowledge, resolve or urgency_change. A pending_action object contains at, the time at which the action will take place. An urgency_change pending_action will contain to, the urgency that the incident will change to.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "priority",
				Description: "Specifies the priority set for this incident.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "resolve_reason",
				Description: "Specifies the reason the incident was resolved.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "service",
				Description: "Specifies the information about the impacted service.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "teams",
				Description: "The teams involved in the incident's lifecycle.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Summary"),
			},
		},
	}
}

//// LIST FUNCTION

func listPagerDutyIncidents(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_incident.listPagerDutyIncidents", "connection_error", err)
		return nil, err
	}

	req := pagerduty.ListIncidentsOptions{}

	// Additional Filters
	if d.EqualsQuals["incident_key"] != nil {
		req.IncidentKey = d.EqualsQuals["incident_key"].GetStringValue()
	}
	if d.EqualsQuals["status"] != nil {
		req.Statuses = []string{d.EqualsQuals["status"].GetStringValue()}
	}
	if d.EqualsQuals["urgency"] != nil {
		req.Urgencies = []string{d.EqualsQuals["urgency"].GetStringValue()}
	}

	quals := d.Quals
	if quals["created_at"] != nil {
		for _, q := range quals["created_at"].Quals {
			givenTime := q.Value.GetTimestampValue().AsTime().UTC()
			beforeTime := givenTime.Add(time.Duration(-1) * time.Second)
			afterTime := givenTime.Add(time.Second * 1)

			// API doesn't supports listing incidents beyond 6 months
			// if the queried range is more than 6 months, set `date_range` attribute to 'all'
			currentTime := time.Now().UTC()
			diffInDays := (currentTime.Sub(givenTime).Hours()) / 24
			if diffInDays > 180 {
				req.DateRange = "all"
				break
			}

			switch q.Operator {
			case ">":
				req.Since = convertTimeString(afterTime)
			case ">=":
				req.Since = convertTimeString(givenTime)
			case "=":
				req.Since = convertTimeString(beforeTime)
				req.Until = convertTimeString(afterTime)
			case "<=":
				req.Until = convertTimeString(afterTime)
			case "<":
				req.Until = convertTimeString(givenTime)
			}
		}
	}

	// Retrieve the list of incidents
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
		incidents, err := client.ListIncidentsWithContext(ctx, req)
		return incidents, err
	}
	for {
		listPageResponse, err := plugin.RetryHydrate(ctx, d, h, listPage, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
		if err != nil {
			plugin.Logger(ctx).Error("pagerduty_incident.listPagerDutyIncidents", "query_error", err)
			return nil, err
		}
		listResponse := listPageResponse.(*pagerduty.ListIncidentsResponse)

		for _, incident := range listResponse.Incidents {
			d.StreamListItem(ctx, incident)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
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

//// HYDRATE FUNCTIONS

func getPagerDutyIncident(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_incident.getPagerDutyIncident", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetStringValue()

	// No inputs
	if id == "" {
		return nil, nil
	}

	getDetails := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		data, err := client.GetIncidentWithContext(ctx, id)
		return data, err
	}
	getResponse, err := plugin.RetryHydrate(ctx, d, h, getDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_incident.getPagerDutyIncident", "query_error", err)

		if isNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	getResp := getResponse.(*pagerduty.Incident)

	return *getResp, nil
}

func convertTimeString(t time.Time) string {
	return t.Format(time.RFC3339)
}
