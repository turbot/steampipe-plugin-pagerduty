# Table: pagerduty_priority

 A priority is a label representing the importance and impact of an incident. This feature is only available on `Standard` and `Enterprise` plans.

 Incident priority levels help you classify the most important incidents from the least important ones. Teams can quickly see which incidents need their immediate attention. If your organization already have priority levels defined, you can customize these default values to match your existing incident classification scheme.

 **Note:** To list the priorities, first [enable and configure your priorities](https://support.pagerduty.com/docs/incident-priority#section-enabling-incident-priority).

## Examples

### Basic info

```sql
select
  name,
  id,
  description,
  self
from
  pagerduty_priority;
```

### List event rules with highest-priority(P1)

```sql
with priority as (
  select
    id
  from
    pagerduty_priority
  where
    name = 'P1'
)
select
  rs.id as rule_id,
  rs.ruleset_id,
  rs.disabled
from
  pagerduty_ruleset_rule as rs,
  priority as p
where
  p.id = rs.actions -> 'priority' ->> 'value';
```
