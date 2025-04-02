package client

import (
	"context"
	"fmt"

	"github.com/ajquack/njalla-dns-go/njalla/schema"
)

type RecordType string

const (
	RecordTypeA       RecordType = "A"
	RecordTypeAAAA    RecordType = "AAAA"
	RecordTypeANAME   RecordType = "ANAME"
	RecordTypeCAA     RecordType = "CAA"
	RecordTypeCNAME   RecordType = "CNAME"
	RecordTypeDynamic RecordType = "DYNAMIC"
	RecordTypeHTTPS   RecordType = "HTTPS"
	RecordTypeMX      RecordType = "MX"
	RecordTypeNAPTR   RecordType = "NAPTR"
	RecordTypeNS      RecordType = "NS"
	RecordTypePTR     RecordType = "PTR"
	RecordTypeSRV     RecordType = "SRV"
	RecordTypeSSHFP   RecordType = "SSHFP"
	RecordTypeSVCB    RecordType = "SVCB"
	RecordTypeTLSA    RecordType = "TLSA"
	RecordTypeTXT     RecordType = "TXT"
)

type RecordClient struct {
	client *Client
}

func (c *RecordClient) ListRecords(ctx context.Context, domain string) ([]schema.RecordResponse, error) {
	const method string = "list-records"
	var responseScheme schema.RecordsListRequestResponse

	body := schema.RecordsListRequest{
		Method: method,
		Params: schema.RecordListParams{
			Domain: domain,
		},
	}
	req, err := c.client.NewRequest(ctx, body)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.DoRequest(req, &responseScheme)
	if err != nil {
		return nil, err
	}
	response := resp.(*schema.RecordsListRequestResponse)
	return response.Records, nil
}

func (c *RecordClient) CreateRecord(ctx context.Context, r schema.RecordCreateParams) (*schema.RecordCreateRequestResponse, error) {
	const method string = "add-record"
	var responseScheme schema.RecordCreateRequestResponse

	// Check if the record already exists
	existingRecords, err := c.ListRecords(context.Background(), r.Domain)
	if err != nil {
		return nil, err
	}
	for _, record := range existingRecords {
		if record.Name == r.Name {
			return nil, fmt.Errorf("record with name %s already exists", r.Name)
		}
	}

	// Create request body
	body := schema.RecordCreateRequest{
		Method: method,
		Params: schema.RecordCreateParams{
			Domain:       r.Domain,
			Type:         r.Type,
			Name:         r.Name,
			Content:      r.Content,
			TTL:          r.TTL,
			Prio:         r.Prio,
			Weight:       r.Weight,
			Port:         r.Port,
			Target:       r.Target,
			SSHAlgorithm: r.SSHAlgorithm,
			SSHType:      r.SSHType,
		},
	}
	// Create request
	req, err := c.client.NewRequest(ctx, body)
	if err != nil {
		return nil, err
	}
	// Send request
	resp, err := c.client.DoRequest(req, &responseScheme)
	if err != nil {
		return nil, err
	}
	// Parse response
	response := resp.(*schema.RecordCreateRequestResponse)
	return response, nil
}

func (c *RecordClient) UpdateRecord(ctx context.Context, r schema.RecordUpdateParams) (*schema.RecordUpdateRequestResponse, error) {
	const method string = "edit-record"
	var responseScheme schema.RecordUpdateRequestResponse
	var exists bool

	// Check if the record exists
	existingRecords, err := c.ListRecords(ctx, r.Domain)
	if err != nil {
		return nil, err
	}
	for _, record := range existingRecords {
		if record.ID == r.ID {
			exists = true
			break
		}
	}

	if !exists {
		return nil, fmt.Errorf("record with ID %s does not exist", r.ID)
	}

	// Create request body
	body := schema.RecordUpdateRequest{
		Method: method,
		Params: schema.RecordUpdateParams{
			ID:           r.ID,
			Domain:       r.Domain,
			Type:         r.Type,
			Name:         r.Name,
			Content:      r.Content,
			TTL:          r.TTL,
			Prio:         r.Prio,
			Weight:       r.Weight,
			Port:         r.Port,
			Target:       r.Target,
			SSHAlgorithm: r.SSHAlgorithm,
			SSHType:      r.SSHType,
		},
	}
	// Create request
	req, err := c.client.NewRequest(ctx, body)
	if err != nil {
		return nil, err
	}
	// Send request
	resp, err := c.client.DoRequest(req, &responseScheme)
	if err != nil {
		return nil, err
	}
	// Parse response
	response := resp.(*schema.RecordUpdateRequestResponse)
	return response, nil
}

func (c *RecordClient) DeleteRecord(ctx context.Context, r schema.RecordDeleteParams) (*schema.RecordDeleteRequestResponse, error) {
	const method string = "remove-record"
	var responseScheme schema.RecordDeleteRequestResponse
	var exists bool

	// Check if the record exists
	existingRecords, err := c.ListRecords(ctx, r.Domain)
	if err != nil {
		return nil, err
	}
	for _, record := range existingRecords {
		if record.ID == r.ID {
			exists = true
			break
		}
	}

	if !exists {
		return nil, fmt.Errorf("record with ID %s does not exist", r.ID)
	}

	// Create request body
	body := schema.RecordDeleteRequest{
		Method: method,
		Params: schema.RecordDeleteParams{
			ID:     r.ID,
			Domain: r.Domain,
		},
	}
	// Create request
	req, err := c.client.NewRequest(ctx, body)
	if err != nil {
		return nil, err
	}
	// Send request
	resp, err := c.client.DoRequest(req, &responseScheme)
	if err != nil {
		return nil, err
	}
	// Parse response
	response := resp.(*schema.RecordDeleteRequestResponse)
	return response, nil
}
