---
title: "Steampipe Table: pagerduty_incident - Query PagerDuty Incidents using SQL"
description: "Allows users to query PagerDuty Incidents, providing comprehensive details about each incident such as status, urgency, and assigned services."
---

# Table: pagerduty_incident - Query PagerDuty Incidents using SQL

PagerDuty Incident Management is a digital operations management platform that combines machine data with human data to improve visibility and agility across organizations. It helps teams to minimize business disruptions and improve the customer experience by providing real-time alerts and incident tracking. PagerDuty Incident Management allows organizations to manage incidents from any source, and it's trusted by thousands of organizations globally to improve their incident response.

## Table Usage Guide

The `pagerduty_incident` table provides detailed insights into incidents managed through the PagerDuty platform. As an Operations or DevOps engineer, explore incident-specific details through this table, including current status, associated services, and urgency level. Utilize it to track and manage incidents, understand their impact, and plan for timely resolution.

## Examples

### List unacknowledged incidents for the last 30 days
Explore the recent incidents that have not been addressed in the past month. This is beneficial for prioritizing urgent tasks and understanding the backlog of unresolved issues.

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
Determine the high urgency incidents from the past week that are still pending acknowledgment. This aids in prioritizing immediate action on urgent matters that have been overlooked.

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
Determine the areas in which urgent issues have not been acknowledged by a specific team member in the last three days. This helps to identify potential bottlenecks and ensures that critical incidents are addressed promptly.

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
Gain insights into all the recent incidents that have not been addressed yet, within the past week. This can help prioritize urgent matters and streamline incident response.

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