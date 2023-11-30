---
title: "Steampipe Table: pagerduty_tag - Query PagerDuty Tags using SQL"
description: "Allows users to query Tags in PagerDuty, specifically the detailed information about each tag, providing insights into tag usage and management."
---

# Table: pagerduty_tag - Query PagerDuty Tags using SQL

PagerDuty is an incident management platform that provides reliable notifications, automatic escalations, on-call scheduling, and other functionality to help teams detect and fix infrastructure problems quickly. Its Tags feature allows users to categorize and filter incidents, services, and teams using custom labels. This aids in organizing resources, prioritizing incidents, and streamlining workflows.

## Table Usage Guide

The `pagerduty_tag` table provides insights into Tags within PagerDuty's incident management platform. As an operations engineer, explore tag-specific details through this table, including associated services, teams, and incidents. Utilize it to uncover information about tags, such as those most frequently used, their associated resources, and the effectiveness of your tagging strategy.

## Examples

### Basic info
Explore which PagerDuty tags are being used. This can help in managing and organizing your PagerDuty services and incidents more effectively.

```sql
select
  id,
  label,
  self
from
  pagerduty_tag;
```

### List unused tags
Identify the tags that are currently not associated with any users, teams, or escalation policies in PagerDuty. This can help streamline your tagging system by removing or reassigning unused tags.

```sql
with associated_tags as (
  select
    t ->> 'id' as id
  from
    pagerduty_user,
    jsonb_array_elements(tags) as t
  union
  select
    t ->> 'id' as id
  from
    pagerduty_team,
    jsonb_array_elements(tags) as t
  union
  select
    t ->> 'id' as id
  from
    pagerduty_escalation_policy,
    jsonb_array_elements(tags) as t
),
distinct_tags as (
  select
    distinct id
  from
    associated_tags
)
select
  t.id,
  t.label,
  t.self
from
  pagerduty_tag as t
  left join distinct_tags as dt on t.id = dt.id
where
  dt.id is null;
```