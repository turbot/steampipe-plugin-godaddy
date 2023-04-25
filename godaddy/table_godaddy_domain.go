package godaddy

import (
	"context"

	"github.com/turbot/go-daddy/daddy"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableGodaddyDomain(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "godaddy_domain",
		Description: "Godaddy Domain",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("domain"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"not found"}),
			},
			Hydrate: getDomain,
		},
		List: &plugin.ListConfig{
			Hydrate: listDomains,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "status",
					Require: plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "domain",
				Description: "The name of the domain.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "domain_id",
				Description: "The ID of the domain.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("DomainID"),
			},
			{
				Name:        "created_at",
				Description: "The date and time when the domain was created as found in the response to a WHOIS query.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "deleted_at",
				Description: "Date and time when this domain was deleted.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("DeletedAt").Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "transfer_away_eligibile_at",
				Description: "Date and time when this domain is eligible to transfer.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("TransferAwayEligibileAt").Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "expires",
				Description: "The date when the registration for the domain is set to expire. The date and time is in Unix time format and Coordinated Universal time (UTC)",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Expires").Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "auth_code",
				Description: "Authorization code for transferring the Domain.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "expiration_protected",
				Description: "Specifies whether the domain is protected from expiration.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "hold_registrar",
				Description: "Specifies whether the domain is on-hold by the registrar.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "locked",
				Description: "Indicates whether a domain is locked from unauthorized transfer to another party.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "privacy",
				Description: "Specifies whether the domain has privacy protection.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "subaccount_id",
				Description: "The sub account ID of the domain.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getDomain,
				Transform:   transform.FromField("SubaccountID"),
			},
			{
				Name:        "renew_auto",
				Description: "Indicates whether the domain is automatically renewed upon expiration.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "renew_deadline",
				Description: "Specifies whether renew deadline is set on the domain.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("RenewDeadline").Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "renewable",
				Description: "Specifies whether the domain is eligible for renewal based on status.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "status",
				Description: "The processing status of the domain.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "transfer_protected",
				Description: "Specifies whether the domain is protected from transfer.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "domain_contacts_admin",
				Description: "The domain admin contact details.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("DomainContacts.Admin"),
			},
			{
				Name:        "domain_contacts_billing",
				Description: "The billing contact details.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("DomainContacts.Billing"),
			},
			{
				Name:        "domain_contacts_registrant",
				Description: "Provides details about the domain registrant.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("DomainContacts.Registrant"),
			},
			{
				Name:        "domain_contacts_tech",
				Description: "Provides details about the domain technical contact.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("DomainContacts.Tech"),
			},
			{
				Name:        "nameservers",
				Description: "Fully-qualified domain names for DNS servers.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "verifications",
				Description: "Fully-qualified domain names for DNS servers.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getDomain,
			},
		},
	}
}

//// LIST FUNCTION

func listDomains(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create Client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_domain.listDomains", "client_error", err)
		return nil, err
	}

	// Reduce the basic request limit down if the user has only requested a small number of rows
	maxLimit := 50
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < int32(maxLimit) {
			maxLimit = int(limit)
		}
	}

	status := d.EqualsQualString("status")

	result, err := client.Domains.List([]string{status}, nil, maxLimit, "", []string{"authCode", "contacts", "nameServers"}, "")
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_domain.listDomains", "api_error", err)
		return nil, err
	}

	for _, item := range result {
		d.StreamListItem(ctx, &daddy.DomainDetail{
			DomainSummary: item,
		})

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATED FUNCTION

func getDomain(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var domain string
	if h.Item != nil {
		domain = h.Item.(*daddy.DomainDetail).Domain
	} else {
		domain = d.EqualsQualString("domain")
	}

	if domain == "" {
		return nil, nil
	}

	// Create Client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_domain.getDomain", "client_error", err)
		return nil, err
	}

	result, err := client.Domains.Get(domain)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_domain.getDomain", "api_error", err)
		return nil, err
	}

	return result, nil
}
