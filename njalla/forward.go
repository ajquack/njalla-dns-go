package client

import (
	"context"
	"fmt"

	"github.com/ajquack/njalla-dns-go/njalla/schema"
)

type ForwardClient struct {
	client *Client
}

// ListForward retrieves a list of forward configurations for a specified domain.
// It sends a request to the API with the provided domain and returns the list
// of forward responses or an error if the operation fails.
//
// Parameters:
//   - ctx: The context for the request, used for cancellation and deadlines.
//   - domain: The domain name for which the forward configurations are to be listed.
//
// Returns:
//   - A slice of schema.ForwardResponse containing the forward configurations.
//   - An error if the request fails or the response cannot be processed.
func (c *ForwardClient) ListForward(ctx context.Context, domain string) ([]schema.ForwardResponse, error) {
	const method string = "list-forwards"
	var responseScheme schema.ForwardListRequestResponse

	body := schema.ForwardListRequest{
		Method: method,
		Params: schema.ForwardListParams{
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
	response := resp.(*schema.ForwardListRequestResponse)
	return response.Forward, nil
}

// CreateForward creates a new forward record for the specified domain.
// It first checks if a forward record with the same "From" and "To" values
// already exists for the domain. If such a record exists, it returns an error.
// Otherwise, it sends a request to create the forward record.
//
// Parameters:
//   - ctx: The context for the request, used for cancellation and deadlines.
//   - forwardParams: The parameters for the forward record, including the
//     domain, "From" address, and "To" address.
//
// Returns:
//   - A pointer to a ForwardCreateRequestResponse containing the details of
//     the created forward record.
//   - An error if the forward record could not be created or if a record with
//     the same "From" and "To" values already exists.
func (c *ForwardClient) CreateForward(ctx context.Context, forwardParams schema.ForwardParams) (*schema.ForwardCreateRequestResponse, error) {
	const method string = "add-forward"
	var responseScheme schema.ForwardCreateRequestResponse

	existingForwards, err := c.ListForward(ctx, forwardParams.Domain)
	if err != nil {
		return nil, err
	}
	for _, forward := range existingForwards {
		if forward.To == forwardParams.To && forward.From == forwardParams.From {
			return nil, fmt.Errorf("Forward record from %s to %s already exists", forwardParams.To, forwardParams.From)
		}
	}

	body := schema.ForwardCreateRequest{
		Method: method,
		Params: forwardParams,
	}
	req, err := c.client.NewRequest(ctx, body)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.DoRequest(req, &responseScheme)
	if err != nil {
		return nil, err
	}
	response := resp.(*schema.ForwardCreateRequestResponse)
	return response, nil
}

// DeleteForward deletes an existing forward record for a given domain.
// It first checks if the forward record exists by comparing the "From" and "To" fields
// in the provided forward parameters with the existing forward records for the domain.
// If the record does not exist, it returns an error.
//
// Parameters:
//   - ctx: The context for the operation, used for cancellation and deadlines.
//   - forwardParams: The parameters specifying the forward record to delete, including
//     the domain, "From" address, and "To" address.
//
// Returns:
//   - A pointer to a ForwardDeleteRequestResponse containing the response from the server.
//   - An error if the forward record does not exist or if there is an issue with the request.
//
// Errors:
//   - Returns an error if the forward record does not exist.
//   - Returns an error if there is an issue creating or executing the request.
func (c *ForwardClient) DeleteForward(ctx context.Context, forwardParams schema.ForwardParams) (*schema.ForwardDeleteRequestResponse, error) {
	const method string = "remove-forward"
	var responseScheme schema.ForwardDeleteRequestResponse
	var exists bool

	// Check if the record exists
	existingForwards, err := c.ListForward(ctx, forwardParams.Domain)
	if err != nil {
		return nil, err
	}
	for _, forward := range existingForwards {
		if (forward.To == forwardParams.To) && (forward.From == forwardParams.From) {
			exists = true
			break
		}
	}

	if !exists {
		return nil, fmt.Errorf("forward record from %s to %s does not exists", forwardParams.To, forwardParams.From)
	}

	body := schema.ForwardDeleteRequest{
		Method: method,
		Params: forwardParams,
	}
	req, err := c.client.NewRequest(ctx, body)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.DoRequest(req, &responseScheme)
	if err != nil {
		return nil, err
	}
	response := resp.(*schema.ForwardDeleteRequestResponse)
	return response, nil
}
