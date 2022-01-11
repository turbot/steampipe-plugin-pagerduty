package pagerduty

import (
	"context"
	"time"

	"github.com/PagerDuty/go-pagerduty"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tablePagerDutyIncidentLog(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pagerduty_incident_log",
		Description: "Records log entries for the specified incident.",
		List: &plugin.ListConfig{
			Hydrate: listPagerDutyIncidentLogs,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "incident_id",
					Require: plugin.Required,
				},
				{
					Name:      "created_at",
					Require:   plugin.Optional,
					Operators: []string{">", ">=", "=", "<", "<="},
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "An unique identifier of the log entry.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromGo(),
			},
			{
				Name:        "incident_id",
				Description: "An unique identifier of the queried incident.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("incident_id"),
			},
			{
				Name:        "created_at",
				Description: "Time at which the log entry was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "summary",
				Description: "A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "acknowledgement_timeout",
				Description: "Specifies the acknowledgement timeout (in seconds) for the incident.",
				Type:        proto.ColumnType_INT,
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
				Name:        "type",
				Description: "The type of object being created.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "agent",
				Description: "The agent (user, service or integration) that created or modified the incident log entry.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "channel",
				Description: "Polymorphic object representation of the means by which the action was channeled. Has different formats depending on type, indicated by channel[type]. Will be one of auto, email, api, nagios, or timeout if agent[type] is service. Will be one of email, sms, website, web_trigger, or note if agent[type] is user.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "contexts",
				Description: "A list of contexts to be included with the trigger such as links to graphs, or images.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "event_details",
				Description: "A list of information about the change events.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "teams",
				Description: "A list of team references unless included.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("id"),
			},
		},
	}
}

//// LIST FUNCTION

func listPagerDutyIncidentLogs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_incident_log.listPagerDutyIncidentLogs", "connection_error", err)
		return nil, err
	}

	incidentID := d.KeyColumnQuals["incident_id"].GetStringValue()

	req := pagerduty.ListIncidentLogEntriesOptions{}

	quals := d.Quals
	if quals["created_at"] != nil {
		for _, q := range quals["created_at"].Quals {
			givenTime := q.Value.GetTimestampValue().AsTime().UTC()
			beforeTime := givenTime.Add(time.Duration(-1) * time.Second)
			afterTime := givenTime.Add(time.Second * 1)

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

	// Retrieve the list of incident logs
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
		incidentLogs, err := client.ListIncidentLogEntriesWithContext(ctx, incidentID, req)
		return incidentLogs, err
	}
	for {
		listPageResponse, err := plugin.RetryHydrate(ctx, d, h, listPage, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
		if err != nil {
			plugin.Logger(ctx).Error("pagerduty_incident_log.listPagerDutyIncidentLogs", "query_error", err)
			return nil, err
		}
		listResponse := listPageResponse.(*pagerduty.ListIncidentLogEntriesResponse)

		for _, logEntry := range listResponse.LogEntries {
			d.StreamListItem(ctx, logEntry)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if !listResponse.APIListObject.More {
			break
		}
		req.APIListObject.Offset = listResponse.Offset + 1
	}

	return nil, nil
}

func convertTimeString(t time.Time) string {
	return t.Format(time.RFC3339)
}
