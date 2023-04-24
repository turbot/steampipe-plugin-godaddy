package godaddy

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableGodaddyDomain(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "godaddy_domain",
		Description: "Godaddy Domain",
		// Get: &plugin.GetConfig{
		// 	KeyColumns: plugin.SingleColumn("app_id"),
		// 	// IgnoreConfig: &plugin.IgnoreConfig{
		// 	// 	ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationException", "NotFoundException"}),
		// 	// },
		// 	Hydrate: getDomain,
		// },
		List: &plugin.ListConfig{
			Hydrate: listDomains,
		},
		// GetMatrixItemFunc: SupportedRegionMatrix(amplifyv1.EndpointsID),
		Columns: []*plugin.Column{
			{
				Name:        "domain",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "domain_id",
				Description: "",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("DomainID"),
			},
			{
				Name:        "created_at",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "deleted_at",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "transfer_away_eligibile_at",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "expires",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "expiration_protected",
				Description: "",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "hold_registrar",
				Description: "",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "locked",
				Description: "",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "privacy",
				Description: "",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "renew_auto",
				Description: "",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "renew_deadline",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "renewable",
				Description: "",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "status",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "transfer_protected",
				Description: "",
				Type:        proto.ColumnType_BOOL,
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

	result, err := client.Domains.List(nil, nil, maxLimit, "", nil, "")
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_domain.listDomains", "api_error", err)
		return nil, err
	}

	for _, item := range result {
		d.StreamListItem(ctx, item)

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
