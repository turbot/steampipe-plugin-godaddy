package godaddy

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/godaddy/godaddysearch-client-go/v3/godaddy/search"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func connect(_ context.Context, d *plugin.QueryData) (*search.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "godaddy"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*search.Client), nil
	}

	// Default to using env vars
	key := os.Getenv("??")
	secret := os.Getenv("??")

	// But prefer the config
	godaddyConfig := GetConfig(d.Connection)
	if &godaddyConfig != nil {
		if godaddyConfig.AppID != nil {
			appID = *godaddyConfig.AppID
		}
		if godaddyConfig.APIKey != nil {
			apiKey = *godaddyConfig.APIKey
		}
	}

	if appID == "" || apiKey == "" {
		// Credentials not set
		return nil, errors.New("app_id and api_key must be configured")
	}

	conn := search.NewClient(appID, apiKey)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}

func unixTimestampToDateTime(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := int64(d.Value.(int))
	if i == 0 {
		return nil, nil
	}
	ts := time.Unix(i, 0)
	return ts, nil
}

func queryParamsToJSON(_ context.Context, d *transform.TransformData) (interface{}, error) {
	q := d.Value.(string)
	m, err := url.ParseQuery(q)
	if err != nil {
		return nil, err
	}
	js, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return string(js), nil
}

func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "GoDaddy API error [404]")
}
