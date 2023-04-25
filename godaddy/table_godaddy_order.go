package godaddy

import (
	"context"

	"github.com/turbot/go-daddy/daddy"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableGodaddyOrder(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "godaddy_order",
		Description: "Godaddy Order",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("order_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"not found"}),
			},
			Hydrate: getOrder,
		},
		List: &plugin.ListConfig{
			Hydrate:       listOrders,
			ParentHydrate: listDomains,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "domain_name",
					Require: plugin.Optional,
				},
				{
					Name:    "parent_order_id",
					Require: plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "order_id",
				Description: "The unique identifier of the order.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("OrderID"),
			},
			{
				Name:        "bill_to_tax_id",
				Description: "The billing tax id that was used at the time of purchase.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("BillTo.TaxID"),
			},
			{
				Name:        "currency",
				Description: "Currency in which the order was placed.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "The date and time when the current order was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "domain_name",
				Description: "The name of the domain for the order.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "parent_order_id",
				Description: "Unique identifier of the parent order. All refund/chargeback orders are tied to the original order. The orginal order's orderId is the parentOrderId of refund/chargeback orders.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "bill_to_contact",
				Description: "The billing contact information that was used at the time of purchase.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("BillTo.Contact"),
			},
			{
				Name:        "bill_item",
				Description: "The sets of two or more line items in the order.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "bill_payment",
				Description: "The payment associated with the order.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "bill_pricing",
				Description: "The pricing associated with the order.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

type OrderInfo struct {
	DomainName    string
	BillTo        daddy.BillTo
	CreatedAt     string
	Currency      string
	Items         []daddy.LineItem
	OrderID       string
	ParentOrderID string
	Payments      []daddy.OrderPayment
	Pricing       daddy.OrderPricing
}

//// LIST FUNCTION

func listOrders(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	domain := h.Item.(*daddy.DomainDetail)

	// Reduce the number of API call, if the domain name has specified in the where clause.
	domainName := d.EqualsQualString("domain_name")
	if domainName != "" && domainName != domain.Domain {
		return nil, nil
	}

	parentOrderId := d.EqualsQualString("parent_order_id")

	// Create Client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_order.listOrders", "client_error", err)
		return nil, err
	}

	// Reduce the basic request limit down if the user has only requested a small number of rows
	maxLimit := 1000
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < int32(maxLimit) {
			maxLimit = int(limit)
		}
	}

	result, err := client.Orders.List("", "", domain.Domain, 0, 0, parentOrderId, 0, maxLimit, "")
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_order.listOrders", "api_error", err)
		return nil, err
	}

	for _, item := range result.Orders {
		d.StreamListItem(ctx, &OrderInfo{
			BillTo:        item.BillTo,
			CreatedAt:     item.CreatedAt,
			Currency:      item.Currency,
			DomainName:    domain.Domain,
			Items:         item.Items,
			OrderID:       item.OrderID,
			ParentOrderID: item.ParentOrderID,
			Payments:      item.Payments,
			Pricing:       item.Pricing,
		})

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATED FUNCTION

func getOrder(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var orderId string
	if h.Item != nil {
		orderId = h.Item.(*OrderInfo).OrderID
	} else {
		orderId = d.EqualsQualString("order_id")
	}

	if orderId == "" {
		return nil, nil
	}

	// Create Client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_domain.getOrder", "client_error", err)
		return nil, err
	}

	result, err := client.Orders.Get(orderId)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_domain.getOrder", "api_error", err)
		return nil, err
	}

	return result, nil
}
