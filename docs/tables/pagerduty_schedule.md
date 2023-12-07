---
title: "Steampipe Table: pagerduty_schedule - Query PagerDuty Schedules using SQL"
description: "Allows users to query PagerDuty Schedules, specifically the schedule layers, rotations and users assigned to each schedule, providing insights into the on-call management and incident response workflow."
---

# Table: pagerduty_schedule - Query PagerDuty Schedules using SQL

PagerDuty is a digital operations management platform that integrates with ITOps and DevOps tools to improve operational reliability and agility. It provides an incident response and on-call management platform to businesses, enabling them to monitor their applications and infrastructure, and alert the right people at the right time. PagerDuty Schedules are used to determine who is on-call when an incident occurs.

## Table Usage Guide

The `pagerduty_schedule` table provides insights into the on-call schedules within PagerDuty. As a DevOps engineer, explore schedule-specific details through this table, including the rotations, layers, and users assigned to each schedule. Utilize it to manage and optimize your on-call schedules, ensuring that incidents are handled promptly and efficiently.

## Examples

### Basic info
Explore which PagerDuty schedules are active to understand the current allocation of resources and manage workflow more effectively. This can help in optimizing resource utilization and ensuring round-the-clock coverage.

```sql+postgres
select
  name,
  id,
  timezone,
  self
from
  pagerduty_schedule;
```

```sql+sqlite
select
  name,
  id,
  timezone,
  self
from
  pagerduty_schedule;
```

### List unused schedules
Discover the schedules that are not linked to any escalation policies. This can help in identifying and cleaning up unused resources, thereby improving the efficiency of your PagerDuty setup.

```sql+postgres
select
  name,
  id,
  timezone,
  self
from
  pagerduty_schedule
where
  jsonb_array_length(escalation_policies) = 0;
```

```sql+sqlite
select
  name,
  id,
  timezone,
  self
from
  pagerduty_schedule
where
  json_array_length(escalation_policies) = 0;
```

### List schedules not assigned to any team
Discover the schedules that are not assigned to any team, helping to identify potential scheduling gaps or unallocated resources within your PagerDuty configuration.

```sql+postgres
select
  name,
  id,
  timezone,
  self
from
  pagerduty_schedule
where
  jsonb_array_length(teams) = 0;
```

```sql+sqlite
select
  name,
  id,
  timezone,
  self
from
  pagerduty_schedule
where
  json_array_length(teams) = 0;
```