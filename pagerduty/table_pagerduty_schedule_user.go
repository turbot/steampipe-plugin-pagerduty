package pagerduty

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePagerDutyScheduleUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pagerduty_schedule_user",
		Description: "List users who are part of a PagerDuty schedule rotation.",
		List: &plugin.ListConfig{
			ParentHydrate: listPagerDutySchedules,
			Hydrate:       listPagerDutyScheduleUsers,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "schedule_id", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Schedule Columns
			{
				Name:        "schedule_id",
				Description: "The ID of the schedule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ScheduleID"),
			},
			{
				Name:        "schedule_name",
				Description: "The name of the schedule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ScheduleName"),
			},

			// User Columns
			{
				Name:        "id",
				Description: "The ID of the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.ID"),
			},
			{
				Name:        "name",
				Description: "The name of the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.Name"),
			},
			{
				Name:        "email",
				Description: "The user's email address.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.Email"),
			},
			{
				Name:        "role",
				Description: "The user's role.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.Role"),
			},
			{
				Name:        "job_title",
				Description: "The user's job title.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.JobTitle"),
			},
			{
				Name:        "description",
				Description: "The user's description.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.Description"),
			},
			{
				Name:        "timezone",
				Description: "The user's timezone.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.Timezone"),
			},
			{
				Name:        "color",
				Description: "The user's color preference.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.Color"),
			},
			{
				Name:        "avatar_url",
				Description: "URL of the user's avatar.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.AvatarURL"),
			},
			{
				Name:        "html_url",
				Description: "URL at which the user can be viewed in the web app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.HTMLURL"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.Name"),
			},
		},
	}
}

type ScheduleUser struct {
	ScheduleID   string
	ScheduleName string
	User         *pagerduty.User
}

func listPagerDutyScheduleUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	schedule := h.Item.(pagerduty.Schedule)

	if d.EqualsQuals["schedule_id"] != nil && d.EqualsQuals["schedule_id"].GetStringValue() != schedule.ID {
		return nil, nil
	}

	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_schedule_user.listPagerDutyScheduleUsers", "connection_error", err)
		return nil, err
	}

	// List on-call users for the schedule
	listPage := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		users, err := client.ListOnCallUsersWithContext(ctx, schedule.ID, pagerduty.ListOnCallUsersOptions{})
		return users, err
	}
	listResponse, err := plugin.RetryHydrate(ctx, d, h, listPage, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_schedule_user.listPagerDutyScheduleUsers", "query_error", err)
		return nil, err
	}

	users := listResponse.([]pagerduty.User)
	for _, user := range users {
		scheduleUser := &ScheduleUser{
			ScheduleID:   schedule.ID,
			ScheduleName: schedule.Name,
			User:         &user,
		}

		d.StreamListItem(ctx, scheduleUser)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
