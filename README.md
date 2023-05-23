![image](https://hub.steampipe.io/images/plugins/turbot/godaddy-social-graphic.png)

# GoDaddy Plugin for Steampipe

Use SQL to query projects, groups, builds and more from GoDaddy.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/godaddy)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/godaddy/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-godaddy/issues)

## Quick start

### Install

Download and install the latest GoDaddy plugin:

```bash
steampipe plugin install godaddy
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/godaddy#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/godaddy#configuration).

Configure the Organization URL and Personal Access Token in `~/.steampipe/config/godaddy.spc`:

```hcl
connection "godaddy" {
  plugin = "godaddy"

  # Authentication information
  api_key = "hkw647irnrhttXW_TmcsFgxJQBvLjE5L1234402"
  secret_key = "DjfrsqEB12345hdsieDShdjs"
  environment_type = "DEV"
}
```

Or through environment variables:

```sh
export GODADDY_API_KEY="hkw64xxxxabchttXW_TmcsFgxJQBvLjE5Lda8402"
export GODADDY_SECRET_KEY="DjfrsqEBA4vVjsdsdsdieDShdjs"
export GODADDY_ENVIRONMENT_TYPE="DEV"
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

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-godaddy/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [GoDaddy Plugin](https://github.com/turbot/steampipe-plugin-godaddy/labels/help%20wanted)
