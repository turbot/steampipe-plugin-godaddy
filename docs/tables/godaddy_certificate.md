---
title: "Steampipe Table: godaddy_certificate - Query GoDaddy Certificates using SQL"
description: "Allows users to query GoDaddy Certificates, specifically providing information about the certificate's common name, issuer, serial number, and validity period."
---

# Table: godaddy_certificate - Query GoDaddy Certificates using SQL

GoDaddy Certificate is a digital certificate used to secure websites and enable HTTPS. These certificates are issued by GoDaddy, a popular web hosting company and domain registrar. GoDaddy Certificates provide encryption, data integrity, and authentication, ensuring secure communication between a client and the server.

## Table Usage Guide

The `godaddy_certificate` table provides insights into GoDaddy Certificates. As a security analyst, you can leverage this table to gather detailed information about each certificate, including its common name, issuer, serial number, and validity period. This can aid in managing and monitoring the security of your web applications and services hosted on GoDaddy.

**Important Notes**
- You must specify the `customer_id` or `certificate_id` in the `where` clause to query this table.

## Examples

### Basic info
Discover the segments that are related to a specific GoDaddy certificate. This query can be used to gain insights into the status, validity period, and revocation details of a certificate, which is beneficial for managing and tracking your SSL certificates.

```sql+postgres
select
  common_name,
  status,
  certificate_id,
  serial_number,
  created_at,
  organization,
  root_type,
  valid_start,
  valid_end,
  revoked_at
from
  godaddy_certificate
where
  certificate_id ='14a42b6a799d4985b5c5e5e5f5d4249b';
```

```sql+sqlite
select
  common_name,
  status,
  certificate_id,
  serial_number,
  created_at,
  organization,
  root_type,
  valid_start,
  valid_end,
  revoked_at
from
  godaddy_certificate
where
  certificate_id ='14a42b6a799d4985b5c5e5e5f5d4249b';
```

### List certificates by customer ID
List certificate for a particular customer.

```sql+postgres
select
  customer_id,
  common_name,
  status,
  certificate_id,
  created_at,
  valid_start,
  valid_end
from
  godaddy_certificate
where
  certificate_id = '295e3bc3-b3b9-4d95-aae5-ede41a994d13';
```

```sql+sqlite
select
  customer_id,
  common_name,
  status,
  certificate_id,
  created_at,
  valid_start,
  valid_end
from
  godaddy_certificate
where
  certificate_id = '295e3bc3-b3b9-4d95-aae5-ede41a994d13';
```

### List of revoked certificates
Discover the segments that contain revoked certificates to maintain the security and integrity of your network. This can be particularly useful in identifying potential vulnerabilities or malicious activities.

```sql+postgres
select
  common_name,
  status,
  certificate_id,
  serial_number,
  created_at,
  valid_start,
  valid_end
from
  godaddy_certificate
where
  certificate_id  in ('14a42b6a799d4985b5c5e5e5f5d4249b','14a42b6a799d4985b5c5e5e5f5d4249f')
  and status = 'REVOKED';
```

```sql+sqlite
select
  common_name,
  status,
  certificate_id,
  serial_number,
  created_at,
  valid_start,
  valid_end
from
  godaddy_certificate
where
  certificate_id  in ('14a42b6a799d4985b5c5e5e5f5d4249b','14a42b6a799d4985b5c5e5e5f5d4249f')
  and status = 'REVOKED';
```

### Get contact details of a specific certificate
This query is used to gather detailed contact information related to a specific certificate from GoDaddy. This can be particularly useful when you need to reach out to the certificate holder for any updates or issues.

```sql+postgres
select
  common_name,
  contact ->> 'Email' as contact_email,
  contact ->> 'Fax' as contact_fax,
  contact ->> 'NameFirst' as contact_first_name,
  contact ->> 'NameLast' as contact_last_name,
  contact ->> 'Organization' as contact_organization,
  contact ->> 'Phone' as contact_phone
from
  godaddy_certificate
where
  certificate_id = '14a42b6a799d4985b5c5e5e5f5d4249b';
```

```sql+sqlite
select
  common_name,
  json_extract(contact, '$.Email') as contact_email,
  json_extract(contact, '$.Fax') as contact_fax,
  json_extract(contact, '$.NameFirst') as contact_first_name,
  json_extract(contact, '$.NameLast') as contact_last_name,
  json_extract(contact, '$.Organization') as contact_organization,
  json_extract(contact, '$.Phone') as contact_phone
from
  godaddy_certificate
where
  certificate_id = '14a42b6a799d4985b5c5e5e5f5d4249b';
```