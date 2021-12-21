# Table: pagerduty_user

An user is a member of a PagerDuty account that have the ability to interact with incidents and other data on the account.

## Examples

### Basic info

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

### List users not belongs to any team

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
