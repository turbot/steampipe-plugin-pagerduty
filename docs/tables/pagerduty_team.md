---
title: "Steampipe Table: pagerduty_team - Query PagerDuty Teams using SQL"
description: "Allows users to query PagerDuty Teams, providing insights into team details, including team members, escalation policies, and associated services."
---

# Table: pagerduty_team - Query PagerDuty Teams using SQL

PagerDuty Teams is a feature within PagerDuty that allows you to group users, escalation policies, and services based on your organization's structure or responsibilities. It provides a centralized way to manage and organize your incident response structure for better visibility and collaboration. PagerDuty Teams helps you streamline incident management and response by ensuring the right information reaches the right people at the right time.

## Table Usage Guide

The `pagerduty_team` table provides insights into teams within PagerDuty. As a DevOps engineer, you can explore team-specific details through this table, including members, escalation policies, and associated services. Use it to manage and organize your incident response structure, ensuring the right information reaches the right people at the right time.

## Examples

### Basic info
Explore the essential details of your PagerDuty team to gain insights into team structure and roles, which can be useful for auditing or restructuring purposes.

```sql+postgres
select
  name,
  id,
  description,
  self
from
  pagerduty_team;
```

```sql+sqlite
select
  name,
  id,
  description,
  self
from
  pagerduty_team;
```

### List teams with no members
Discover the teams that currently have no members assigned to them. This can be useful for identifying and managing unallocated resources or underutilized teams within your organization.

```sql+postgres
select
  name,
  id,
  description,
  self
from
  pagerduty_team
where
  members is null;
```

```sql+sqlite
select
  name,
  id,
  description,
  self
from
  pagerduty_team
where
  members is null;
```

### List members with pending invitation
Explore which team members have yet to accept their invitations. This is useful in monitoring the status of team onboarding and identifying any potential issues or delays.

```sql+postgres
select
  t.name as team_name,
  member -> 'user' ->> 'summary' as user_name,
  member ->> 'role' as role,
  u.invitation_sent
from
  pagerduty_team as t,
  jsonb_array_elements(members) as member,
  pagerduty_user as u
where
  member -> 'user' ->> 'id' = u.id
  and u.invitation_sent;
```

```sql+sqlite
select
  t.name as team_name,
  json_extract(member.value, '$.user.summary') as user_name,
  json_extract(member.value, '$.role') as role,
  u.invitation_sent
from
  pagerduty_team as t,
  json_each(members) as member,
  pagerduty_user as u
where
  json_extract(member.value, '$.user.id') = u.id
  and u.invitation_sent;
```