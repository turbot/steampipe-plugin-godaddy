![image](https://hub.steampipe.io/images/plugins/turbot/godaddy-social-graphic.png)

# GoDaddy Plugin for Steampipe

Use SQL to query projects, groups, builds and more from GoDaddy.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/godaddy)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/godaddy/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-godaddy/issues)

## Quick start

### Install

Download and install the latest GoDaddy plugin:

```bash
steampipe plugin install godaddy
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/godaddy#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/godaddy#configuration).

Configure the Godaddy API Key, API Secret and Environment in `~/.steampipe/config/godaddy.spc`:

```hcl
connection "godaddy" {
  plugin = "godaddy"

  # Authentication information
  api_key = "hkw647irnrhttXW_TmcsFgxJQBvLjE5L1234402"
  api_secret = "DjfrsqEB12345hdsieDShdjs"
  environment = "PROD"
}
```

Or through environment variables:

```sh
export GODADDY_API_KEY=hkw64xxxxabchttXW_TmcsFgxJQBvLjE5Lda8402
export GODADDY_API_SECRET=DjfrsqEBA4vVjsdsdsdieDShdjs
export GODADDY_ENVIRONMENT=PROD
```

Run steampipe:

```shell
steampipe query
```

List your GoDaddy domains:

```sql
select
  domain,
  status,
  nameservers,
  privacy
from
  godaddy_domain;
```

```
+-----------------+-----------------------------------------------------+--------+---------+
| domain          | nameservers                                         | status | privacy |
+-----------------+-----------------------------------------------------+--------+---------+
| mycloudlab.in   | ["ns21.domaincontrol.com","ns22.domaincontrol.com"] | ACTIVE | false   |
+-----------------+-----------------------------------------------------+--------+---------+
```

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/index) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs//steampipe_sqlite/index) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/index) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-godaddy.git
cd steampipe-plugin-godaddy
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/godaddy.spc
```

Try it!

```
steampipe query
> .inspect godaddy
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-godaddy/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-godaddy/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [GoDaddy Plugin](https://github.com/turbot/steampipe-plugin-godaddy/labels/help%20wanted)
