# Table: pagerduty_incident

An incident represents a problem or an issue that needs to be addressed and resolved.

## Examples

### List unacknowledged incidents

```sql
select
  incident_number,
  summary,
  urgency,
  created_at,
  assignments
from
  pagerduty_incident
where
  status = 'triggered';
```

### List unacknowledged incidents with high urgency

```sql
select
  incident_number,
  summary,
  urgency,
  created_at,
  assignments
from
  pagerduty_incident
where
  status = 'triggered'
  and urgency = 'high';
```

### List unacknowledged incidents assigned to a specific user

```sql
select
  i.incident_number,
  i.summary,
  i.urgency,
  i.created_at,
  p.email as assigned_to
from
  pagerduty_incident as i,
  jsonb_array_elements(i.assignments) as a
  join pagerduty_user as p on p.id = a -> 'assignee' ->> 'id'
where
  p.id = 'P5ISTE8'
  and status = 'triggered';
```

### List all unacknowledged incidents in last 7 days

```sql
select
  incident_number,
  summary,
  urgency,
  created_at,
  assignments
from
  pagerduty_incident
where
  status = 'triggered'
  and created_at >= now() - interval '7 days';
```
