package godaddy

import (
	"context"

	"github.com/turbot/go-daddy/daddy"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/query_cache"
)

//// TABLE DEFINITION

func tableGodaddyCertificate(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "godaddy_certificate",
		Description: "Returns information about the GoDaddy certificate by customer.",
		List: &plugin.ListConfig{
			Hydrate: getCertificate,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:       "customer_id",
					Require:    plugin.Required,
					CacheMatch: query_cache.CacheMatchExact,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate: getCertificate,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"not found"}),
			},
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "certificate_id",
					Require: plugin.Required,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "certificate_id",
				Description: "The unique identifier of the certificate request. Only present if no errors returned.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "customer_id",
				Description: "The unique identifier for a customer",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("customer_id"),
			},
			{
				Name:        "common_name",
				Description: "The common name of certificate.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "The date the certificate was ordered.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("CreatedAt").Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "denied_reason",
				Description: "Reason specified if certificate order has been denied.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getCertificate,
			},
			{
				Name:        "period",
				Description: "Validity period of order. Specified in years.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "product_type",
				Description: "Certificate product type.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getCertificate,
			},
			{
				Name:        "progress",
				Description: "Percentage of completion for certificate vetting.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "revoked_at",
				Description: "The revocation date of certificate (if revoked).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("RevokedAt").Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "root_type",
				Description: "Root type of the certificate.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getCertificate,
			},
			{
				Name:        "serial_number",
				Description: "The unique number of certificate (if issued or revoked).",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "serial_number_hex",
				Description: "The hexadecmial format of serial number of the certificate(if issued or revoked).",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getCertificate,
			},
			{
				Name:        "slot_size",
				Description: "Number of subject alternative names(SAN) to be included in certificate.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "Status of certificate. Valid values are PENDING_ISSUANCE, ISSUED, REVOKED, CANCELED, DENIED, PENDING_REVOCATION, PENDING_REKEY, UNUSED, and EXPIRED.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "valid_end",
				Description: "The end date of the certificate's validity (if issued or revoked).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("ValidEnd", "ValidEndAt").Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "valid_start",
				Description: "The start date of the certificate's validity (if issued or revoked).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("ValidStart", "ValidStartAt").Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "contact",
				Description: "Contact details of the certificate.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getCertificate,
			},
			{
				Name:        "organization",
				Description: "Details of the organization that owns the certificate.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getCertificate,
			},
			{
				Name:        "subject_alternative_names",
				Description: "Specifies subject alternative names set for the certificate.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getCertificate,
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("CommonName"),
			},
		},
	}
}

//// LIST FUNCTION

func listCertificates(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// Create Client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_certificate.listCertificates", "client_error", err)
		return nil, err
	}

	customerId := d.EqualsQualString("customer_id")

	// Empty check
	if customerId == "" {
		return nil, nil
	}

	maxLimit := 100
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < int32(maxLimit) {
			maxLimit = int(limit)
		}
	}
	offset := 0

	result, err := client.Certificates.ListByCustomer(customerId, maxLimit, offset)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_certificate.listCertificates", "api_error", err.Error())
		return nil, err
	}

	for _, item := range result.Certificates {
		d.StreamListItem(ctx, item)

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	for result.Pagination.Next != "" {
		offset += maxLimit
		res, err := client.Certificates.ListByCustomer(customerId, maxLimit, offset)

		if err != nil {
			plugin.Logger(ctx).Error("godaddy_certificate.listCertificates", "api_paging_error", err.Error())
			return nil, err
		}
		for _, item := range res.Certificates {
			d.StreamListItem(ctx, item)

			// Context may get cancelled due to manual cancellation or if the limit has been reached
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		result.Pagination = res.Pagination
	}

	return nil, nil
}

//// HYDRATED FUNCTION

func getCertificate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var certId string
	if h.Item != nil {
		data := h.Item.(daddy.CertificateListResponse)
		certId = data.CertificateID
	} else {
		certId = d.EqualsQualString("certificate_id")
	}

	// Create Client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_certificate.getCertificate", "client_error", err)
		return nil, err
	}

	// Empty check
	if certId == "" {
		return nil, nil
	}

	result, err := client.Certificates.Get(certId)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_certificate.getCertificate", "api_error", err.Error())
		return nil, err
	}

	d.StreamListItem(ctx, result)

	return nil, nil
}
