---
title: "Steampipe Table: pagerduty_escalation_policy - Query PagerDuty Escalation Policies using SQL"
description: "Allows users to query PagerDuty Escalation Policies, specifically the details about each policy, including its assigned teams, services, and escalation rules."
---

# Table: pagerduty_escalation_policy - Query PagerDuty Escalation Policies using SQL

PagerDuty Escalation Policies are a crucial component of the PagerDuty incident management platform. They define the sequence in which alerts are sent to different individuals or teams until the incident is acknowledged or resolved. These policies play a vital role in ensuring timely response to incidents and maintaining service reliability.

## Table Usage Guide

The `pagerduty_escalation_policy` table provides insights into Escalation Policies within PagerDuty's incident management platform. As an incident manager or DevOps engineer, explore policy-specific details through this table, including the associated teams, services, and escalation rules. Utilize it to gain a comprehensive understanding of your incident response workflow, and to ensure an effective and timely response to incidents.

## Examples

### Basic info
Explore the fundamental details of PagerDuty's escalation policies to understand their structure and access points. This can be useful in quickly assessing the policies and identifying the specific ones for review or modification.

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
Explore the default escalation policy within your system. This allows you to understand and manage the standard procedures in place for escalating issues.

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
Discover the escalation policies that are not currently linked to any services or teams, which could help optimize resource allocation and streamline incident management processes.

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
Explore which escalation policies are set to not repeat if incidents are not acknowledged. This can be useful to identify potential gaps in incident management where critical alerts may be missed.

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