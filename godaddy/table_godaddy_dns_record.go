package godaddy

import (
	"context"
	"fmt"

	"github.com/turbot/go-daddy/daddy"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableGodaddyDNSRecord(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "godaddy_dns_record",
		Description: "Returns information about the DNS record.",
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
				Description: "The name of the domain.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The name of the record.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "type",
				Description: "The record type. Valid values are A, AAAA, CAA, CNAME, MX, NAPTR, NS, PTR, SOA, SPF, SRV and TXT.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "data",
				Description: "Refers to the value of the DNS record being added or updated. The value of this field will depend on the type of DNS record being added or updated.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "port",
				Description: "The port for the dns record.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "priority",
				Description: "The priority of the record. The valid values are MX and SRV only.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "protocol",
				Description: "The protocol being used for a particular type of DNS record.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "service",
				Description: "The type of the service.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "ttl",
				Description: "The resource record cache time to live (TTL), in seconds.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("TTL"),
			},
			{
				Name:        "weight",
				Description: "Weighted resource record sets only: Among resource record sets that have the same combination of DNS name and type, a value that determines the proportion of DNS queries that Amazon Route 53 responds to using the current resource record set. Route 53 calculates the sum of the weights for the resource record sets that have the same combination of DNS name and type. Route 53 then responds to queries based on the ratio of a resource's weight to the total.",
				Type:        proto.ColumnType_INT,
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

// The response data should include the domain name as it is the parent of the DNS record. However, the API response does not provide the domain name. Therefore, we need to create a custom struct to accommodate it.
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

	domain := h.Item.(*daddy.DomainDetail)

	// Create Client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("godaddy_dns_record.listDNSRecords", "client_error", err)
		return nil, err
	}

	// Reduce the basic request limit down if the user has only requested a small number of rows.
	// The API throws error (Page size may not exceed 500 records) if we are providing the max limit value grater than 500.
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

	// To avoid errors, both the input parameters "name" and "DNS type" must be provided simultaneously in the where clause. Passing only DNS name will result in an error being thrown by the API.
	if name != "" && dnsType == "" {
		return nil, fmt.Errorf("Name and DNS type must be pass in where clause alltogether if we are passing DNS record name")
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
