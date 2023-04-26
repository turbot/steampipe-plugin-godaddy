# Table: godaddy_domain

GoDaddy is a popular domain name registrar and web hosting company that allows individuals and businesses to register domain names for their websites. A domain name is the address that people use to access your website on the internet, and GoDaddy offers a wide variety of domain extensions to choose from, such as .com, .org, .net, and many others. In addition to domain registration.

## Examples

### Basic info

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

### List domains that are eligible for transfer away

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
  transfer_away_eligibile_at < NOW();
```

### List domains that expire after a certain date

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

### List domains that are created in the last 30 days

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

### Get nameservers for each domain

```sql
select
  domain,
  domain_id,
  n as nameserver
from
  godaddy_domain,
  jsonb_array_elements_text(nameservers) as n;
```

### Get domain admin contact details for each domain

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
  domain_contacts_admin -> 'AddressMailing' as domain_contacts_admin_address_mailing
from
  godaddy_domain;
```

### Get verification details for each domain

```sql
select
  domain,
  verifications -> 'DomainName' ->> 'Status' as verfivication_domain_name,
  verifications -> 'RealName' ->> 'Status' as verfivication_real_name
from
  godaddy_domain;
```