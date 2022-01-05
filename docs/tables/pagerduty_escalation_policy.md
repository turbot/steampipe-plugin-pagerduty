# Table: pagerduty_escalation_policy

An escalation policy determines what user or schedule will be notified first, second, and so on when an incident is triggered. Escalation policies are used by one or more services.

## Examples

### Basic info

```sql
select
  name,
  id,
  self,
  html_url
from
  pagerduty_escalation_policy;
```

### List default escalation policy

```sql
select
  name,
  id,
  self,
  html_url
from
  pagerduty_escalation_policy
where
  name = 'Default';
```

### List unused escalation policies

```sql
select
  name,
  id,
  self,
  html_url
from
  pagerduty_escalation_policy
where
  jsonb_array_length(services) < 1
  and jsonb_array_length(teams) < 1;
```

### List policies that do not repeat if incidents are not acknowledged

```sql
select
  name,
  id,
  self,
  html_url
from
  pagerduty_escalation_policy
where
  num_loops = 0;
```
