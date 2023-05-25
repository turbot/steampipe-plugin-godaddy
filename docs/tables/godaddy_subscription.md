# Table: godaddy_subscription

GoDaddy is a popular web hosting and domain registrar company that offers various subscription plans for individuals and businesses. Their subscriptions typically include web hosting services, a domain name registration, email hosting, website building tools, and marketing tools. The specific features and pricing of the subscription plans vary depending on the type of subscription and the length of the subscription term. Subscribers can choose from a range of options that cater to their specific needs and budget

## Examples

### Basic info

```sql
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

```sql
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

### Count subscriptions by product key

```sql
select
  product_group_key,
  count(*) as subscription_count
from
 godaddy_subscription
group by
  product_group_key;
```

### Get subscription details that have addons

```sql
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

```sql
select
  subscription_id,
  created_at,
  cancelable
from
  godaddy_subscription
where
  not cancelable;
```

### List subscriptions that are created between two specific dates

```sql
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

```sql
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

### List active subscriptions

```sql
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

```sql
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

### Get product details for each subscription

```sql
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

### Get the billing details for each subscription

```sql
select
  subscription_id,
  billing ->> 'Status' as billing_status,
  billing ->> 'RenewAt' as billing_renew_at,
  billing ->> 'Commitment' as billing_commitment,
  billing ->> 'PastDueTypes' as billing_past_due_types
from
  godaddy_subscription;
```
