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
		Description: "Godaddy Shopper",
		List: &plugin.ListConfig{
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
				Description: "The unique ID of the shopper.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ShopperID"),
			},
			{
				Name:        "customer_id",
				Description: "Identifier for the Customer record associated with this Shopper record. This is an alternate identifier that some systems use to identify an individual shopper record.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("CustomerID"),
			},
			{
				Name:        "email",
				Description: "The mail id of the customer.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "external_id",
				Description: "The external ID of the shopper.",
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
				Description: "The first name of the customer.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name_last",
				Description: "The last name of the shopper.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

//// LIST FUNCTION

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

	d.StreamListItem(ctx, result)

	return nil, nil
}
