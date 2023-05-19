# Table: godaddy_certificate

GoDaddy is a popular Certificate Authority (CA) that issues digital certificates for websites and other online services. These certificates provide a secure communication channel between users and websites by encrypting the data being transmitted. GoDaddy issues various types of certificates, including Domain Validated (DV), Organization Validated (OV), and Extended Validation (EV) certificates, depending on the level of validation required by the user. These certificates have varying levels of trust and security features, and they can be used for a range of purposes, including securing e-commerce transactions, protecting sensitive information, and establishing a trustworthy online presence. Overall, GoDaddy certificates play a critical role in online security, helping to safeguard sensitive data and build trust between businesses and their customers.

Note: The used [Go API](https://pkg.go.dev/github.com/alyx/go-daddy/daddy#CertificatesService.Get) limitation, you must provide valid `certificate_id` in the `where` clause of the queries to fetch certificate details.

## Examples

### Basic info

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
where certificate_id ='14a42b6a799d4985b5c5e5e5f5d4249b';
```

### List of revoked certificates

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
  certificate_id  in ('14a42b6a799d4985b5c5e5e5f5d4249b','14a42b6a799d4985b5c5e5e5f5d4249f') and status = 'REVOKED';
```

### Get contact details of a specific certificate

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
