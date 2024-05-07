package pagerduty

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablePagerDutyBusinessService(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pagerduty_business_service",
		Description: "Business services model capabilities that span multiple technical services and that may be owned by several different teams.",
		List: &plugin.ListConfig{
			Hydrate: listPagerDutyBusinessServices,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getPagerDutyBusinessService,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the business service.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "An unique identifier of a business service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "self",
				Description: "The API show URL at which the object is accessible.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The user-provided description of the business service.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "html_url",
				Description: "An URL at which the entity is uniquely displayed in the Web app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("HTMLURL").NullIfZero(),
			},
			{
				Name:        "point_of_contact",
				Description: "The point of contact assigned to this service.\n\n",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "team",
				Description: "Reference to the team that owns the business service.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "dependencies",
				Description: "Immediate dependencies of the business service.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     hydrateBusinessServiceDependencies,
				Transform:   transform.FromValue(),
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

func listPagerDutyBusinessServices(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_service.listPagerDutyBusinessServices", "connection_error", err)
		return nil, err
	}

	req := pagerduty.ListBusinessServiceOptions{}

	// Retrieve the list of services
	maxResult := uint(100)

	// Reduce the basic request limit down if the user has only requested a small number of rows
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if uint(*limit) < maxResult {
			maxResult = uint(*limit)
		}
	}
	req.APIListObject.Limit = maxResult

	resp, err := client.ListBusinessServicesPaginated(ctx, req)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_service.listPagerDutyBusinessServices", "query_error", err)
		return nil, err
	}

	for _, service := range resp {
		d.StreamListItem(ctx, service)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getPagerDutyBusinessService(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_service.getPagerDutyBusinessService", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetStringValue()

	// No inputs
	if id == "" {
		return nil, nil
	}

	data, err := client.GetBusinessServiceWithContext(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_service.getPagerDutyBusinessService", "query_error", err)

		if isNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return data, nil
}

func hydrateBusinessServiceDependencies(ctx context.Context, queryData *plugin.QueryData, hydrateData *plugin.HydrateData) (interface{}, error) {
	client, err := getSessionConfig(ctx, queryData)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_service.hydrateBusinessServiceDependencies", "connection_error", err)
		return nil, err
	}

	resp, err := client.GetBusinessServiceDependencies(ctx, hydrateData.Item.(*pagerduty.BusinessService).ID)
	if err != nil {
		return nil, err
	}

	return resp["relationships"], nil
}
