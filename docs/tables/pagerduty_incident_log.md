# Table: pagerduty_incident_log

Retrieves the record of all log entries for the specified incident.

## Examples

### List all activities on specified incident in last 3 days

```sql
select
  id,
  incident_id,
  created_at,
  agent ->> 'summary' as created_by
from
  pagerduty_incident_log
where
  incident_id = 'Q0FH5K82AJ101C'
  and created_at > now() - interval '3 days';
```

### List all activities performed by users

```sql
select
  id,
  incident_id,
  created_at,
  jsonb_pretty(channel) as action,
  agent ->> 'summary' as user_name
from
  pagerduty_incident_log
where
  incident_id = 'Q0FH5K82AJ101C'
  and created_at > now() - interval '3 days'
  and agent ->> 'type' = 'user_reference';
```

### List all activities performed by service

```sql
select
  id,
  incident_id,
  created_at,
  jsonb_pretty(channel) as action,
  agent ->> 'summary' as user_name
from
  pagerduty_incident_log
where
  incident_id = 'Q0FH5K82AJ101C'
  and created_at > now() - interval '3 days'
  and agent ->> 'type' = 'service_reference';
```
