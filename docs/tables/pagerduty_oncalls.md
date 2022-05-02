# Table: pagerduty_oncalls

An on-call represents a contiguous unit of time for which a User will be on call for a given Escalation Policy and Escalation Rules.

## Examples

### Basic info

```sql
select
  escalation_policy,
  "user",
  schedule,
  escalation_level,
  start,
  end
from
  pagerduty_oncalls;
```

### Get the current on call user's name for a given schedule name

```sql
select
  "user" ->> 'summary' as "User"
from
  pagerduty_oncalls
where
  schedule ->> 'summary' = 'Schedule Name'
```
