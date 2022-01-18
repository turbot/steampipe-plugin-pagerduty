# Table: pagerduty_service_integration

A service integration is an integration that belongs to a Pagerduty service.It allows you to represent the actual entities you are monitoring, managing, and operating as services in PagerDuty.

## Examples

### Basic info

```sql
select
  name,
  id,
  service_id,
  created_at,
  jsonb_pretty(vendor) as vendor
from
  pagerduty_service_integration;
```

### List all vendor specific integrations of a service

```sql
select
  name,
  id,
  service_id,
  created_at,
  vendor ->> 'summary' as vendor_name
from
  pagerduty_service_integration
where
  vendor is not null;
```
