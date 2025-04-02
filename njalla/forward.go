package client

import (
	"context"
	"fmt"

	"github.com/ajquack/njalla-dns-go/njalla/schema"
)

type ForwardClient struct {
	client *Client
}

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
