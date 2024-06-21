package pagerduty

import (
	"context"
	"fmt"
	"os"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// getSessionConfig :: returns PagerDuty client to perform API requests
func getSessionConfig(ctx context.Context, d *plugin.QueryData) (*ExtendedPagerDutyClient, error) {
	// Load clientOptions from cache
	sessionCacheKey := "pagerduty.clientoption"
	if cachedData, ok := d.ConnectionManager.Cache.Get(sessionCacheKey); ok {
		return cachedData.(*ExtendedPagerDutyClient), nil
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
	client := NewExtendedPagerDutyClient(pagerduty.NewClient(token), token)

	// save clientOptions in cache
	d.ConnectionManager.Cache.Set(sessionCacheKey, client)

	return client, nil
}
