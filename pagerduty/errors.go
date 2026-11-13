package pagerduty

import (
	"context"
	"errors"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func shouldRetryError(_ context.Context, _ *plugin.QueryData, _ *plugin.HydrateData, err error) bool {
	var aerr pagerduty.APIError

	if errors.As(err, &aerr) {
		return aerr.RateLimited()
	}
	return false
}

func isNotFoundError(err error) bool {
	var aerr pagerduty.APIError

	if errors.As(err, &aerr) {
		return aerr.NotFound()
	}
	return false
}
