---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/pagerduty.svg"
brand_color: "#06ac38"
display_name: "PagerDuty"
short_name: "pagerduty"
description: "Steampipe plugin to query services, teams, escalation policies and more from your PagerDuty account."
og_description: "Query PagerDuty with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/pagerduty-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# PagerDuty + Steampipe

[PagerDuty](https://www.pagerduty.com/) is a platform for agile incident management.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List disabled services in your PagerDuty account:

```sql
select
  name,
  id,
  status
from
  pagerduty_service
where
  status = 'disabled';
```

```
+-----------+---------+--------+
| name      | id      | status |
+-----------+---------+--------+
| Steampipe | PE0PJEP | active |
+-----------+---------+--------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/pagerduty/tables)**

## Get started

### Install

Download and install the latest PagerDuty plugin:

```bash
steampipe plugin install pagerduty
```

### Credentials

| Item | Description |
| - | - |
| Credentials | [Get your user token](https://support.pagerduty.com/docs/generating-api-keys#generating-a-personal-rest-api-key) or if you have `Admin`, `Global Admin` or `Account Owner` access within your PagerDuty account, [generate a general authorization token](https://support.pagerduty.com/docs/generating-api-keys#generating-a-general-access-rest-api-key). |
| Resolution | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/pagerduty.spc`).<br />2. Credentials specified in environment variables, e.g., `PAGERDUTY_TOKEN`. |

### Configuration

Installing the latest pagerduty plugin will create a config file (`~/.steampipe/config/pagerduty.spc`) with a single connection named `pagerduty`:

```hcl
connection "pagerduty" {
  plugin = "pagerduty"

  # Account or user API token
  # This can also be set via the `PAGERDUTY_TOKEN` environment variable.
  # token = "u+AtBdqvNtestTokeNcg"
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-pagerduty
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
