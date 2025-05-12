---
title: "Steampipe Table: pagerduty_schedule_user - Query PagerDuty Schedule Users using SQL"
description: "Allows users to query Schedule Users in PagerDuty, specifically providing details about users assigned to each schedule, including their roles, contact information, and schedule assignments."
---

# Table: pagerduty_schedule_user - Query PagerDuty Schedule Users using SQL

PagerDuty Schedule Users represent the individuals who are assigned to on-call rotations within a schedule. These assignments determine who will be notified when incidents occur during specific time periods. Understanding schedule user assignments is crucial for maintaining effective incident response coverage.

## Table Usage Guide

The `pagerduty_schedule_user` table provides insights into user assignments within PagerDuty schedules. As an incident response manager or DevOps engineer, explore user-specific details through this table, including their roles, contact information, and schedule assignments. Utilize it to manage on-call rotations, ensure proper coverage, and verify user availability across different schedules.

## Examples

### Basic info
Explore which users are assigned to different schedules in PagerDuty to understand the distribution of on-call responsibilities.

```sql+postgres
select
  schedule_name,
  name,
  email,
  role,
  timezone
from
  pagerduty_schedule_user;
```

```sql+sqlite
select
  schedule_name,
  name,
  email,
  role,
  timezone
from
  pagerduty_schedule_user;
```

### List users for a specific schedule
Identify the users assigned to a particular schedule to understand the on-call rotation for that specific schedule.

```sql+postgres
select
  name,
  email,
  role,
  job_title,
  timezone
from
  pagerduty_schedule_user
where
  schedule_id = 'P123ABC';
```

```sql+sqlite
select
  name,
  email,
  role,
  job_title,
  timezone
from
  pagerduty_schedule_user
where
  schedule_id = 'P123ABC';
```

### Find users assigned to multiple schedules
Discover users who are part of multiple schedule rotations to help identify potential overload situations.

```sql+postgres
select
  name,
  email,
  count(distinct schedule_id) as schedule_count,
  array_agg(schedule_name) as schedule_names
from
  pagerduty_schedule_user
group by
  name,
  email
having
  count(distinct schedule_id) > 1
order by
  schedule_count desc;
```

```sql+sqlite
select
  name,
  email,
  count(distinct schedule_id) as schedule_count,
  group_concat(distinct schedule_name) as schedule_names
from
  pagerduty_schedule_user
group by
  name,
  email
having
  count(distinct schedule_id) > 1
order by
  schedule_count desc;
```

### List users by timezone
Analyze the distribution of on-call users across different timezones to ensure global coverage.

```sql+postgres
select
  timezone,
  count(*) as user_count,
  array_agg(distinct name) as users
from
  pagerduty_schedule_user
group by
  timezone
order by
  user_count desc;
```

```sql+sqlite
select
  timezone,
  count(*) as user_count,
  group_concat(distinct name) as users
from
  pagerduty_schedule_user
group by
  timezone
order by
  user_count desc;
```

### Find schedules and their associated escalation policies
Identify which escalation policies are using each schedule and who the users are in those schedules.

```sql+postgres
select
  su.schedule_name,
  su.name as user_name,
  ep.name as escalation_policy_name,
  ep.description as escalation_policy_description
from
  pagerduty_schedule_user su
  left join pagerduty_schedule s on s.id = su.schedule_id
  left join jsonb_array_elements(s.escalation_policies) as ep_ref on true
  left join pagerduty_escalation_policy ep on ep.id = ep_ref->>'id'
order by
  su.schedule_name, su.name;
```

```sql+sqlite
select
  su.schedule_name,
  su.name as user_name,
  ep.name as escalation_policy_name,
  ep.description as escalation_policy_description
from
  pagerduty_schedule_user su
  left join pagerduty_schedule s on s.id = su.schedule_id
  left join json_each(s.escalation_policies) as ep_ref
  left join pagerduty_escalation_policy ep on ep.id = json_extract(ep_ref.value, '$.id')
order by
  su.schedule_name, su.name;
```

### List users with their contact methods and notification rules
Get detailed information about how schedule users can be contacted and their notification preferences.

```sql+postgres
select
  su.name,
  su.email,
  su.schedule_name,
  cm->>'type' as contact_method_type,
  cm->>'address' as contact_method_address,
  nr->>'summary' as notification_rule
from
  pagerduty_schedule_user su
  left join pagerduty_user u on u.id = su.id
  left join jsonb_array_elements(u.contact_methods) as cm on true
  left join jsonb_array_elements(u.notification_rules) as nr on true
order by
  su.name, su.schedule_name;
```

```sql+sqlite
select
  su.name,
  su.email,
  su.schedule_name,
  json_extract(cm.value, '$.type') as contact_method_type,
  json_extract(cm.value, '$.address') as contact_method_address,
  json_extract(nr.value, '$.summary') as notification_rule
from
  pagerduty_schedule_user su
  left join pagerduty_user u on u.id = su.id
  left join json_each(u.contact_methods) as cm
  left join json_each(u.notification_rules) as nr
order by
  su.name, su.schedule_name;
```

### Find schedules with their services and incidents
Analyze which services are associated with schedules through escalation policies and list any recent incidents.

```sql+postgres
select
  su.schedule_name,
  su.name as user_name,
  s.name as service_name,
  i.summary as incident_summary,
  i.status as incident_status,
  i.created_at as incident_created
from
  pagerduty_schedule_user su
  left join pagerduty_schedule sch on sch.id = su.schedule_id
  left join jsonb_array_elements(sch.escalation_policies) as ep_ref on true
  left join pagerduty_escalation_policy ep on ep.id = ep_ref->>'id'
  left join pagerduty_service s on s.escalation_policy->>'id' = ep.id
  left join pagerduty_incident i on i.service->>'id' = s.id
where
  i.created_at >= now() - interval '7 days'
  or i.created_at is null
order by
  su.schedule_name, i.created_at desc;
```

```sql+sqlite
select
  su.schedule_name,
  su.name as user_name,
  s.name as service_name,
  i.summary as incident_summary,
  i.status as incident_status,
  i.created_at as incident_created
from
  pagerduty_schedule_user su
  left join pagerduty_schedule sch on sch.id = su.schedule_id
  left join json_each(sch.escalation_policies) as ep_ref
  left join pagerduty_escalation_policy ep on ep.id = json_extract(ep_ref.value, '$.id')
  left join pagerduty_service s on json_extract(s.escalation_policy, '$.id') = ep.id
  left join pagerduty_incident i on json_extract(i.service, '$.id') = s.id
where
  datetime(i.created_at) >= datetime('now', '-7 days')
  or i.created_at is null
order by
  su.schedule_name, i.created_at desc;
```
