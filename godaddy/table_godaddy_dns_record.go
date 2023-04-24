package godaddy

import (
	"context"

	"github.com/alyx/go-daddy/daddy"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableGodaddyDNSRecord(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "godaddy_dns_record",
		Description: "Godaddy DNS Record",
		List: &plugin.ListConfig{
			ParentHydrate: listDomains,
			Hydrate:       listDNSRecords,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "domain_name",
					Require: plugin.Optional,
				},
				{
					Name:    "name",
					Require: plugin.Optional,
				},
				{
					Name:    "type",
					Require: plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "domain_name",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "type",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "data",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "port",
				Description: "",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "priority",
				Description: "",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "protocol",
				Description: "",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "service",
				Description: "",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "ttl",
				Description: "",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("TTL"),
			},
			{
				Name:        "weight",
				Description: "",
				Type:        proto.ColumnType_INT,
			},
		},
	}
}

type DNSRecordInfo struct {
	DomainName string
	Data       string
	Name       string
	Port       int
	Priority   int
	Protocol   int
	Service    int
	TTL        int
	Type       string
	Weight     int
}

//// LIST FUNCTION

func listDNSRecords(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	domain := h.Item.(daddy.DomainSummary)

	// Create Client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_dns_record.listDNSRecords", "client_error", err)
		return nil, err
	}

	// Reduce the basic request limit down if the user has only requested a small number of rows
	maxLimit := 500
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < int32(maxLimit) {
			maxLimit = int(limit)
		}
	}

	// Reduce the API calls for a given domain name in where clause.
	domainName := domain.Domain
	if d.EqualsQualString("domain_name") != "" && d.EqualsQualString("domain_name") != domainName {
		return nil, nil
	}

	dnsType := ""
	if d.EqualsQualString("type") != "" {
		dnsType = d.EqualsQualString("type")
	}

	name := ""
	if d.EqualsQualString("name") != "" {
		name = d.EqualsQualString("name")
	}

	if name == "" || dnsType == "" {
		name, dnsType = "", ""
	}

	result, err := client.Domains.GetRecords(domainName, dnsType, name, 0, maxLimit)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_dns_record.listDNSRecords", "api_error", err)
		return nil, err
	}

	for _, item := range result {
		d.StreamListItem(ctx, &DNSRecordInfo{
			DomainName: domainName,
			Data:       item.Data,
			Name:       item.Name,
			Port:       item.Port,
			Priority:   item.Priority,
			Protocol:   item.Protocol,
			Service:    item.Service,
			TTL:        item.TTL,
			Type:       item.Type,
			Weight:     item.Weight,
		})

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTION

func getDNSRecord(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// Create Client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_dns_record.listDNSRecords", "client_error", err)
		return nil, err
	}

	// Reduce the API calls for a given domain name in where clause.
	domainName := d.EqualsQualString("domain_name")
	name := d.EqualsQualString("name")
	recordType := d.EqualsQualString("type")

	// Empty check
	if domainName == "" || name == "" {
		return nil, nil
	}

	result, err := client.Domains.GetRecords(domainName, recordType, name, 0, 2)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_dns_record.listDNSRecords", "api_error", err)
		return nil, err
	}

	if len(result) > 0 {
		return &DNSRecordInfo{
			DomainName: domainName,
			Data:       result[0].Data,
			Name:       result[0].Name,
			Port:       result[0].Port,
			Priority:   result[0].Priority,
			Protocol:   result[0].Protocol,
			Service:    result[0].Service,
			TTL:        result[0].TTL,
			Type:       result[0].Type,
			Weight:     result[0].Weight,
		}, nil

	}

	return nil, nil
}
