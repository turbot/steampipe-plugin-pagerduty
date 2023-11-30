---
title: "Steampipe Table: pagerduty_user - Query PagerDuty Users using SQL"
description: "Allows users to query PagerDuty Users, providing detailed information on each user's profile such as email, role, job title, time zone, and more."
---

# Table: pagerduty_user - Query PagerDuty Users using SQL

PagerDuty is a real-time operations platform that integrates machine data & human intelligence to improve visibility & agility across organizations. It is used for incident response, on-call scheduling, escalation policies, and analytics. Users in PagerDuty are individuals who have access to the platform and their profile information includes details like email, role, job title, time zone, etc.

## Table Usage Guide

The `pagerduty_user` table provides insights into user profiles within PagerDuty. As an IT manager or DevOps engineer, explore user-specific details through this table, including their roles, job titles, and time zones. Utilize it to manage user access, understand user responsibilities, and ensure appropriate on-call schedules are in place.

## Examples

### Basic info
Explore the user base in your PagerDuty account to understand their roles and contact information. This can help in managing team responsibilities and communication channels effectively.

```sql
select
  name,
  id,
  email,
  role
from
  pagerduty_user;
```

### List invited users
Discover the details of users who have been invited to join your PagerDuty team. This can help you track pending invitations and understand the roles assigned to each invitee.

```sql
select
  name,
  id,
  email,
  role
from
  pagerduty_user
where
  invitation_sent;
```

### List users not in any team
Discover the segments that consist of users who are not part of any team, a useful approach for identifying potential areas for team expansion or redistribution of tasks.

```sql
select
  name,
  id,
  email,
  role
from
  pagerduty_user
where
  jsonb_array_length(teams) < 1;
```

### List users with `owner` tags
Discover the segments that consist of users tagged as 'owners' in your system. This allows you to quickly identify and communicate with the responsible parties for specific tasks or issues.

```sql
select
  name,
  id,
  email,
  role
from
  pagerduty_user,
  jsonb_array_elements(tags) as t
where
  t ->> 'label' ilike 'owner';
```