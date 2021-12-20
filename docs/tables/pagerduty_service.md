# Table: pagerduty_service

A service represents something you monitor (like a web service, email service, or database service). It is a container for related incidents that associates them with escalation policies.

## Examples

### Basic info

```sql
select
  name,
  id,
  status,
  self
from
  pagerduty_service;
```

### List disabled services

```sql
select
  name,
  id,
  status,
  self
from
  pagerduty_service
where
  status = 'disabled';
```

### List services not associated with any team

```sql
select
  name,
  id,
  status,
  self
from
  pagerduty_service
where
  jsonb_array_length(teams) < 1;
```
