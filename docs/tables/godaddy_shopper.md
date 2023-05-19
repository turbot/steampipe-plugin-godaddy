# Table: godaddy_shopper

GoDaddy Shopper feature allows customers to create a centralized account to manage all of their GoDaddy products and services in one place.

Note: The used [Go API](https://pkg.go.dev/github.com/alyx/go-daddy/daddy#ShoppersService.Get) limitation, you must provide valid `shopper_id` in the `where` clause of the queries to fetch shopping details.

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
  shopper_id ='568421234'
```
