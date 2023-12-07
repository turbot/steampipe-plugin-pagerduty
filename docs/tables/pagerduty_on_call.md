---
title: "Steampipe Table: pagerduty_on_call - Query PagerDuty On-calls using SQL"
description: "Allows users to query On-calls in PagerDuty, specifically the on-call schedules, providing insights into who is currently on-call and when."
---

# Table: pagerduty_on_call - Query PagerDuty On-calls using SQL

PagerDuty On-call is a feature within PagerDuty that allows you to manage and view who is currently on-call and when. It provides a centralized way to set up and manage on-call schedules for various teams and individuals. PagerDuty On-call helps you stay informed about the on-call status and take appropriate actions when needed.

## Table Usage Guide

The `pagerduty_on_call` table provides insights into on-call schedules within PagerDuty. As a DevOps engineer, explore on-call details through this table, including who is currently on-call, when they started, and when they will end. Utilize it to uncover information about on-call schedules, such as overlapping schedules, and the verification of on-call rotations.

## Examples

### Basic info
Explore which users are currently on call and the associated escalation policies and schedules. This can help in understanding the current on-call management setup and in planning future on-call schedules.

```sql+postgres
select
  escalation_policy,
  user_on_call,
  schedule,
  escalation_level,
  start,
  "end"
from
  pagerduty_on_call;
```

```sql+sqlite
select
  escalation_policy,
  user_on_call,
  schedule,
  escalation_level,
  start,
  "end"
from
  pagerduty_on_call;
```

### Get the current on call user's name for a given schedule name
Determine the current on-call individual for a specific schedule. This is useful for identifying who is responsible for handling urgent issues during a particular time frame.

```sql+postgres
select
  user_on_call ->> 'summary' as "User"
from
  pagerduty_on_call
where
  schedule ->> 'summary' = 'Schedule Name';
```

```sql+sqlite
select
  json_extract(user_on_call, '$.summary') as "User"
from
  pagerduty_on_call
where
  json_extract(schedule, '$.summary') = 'Schedule Name';
```