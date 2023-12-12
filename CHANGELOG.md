## v0.2.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#17](https://github.com/turbot/steampipe-plugin-godaddy/pull/17))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#17](https://github.com/turbot/steampipe-plugin-godaddy/pull/17))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-godaddy/blob/main/docs/LICENSE). ([#17](https://github.com/turbot/steampipe-plugin-godaddy/pull/17))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#16](https://github.com/turbot/steampipe-plugin-godaddy/pull/16))

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
