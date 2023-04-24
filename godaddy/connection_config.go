package godaddy

import (
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/schema"
)

type godaddyConfig struct {
	AppID  *string `cty:"app_id"`
	APIKey *string `cty:"api_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"app_id": {
		Type: schema.TypeString,
	},
	"api_key": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &godaddyConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) godaddyConfig {
	if connection == nil || connection.Config == nil {
		return godaddyConfig{}
	}
	config, _ := connection.Config.(godaddyConfig)
	return config
}
