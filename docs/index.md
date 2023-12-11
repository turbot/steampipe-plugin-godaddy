---
organization: Turbot
category: ["internet"]
icon_url: "/images/plugins/turbot/godaddy.svg"
brand_color: "#1CDBDB"
display_name: "GoDaddy"
short_name: "godaddy"
description: "Steampipe plugin to query domains, orders, certificates and more from GoDaddy."
og_description: "Query GoDaddy with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/godaddy-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
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
| Credentials | GoDaddy requires an  `API Key` and `API Secret` for all requests. Create [API keys](https://developer.godaddy.com/keys) and add to `~/.steampipe/config/godaddy.spc`.  GoDaddy API access is segregated for OTE (Operational Test Environment) & Production environment. Also, note that the API and secret keys used to authenticate with the production environment will not work in the OTE environment. You must generate a separate set of API keys for OTE testing. See [here](https://developer.godaddy.com/getstarted) for information. |                                                                                                                                                                                 |                                                                                                                                                                                                                                                                   |
| Radius  | Each connection represents a single GoDaddy account.                                                                                                                             |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/godaddy.spc`)<br />2. Credentials specified in environment variables, e.g., `GODADDY_API_KEY`, `GODADDY_API_SECRET`and `GODADDY_ENVIRONMENT`.

### Configuration

Installing the latest godaddy plugin will create a config file (`~/.steampipe/config/godaddy.spc`) with a single connection named `godaddy`:

```hcl
connection "godaddy" {
  plugin = "godaddy"

  # For setting API keys see instructions at https://developer.godaddy.com/keys
  # `api_key`: The API key of the GoDaddy account. (Required).
  # This can also be set via the `GODADDY_API_KEY` environment variable.
  # api_key = "hkw647irnrhttXW_TmcsFgxJQBvLjE5L1234402"

  # `api_secret`: The secret key of the GoDaddy account. (Required).
  # This can also be set via the `GODADDY_API_SECRET` environment variable.
  # api_secret = "DjfrsqEB12345hdsieDShdjs"

  # `environment`: The type of the environment, based on the value the API endpoint will change. Possible values are: DEV | PROD. (Optional)
  # Defaults to PROD
  # This can also be set via the `GODADDY_ENVIRONMENT` environment variable.
  # environment = "PROD"
}
```

Alternatively, you can also use the GoDaddy environment variables to obtain credentials **only if other arguments `api_key`  `api_secret` and `environment` are not specified** in the connection:

```sh
export GODADDY_API_KEY=hkw64xxxxabchttXW_TmcsFgxJQBvLjE5Lda8402
export GODADDY_API_SECRET=DjfrsqEBA4vVjsdsdsdieDShdjs
export GODADDY_ENVIRONMENT=DEV
```


