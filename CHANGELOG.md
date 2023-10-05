## v0.1.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#9](https://github.com/turbot/steampipe-plugin-godaddy/pull/9))

## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#6](https://github.com/turbot/steampipe-plugin-godaddy/pull/6))
- Recompiled plugin with Go version `1.21`. ([#6](https://github.com/turbot/steampipe-plugin-godaddy/pull/6))

## v0.0.1 [2023-05-25]

_What's new?_

- New tables added
  - [godaddy_certificate](https://hub.steampipe.io/plugins/turbot/godaddy/tables/godaddy_certificate)
  - [godaddy_dns_record](https://hub.steampipe.io/plugins/turbot/godaddy/tables/godaddy_dns_record)
  - [godaddy_domain](https://hub.steampipe.io/plugins/turbot/godaddy/tables/godaddy_domain)
  - [godaddy_order](https://hub.steampipe.io/plugins/turbot/godaddy/tables/godaddy_order)
  - [godaddy_shopper](https://hub.steampipe.io/plugins/turbot/godaddy/tables/godaddy_shopper)
  - [godaddy_subscription](https://hub.steampipe.io/plugins/turbot/godaddy/tables/godaddy_subscription)
