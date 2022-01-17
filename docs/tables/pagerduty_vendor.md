# Table: pagerduty_vendor

A PagerDuty vendor represents a specific type of integration, e.g., AWS Cloudwatch, Splunk, Datadog.

**Note:** It is recommended that queries specify `name` or `id` in order to limit results and produce faster queries.

## Examples

### List vendor for AWS CloudWatch integration

```sql
select
  name,
  id,
  description,
  website_url,
  alert_creation_default
from
  pagerduty_vendor
where
  name = 'Amazon CloudWatch';
```

### Get count of services using AWS CloudTrail integrations

```sql
with cloudtrail_vendor as (
  select
    id as vendor_id,
    name as vendor_name
  from
    pagerduty_vendor
  where
    name = 'AWS CloudTrail'
),
service_integrations as (
  select
    service_id,
    name as integration_name,
    vendor ->> 'id' as vendor_id
  from
    pagerduty_service_integration
  where
    service_id in (select id from pagerduty_service)
)
select
  cv.vendor_name,
  count(si.service_id) as service_count
from
  cloudtrail_vendor as cv
  left join service_integrations as si on cv.vendor_id = si.vendor_id
group by
  cv.vendor_name, 
  si.service_id;
```
