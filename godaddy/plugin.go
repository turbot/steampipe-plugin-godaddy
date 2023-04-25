/*
Package aws implements a steampipe plugin for godaddy.
This plugin provides data that Steampipe uses to present foreign
tables that represent GoDaddy resource.
*/
package godaddy

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-godaddy"

// Plugin creates this (aws) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),

		// TODO //
		// DefaultGetConfig: &plugin.GetConfig{
		// 	IgnoreConfig: &plugin.IgnoreConfig{
		// 		ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{
		// 			"NotFoundException",
		// 		}),
		// 	},
		// },

		// TODO //
		// Default ignore config for the plugin
		// DefaultIgnoreConfig: &plugin.IgnoreConfig{
		// 	ShouldIgnoreErrorFunc: shouldIgnoreErrorPluginDefault(),
		// },

		// TODO //
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},

		TableMap: map[string]*plugin.Table{
			"godaddy_domain":     tableGodaddyDomain(ctx),
			"godaddy_dns_record": tableGodaddyDNSRecord(ctx),
			"godaddy_order":      tableGodaddyOrder(ctx),
			"godaddy_subscription": tableGodaddySubscription(ctx),
		},
	}

	return p
}
