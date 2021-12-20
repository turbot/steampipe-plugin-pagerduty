package pagerduty

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tablePagerDutyTeam(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pagerduty_team",
		Description: "Members of a PagerDuty account that have the ability to interact with incidents and other data on the account.",
		List: &plugin.ListConfig{
			Hydrate: listPagerDutyTeams,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getPagerDutyTeam,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the team.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "An unique identifier of a team.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "description",
				Description: "The description of the team.",
				Type:        proto.ColumnType_STRING,
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
				Name:        "members",
				Description: "A list of members of a team.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     listPagerDutyTeamMembers,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "tags",
				Description: "A list of tags applied on team.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     listPagerDutyTeamTags,
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

func listPagerDutyTeams(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_team.listPagerDutyTeams", "connection_error", err)
		return nil, err
	}

	req := pagerduty.ListTeamOptions{}

	// Additional Filters
	if d.KeyColumnQuals["name"] != nil {
		req.Query = d.KeyColumnQuals["name"].GetStringValue()
	}

	// Retrieve the list of teams
	maxResult := uint(100)

	// Reduce the basic request limit down if the user has only requested a small number of rows
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if uint(*limit) < maxResult {
			maxResult = uint(*limit)
		}
	}
	req.APIListObject.Limit = maxResult

	for {
		resp, err := client.ListTeamsWithContext(ctx, req)
		if err != nil {
			plugin.Logger(ctx).Error("pagerduty_team.listPagerDutyTeams", "query_error", err)
		}

		for _, team := range resp.Teams {
			d.StreamListItem(ctx, team)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if !resp.APIListObject.More {
			break
		}
		req.APIListObject.Offset = resp.Offset + 1
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getPagerDutyTeam(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_team.getPagerDutyTeam", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetStringValue()

	// No inputs
	if id == "" {
		return nil, nil
	}

	data, err := client.GetTeamWithContext(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_team.getPagerDutyTeam", "query_error", err)

		if isNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return *data, nil
}

func listPagerDutyTeamMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data := h.Item.(pagerduty.Team)

	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_team.listPagerDutyTeamMembers", "connection_error", err)
		return nil, err
	}

	resp, err := client.ListMembersPaginated(ctx, data.ID)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_team.listPagerDutyTeams", "query_error", err)
		return nil, err
	}

	return resp, nil
}

func listPagerDutyTeamTags(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data := h.Item.(pagerduty.Team)

	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_team.listPagerDutyTeamTags", "connection_error", err)
		return nil, err
	}

	resp, err := client.GetTagsForEntityPaginated(ctx, "teams", data.ID, pagerduty.ListTagOptions{})
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_team.listPagerDutyTeamTags", "query_error", err)
		return nil, err
	}

	return resp, nil
}
