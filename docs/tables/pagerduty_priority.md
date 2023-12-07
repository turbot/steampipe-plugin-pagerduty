---
title: "Steampipe Table: pagerduty_priority - Query PagerDuty Priorities using SQL"
description: "Allows users to query PagerDuty Priorities, providing insights into the different levels of urgency that can be assigned to incidents."
---

# Table: pagerduty_priority - Query PagerDuty Priorities using SQL

PagerDuty is an incident management platform that provides reliable incident notifications via email, push, SMS, and phone, as well as automatic escalations, on-call scheduling, and other functionality to help teams detect and fix infrastructure problems quickly. The PagerDuty Priorities feature allows users to define the relative urgency of incidents, which can help streamline response and resolution processes.

## Table Usage Guide

The `pagerduty_priority` table provides insights into the different levels of urgency that can be assigned to incidents in PagerDuty. As an Incident Manager or DevOps engineer, explore priority-specific details through this table, including the name and description of each priority level, and the time frame in which incidents at each priority level should be resolved. Utilize it to better understand and manage your team's response to incidents of varying urgency.

**Important Notes**
- To list the priorities, first [enable and configure your priorities](https://support.pagerduty.com/docs/incident-priority#section-enabling-incident-priority).

## Examples

### Basic info
Explore the priorities in your PagerDuty setup to gain insights into their names, IDs, and descriptions, helping you better understand and manage your incident response hierarchy.

```sql+postgres
select
  name,
  id,
  description,
  self
from
  pagerduty_priority;
```

```sql+sqlite
select
  name,
  id,
  description,
  self
from
  pagerduty_priority;
```

### List event rules with highest priority (P1)
Discover the segments that have the highest priority (P1) in the event rules. This can be useful for identifying and prioritizing the most critical rules for incident management.

```sql+postgres
with priority as (
  select
    id
  from
    pagerduty_priority
  where
    name = 'P1'
)
select
  rs.id as rule_id,
  rs.ruleset_id,
  rs.disabled
from
  pagerduty_ruleset_rule as rs,
  priority as p
where
  p.id = rs.actions -> 'priority' ->> 'value';
```

```sql+sqlite
select
  rs.id as rule_id,
  rs.ruleset_id,
  rs.disabled
from
  pagerduty_ruleset_rule as rs,
  (select id from pagerduty_priority where name = 'P1') as p
where
  p.id = json_extract(rs.actions, '$.priority.value');
```