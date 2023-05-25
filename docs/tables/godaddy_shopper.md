# Table: godaddy_shopper

GoDaddy Shopper feature allows customers to create a centralized account to manage all of their GoDaddy products and services in one place.

Note: To query this table you must provide a valid `shopper_id` in the `where` clause.

## Examples

### Basic info

```sql
select
  customer_id,
  market_id,
  name_first,
  name_last,
  email,
  external_id
from
  godaddy_shopper
where
  shopper_id ='568421234';
```
