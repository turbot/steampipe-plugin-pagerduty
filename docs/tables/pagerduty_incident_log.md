---
title: "Steampipe Table: pagerduty_incident_log - Query PagerDuty Incident Logs using SQL"
description: "Allows users to query PagerDuty Incident Logs, specifically the logs for each incident, providing insights into incident history and resolution details."
---

# Table: pagerduty_incident_log - Query PagerDuty Incident Logs using SQL

PagerDuty Incident Logs are a record of all the activities related to an incident within PagerDuty. These logs provide a detailed timeline of incident activity, including status changes, escalations, and notes added by users. They are crucial for postmortem analysis and understanding the incident's lifecycle.

## Table Usage Guide

The `pagerduty_incident_log` table provides insights into incident logs within PagerDuty. As an incident responder or a DevOps engineer, explore the detailed timeline of incident activity through this table, including status changes, escalations, and user-added notes. Utilize it to facilitate postmortem analysis and gain a comprehensive understanding of the incident's lifecycle.

**Important Notes**
- You must specify the `incident_id` in the `where` or join clause (`where incident_id=`, `join pagerduty_incident_log l on l.incident_id=`) to query this table.
- It is recommended that queries specify `created_at` (usually in the `where` clause) to filter the log entries within a specific time range.

## Examples

### List log entries for all incident in last 24 hrs
Explore the recent incidents in the last 24 hours for a comprehensive understanding of the situation, including the incident's summary and who created it. This query can assist in identifying patterns and trends in incidents, which can be useful for troubleshooting and improving system stability.

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
Explore recent incident logs to gain insights into the activities and changes made within the last three days. This is beneficial in understanding the sequence of events or actions taken for a specific incident, aiding in incident management and resolution.

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
Discover the segments that highlight user activities within specific incidents. This can be beneficial in tracking user actions, identifying patterns, and managing incident responses more effectively.

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
Explore which activities have been performed by different services in specific incidents, allowing you to gain insights into the actions taken and who performed them for better incident management.

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