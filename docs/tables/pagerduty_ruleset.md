# Table: pagerduty_ruleset

Rulesets allow you to route events to an endpoint and create collections of Event Rules, which define sets of actions to take based on event content.

## Examples

### Basic info

```sql
select
  name,
  id,
  type
from
  pagerduty_ruleset;
```

### List default global rulesets

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

```sql
select
  rs.id as ruleset_id,
  count(r.id)
from
  pagerduty_ruleset as rs
  left join pagerduty_ruleset_rule as r on rs.id = r.ruleset_id
group by rs.id;
```
