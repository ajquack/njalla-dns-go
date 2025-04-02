package client

import (
	"context"
	"fmt"

	"github.com/ajquack/njalla-dns-go/njalla/schema"
)

type GlueClient struct {
	client *Client
}

// ListGlue retrieves a list of glue records for the specified domain.
// It sends a request to the API with the provided domain name and returns
// a slice of GlueResponse objects or an error if the operation fails.
//
// Parameters:
//   - ctx: The context for the request, used for cancellation and timeouts.
//   - domain: The domain name for which to retrieve glue records.
//
// Returns:
//   - ([]schema.GlueResponse): A slice of GlueResponse objects containing
//     the glue records for the specified domain.
//   - (error): An error if the request fails or the response cannot be processed.
func (c *GlueClient) ListGlue(ctx context.Context, domain string) ([]schema.GlueResponse, error) {
	const method string = "list-glue"
	var responseScheme schema.GlueListRequestResponse

	body := schema.GlueListRequest{
		Method: method,
		Params: schema.GlueListParams{
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
	response := resp.(*schema.GlueListRequestResponse)
	return response.Glue, nil
}

// CreateGlue creates a new glue record for the specified domain.
// It first checks if a glue record with the same name already exists
// for the domain by calling the ListGlue method. If a duplicate is found,
// an error is returned.
//
// Parameters:
//   - ctx: The context for the request, used for cancellation and deadlines.
//   - glueParams: The parameters for the glue record, including the domain
//     and the glue name.
//
// Returns:
//   - A pointer to a GlueCreateRequestResponse containing the response data
//     from the glue creation request.
//   - An error if the glue record already exists or if there is an issue
//     with the request or response processing.
func (c *GlueClient) CreateGlue(ctx context.Context, glueParams schema.GlueParams) (*schema.GlueCreateRequestResponse, error) {
	const method string = "add-glue"
	var responseScheme schema.GlueCreateRequestResponse

	existingGlues, err := c.ListGlue(ctx, glueParams.Domain)
	if err != nil {
		return nil, err
	}
	for _, glue := range existingGlues {
		if glue.Name == glueParams.Name {
			return nil, fmt.Errorf("Glue record %s already exists", glueParams.Name)
		}
	}

	body := schema.GlueCreateRequest{
		Method: method,
		Params: glueParams,
	}
	req, err := c.client.NewRequest(ctx, body)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.DoRequest(req, &responseScheme)
	if err != nil {
		return nil, err
	}
	response := resp.(*schema.GlueCreateRequestResponse)
	return response, nil
}

// UpdateGlue updates an existing glue record for a given domain.
// It first checks if the glue record exists by listing all glue records for the domain.
// If the record does not exist, it returns an error.
// If the record exists, it sends a request to update the glue record with the provided parameters.
//
// Parameters:
//   - ctx: The context for the request, used for cancellation and deadlines.
//   - glueParams: The parameters for the glue record to be updated, including domain and record details.
//
// Returns:
//   - A pointer to a GlueUpdateRequestResponse containing the response from the update operation.
//   - An error if the operation fails or the glue record does not exist.
func (c *GlueClient) UpdateGlue(ctx context.Context, glueParams schema.GlueParams) (*schema.GlueUpdateRequestResponse, error) {
	const method string = "edit-glue"
	var responseScheme schema.GlueUpdateRequestResponse
	var exists bool

	// Check if the record exists
	existingRecords, err := c.ListGlue(ctx, glueParams.Domain)
	if err != nil {
		return nil, err
	}
	for _, record := range existingRecords {
		if record.Name == glueParams.Name {
			exists = true
			break
		}
	}

	if !exists {
		return nil, fmt.Errorf("glue record with ID %s does not exist", glueParams.Name)
	}

	// Create request body
	body := schema.GlueUpdateRequest{
		Method: method,
		Params: glueParams,
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
	response := resp.(*schema.GlueUpdateRequestResponse)
	return response, nil
}

// DeleteGlue deletes a glue record for a given domain and name.
// It first checks if the glue record exists by listing all glue records for the specified domain.
// If the record does not exist, it returns an error.
// If the record exists, it sends a request to remove the glue record.
//
// Parameters:
//   - ctx: The context for the request, used for cancellation and deadlines.
//   - glueParams: The parameters required to identify the glue record to delete, including the domain and name.
//
// Returns:
//   - A pointer to a GlueDeleteRequestResponse containing the response from the server.
//   - An error if the glue record does not exist or if there is an issue with the request.
//
// Errors:
//   - Returns an error if the glue record with the specified name does not exist.
//   - Returns an error if there is an issue creating or sending the request.
func (c *GlueClient) DeleteGlue(ctx context.Context, glueParams schema.GlueDeleteParams) (*schema.GlueDeleteRequestResponse, error) {
	const method string = "remove-glue"
	var responseScheme schema.GlueDeleteRequestResponse
	var exists bool

	// Check if the record exists
	existingRecords, err := c.ListGlue(ctx, glueParams.Domain)
	if err != nil {
		return nil, err
	}
	for _, record := range existingRecords {
		if record.Name == glueParams.Name {
			exists = true
			break
		}
	}

	if !exists {
		return nil, fmt.Errorf("glue record with name %s does not exist", glueParams.Name)
	}

	// Create request body
	body := schema.GlueDeleteRequest{
		Method: method,
		Params: glueParams,
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
	response := resp.(*schema.GlueDeleteRequestResponse)
	return response, nil
}
