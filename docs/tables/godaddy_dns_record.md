# Table: godaddy_dns_record

GoDaddy is a popular web hosting and domain registrar company that offers various subscription plans for individuals and businesses. A DNS record in GoDaddy specifies the mapping between a domain name and an IP address or other information. DNS records are used to translate human-friendly domain names, such as example.com, into machine-readable IP addresses, such as 192.0.2.1, that are used to route traffic over the internet. Different types of DNS records serve different purposes, such as mapping a domain name to an IP address (A record), specifying a mail server for the domain (MX record), or defining aliases for the domain (CNAME record).

**Note**: To filter the resource using `name`; you must set `type` in the where clause.

## Examples

### Basic info

```sql
select
  name,
  domain_name,
  data,
  protocol,
  type,
  ttl
from
  godaddy_dns_record;
```

### List all DNS records for a specific domain name

```sql
select
  name,
  data,
  protocol,
  type,
  service,
  ttl
from
  godaddy_dns_record
where
  domain_name = 'example.com';
```

### List all DNS records of a specific type for a specific domain name

```sql
select
  name,
  data,
  protocol,
  type,
  service,
  ttl
from
  godaddy_dns_record
where
  domain_name = 'example.com'
  and type = 'A';
```

### List all DNS records with a TTL less than or equal to a certain value

```sql
select
  name,
  domain_name,
  data,
  protocol,
  type,
  service,
  ttl
from
  godaddy_dns_record
where
  ttl <= '3600';
```

### List all DNS records with a specific data value

```sql
select
  name,
  domain_name,
  data,
  protocol,
  type,
  service,
  ttl
from
  godaddy_dns_record
where
  data = '192.168.1.1';
```

### List all DNS records with a specific priority value

```sql
select
  name,
  domain_name,
  data,
  protocol,
  type,
  service,
  ttl
from
  godaddy_dns_record
where
  priority = '10';
```
