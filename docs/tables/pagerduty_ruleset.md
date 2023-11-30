---
title: "Steampipe Table: pagerduty_ruleset - Query PagerDuty Rulesets using SQL"
description: "Allows users to query PagerDuty Rulesets, providing insights into the rules defined for incident management and notifications."
---

# Table: pagerduty_ruleset - Query PagerDuty Rulesets using SQL

A PagerDuty Ruleset is a collection of rules that determine what actions to take when an event is triggered. Rulesets help to control the routing, suppression, and grouping of incidents. They are a crucial component in managing the incident response process in PagerDuty.

## Table Usage Guide

The `pagerduty_ruleset` table provides insights into the rulesets within PagerDuty's incident management system. As an incident response manager or system administrator, explore ruleset-specific details through this table, including the rules defined, their conditions, and actions. Utilize it to manage and optimize the incident response process, such as routing incidents to appropriate responders, suppressing non-critical incidents, and grouping related incidents together.

## Examples

### Basic info
Explore the different rulesets in your PagerDuty account to understand their types and identifiers. This can help in managing and organizing your incident response workflows effectively.

```sql
select
  name,
  id,
  type
from
  pagerduty_ruleset;
```

### List default global rulesets
Explore which rulesets are set as the default global rules within your PagerDuty configuration. This is useful for understanding the overarching rules that apply to all incidents, allowing for better incident management and response.

```sql
select
  name,
  id,
  type
from
  pagerduty_ruleset
where
  type = 'default_global';
```

### List rulesets not owned by any team
Explore which rulesets in your PagerDuty configuration aren't associated with any team. This can help identify potential gaps in your incident management workflows.

```sql
select
  name,
  id,
  type
from
  pagerduty_ruleset
where
  team is null;
```

### Count event rules per ruleset
Analyze the settings to understand the distribution of event rules across different rulesets. This can help you optimize event management by identifying rulesets with an excessive or insufficient number of rules.

```sql
select
  rs.id as ruleset_id,
  count(r.id)
from
  pagerduty_ruleset as rs
  left join pagerduty_ruleset_rule as r on rs.id = r.ruleset_id
group by rs.id;
```