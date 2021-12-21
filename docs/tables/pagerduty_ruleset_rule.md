# Table: pagerduty_ruleset_rule

An event rule allows you to set actions that should be taken on events that meet your designated rule criteria.

## Examples

### Basic info

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

```sql
select
  rs.id as ruleset_id,
  count(r.id)
from
  pagerduty_ruleset as rs
  left join pagerduty_ruleset_rule as r on rs.id = r.ruleset_id
group by rs.id;
```

### List rules with no priority

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
