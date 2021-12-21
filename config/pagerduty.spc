connection "pagerduty" {
  plugin = "pagerduty"
  
  # Account/user API token authentication
  # This can also be set via the `PAGERDUTY_TOKEN` environment variable.
  # token = "YOUR_AUTH_TOKEN"

  # A custom proxy endpoint as PagerDuty client API URL overriding 'service_region' setup.
  # Default is "https://api.pagerduty.com"
  # api_url_override = "YOUR_CUSTOM_ENDPOINT"
}
