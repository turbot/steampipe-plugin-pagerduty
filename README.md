![image](https://hub.steampipe.io/images/plugins/turbot/pagerduty-social-graphic.png)

# PagerDuty Plugin for Steampipe

Use SQL to query infrastructure services, teams, escalation policies and more from your PagerDuty account.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/pagerduty)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/pagerduty/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-pagerduty/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install pagerduty
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/pagerduty#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/pagerduty#configuration).

Run a query:

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-pagerduty.git
cd steampipe-plugin-pagerduty
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```sh
make
```

Configure the plugin:

```sh
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/pagerduty.spc
```

Try it!

```shell
steampipe query
> .inspect pagerduty
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-pagerduty/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [PagerDuty Plugin](https://github.com/turbot/steampipe-plugin-pagerduty/labels/help%20wanted)
