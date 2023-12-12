---
title: "Steampipe Table: pagerduty_service_integration - Query PagerDuty Service Integrations using SQL"
description: "Allows users to query PagerDuty Service Integrations, providing insights into the connected systems that trigger incidents within the PagerDuty platform."
---

# Table: pagerduty_service_integration - Query PagerDuty Service Integrations using SQL

PagerDuty Service Integrations are connections between PagerDuty and various other systems, tools, and services. These integrations allow incidents to be triggered, acknowledged, and resolved within PagerDuty based on events occurring in the connected systems. They form a crucial part of incident management workflows, helping teams to respond quickly and effectively to operational issues.

## Table Usage Guide

The `pagerduty_service_integration` table provides insights into the integrations between PagerDuty and other systems. As an incident manager, you can use this table to understand the sources of incidents within your PagerDuty environment, including the types of integrations, their configurations, and their associated services. This information can help you to optimize your incident response processes and ensure that your team has the necessary context to resolve incidents efficiently.

## Examples

### Basic info
Explore which PagerDuty service integrations have been created by analyzing their names, IDs, and creation dates. This can provide insights into your system's integration history and help you understand the vendor details associated with each service.

```sql+postgres
select
  name,
  id,
  service_id,
  created_at,
  jsonb_pretty(vendor) as vendor
from
  pagerduty_service_integration;
```

```sql+sqlite
select
  name,
  id,
  service_id,
  created_at,
  vendor
from
  pagerduty_service_integration;
```

### List all vendor specific integrations of a service
Discover the segments that have specific integrations with vendors for a service. This can be useful to identify and manage third-party dependencies and assess potential risks associated with vendor-specific integrations.

```sql+postgres
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

```sql+sqlite
select
  name,
  id,
  service_id,
  created_at,
  json_extract(vendor, '$.summary') as vendor_name
from
  pagerduty_service_integration
where
  vendor is not null;
```