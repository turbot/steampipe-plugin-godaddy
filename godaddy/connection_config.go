package godaddy

import (
	"context"
	"errors"
	"os"

	"github.com/turbot/go-daddy/daddy"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type godaddyConfig struct {
	ApiKey      *string `cty:"api_key"`
	ApiSecret   *string `cty:"api_secret"`
	Environment *string `cty:"environment"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
		Type: schema.TypeString,
	},
	"api_secret": {
		Type: schema.TypeString,
	},
	"environment": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &godaddyConfig{}
}

func GetConfig(connection *plugin.Connection) godaddyConfig {
	if connection == nil || connection.Config == nil {
		return godaddyConfig{}
	}

	config, _ := connection.Config.(godaddyConfig)

	return config
}

func getClient(ctx context.Context, d *plugin.QueryData) (*daddy.Client, error) {
	godaddyConfig := GetConfig(d.Connection)

	apiKey := os.Getenv("GODADDY_API_KEY")
	secretKey := os.Getenv("GODADDY_API_SECRET")
	envType := os.Getenv("GODADDY_ENVIRONMENT")

	if godaddyConfig.ApiKey != nil {
		apiKey = *godaddyConfig.ApiKey
	}
	if godaddyConfig.ApiSecret != nil {
		secretKey = *godaddyConfig.ApiSecret
	}
	if godaddyConfig.Environment != nil {
		envType = *godaddyConfig.Environment
	}

	if apiKey == "" || secretKey == "" {
		return nil, errors.New("'api_key' and 'api_secret' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
	}

	// Based on the environment we need to pass the endpoint URL
	switch envType {
	case "DEV":
		return daddy.NewClient(apiKey, secretKey, true)
	case "PROD", "": // Default to production if no environment type is set.
		return daddy.NewClient(apiKey, secretKey, false)
	default:
		return nil, errors.New("'environment' must be set in the connection configuration to 'DEV' or 'PROD'. Edit your connection configuration file and then restart Steampipe.")
	}
}
