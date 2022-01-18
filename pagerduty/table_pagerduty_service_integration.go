package pagerduty

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tablePagerDutyServiceIntegration(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pagerduty_service_integration",
		Description: "A service integration is an integration that belongs to a Pagerduty service.",
		List: &plugin.ListConfig{
			ParentHydrate: listPagerDutyServices,
			Hydrate:       listPagerDutyServiceIntegrations,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "service_id",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getPagerDutyServiceIntegration,
			KeyColumns: plugin.AllColumns([]string{"service_id", "id"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of this integration.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "An unique identifier of the integration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "service_id",
				Description: "An unique identifier of the queried service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Service.ID"),
			},
			{
				Name:        "created_at",
				Description: "The date/time when this integration was created.",
				Type:        proto.ColumnType_TIMESTAMP,
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
				Name:        "integration_key",
				Description: "Specify the integration key for the service integration.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "integration_email",
				Description: "Specify for generic_email_inbound_integration. Must be set to an email address @your-subdomain.pagerduty.com.",
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
				Name:        "service",
				Description: "Describes the information about the queried service.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "vendor",
				Description: "Describes the information about a specific type of integration.",
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

func listPagerDutyServiceIntegrations(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Get service details
	serviceData := h.Item.(pagerduty.Service)

	// Return if the service doesn't match with the specified one
	if d.KeyColumnQuals["service_id"] != nil && d.KeyColumnQuals["service_id"].GetStringValue() != serviceData.ID {
		return nil, nil
	}

	for _, integration := range serviceData.Integrations {
		d.StreamListItem(ctx, integration)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getPagerDutyServiceIntegration(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_service_integration.getPagerDutyServiceIntegration", "connection_error", err)
		return nil, err
	}
	serviceID := d.KeyColumnQuals["service_id"].GetStringValue()
	id := d.KeyColumnQuals["id"].GetStringValue()

	// No inputs
	if id == "" || serviceID == "" {
		return nil, nil
	}

	data, err := client.GetIntegrationWithContext(ctx, serviceID, id, pagerduty.GetIntegrationOptions{})
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_service_integration.getPagerDutyServiceIntegration", "query_error", err)

		if isNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return *data, nil
}
