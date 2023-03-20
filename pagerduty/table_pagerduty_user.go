package pagerduty

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablePagerDutyUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pagerduty_user",
		Description: "Members of a PagerDuty account that have the ability to interact with incidents and other data on the account.",
		List: &plugin.ListConfig{
			Hydrate: listPagerDutyUsers,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "email",
					Require: plugin.Optional,
				},
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getPagerDutyUser,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "An unique identifier of an user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "email",
				Description: "The user's email address.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "role",
				Description: "The user role. Account must have the 'read_only_users' ability to set a user as a 'read_only_user' or a 'read_only_limited_user', and must have advanced permissions abilities to set a user as 'observer' or 'restricted_access'.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "invitation_sent",
				Description: "If true, the user has an outstanding invitation.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "avatar_url",
				Description: "The URL of the user's avatar.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AvatarURL").NullIfZero(),
			},
			{
				Name:        "color",
				Description: "The schedule color.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The user's bio.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "html_url",
				Description: "An URL at which the entity is uniquely displayed in the Web app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("HTMLURL").NullIfZero(),
			},
			{
				Name:        "job_title",
				Description: "The user's job title.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "self",
				Description: "The API show URL at which the object is accessible.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "summary",
				Description: "A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to 'name', though it is not intended to be an identifier.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "timezone",
				Description: "The preferred time zone name. If null, the account's time zone will be used.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "type",
				Description: "The type of object being created.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "contact_methods",
				Description: "A list of contact methods for the user.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "notification_rules",
				Description: "A list of notification rules for the user.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tags",
				Description: "A list of tags applied on user.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     listPagerDutyUserTags,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "teams",
				Description: "A list of teams to which the user belongs. Account must have the teams ability to set this.",
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

func listPagerDutyUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_user.listPagerDutyUsers", "connection_error", err)
		return nil, err
	}

	req := pagerduty.ListUsersOptions{}

	// Additional Filters
	if d.EqualsQuals["email"] != nil {
		req.Query = d.EqualsQuals["email"].GetStringValue()
	}
	if d.EqualsQuals["name"] != nil {
		req.Query = d.EqualsQuals["name"].GetStringValue()
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

	// Check for additional models to include in response
	// for example, contact_methods, notification_rules, teams
	givenColumns := d.QueryContext.Columns
	includeFields := buildUserRequestFields(ctx, givenColumns)

	if len(includeFields) > 0 {
		req.Includes = includeFields
	}

	listPage := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		users, err := client.ListUsersWithContext(ctx, req)
		return users, err
	}
	for {
		listPageResponse, err := plugin.RetryHydrate(ctx, d, h, listPage, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
		if err != nil {
			plugin.Logger(ctx).Error("pagerduty_user.listPagerDutyUsers", "query_error", err)
			return nil, err
		}
		listResponse := listPageResponse.(*pagerduty.ListUsersResponse)

		for _, user := range listResponse.Users {
			d.StreamListItem(ctx, user)

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

func getPagerDutyUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_user.getPagerDutyUser", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetStringValue()

	// No inputs
	if id == "" {
		return nil, nil
	}

	getDetails := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		data, err := client.GetUserWithContext(ctx, id, pagerduty.GetUserOptions{})
		return data, err
	}
	getResponse, err := plugin.RetryHydrate(ctx, d, h, getDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_user.getPagerDutyUser", "query_error", err)

		if isNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	getResp := getResponse.(*pagerduty.User)

	return *getResp, nil
}

func listPagerDutyUserTags(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data := h.Item.(pagerduty.User)

	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_user.listPagerDutyUserTags", "connection_error", err)
		return nil, err
	}

	getDetails := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		data, err := client.GetTagsForEntityPaginated(ctx, "users", data.ID, pagerduty.ListTagOptions{})
		return data, err
	}
	getResponse, err := plugin.RetryHydrate(ctx, d, h, getDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_user.listPagerDutyUserTags", "query_error", err)
		return nil, err
	}
	getResp := getResponse.([]*pagerduty.Tag)

	return getResp, nil
}

func buildUserRequestFields(ctx context.Context, queryColumns []string) []string {
	var fields []string
	for _, columnName := range queryColumns {
		if columnName == "contact_methods" || columnName == "notification_rules" || columnName == "teams" {
			fields = append(fields, columnName)
		}
	}
	return fields
}
