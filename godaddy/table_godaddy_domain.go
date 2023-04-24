package godaddy

import (
	"context"

	"github.com/godaddy/godaddysearch-client-go/v3/godaddy/search"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableGoDaddyIndex(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "godaddy_index",
		Description: "Indices in the account and their metadata.",
		List: &plugin.ListConfig{
			Hydrate: listIndex,
		},
		Columns: []*plugin.Column{

			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Index name."},
			{Name: "entries", Type: proto.ColumnType_INT, Description: "Number of records contained in the index."},

			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Index creation date. If empty then the index has no records."},
			{Name: "data_size", Type: proto.ColumnType_INT, Description: "Number of bytes of the index in minified format."},
			{Name: "file_size", Type: proto.ColumnType_INT, Description: "Number of bytes of the index binary file."},
			{Name: "last_build_time_secs", Type: proto.ColumnType_INT, Description: "Last build time in seconds."},
			// Deprecated, do not use - {Name: "number_of_pending_tasks", Type: proto.ColumnType_INT},
			// Deprecated, do not use - {Name: "pending_task", Type: proto.ColumnType_BOOL},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Date of last update. An empty string means that the index has no records."},
			{Name: "primary", Type: proto.ColumnType_STRING, Description: "Only present if the index is a replica. Contains the name of the related primary index."},
			{Name: "replicas", Type: proto.ColumnType_JSON, Description: "Only present if the index is a primary index with replicas. Contains the names of all linked replicas."},
			{Name: "settings", Type: proto.ColumnType_JSON, Hydrate: getIndexSettings, Transform: transform.FromValue(), Description: "Index settings."},
		},
	}
}

func listIndex(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	result, err := conn.ListIndices()
	if err != nil {
		return nil, err
	}
	for _, i := range result.Items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getIndexSettings(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	item := h.Item.(search.IndexRes)
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	index := conn.InitIndex(item.Name)
	settings, err := index.GetSettings()
	if err != nil {
		return nil, err
	}
	return settings, nil
}
