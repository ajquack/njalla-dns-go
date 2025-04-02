package client

import (
	"context"
	"fmt"

	"github.com/ajquack/njalla-dns-go/njalla/schema"
)

type DomainClient struct {
	client *Client
}

// GetDomain retrieves information about a specific domain using the provided domain parameters.
// It sends a request to the Njalla API with the "get-domain" method and returns the response.
//
// Parameters:
//   - ctx: The context for the request, used for cancellation and deadlines.
//   - domainParams: The parameters required to identify the domain.
//
// Returns:
//   - A pointer to a GetDomainRequestResponse struct containing the domain information.
//   - An error if the request fails or the response cannot be processed.
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

// ListDomains retrieves a list of domains associated with the account.
// It sends a request to the "list-domains" API endpoint and parses the response.
//
// Parameters:
//   - ctx: The context for the request, which can be used to control timeouts or cancellations.
//
// Returns:
//   - A slice of schema.ListDomainResponse containing the domain details.
//   - An error if the request fails or the response cannot be parsed.
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

// EditDomain updates the settings of an existing domain.
//
// This method allows you to modify the mail forwarding, DNSSEC, and lock
// settings of a specified domain. It first checks if the domain exists
// by retrieving the list of existing domains. If the domain does not exist,
// an error is returned.
//
// Parameters:
//   - ctx: The context for the request, used for cancellation and deadlines.
//   - domain: The name of the domain to be updated.
//   - mailForwarding: A boolean indicating whether mail forwarding should be enabled.
//   - dnssec: A boolean indicating whether DNSSEC should be enabled.
//   - lock: A boolean indicating whether the domain should be locked.
//
// Returns:
//   - A pointer to a schema.UpdateDomainRequestResponse containing the updated
//     domain information if the operation is successful.
//   - An error if the domain does not exist or if the update operation fails.
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
