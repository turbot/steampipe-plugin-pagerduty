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

### Get a team by name

```sql
select
  name,
  id,
  description,
  self
from
  pagerduty_team
where
  name = 'developer';
```
