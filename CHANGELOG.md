## v0.4.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#38](https://github.com/turbot/steampipe-plugin-pagerduty/pull/38))

## v0.4.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#35](https://github.com/turbot/steampipe-plugin-pagerduty/pull/35))
- Recompiled plugin with Go version `1.21`. ([#35](https://github.com/turbot/steampipe-plugin-pagerduty/pull/35))

## v0.3.0 [2023-03-22]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#30](https://github.com/turbot/steampipe-plugin-pagerduty/pull/30))

## v0.2.1 [2022-09-28]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#26](https://github.com/turbot/steampipe-plugin-pagerduty/pull/26))

## v0.2.0 [2022-08-30]

_What's new?_

- New tables added
  - [pagerduty_on_call](https://hub.steampipe.io/plugins/turbot/pagerduty/tables/pagerduty_on_call) ([#15](https://github.com/turbot/steampipe-plugin-pagerduty/pull/15)) (Thanks to [@coop182](https://github.com/coop182) for the contribution!)

_Bug fixes_

- Fixed offset calculation in all tables' list functions' paging. ([#19](https://github.com/turbot/steampipe-plugin-pagerduty/pull/19)) (Thanks to [@janritter](https://github.com/janritter) for the contribution!)
- Fixed `title` column in `pagerduty_incident` and `pagerduty_incident_log` tables. ([#23](https://github.com/turbot/steampipe-plugin-pagerduty/pull/23))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.4](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v414-2022-08-26) which includes several caching and memory management improvements. ([#20](https://github.com/turbot/steampipe-plugin-pagerduty/pull/20))
- Recompiled plugin with Go version `1.19`. ([#20](https://github.com/turbot/steampipe-plugin-pagerduty/pull/20))

## v0.1.0 [2022-04-27]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#12](https://github.com/turbot/steampipe-plugin-pagerduty/pull/12))
- Added support for native Linux ARM and Mac M1 builds. ([#13](https://github.com/turbot/steampipe-plugin-pagerduty/pull/13))

## v0.0.2 [2022-01-19]

_What's new?_

- New tables added
  - [pagerduty_incident](https://hub.steampipe.io/plugins/turbot/pagerduty/tables/pagerduty_incident) ([#4](https://github.com/turbot/steampipe-plugin-pagerduty/pull/4))
  - [pagerduty_incident_log](https://hub.steampipe.io/plugins/turbot/pagerduty/tables/pagerduty_incident_log) ([#6](https://github.com/turbot/steampipe-plugin-pagerduty/pull/6))
  - [pagerduty_service_integration](https://hub.steampipe.io/plugins/turbot/pagerduty/tables/pagerduty_service_integration) ([#8](https://github.com/turbot/steampipe-plugin-pagerduty/pull/8))
  - [pagerduty_vendor](https://hub.steampipe.io/plugins/turbot/pagerduty/tables/pagerduty_vendor) ([#10](https://github.com/turbot/steampipe-plugin-pagerduty/pull/10))

## v0.0.1 [2022-01-05]

_What's new?_

- New tables added
  - [pagerduty_escalation_policy](https://hub.steampipe.io/plugins/turbot/pagerduty/tables/pagerduty_escalation_policy)
  - [pagerduty_priority](https://hub.steampipe.io/plugins/turbot/pagerduty/tables/pagerduty_priority)
  - [pagerduty_ruleset](https://hub.steampipe.io/plugins/turbot/pagerduty/tables/pagerduty_ruleset)
  - [pagerduty_ruleset_rule](https://hub.steampipe.io/plugins/turbot/pagerduty/tables/pagerduty_ruleset_rule)
  - [pagerduty_schedule](https://hub.steampipe.io/plugins/turbot/pagerduty/tables/pagerduty_schedule)
  - [pagerduty_service](https://hub.steampipe.io/plugins/turbot/pagerduty/tables/pagerduty_service)
  - [pagerduty_tag](https://hub.steampipe.io/plugins/turbot/pagerduty/tables/pagerduty_tag)
  - [pagerduty_team](https://hub.steampipe.io/plugins/turbot/pagerduty/tables/pagerduty_team)
  - [pagerduty_user](https://hub.suserpipe.io/plugins/turbot/pagerduty/tables/pagerduty_user)
