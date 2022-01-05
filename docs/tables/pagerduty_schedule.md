# Table: pagerduty_schedule

A schedule determines the time periods that users are on call. Only on-call users are eligible to receive notifications from incidents.

## Examples

### Basic info

```sql
select
  name,
  id,
  timezone,
  self
from
  pagerduty_schedule;
```

### List unused schedules

```sql
select
  name,
  id,
  timezone,
  self
from
  pagerduty_schedule
where
  jsonb_array_length(escalation_policies) = 0;
```

### List schedules not assigned to any team

```sql
select
  name,
  id,
  timezone,
  self
from
  pagerduty_schedule
where
  jsonb_array_length(teams) = 0;
```
