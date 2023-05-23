package godaddy

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableGodaddyShopper(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "godaddy_shopper",
		Description: "Returns information about the GoDaddy shopper.",
		Get: &plugin.GetConfig{
			Hydrate: getShopper,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"Not Found"}),
			},
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "shopper_id",
					Require: plugin.Required,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "shopper_id",
				Description: "The unique identifier of the shopper.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ShopperID"),
			},
			{
				Name:        "customer_id",
				Description: "Unique identifier for the customer record associated with this shopper record. This is an alternate identifier that some systems use to identify an individual shopper record.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("CustomerID"),
			},
			{
				Name:        "email",
				Description: "The email address of the customer.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "external_id",
				Description: "Unique identifier for a shopper provided by the calling application.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ExternalID"),
			},
			{
				Name:        "market_id",
				Description: "The market ID for the shopper.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("MarketID"),
			},
			{
				Name:        "name_first",
				Description: "The first name of the shopper.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name_last",
				Description: "The last name of the shopper.",
				Type:        proto.ColumnType_STRING,
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ShopperID"),
			},
		},
	}
}

type ShopperInfo struct {
	CustomerID string
	ShopperID  string
	Email      string
	ExternalID int
	MarketID   string
	NameFirst  string
	NameLast   string
}

//// HYDRATED FUNCTION

func getShopper(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create Client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_shopper.getShopper", "client_error", err)
		return nil, err
	}

	shopperId := d.EqualsQualString("shopper_id")

	result, err := client.Shoppers.Get(shopperId, []string{"customerId"})
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_shopper.getShopper", "api_error", err)
		return nil, err
	}

	if result != nil {
		d.StreamListItem(ctx, &ShopperInfo{
			CustomerID: result.ShopperID.CustomerID,
			ShopperID:  result.ShopperID.ShopperID,
			Email:      result.ShopperUpdate.Email,
			ExternalID: result.ShopperUpdate.ExternalID,
			MarketID:   result.ShopperUpdate.MarketID,
			NameFirst:  result.ShopperUpdate.NameFirst,
			NameLast:   result.ShopperUpdate.NameLast,
		})
	}

	return nil, nil
}
