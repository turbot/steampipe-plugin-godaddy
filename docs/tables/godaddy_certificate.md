---
title: "Steampipe Table: godaddy_certificate - Query GoDaddy Certificates using SQL"
description: "Allows users to query GoDaddy Certificates, specifically providing information about the certificate's common name, issuer, serial number, and validity period."
---

# Table: godaddy_certificate - Query GoDaddy Certificates using SQL

GoDaddy Certificate is a digital certificate used to secure websites and enable HTTPS. These certificates are issued by GoDaddy, a popular web hosting company and domain registrar. GoDaddy Certificates provide encryption, data integrity, and authentication, ensuring secure communication between a client and the server.

## Table Usage Guide

The `godaddy_certificate` table provides insights into GoDaddy Certificates. As a security analyst, you can leverage this table to gather detailed information about each certificate, including its common name, issuer, serial number, and validity period. This can aid in managing and monitoring the security of your web applications and services hosted on GoDaddy.

## Examples

### Basic info
Discover the segments that are related to a specific GoDaddy certificate. This query can be used to gain insights into the status, validity period, and revocation details of a certificate, which is beneficial for managing and tracking your SSL certificates.

```sql
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

### List of revoked certificates
Discover the segments that contain revoked certificates to maintain the security and integrity of your network. This can be particularly useful in identifying potential vulnerabilities or malicious activities.

```sql
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

```sql
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