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

// ListRecords retrieves a list of DNS records for the specified domain.
// It sends a request to the Njalla API using the "list-records" method and
// returns the records as a slice of schema.RecordResponse.
//
// Parameters:
//   - ctx: The context for the request, used for cancellation and deadlines.
//   - domain: The domain name for which to retrieve DNS records.
//
// Returns:
//   - []schema.RecordResponse: A slice containing the DNS records for the domain.
//   - error: An error if the request fails or the response cannot be processed.
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

// CreateRecord creates a new DNS record for the specified domain.
// It first checks if a record with the same name already exists for the domain,
// and returns an error if a duplicate is found. If no duplicate exists, it sends
// a request to create the record with the provided parameters.
//
// Parameters:
//   - ctx: The context for the request, used for cancellation and timeouts.
//   - r: The parameters for the record to be created, including domain, type,
//     name, content, TTL, priority, weight, port, target, SSH algorithm, and SSH type.
//
// Returns:
//   - A pointer to a RecordCreateRequestResponse containing the details of the created record.
//   - An error if the record creation fails or if a record with the same name already exists.
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

// UpdateRecord updates an existing DNS record for a given domain.
// It first checks if the record with the specified ID exists in the domain's records.
// If the record exists, it sends a request to update the record with the provided parameters.
//
// Parameters:
//   - ctx: The context for the request, used for cancellation and deadlines.
//   - r: A schema.RecordUpdateParams object containing the details of the record to update.
//
// Returns:
//   - A pointer to schema.RecordUpdateRequestResponse containing the response from the update operation.
//   - An error if the record does not exist or if there is an issue during the update process.
//
// Errors:
//   - Returns an error if the record with the specified ID does not exist.
//   - Returns an error if there is an issue creating or sending the update request.
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

// DeleteRecord deletes a DNS record for a given domain.
// It first checks if the record exists by listing all records for the domain.
// If the record does not exist, it returns an error.
// If the record exists, it sends a request to delete the record.
//
// Parameters:
//   - ctx: The context for the request, used for cancellation and deadlines.
//   - r: The parameters for deleting the record, including the record ID and domain.
//
// Returns:
//   - A pointer to a RecordDeleteRequestResponse containing the response from the server.
//   - An error if the operation fails, such as if the record does not exist or the request fails.
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
