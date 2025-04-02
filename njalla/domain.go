package client

import (
	"context"
	"fmt"

	"github.com/ajquack/njalla-dns-go/njalla/schema"
)

type DomainClient struct {
	client *Client
}

func (c *DomainClient) GetDomain(ctx context.Context, domainParams schema.GetDomainParams) (*schema.GetDomainRequestResponse, error) {
	const method string = "get-domain"
	var responseScheme schema.GetDomainRequestResponse

	body := schema.GetDomainRequest{
		Method: method,
		Params: domainParams,
	}
	req, err := c.client.NewRequest(ctx, body)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.DoRequest(req, &responseScheme)
	if err != nil {
		return nil, err
	}
	response := resp.(*schema.GetDomainRequestResponse)
	return response, nil
}

func (c *DomainClient) ListDomains(ctx context.Context) ([]schema.ListDomainResponse, error) {
	const method string = "list-domains"
	var responseScheme schema.ListDomainsRequestResponse

	body := schema.ListDomainsRequest{
		Method: method,
	}
	req, err := c.client.NewRequest(ctx, body)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.DoRequest(req, &responseScheme)
	if err != nil {
		return nil, err
	}
	response := resp.(*schema.ListDomainsRequestResponse)
	return response.Domains, nil
}

func (c *DomainClient) EditDomain(ctx context.Context, domain string, mailForwarding bool, dnssec bool, lock bool) (*schema.UpdateDomainRequestResponse, error) {
	const method string = "edit-domain"
	var responseScheme schema.UpdateDomainRequestResponse
	var exists bool

	// Check if the domain exists
	existingDomains, err := c.ListDomains(ctx)
	if err != nil {
		return nil, err
	}
	for _, d := range existingDomains {
		if d.Name == domain {
			exists = true
			break
		}
	}

	if !exists {
		return nil, fmt.Errorf("domain %s does not exists", domain)
	}

	body := schema.UpdateDomainRequest{
		Method: method,
		Params: schema.UpdateDomainParams{
			Domain:         domain,
			MailForwarding: mailForwarding,
			DNSSEC:         dnssec,
			Lock:           lock,
		},
	}
	req, err := c.client.NewRequest(ctx, body)
	if err != nil {
		return nil, err
	}
	fmt.Println(req.Body)
	resp, err := c.client.DoRequest(req, &responseScheme)
	if err != nil {
		return nil, err
	}
	response := resp.(*schema.UpdateDomainRequestResponse)
	return response, nil
}
