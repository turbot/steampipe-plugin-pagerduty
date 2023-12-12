---
title: "Steampipe Table: pagerduty_vendor - Query PagerDuty Vendors using SQL"
description: "Allows users to query PagerDuty Vendors, specifically the vendor details including id, name, summary, and other relevant information."
---

# Table: pagerduty_vendor - Query PagerDuty Vendors using SQL

PagerDuty is an incident management platform that provides reliable notifications, automatic escalations, on-call scheduling, and other functionality to help teams detect and fix infrastructure problems quickly. The Vendor resource in PagerDuty represents the third-party services that are integrated with PagerDuty to create, update, and resolve incidents. It includes details such as vendor id, name, summary, and other relevant information.

## Table Usage Guide

The `pagerduty_vendor` table provides insights into third-party services integrated with PagerDuty. As a DevOps engineer, explore vendor-specific details through this table, including vendor id, name, summary, and other relevant information. Utilize it to uncover information about vendors, such as their integration status with PagerDuty, and the services they offer.

**Important Notes**
- It is recommended that queries specify `name` or `id` in order to limit results due to the large number of vendors.

## Examples

### Get AWS CloudWatch integration vendor
Explore the integration between your system and Amazon CloudWatch, a monitoring service for AWS resources and the applications you run on AWS. This query is useful for understanding the integration's details and potential alert configurations, which can aid in system management and troubleshooting.

```sql+postgres
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

```sql+sqlite
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

### Get count of services using AWS CloudTrail integration
This query helps you determine the number of services integrated with AWS CloudTrail in your PagerDuty account. It's useful for assessing the extent of your CloudTrail usage and ensuring all necessary services are properly connected for optimal incident management.

```sql+postgres
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

```sql+sqlite
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
    json_extract(vendor, '$.id') as vendor_id
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