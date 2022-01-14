# Table: pagerduty_incident

An incident represents a problem or an issue that needs to be addressed and resolved.

**Note:** If no `created_at` key qual is specified, incidents from the last 30 days will be returned by default.

## Examples

### List unacknowledged incidents for the last 30 days

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
  and created_at >= now() - interval '30 days';
```

### List unacknowledged incidents with high urgency for the last 1 week

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
  and urgency = 'high'
  and created_at >= now() - interval '7 days';
```

### List unacknowledged incidents assigned to a specific user for the last 3 days

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
  and status = 'triggered'
  and created_at >= now() - interval '3 days';
```

### List all unacknowledged incidents for the last 7 days

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
