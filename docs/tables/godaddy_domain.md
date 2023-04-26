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