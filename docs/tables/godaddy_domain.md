---
title: "Steampipe Table: godaddy_domain - Query GoDaddy Domains using SQL"
description: "Allows users to query Domains in GoDaddy, specifically the domain name, status, and expiration date, providing insights into domain management and potential renewals."
---

# Table: godaddy_domain - Query GoDaddy Domains using SQL

GoDaddy is a popular web hosting and domain registrar company. Domains in GoDaddy represent the unique web addresses or URLs purchased by users. These domains are essential for establishing an online presence, and their management includes details such as domain name, status, and expiration date.

## Table Usage Guide

The `godaddy_domain` table provides insights into Domains within GoDaddy. As a Web Administrator, explore domain-specific details through this table, including domain name, status, and expiration date. Use it to manage your domains effectively, such as those nearing expiration, and to stay ahead of domain renewals.

## Examples

### Basic info
Gain insights into the creation, deletion, and expiration timelines of domains. This can be useful for managing domain lifecycles and identifying when domains are eligible for transfer.

```sql
select
  domain,
  domain_id,
  created_at,
  deleted_at,
  transfer_away_eligibile_at,
  expires
from
  godaddy_domain;
```

### List domains that have privacy protection enabled
Determine the areas in which domains have privacy protection enabled. This can be useful for assessing the security measures in place and ensuring the safeguarding of sensitive information.

```sql
select
  domain,
  domain_id,
  created_at,
  expiration_protected,
  privacy
from
  godaddy_domain
where
  privacy;
```

### List domains that are locked
Identify domains that have been locked, which can be useful for maintaining security and control over domain ownership and transfers.

```sql
select
  domain,
  expiration_protected,
  hold_registrar,
  locked
from
  godaddy_domain
where
  locked;
```

### List renewable domains
Explore which domains are renewable to manage and maintain your online presence efficiently. This is beneficial for planning renewals in advance and avoiding unexpected expiration of important domains.

```sql
select
  domain,
  domain_id,
  status,
  created_at
from
  godaddy_domain
where
  renewable;
```

### List inactive domains
Discover the segments that consist of inactive domains, which can help in identifying those that are not currently in use or have expired. This information can be beneficial for domain management, enabling you to assess the status of your domains and take necessary actions such as renewals or transfers.

```sql
select
  domain,
  auth_code,
  transfer_away_eligibile_at,
  expires,
  status
from
  godaddy_domain
where
  status <> 'ACTIVE';
```

### List domains that are eligible for transfer at the moment
Explore domains that are currently eligible for transfer. This query can be used to assess which domains can be moved away from their current registrar, providing insights for strategic domain management.

```sql
select
  domain,
  domain_id,
  created_at,
  transfer_away_eligibile_at,
  renew_auto
from
  godaddy_domain
where
  transfer_away_eligibile_at < now();
```

### List domains that will expire after a particular date
Discover the segments that have domain registrations set to expire after a specific date. This is useful for planning renewals and preventing accidental domain loss.

```sql
select
  domain,
  transfer_protected,
  renew_deadline,
  renew_auto,
  expires
from
  godaddy_domain
where
  expires > '2023-12-31';
```

### List domains created in the last 30 days
Explore which domains have been established in the recent month. This can be useful for monitoring the creation of new domains and ensuring they are properly protected and registered.

```sql
select
  domain,
  domain_id,
  created_at,
  expiration_protected,
  hold_registrar,
  subaccount_id
from
  godaddy_domain
where
  created_at >= now() - interval '30' day;
```

### Get nameservers of each domain
Discover the nameservers associated with each domain to better understand your domain's configuration and manage DNS settings more effectively. This can be particularly useful in identifying potential issues or inconsistencies in your DNS setup.

```sql
select
  domain,
  domain_id,
  n as nameserver
from
  godaddy_domain,
  jsonb_array_elements_text(nameservers) as n;
```

### Get domain admin contact details of each domain
Discover the contact details for the administrators of each domain, including their job title, name, organization, and mailing address. This information is useful for direct communication with the domain admin or for auditing purposes.

```sql
select
  domain,
  domain_contacts_admin ->> 'Fax' as domain_contacts_admin_fax,
  domain_contacts_admin ->> 'Email' as domain_contacts_admin_email,
  domain_contacts_admin ->> 'Phone' as domain_contacts_admin_phone,
  domain_contacts_admin ->> 'JobTitle' as domain_contacts_admin_job_title,
  domain_contacts_admin ->> 'NameLast' as domain_contacts_admin_name_last,
  domain_contacts_admin ->> 'NameFirst' as domain_contacts_admin_name_first,
  domain_contacts_admin ->> 'NameMiddle' as domain_contacts_admin_name_middle,
  domain_contacts_admin ->> 'Organization' as domain_contacts_admin_Organization,
  domain_contacts_admin ->> 'AddressMailing' as domain_contacts_admin_address_mailing
from
  godaddy_domain;
```

### Get verification details of each domain
Explore the verification status of each domain to understand which domains have completed the verification process and which are still pending. This can be useful in maintaining domain authenticity and ensuring all domains are verified as per the requirements.

```sql
select
  domain,
  domain_id,
  created_at,
  verifications -> 'DomainName' ->> 'Status' as verfivication_domain_name,
  verifications -> 'RealName' ->> 'Status' as verfivication_real_name
from
  godaddy_domain;
```