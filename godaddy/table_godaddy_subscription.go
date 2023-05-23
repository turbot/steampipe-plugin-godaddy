package godaddy

import (
	"context"

	"github.com/turbot/go-daddy/daddy"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableGodaddySubscription(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "godaddy_subscription",
		Description: "Returns information about the GoDaddy product subscription.",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("subscription_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NOT_FOUND", "Invalid Subscription Id"}),
			},
			Hydrate: getSubscription,
		},
		List: &plugin.ListConfig{
			Hydrate: listSubscriptions,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "product_group_key",
					Require: plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "subscription_id",
				Description: "Unique identifier of the subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SubscriptionID"),
			},
			{
				Name:        "created_at",
				Description: "The date and time when the subscription was created.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cancelable",
				Description: "Whether or not the Subscription is allowed to be canceled.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "label",
				Description: "A human readable description of the subscription.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "launch_url",
				Description: "The url to use or manage this subscription's active product.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("LaunchURL"),
			},
			{
				Name:        "payment_profile_id",
				Description: "The unique identifier of the payment profile that will be used to automatically renew this subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("PaymentProfileID"),
			},
			{
				Name:        "price_locked",
				Description: "Whether the renewal price will be based from the list price or a locked-in price for this shopper.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "expires_at",
				Description: "The date when the registration for the domain is set to expire. The date and time is in Unix time format and coordinated universal time (UTC).",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "renew_auto",
				Description: "Whether or not the subscription is set to be automatically renewed via the billing agent.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "renewable",
				Description: "Specifies whether subscription is allowed to be renewed.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "product_renewal_period",
				Description: "The number of renewal period units that will be added for each subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Product.RenewalPeriod"),
			},
			{
				Name:        "product_renewal_period_unit",
				Description: "Specifies whether subscription is allowed to be renewed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Product.RenewalPeriodUnit"),
			},
			{
				Name:        "status",
				Description: "Specifies whether the subscription is active or the specific non-active state.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "product_group_key",
				Description: "Primary key of a grouping of related subscriptions.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Product.ProductGroupKey"),
			},
			{
				Name:        "upgradable",
				Description: "Specifies whether or not the subscription is allowed to be upgraded.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "addons",
				Description: "An array of additional products that have been purchased to augment the subscription.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "billing",
				Description: "The billing details of the subscription.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "product",
				Description: "The product details associated to the subscription.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "relations",
				Description: "The relationship between the subscriptions that will be canceled automatically, when current subscription is canceled.",
				Type:        proto.ColumnType_JSON,
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SubscriptionID"),
			},
		},
	}
}

//// LIST FUNCTION

func listSubscriptions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	productGroupKey := d.EqualsQualString("product_group_key")

	// Create Client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_subscription.listSubscriptions", "client_error", err)
		return nil, err
	}

	// Reduce the basic request limit down if the user has only requested a small number of rows
	// As per the testing we have added it as 1000, the API has been tested with the max limit value grater than 1000.
	// The default value for the max limit is 25, doc: https://developer.godaddy.com/doc/endpoint/subscriptions#/v1/list.
	maxLimit := 1000
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < int32(maxLimit) {
			maxLimit = int(limit)
		}
	}

	offset := 0

	result, err := client.Subscriptions.List([]string{productGroupKey}, []string{}, offset, maxLimit, "")
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_subscription.listSubscriptions", "api_error", err)
		return nil, err
	}

	for _, item := range result.Subscriptions {
		d.StreamListItem(ctx, item)

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	for result.Pagination.Next != "" {
		offset += maxLimit
		res, err := client.Subscriptions.List([]string{productGroupKey}, []string{}, offset, maxLimit, "")
		if err != nil {
			plugin.Logger(ctx).Error("godaddy_subscription.listSubscriptions", "api_pagging_error", err)
			return res, err
		}

		for _, item := range res.Subscriptions {
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

func getSubscription(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var subscription string
	if h.Item != nil {
		subscription = h.Item.(daddy.Subscription).SubscriptionID
	} else {
		subscription = d.EqualsQualString("subscription_id")
	}

	// Empty Check
	if subscription == "" {
		return nil, nil
	}

	// Create Client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_subscription.getSubscription", "client_error", err)
		return nil, err
	}

	result, err := client.Subscriptions.Get(subscription)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_subscription.getSubscription", "api_error", err)
		return nil, err
	}

	return result, nil
}
