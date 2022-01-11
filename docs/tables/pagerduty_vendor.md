# Table: pagerduty_vendor

A PagerDuty vendor represents a specific type of integration, i.e. AWS Cloudwatch, Splunk, Datadog etc.

**Note:**It is recommended that queries to this table should include (usually in the `where` clause) at least one of these columns: `name` or `id`.

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

### Get count of service using AWS CloudTrail integrations

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

### List all available vendors

WARNING - This is a large query and may take minutes to run. It is not recommended and may timeout. It's included here as a reference for those who need to extract all data.

```sql
select
  name,
  id,
  description,
  website_url,
  alert_creation_default
from
  pagerduty_vendor;
```
