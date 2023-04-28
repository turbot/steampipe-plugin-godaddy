---
organization: Turbot
category: ["internet"]
icon_url: "/images/plugins/turbot/godaddy.svg"
brand_color: "#1BDBDB"
display_name: "GoDaddy"
short_name: "godaddy"
description: "Steampipe plugin to query projects, groups, builds and more from GoDaddy."
og_description: "Query GoDaddy with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/godaddy-social-graphic.png"
---

# GoDaddy + Steampipe

[GoDaddy](https://godaddy.com) is a publicly traded internet domain registrar and web hosting company. It provides services such as domain registration, website hosting, website building, and email hosting for individuals and businesses.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

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

## Documentation

- **[Table definitions & examples →](/plugins/turbot/godaddy/tables)**

## Quick start

### Install

Download and install the latest GoDaddy plugin:

```sh
steampipe plugin install godaddy
```

### Credentials

| Item        | Description                                                                                                                                                                                                                                                                                                                                           |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | GoDaddy requires an  `API Key` and `Secret Key` for all requests. Create [API keys](https://developer.godaddy.com/keys) and add to `~/.steampipe/config/godaddy.spc`.  GoDaddy API access is segregated for OTE (Operational Test Environment) & Production environment. Also, note that the API and secret keys used to authenticate with the production environment will not work in the OTE environment. You must generate a separate set of API keys for OTE testing. See [here](https://developer.godaddy.com/getstarted) for information. |                                                                                                                                                                                 |                                                                                                                                                                                                                                                                   |
| Radius  | Each connection represents a single GoDaddy account.                                                                                                                             |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/godaddy.spc`)<br />2. Credentials specified in environment variables, e.g., `GODADDY_ENVIRONMENT_TYPE`, `GODADDY_API_KEY` and `GODADDY_SECRET_KEY`.                                                                                                                               |

### Configuration

Installing the latest godaddy plugin will create a config file (`~/.steampipe/config/godaddy.spc`) with a single connection named `godaddy`:

```hcl
connection "godaddy" {
  plugin = "godaddy"

  # For setting API keys see instructions at https://developer.godaddy.com/keys
  # The API key for the GoDaddy account.
  # api_key = "hkw64xxxxabchttXW_TmcsFgxJQBvLjE5Lda8402"

  # The secret key for the GoDaddy account.
  # secret_key = "DjfrsqEBA4vVjsdsdsdieDShdjs"

  # The type of the environment, based on the value the endpoint will be change. Possible values are: DEV | PROD.
  # environment_type = "DEV"
}
```

Alternatively, you can also use the GoDaddy environment variables to obtain credentials **only if other arguments `api_key`  `secret_key` and `environment_type` are not specified** in the connection:

```sh
export GODADDY_ENVIRONMENT_TYPE="DEV"
export GODADDY_API_KEY="hkw64xxxxabchttXW_TmcsFgxJQBvLjE5Lda8402"
export GODADDY_SECRET_KEY="DjfrsqEBA4vVjsdsdsdieDShdjs"
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-godaddy
- Community: [Slack Channel](https://steampipe.io/community/join)
