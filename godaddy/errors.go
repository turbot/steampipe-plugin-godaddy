package godaddy

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// shouldIgnoreErrors:: function which returns an ErrorPredicate for GoDaddy API calls
func shouldIgnoreErrors(notFoundErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {

		// Added to support regex in not found errors
		for _, pattern := range notFoundErrors {
			if ok := strings.Contains(err.Error(), pattern); ok {
				return true
			}
		}

		return false
	}
}
