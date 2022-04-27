package pagerduty

import (
	"context"
	"fmt"
	"os"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

// getSessionConfig :: returns PagerDuty client to perform API requests
func getSessionConfig(ctx context.Context, d *plugin.QueryData) (*pagerduty.Client, error) {
	// Load clientOptions from cache
	sessionCacheKey := "pagerduty.clientoption"
	if cachedData, ok := d.ConnectionManager.Cache.Get(sessionCacheKey); ok {
		return cachedData.(*pagerduty.Client), nil
	}

	// Get pagerduty config
	pagerDutyConfig := GetConfig(d.Connection)

	// Get the authorization token
	token := os.Getenv("PAGERDUTY_TOKEN")
	if pagerDutyConfig.Token != nil {
		token = *pagerDutyConfig.Token
	}

	// No creds
	if token == "" {
		return nil, fmt.Errorf("token must be configured")
	}

	// Create client
	client := pagerduty.NewClient(token)

	// save clientOptions in cache
	d.ConnectionManager.Cache.Set(sessionCacheKey, client)

	return client, nil
}
