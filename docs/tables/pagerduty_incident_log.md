# Table: pagerduty_incident_log

Retrieves the log entries for the specified incident.

The `pagerduty_incident_log` table can be used to query log entries for ANY incident, and **you must specify the incident ID** in the where or join clause (`where incident_id=`, `join pagerduty_incident_log on incident_id=`).

**Note:** It is recommended that queries specify `created_at` (usually in the `where` clause) to filter the log entries within a specific time range.

## Examples

### List log entries for all incident in last 24 hrs

```sql
select
  i.summary as incident_summary,
  l.id as log_entry_id,
  l.created_at,
  l.agent ->> 'summary' as created_by
from
  pagerduty_incident_log as l,
  pagerduty_incident as i
where
  l.incident_id = i.id
  and l.created_at > now() - interval '24 hrs';
```

### List incident logs for an incident from the last 3 days

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

### List incident log entries for activities performed by users

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
  and agent ->> 'type' = 'user_reference';
```

### List incident log entries for activities performed by services

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
  and agent ->> 'type' = 'service_reference';
```
