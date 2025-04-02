package client

import (
	"context"
	"fmt"

	"github.com/ajquack/njalla-dns-go/njalla/schema"
)

type GlueClient struct {
	client *Client
}

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
