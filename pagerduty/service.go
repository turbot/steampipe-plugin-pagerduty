package pagerduty

import (
	"context"
	"fmt"
	"os"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

// getSessionConfig :: returns PagerDuty client to perform API requests
func getSessionConfig(ctx context.Context, d *plugin.QueryData) (*pagerduty.Client, error) {
	// Load clientOptions from cache
	sessionCacheKey := "pagerduty.clientoption"
	if cachedData, ok := d.ConnectionManager.Cache.Get(sessionCacheKey); ok {
		return cachedData.(*pagerduty.Client), nil
	}

	// Get scaleway config
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

	opts := []pagerduty.ClientOptions{}
	
	// Override default PagerDuty Base URL. Default is "https://api.pagerduty.com"
	if pagerDutyConfig.ApiUrlOverride != nil {
		opts = append(opts, pagerduty.WithAPIEndpoint(*pagerDutyConfig.ApiUrlOverride))
	}

	// Create client
	client := pagerduty.NewClient(token, opts...)

	// save clientOptions in cache
	d.ConnectionManager.Cache.Set(sessionCacheKey, client)

	return client, nil
}
