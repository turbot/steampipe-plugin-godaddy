---
title: "Steampipe Table: godaddy_dns_record - Query GoDaddy DNS Records using SQL"
description: "Allows users to query GoDaddy DNS Records, providing a complete overview of all the DNS records associated with the user's domain."
---

# Table: godaddy_dns_record - Query GoDaddy DNS Records using SQL

GoDaddy DNS Records are a critical part of the GoDaddy domain hosting service. They allow users to control the domain's email settings, website location, and other important aspects of their internet presence. By managing DNS records, users can direct web traffic to their desired location, configure email routing, and ensure secure and reliable operation of their domain.

## Table Usage Guide

The `godaddy_dns_record` table provides insights into DNS records within GoDaddy's domain hosting service. As a network administrator or website owner, explore specific details through this table, including record type, name, and data. Utilize it to manage and monitor your DNS records, ensuring your domain's web traffic and email routing are correctly configured and operating as intended.

**Important Notes**
- To filter the resource using `name`, you must set `type` in the where clause.

## Examples

### Basic info
Explore which domain names are associated with your GoDaddy DNS records. This query can help you gain insights into your DNS configuration, including the protocol and type of each record, which could be beneficial for managing and troubleshooting your network.

```sql+postgres
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

```sql+sqlite
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
Explore all DNS records associated with a specific domain name. This can be useful for understanding and managing the various services and protocols tied to your domain.

```sql+postgres
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

```sql+sqlite
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
Determine the specific DNS records associated with a particular domain name. This query can be used to gain insights into the configuration and settings of a domain, which can be beneficial for troubleshooting or optimization purposes.

```sql+postgres
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

```sql+sqlite
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
Explore DNS records that have a time-to-live (TTL) value of 3600 seconds or less. This can be useful for identifying domains that may require frequent updates or are more susceptible to potential caching issues.

```sql+postgres
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

```sql+sqlite
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
  ttl <= 3600;
```

### List all DNS records with a specific data value
Explore DNS records that are associated with a specific IP address. This is useful for identifying potential anomalies or inconsistencies in your DNS configuration.

```sql+postgres
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

```sql+sqlite
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
Determine the areas in which DNS records have been assigned a specific priority value. This is useful for assessing the configuration of your domain names and ensuring that the most important records have the correct priority level.

```sql+postgres
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

```sql+sqlite
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