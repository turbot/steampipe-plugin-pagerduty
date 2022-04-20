package pagerduty

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

//// TABLE DEFINITION

func tablePagerDutyEscalationPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pagerduty_escalation_policy",
		Description: "An escalation policy determines what user or schedule will be notified first, second, and so on when an incident is triggered.",
		List: &plugin.ListConfig{
			Hydrate: listPagerDutyEscalationPolicies,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getPagerDutyEscalationPolicy,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the escalation policy.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "An unique identifier of an escalation policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "description",
				Description: "A shortened description of escalation policy.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "html_url",
				Description: "An URL at which the entity is uniquely displayed in the Web app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("HTMLURL").NullIfZero(),
			},
			{
				Name:        "num_loops",
				Description: "The number of times the escalation policy will repeat after reaching the end of its escalation.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("NumLoops"),
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
				Name:        "escalation_rules",
				Description: "A list of escalation rules.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "services",
				Description: "A list of services associated with the policy.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "teams",
				Description: "A list of teams associated with the policy.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tags",
				Description: "A list of tags applied on escalation policy.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     listPagerDutyEscalationPolicyTags,
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

func listPagerDutyEscalationPolicies(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_escalation_policy.listPagerDutyEscalationPolicies", "connection_error", err)
		return nil, err
	}
	req := pagerduty.ListEscalationPoliciesOptions{}

	// Additional Filters
	if d.KeyColumnQuals["name"] != nil {
		req.Query = d.KeyColumnQuals["name"].GetStringValue()
	}

	// Retrieve the list of policies
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
		policies, err := client.ListEscalationPoliciesWithContext(ctx, req)
		return policies, err
	}
	for {
		listPageResponse, err := plugin.RetryHydrate(ctx, d, h, listPage, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
		if err != nil {
			plugin.Logger(ctx).Error("pagerduty_escalation_policy.listPagerDutyEscalationPolicies", "query_error", err)
			return nil, err
		}
		listResponse := listPageResponse.(*pagerduty.ListEscalationPoliciesResponse)

		for _, policy := range listResponse.EscalationPolicies {
			d.StreamListItem(ctx, policy)

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

//// HYDRATE FUNCTIONS

func getPagerDutyEscalationPolicy(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_escalation_policy.getPagerDutyEscalationPolicy", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetStringValue()

	// No inputs
	if id == "" {
		return nil, nil
	}

	getDetails := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		data, err := client.GetEscalationPolicyWithContext(ctx, id, &pagerduty.GetEscalationPolicyOptions{})
		return data, err
	}
	getResponse, err := plugin.RetryHydrate(ctx, d, h, getDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_escalation_policy.getPagerDutyEscalationPolicy", "query_error", err)

		if isNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	getResp := getResponse.(*pagerduty.EscalationPolicy)

	return *getResp, nil
}

func listPagerDutyEscalationPolicyTags(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data := h.Item.(pagerduty.EscalationPolicy)

	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_escalation_policy.listPagerDutyEscalationPolicyTags", "connection_error", err)
		return nil, err
	}

	getDetails := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		data, err := client.GetTagsForEntityPaginated(ctx, "escalation_policies", data.ID, pagerduty.ListTagOptions{})
		return data, err
	}
	getResponse, err := plugin.RetryHydrate(ctx, d, h, getDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_escalation_policy.listPagerDutyEscalationPolicyTags", "query_error", err)
		return nil, err
	}
	getResp := getResponse.([]*pagerduty.Tag)

	return getResp, nil
}
