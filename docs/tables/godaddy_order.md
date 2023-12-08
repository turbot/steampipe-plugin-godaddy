---
title: "Steampipe Table: godaddy_order - Query GoDaddy Orders using SQL"
description: "Allows users to query GoDaddy Orders, providing comprehensive details about each order placed."
---

# Table: godaddy_order - Query GoDaddy Orders using SQL

GoDaddy is a widely used domain registrar and web hosting company. It offers a variety of services including domain registration, website hosting, and email hosting. An order in GoDaddy represents a transaction for purchasing one or more of these services.

## Table Usage Guide

The `godaddy_order` table provides insights into orders placed within GoDaddy. As a DevOps engineer, explore order-specific details through this table, including currency, items ordered, and status of the order. Utilize it to track and manage orders, such as those pending, completed, or cancelled.

## Examples

### Basic info
Determine the areas in which specific orders were placed, including the currency used and the date of order placement. This can be useful for understanding your customer base and tracking order trends.

```sql+postgres
select
  order_id,
  currency,
  created_at,
  domain_name,
  parent_order_id
from
  godaddy_order;
```

```sql+sqlite
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
Determine the total payment made for each order. This can be used to monitor financial transactions and ensure accurate tracking of payments.

```sql+postgres
select
  order_id,
  SUM((pricing ->> 'Total')::numeric) as total_paid
from
  godaddy_order
group by
  order_id;
```

```sql+sqlite
select
  order_id,
  SUM(CAST(json_extract(pricing, '$.Total') AS REAL)) as total_paid
from
  godaddy_order
group by
  order_id;
```

### Count the number of orders created per month
Determine the frequency of orders on a monthly basis to understand the sales trend and identify peak periods. This can aid in planning inventory and sales strategies accordingly.

```sql+postgres
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

```sql+sqlite
select
  order_id,
  strftime('%Y-%m', created_at) as month,
  count(*) as num_orders
from
  godaddy_order
group by
  order_id,
  strftime('%Y-%m', created_at);
```

### List orders by parent order ID
Explore which orders are related to a specific parent order, helping to track and manage grouped transactions in a more organized manner. This is particularly useful in scenarios where you need to understand the relationship between different orders for better order management.

```sql+postgres
select
  order_id,
  parent_order_id,
  domain_name
from
  godaddy_order
where
  parent_order_id = '12345';
```

```sql+sqlite
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
Uncover the details of each order line item to gain insights into various aspects like the domains involved, the duration of the order, the quantity of items, and the tax collector's details. This can be useful for auditing purposes, inventory management, or financial analysis.

```sql+postgres
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

```sql+sqlite
select
  order_id,
  json_extract(i.value, '$.Domains') as domains,
  json_extract(i.value, '$.Period') as period,
  json_extract(i.value, '$.PeriodUnit') as period_unit,
  json_extract(i.value, '$.Pfid') as pfid,
  json_extract(i.value, '$.Quantity') as quantity,
  json_extract(json_extract(i.value, '$.TaxCollector'), '$.TaxCollectorID') as tax_collector_id
from
  godaddy_order,
  json_each(items) as i;
```

### Get pricing details of an order
Determine the financial aspects of a specific order, including discounts, fees, savings, and taxes. This is useful for a comprehensive review of an order's cost breakdown and can assist in financial planning and budgeting.

```sql+postgres
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

```sql+sqlite
select
  order_id,
  domain_name,
  json_extract(pricing, '$.Discount') as discount,
  pricing as fees,
  json_extract(pricing, '$.ID') as pricing_id,
  json_extract(pricing, '$.List') as pricing_list,
  json_extract(pricing, '$.Savings') as savings,
  json_extract(pricing, '$.Subtotal') as subtotal,
  json_extract(pricing, '$.Taxes') as taxes,
  pricing as tax_details,
  json_extract(pricing, '$.Total') as total
from
  godaddy_order
where
  order_id = '123456';
```

### Get domain details for each order
Determine the chronological relationship between order creation and domain registration, and assess when the domain will expire. This is useful for understanding the lifecycle of an order and its associated domain.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that have the highest total payments to prioritize customer follow-ups or assess the performance of top-selling products. This query helps in financial analysis and strategic planning by identifying the top 5 orders with the highest total amount paid.

```sql+postgres
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

```sql+sqlite
select
  order_id,
  cast(json_extract(pricing, '$.Total') as real) as total_paid
from
  godaddy_order
order by
  total_paid desc limit 5;
```

### Count orders per domain
Determine the areas in which your domains have the most orders to gain insights into your most popular domains.

```sql+postgres
select
  domain_name,
  count(*) as num_orders
from
  godaddy_order
group by
  domain_name;
```

```sql+sqlite
select
  domain_name,
  count(*) as num_orders
from
  godaddy_order
group by
  domain_name;
```