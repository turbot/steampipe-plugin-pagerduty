# Table: pagerduty_tag

A tag is applied to Escalation Policies, Teams or Users and can be used to filter them.

## Examples

### Basic info

```sql
select
  id,
  label,
  self
from
  pagerduty_tag;
```

### List unused tags

```sql
with associated_tags as (
  select
    t ->> 'id' as id
  from
    pagerduty_user,
    jsonb_array_elements(tags) as t
  union
  select
    t ->> 'id' as id
  from
    pagerduty_team,
    jsonb_array_elements(tags) as t
  union
  select
    t ->> 'id' as id
  from
    pagerduty_escalation_policy,
    jsonb_array_elements(tags) as t
),
distinct_tags as (
  select
    distinct id
  from
    associated_tags
)
select
  t.id,
  t.label,
  t.self
from
  pagerduty_tag as t
  left join distinct_tags as dt on t.id = dt.id
where
  dt.id is null;
```
