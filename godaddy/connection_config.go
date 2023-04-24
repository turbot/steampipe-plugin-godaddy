package godaddy

import (
	"context"
	"errors"
	"os"

	"github.com/alyx/go-daddy/daddy"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type godaddyConfig struct {
	API_KEY         *string `cty:"api_key"`
	SECRET_KEY      *string `cty:"secret_key"`
	EVIRONMENT_TYPE *string `cty:"environment_type"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
		Type: schema.TypeString,
	},
	"secret_key": {
		Type: schema.TypeString,
	},
	"environment_type": {
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

	envType := os.Getenv("GODADDY_ENVIRONMENT_TYPE")
	apiKey := os.Getenv("GODADDY_API_KEY")
	secretKey := os.Getenv("GODADDY_SECRET_KEY")

	if godaddyConfig.API_KEY != nil {
		apiKey = *godaddyConfig.API_KEY
	}
	if godaddyConfig.SECRET_KEY != nil {
		secretKey = *godaddyConfig.SECRET_KEY
	}
	if godaddyConfig.EVIRONMENT_TYPE != nil {
		envType = *godaddyConfig.EVIRONMENT_TYPE
	}

	// Set the default environment type to developement
	if envType == "" {
		envType = "DEV"
	}

	// Based on the environment type we need to pass the
	switch envType {
	case "DEV":
		return daddy.NewClient(apiKey, secretKey, true)
	case "PROD":
		return daddy.NewClient(apiKey, secretKey, false)
	}

	return nil, errors.New("'api_key' and 'secret_key' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
}
