---
title: "Steampipe Table: pagerduty_ruleset_rule - Query PagerDuty Ruleset Rules using SQL"
description: "Allows users to query Ruleset Rules in PagerDuty, specifically the rule conditions, actions, and associated ruleset data, providing insights into incident management rules and patterns."
---

# Table: pagerduty_ruleset_rule - Query PagerDuty Ruleset Rules using SQL

PagerDuty Ruleset Rules are a part of the incident management service within PagerDuty that allows you to define and manage rules for incident notifications. It provides a way to set up and manage rules for various incident scenarios, including specific conditions and actions. PagerDuty Ruleset Rules help you stay informed about the incident management rules and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `pagerduty_ruleset_rule` table provides insights into Ruleset Rules within PagerDuty's incident management service. As a DevOps engineer, explore rule-specific details through this table, including conditions, actions, and associated ruleset data. Utilize it to uncover information about rules, such as those with specific conditions and actions, the relationships between rules, and the verification of rule actions.

## Examples

### Basic info
Explore the status of different rules within your PagerDuty ruleset to understand which are currently in use and which are disabled. This can aid in streamlining your incident management process by ensuring only necessary rules are active.

```sql
select
  id,
  ruleset_id,
  disabled,
  self
from
  pagerduty_ruleset_rule;
```

### List disabled rules
Explore which rules within your PagerDuty ruleset are currently disabled. This can be beneficial in understanding and managing the active rules and alerts in your system.

```sql
select
  id,
  ruleset_id,
  disabled,
  self
from
  pagerduty_ruleset_rule
where
  disabled;
```

### Count event rules per ruleset
Assess the number of event rules within each ruleset to better understand their complexity and manageability. This is useful for optimizing ruleset configurations and enhancing system efficiency.

```sql
select
  rs.id as ruleset_id,
  count(r.id)
from
  pagerduty_ruleset as rs
  left join pagerduty_ruleset_rule as r on rs.id = r.ruleset_id
group by rs.id;
```

### List rules without any priority
Explore which PagerDuty rules lack a set priority, helping to identify potential gaps in your incident management process. This is beneficial in ensuring all rules are appropriately prioritized for effective incident response.

```sql
select
  id,
  ruleset_id,
  disabled,
  self
from
  pagerduty_ruleset_rule
where
  actions -> 'priority' is null;
```