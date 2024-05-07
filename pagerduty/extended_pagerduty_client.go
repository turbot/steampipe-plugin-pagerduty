package pagerduty

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/PagerDuty/go-pagerduty"
)

const (
	apiEndpoint = "https://api.pagerduty.com"
)

type ExtendedPagerDutyClient struct {
	*pagerduty.Client
	authToken string
}

func NewExtendedPagerDutyClient(client *pagerduty.Client, token string) *ExtendedPagerDutyClient {
	return &ExtendedPagerDutyClient{client, token}
}

func (c *ExtendedPagerDutyClient) GetIncidentCustomFields(ctx context.Context, incidentID string) (map[string]interface{}, error) {
	resp, err := c.get(ctx, fmt.Sprintf("/incidents/%s/custom_fields/values", incidentID))
	if err != nil {
		return nil, err
	}

	return c.decodeJSON(resp)
}

func (c *ExtendedPagerDutyClient) GetIncidentBusinessServicesImpacts(ctx context.Context, incidentID string) (map[string]interface{}, error) {
	resp, err := c.get(ctx, fmt.Sprintf("/incidents/%s/business_services/impacts", incidentID))
	if err != nil {
		return nil, err
	}

	return c.decodeJSON(resp)
}

func (c *ExtendedPagerDutyClient) GetServiceDependencies(ctx context.Context, serviceID string) (map[string]interface{}, error) {
	resp, err := c.get(ctx, fmt.Sprintf("/service_dependencies/technical_services/%s", serviceID))
	if err != nil {
		return nil, err
	}

	return c.decodeJSON(resp)
}

func (c *ExtendedPagerDutyClient) GetBusinessServiceDependencies(ctx context.Context, serviceID string) (map[string]interface{}, error) {
	resp, err := c.get(ctx, fmt.Sprintf("/service_dependencies/business_services/%s", serviceID))
	if err != nil {
		return nil, err
	}

	return c.decodeJSON(resp)
}

func (c *ExtendedPagerDutyClient) decodeJSON(resp *http.Response) (map[string]interface{}, error) {
	defer func() { _ = resp.Body.Close() }() // explicitly discard error
	customFields := make(map[string]interface{})

	defer func() { _ = resp.Body.Close() }() // explicitly discard error
	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&customFields); err != nil {
		return nil, err
	}

	return customFields, nil
}

func (c *ExtendedPagerDutyClient) get(ctx context.Context, path string) (*http.Response, error) {
	return c.do(ctx, apiEndpoint, http.MethodGet, path, nil)
}

func (c *ExtendedPagerDutyClient) do(ctx context.Context, endpoint, method, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, endpoint+path, body)
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.pagerduty+json;version=2")
	req.Header.Set("Authorization", "Token token="+c.authToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	return c.checkResponse(resp, err)
}

func (c *ExtendedPagerDutyClient) checkResponse(resp *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return resp, fmt.Errorf("error calling the API endpoint: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return resp, fmt.Errorf("%d error returned from endpoint", resp.StatusCode)
	}

	return resp, nil
}
