---
title: "Steampipe Table: pagerduty_service - Query PagerDuty Services using SQL"
description: "Allows users to query PagerDuty Services, providing detailed information about each service and its current status."
---

# Table: pagerduty_service - Query PagerDuty Services using SQL

PagerDuty is an incident response platform that provides reliable notifications, automatic escalations, on-call scheduling, and other functionality to help teams detect and fix infrastructure problems quickly. The PagerDuty Service is a component of the PagerDuty system that represents something you monitor (like a server, an application, or a database). Each service has its own set of integrations, escalation policies, and notification settings.

## Table Usage Guide

The `pagerduty_service` table provides insights into services within PagerDuty. As a DevOps engineer or system administrator, explore service-specific details through this table, including integrations, escalation policies, and notification settings. Utilize it to uncover information about services, such as their current status, associated teams, and the incidents related to each service.

## Examples

### Basic info
Explore which services are currently active on your PagerDuty account. This is useful for understanding your overall service usage and identifying any services that may be inactive or unused.

```sql+postgres
select
  name,
  id,
  status,
  self
from
  pagerduty_service;
```

```sql+sqlite
select
  name,
  id,
  status,
  self
from
  pagerduty_service;
```

### List disabled services
Uncover the details of inactive services on PagerDuty. This query is useful for maintaining system health by identifying services that are no longer in use.

```sql+postgres
select
  name,
  id,
  status,
  self
from
  pagerduty_service
where
  status = 'disabled';
```

```sql+sqlite
select
  name,
  id,
  status,
  self
from
  pagerduty_service
where
  status = 'disabled';
```

### List services not associated with any team
Identify services in your PagerDuty account that are not linked to any team. This is useful for ensuring all services are properly assigned for effective incident management.

```sql+postgres
select
  name,
  id,
  status,
  self
from
  pagerduty_service
where
  jsonb_array_length(teams) < 1;
```

```sql+sqlite
select
  name,
  id,
  status,
  self
from
  pagerduty_service
where
  json_array_length(teams) < 1;
```