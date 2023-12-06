---
title: "Steampipe Table: godaddy_shopper - Query GoDaddy Shoppers using SQL"
description: "Allows users to query GoDaddy Shoppers, returning detailed information about the registered shoppers on GoDaddy."
---

# Table: godaddy_shopper - Query GoDaddy Shoppers using SQL

GoDaddy is a widely-used domain registrar and web hosting company. The 'Shopper' is a resource in GoDaddy that represents a registered user who has the ability to purchase and manage domains. It provides details about the shopper's profile, including their shopper ID, name, and email.

## Table Usage Guide

The `godaddy_shopper` table provides insights into registered shoppers within GoDaddy. As a system administrator or a DevOps engineer, you can explore shopper-specific details through this table, including shopper ID, name, and email. Utilize it to uncover information about shoppers, such as their contact details and the status of their accounts.

**Important Notes**
- You must specify the `shopper_id` in the `where` clause to query this table.

## Examples

### Basic info
Discover the segments that include specific customer data to understand their market affiliation and contact details. This can be useful to gain insights into individual customer profiles for personalized marketing or customer service efforts.

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