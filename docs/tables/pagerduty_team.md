# Table: pagerduty_team

A team is a collection of Users and Escalation Policies that represent a group of people within an organization.

## Examples

### Basic info

```sql
select
  name,
  id,
  description,
  self
from
  pagerduty_team;
```

### List teams with no members

```sql
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

```sql
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
