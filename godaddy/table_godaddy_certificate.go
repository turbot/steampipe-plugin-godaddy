package godaddy

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableGodaddyCertificate(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "godaddy_certificate",
		Description: "Godaddy Certificate",
		List: &plugin.ListConfig{
			Hydrate: listCertificates,
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
				Name:        "common_name",
				Description: "The common name of certificate.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "The date the certificate was ordered.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "denied_reason",
				Description: "Reason specified if certificate order has been denied.",
				Type:        proto.ColumnType_STRING,
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
			},
			{
				Name:        "progress",
				Description: "Percentage of completion for certificate vetting.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "revoked_at",
				Description: "The revocation date of certificate (if revoked).",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "root_type",
				Description: "Root type of the certificate.",
				Type:        proto.ColumnType_STRING,
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
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ValidEnd").Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "valid_start",
				Description: "The start date of the certificate's validity (if issued or revoked).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("ValidStart").Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "organization",
				Description: "Details of the organization that owns the certificate.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "subject_alternative_names",
				Description: "Specifies subject alternative names set for the certificate.",
				Type:        proto.ColumnType_JSON,
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

	certId := d.EqualsQualString("certificate_id")

	// Empty check
	if certId == "" {
		return nil, nil
	}

	result, _ := client.Certificates.Get(certId)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_certificate.listCertificates", "api_error", err.Error())
		return nil, err
	}

	d.StreamListItem(ctx, result)

	return nil, nil
}
