# Table: godaddy_order

The Order feature of GoDaddy is an API that allows developers to programmatically manage domain name orders and other related products.

### Basic info

```sql
select
  order_id,
  currency,
  created_at,
  domain_name,
  parent_order_id
from
  godaddy_order;
```

### Get the total amount paid for all orders

```sql
select
  order_id,
  SUM((pricing ->> 'Total')::numeric) as total_paid
from
  godaddy_order
group by
  order_id;
```

### Count the number of orders created per month

```sql
select
  order_id,
  date_trunc('month', created_at) as month,
  count(*) as num_orders
from
  godaddy_order
group by
  order_id,
  created_at;
```

### List orders by parent order ID

```sql
select
  order_id,
  parent_order_id,
  domain_name
from
  godaddy_order
where
  parent_order_id = '12345';
```

### Get line item details of each order

```sql
select
  order_id,
  i -> 'Domains' as domains,
  i -> 'Period' as period,
  i ->> 'PeriodUnit' as period_unit,
  i -> 'Pfid' as pfid,
  i -> 'Quantity' as quantity,
  i -> 'TaxCollector' ->> 'TaxCollectorID' as tax_collector_id
from
  godaddy_order,
  jsonb_array_elements(items) as i;
```

### Get pricing details of an order

```sql
select
  order_id,
  domain_name,
  pricing ->> 'Discount' as discount,
  pricing -> 'Fees' as fees,
  pricing ->> 'ID' as pricing_id,
  pricing ->> 'List' as pricing_list,
  pricing ->> 'Savings' as savings,
  pricing ->> 'Subtotal' as subtotal,
  pricing ->> 'Taxes' as taxes,
  pricing -> 'TaxDetails' as tax_details,
  pricing ->> 'Total' as total
from
  godaddy_order
where
  order_id = '123456';
```

### Get domain details for each order

```sql
select
  o.order_id,
  o.created_at,
  o.domain_name,
  d.domain_id,
  d.created_at as domain_created_at,
  d.expires as domain_expires
from
  godaddy_order as o,
  godaddy_domain as d
where
  o.domain_name = d.domain;
```

### List the top 5 orders with the highest total amount paid

```sql
select
  order_id,
  (
    pricing ->> 'Total'
  )
  ::numeric as total_paid
from
  godaddy_order
order by
  total_paid desc limit 5;
```

### Count orders per domain

```sql
select
  domain_name,
  count(*) as num_orders
from
  godaddy_order
group by
  domain_name;
```
