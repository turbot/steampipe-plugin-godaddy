---
title: "Steampipe Table: godaddy_subscription - Query GoDaddy Subscriptions using SQL"
description: "Allows users to query GoDaddy Subscriptions, specifically their details and statuses, providing insights into subscription patterns and potential anomalies."
---

# Table: godaddy_subscription - Query GoDaddy Subscriptions using SQL

A GoDaddy Subscription is a service that allows users to avail various GoDaddy services on a recurring basis. It provides a centralized way to manage and maintain subscriptions for various GoDaddy resources, including domains, hosting, and more. GoDaddy Subscription helps you stay informed about the status and details of your GoDaddy resources and take appropriate actions when renewal or cancellation is required.

## Table Usage Guide

The `godaddy_subscription` table provides insights into subscriptions within GoDaddy. As a DevOps engineer or a domain manager, explore subscription-specific details through this table, including renewal status, expiration dates, and associated metadata. Utilize it to uncover information about subscriptions, such as those nearing expiration, the renewal status of subscriptions, and verification of subscription details.

## Examples

### Basic info
Explore which GoDaddy subscriptions are cancelable and when they are due to expire. This is useful for assessing potential changes in your domain portfolio costs and planning ahead.

```sql+postgres
select
  subscription_id,
  created_at,
  cancelable,
  label,
  launch_url,
  payment_profile_id,
  price_locked,
  expires_at
from
  godaddy_subscription;
```

```sql+sqlite
select
  subscription_id,
  created_at,
  cancelable,
  label,
  launch_url,
  payment_profile_id,
  price_locked,
  expires_at
from
  godaddy_subscription;
```

### List subscriptions with price locked enabled
Analyze your GoDaddy subscriptions to identify which ones have a locked price. This can be useful to understand your financial commitments and plan for future costs.

```sql+postgres
select
  subscription_id,
  created_at,
  cancelable,
  price_locked,
  expires_at
from
  godaddy_subscription
where
  price_locked;
```

```sql+sqlite
select
  subscription_id,
  created_at,
  cancelable,
  price_locked,
  expires_at
from
  godaddy_subscription
where
  price_locked = 1;
```

### Count subscriptions by product key
Analyze your product subscriptions to understand their distribution. This is useful in identifying the most subscribed products, aiding in decision-making for marketing strategies or product development.

```sql+postgres
select
  product_group_key,
  count(*) as subscription_count
from
 godaddy_subscription
group by
  product_group_key;
```

```sql+sqlite
select
  product_group_key,
  count(*) as subscription_count
from
 godaddy_subscription
group by
  product_group_key;
```

### Get subscription details that have addons
Explore which subscription details include addons to understand the pricing and cancellation options. This can be useful for assessing the costs and flexibility of different subscriptions.

```sql+postgres
select
  subscription_id,
  created_at,
  cancelable,
  price_locked
from
  godaddy_subscription
where
  addons is not null;
```

```sql+sqlite
select
  subscription_id,
  created_at,
  cancelable,
  price_locked
from
  godaddy_subscription
where
  addons is not null;
```

### List non-cancelable subscriptions
Determine the areas in which you have non-cancelable subscriptions. This can be beneficial for understanding your ongoing commitments and potential liabilities.

```sql+postgres
select
  subscription_id,
  created_at,
  cancelable
from
  godaddy_subscription
where
  not cancelable;
```

```sql+sqlite
select
  subscription_id,
  created_at,
  cancelable
from
  godaddy_subscription
where
  cancelable = 0;
```

### List subscriptions that were created between two specific dates
Explore which subscriptions were initiated within a particular year to gain insights into customer engagement and subscription trends during that period. This can assist in understanding the effectiveness of marketing strategies implemented within that timeframe.

```sql+postgres
select
  subscription_id,
  created_at,
  expires_at,
  renew_auto
from
  godaddy_subscription
where
  created_at between '2022-01-01' and '2022-12-31';
```

```sql+sqlite
select
  subscription_id,
  created_at,
  expires_at,
  renew_auto
from
  godaddy_subscription
where
  created_at between '2022-01-01' and '2022-12-31';
```

### List subscriptions that have auto-renew enabled
Explore which subscriptions have the auto-renew feature enabled. This is useful for understanding which customers will automatically renew their subscriptions, aiding in revenue forecasting and customer retention strategies.

```sql+postgres
select
  subscription_id,
  created_at,
  product_renewal_period,
  product_renewal_period_unit,
  status,
  upgradable
from
  godaddy_subscription
where
  renew_auto;
```

```sql+sqlite
select
  subscription_id,
  created_at,
  product_renewal_period,
  product_renewal_period_unit,
  status,
  upgradable
from
  godaddy_subscription
where
  renew_auto = 1;
```

### List active subscriptions
Explore which subscriptions are currently active to manage and monitor your GoDaddy account effectively. This helps in assessing the status of your subscriptions, allowing for better resource allocation and planning.

```sql+postgres
select
  subscription_id,
  status,
  created_at
from
  godaddy_subscription
where
  status = 'ACTIVE';
```

```sql+sqlite
select
  subscription_id,
  status,
  created_at
from
  godaddy_subscription
where
  status = 'ACTIVE';
```

### List subscriptions that will expire within the next 10 days
Discover the segments that are due for renewal in the coming days. This helps in proactive renewal management and avoids service disruption due to expired subscriptions.

```sql+postgres
select
  subscription_id,
  status,
  created_at,
  expires_at
from
  godaddy_subscription
where
  expires_at <= now() + interval '10' day;
```

```sql+sqlite
select
  subscription_id,
  status,
  created_at,
  expires_at
from
  godaddy_subscription
where
  expires_at <= datetime('now', '+10 day');
```

### Get product details of each subscription
Explore the specific details of each subscription, such as product label, namespace, renewal period, and billing information. This can help in understanding the subscription structure, tracking renewals, and managing billing cycles.

```sql+postgres
select
  subscription_id,
  product -> 'PfID' as pf_id,
  product ->> 'Label' as product_label,
  product ->> 'Namespace' as product_namespace,
  product -> 'RenewalPfID' as product_renewal_pf_id,
  product -> 'PRenewalPeriodfID' as product_renewal_period,
  product -> 'SupportBillOn' as product_support_bill_on,
  product ->> 'ProductGroupKey' as product_group_key,
  product ->> 'RenewalPeriodUnit' as product_renewal_period_unit
from
  godaddy_subscription;
```

```sql+sqlite
select
  subscription_id,
  json_extract(product, '$.PfID') as pf_id,
  json_extract(product, '$.Label') as product_label,
  json_extract(product, '$.Namespace') as product_namespace,
  json_extract(product, '$.RenewalPfID') as product_renewal_pf_id,
  json_extract(product, '$.PRenewalPeriodfID') as product_renewal_period,
  json_extract(product, '$.SupportBillOn') as product_support_bill_on,
  json_extract(product, '$.ProductGroupKey') as product_group_key,
  json_extract(product, '$.RenewalPeriodUnit') as product_renewal_period_unit
from
  godaddy_subscription;
```

### Get the billing details of each subscription
Discover the billing status and renewal date of each subscription to understand their commitment and potential past due issues. This can be used to manage financial aspects and ensure timely renewals.

```sql+postgres
select
  subscription_id,
  billing ->> 'Status' as billing_status,
  billing ->> 'RenewAt' as billing_renew_at,
  billing ->> 'Commitment' as billing_commitment,
  billing ->> 'PastDueTypes' as billing_past_due_types
from
  godaddy_subscription;
```

```sql+sqlite
select
  subscription_id,
  json_extract(billing, '$.Status') as billing_status,
  json_extract(billing, '$.RenewAt') as billing_renew_at,
  json_extract(billing, '$.Commitment') as billing_commitment,
  json_extract(billing, '$.PastDueTypes') as billing_past_due_types
from
  godaddy_subscription;
```