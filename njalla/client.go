// Package client provides a client implementation for interacting with the Njalla DNS API.
// It includes functionality for managing domains, DNS records, forwards, glue records, and DNSSEC.
//
// The Client struct serves as the main entry point for making API requests. It supports
// customization through various options, such as setting an API key, application name, and version.
//
// The package also includes helper methods for creating and executing HTTP requests, handling
// API responses, and managing user-agent headers.
//
// Usage:
//   - Use NewClient to create a new client instance with optional configurations.
//   - Use the sub-clients (Domain, Record, Forward, Glue, DNSSEC) for specific API operations.
//
// Example:
//
//	client := client.NewClient(client.APIKey("your-api-key"), client.Application("MyApp", "1.0"))
//	domainInfo, err := client.Domain.Get("example.com")
//
// Constants:
//   - HTTPMethod: The HTTP method used for API requests (default: "POST").
//
// Types:
//   - Client: Represents the main client for interacting with the API.
//   - ClientOption: A function type for customizing the client configuration.
//
// Functions:
//   - APIKey: Sets the API key for the client.
//   - Application: Sets the application name and version for the client.
//   - NewClient: Creates a new client instance with optional configurations.
//   - (Client) NewRequest: Creates a new HTTP request for the API.
//   - (Client) DoRequest: Executes an HTTP request and processes the response.
//   - (Client) UserAgent: Configures the user-agent header for the client.
//
// Errors:
//   - Returns an error if the API key is invalid or if the server responds with an error status code.
//
// Notes:
//   - The API key must be a 40-character alphanumeric string.
//   - The client automatically sets the "Authorization" header if a valid API key is provided.
//   - The "User-Agent" header is customizable based on the application name and version.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

const httpMethod string = "POST"

type Client struct {
	httpClient         *http.Client
	apiKey             string
	endpoint           string
	apiKeyValid        bool
	applicationName    string
	applicationVersion string
	userAgent          string

	Domain  *DomainClient
	Record  *RecordClient
	Forward *ForwardClient
	Glue    *GlueClient
	DNSSEC  *DNSSECClient
}

var validApiKey = regexp.MustCompile("[a-z0-9]{40}")

type ClientOption func(*Client)

func APIKey(apiKey string) ClientOption {
	return func(client *Client) {
		client.apiKey = apiKey
		client.apiKeyValid = validApiKey.MatchString(apiKey)
	}
}

// Application sets the application name and version for the client.
// This function returns a ClientOption, which is a function that modifies
// the Client instance by assigning the provided application name and version.
//
// Parameters:
//   - name: The name of the application.
//   - version: The version of the application.
//
// Returns:
//
//	A ClientOption that applies the specified application name and version
//	to a Client instance.
//
// Notes:
//   - If the API key is invalid, the method returns an error.
//   - If the body is nil, the Content-Type and Accept headers are not set.
func Application(name, version string) ClientOption {
	return func(client *Client) {
		client.applicationName = name
		client.applicationVersion = version
	}
}

// NewClient creates a new instance of the Client with the provided options.
// It initializes the client with default values and applies any ClientOption
// functions passed as arguments to customize the client configuration.
//
// The function also sets up various sub-clients for managing domains, records,
// forwards, glue records, and DNSSEC.
//
// Parameters:
//
//	options - A variadic list of ClientOption functions to customize the client.
//
// Returns:
//
//	*Client - A pointer to the newly created Client instance.
func NewClient(options ...ClientOption) *Client {
	client := &Client{
		endpoint:    Endpoint,
		apiKeyValid: true,
		httpClient:  http.DefaultClient,
	}

	for _, option := range options {
		option(client)
	}

	client.UserAgent()

	client.Domain = &DomainClient{client}
	client.Record = &RecordClient{client}
	client.Forward = &ForwardClient{client}
	client.Glue = &GlueClient{client}
	client.DNSSEC = &DNSSECClient{client}

	return client
}

// NewRequest creates a new HTTP request for the Njalla API client.
// It accepts a context.Context and a request body of any type, which will be
// marshaled into JSON format. The method sets appropriate headers, including
// User-Agent, Authorization (if the API key is valid), and Content-Type/Accept
// for JSON requests.
//
// Parameters:
//   - ctx: The context for the request, used for cancellation and deadlines.
//   - body: The request payload to be sent in the body of the HTTP request.
//
// Returns:
//   - *http.Request: The constructed HTTP request.
//   - error: An error if the request could not be created or if the API key is invalid.
//
// Notes:
//   - If the API key is invalid, the method returns an error.
//   - If the body is nil, the Content-Type and Accept headers are not set.
func (c *Client) NewRequest(ctx context.Context, body any) (*http.Request, error) {
	url := c.endpoint
	reqBody, err := json.Marshal(body)
	req, err := http.NewRequest(HTTPMethod, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.userAgent)
	if !c.apiKeyValid {
		return nil, errors.New("invalid API key")
	} else if c.apiKey != "" {
		req.Header.Set("Authorization", "Njalla "+c.apiKey)
	}
	if body != nil {
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
	}
	req = req.WithContext(ctx)
	return req, nil
}

// DoRequest sends an HTTP request using the client's HTTP client and processes the response.
// It expects the response body to be a JSON object containing a "result" field.
// The "result" field is unmarshaled into the provided `v` parameter if it is not nil.
//
// Parameters:
//   - r: The HTTP request to be sent.
//   - v: A pointer to a variable where the "result" field of the response will be unmarshaled.
//
// Returns:
//   - any: The unmarshaled value of the "result" field if `v` is provided, or nil otherwise.
//   - error: An error if the request fails, the response status code is not 200,
//     the response body cannot be decoded, or the "result" field is missing.
//
// Example usage:
//
//	var result MyStruct
//	err := client.DoRequest(request, &result)
//	if err != nil {
//	    log.Fatalf("Request failed: %v", err)
//	}
func (c *Client) DoRequest(r *http.Request, v any) (any, error) {
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Server responded with status code %d", resp.StatusCode)
	}

	wrapper := struct {
		Result json.RawMessage `json:"result"`
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// If no result field was found
	if len(wrapper.Result) == 0 {
		return nil, fmt.Errorf("missing result field in response")
	}

	if v != nil {
		if err := json.Unmarshal(wrapper.Result, v); err != nil {
			return nil, fmt.Errorf("failed to unmarshal result: %w", err)
		}
		return v, nil
	}
	return nil, err
}

// UserAgent constructs and sets the user agent string for the Client instance.
// It uses the application's name and version if provided, appending them to
// the default UserAgent string. The logic is as follows:
//   - If both applicationName and applicationVersion are set, the user agent
//     will be formatted as "applicationName/applicationVersion UserAgent".
//   - If only applicationName is set, the user agent will be formatted as
//     "applicationName UserAgent".
//   - If neither is set, the user agent will default to the value of UserAgent.
func (c *Client) UserAgent() {
	switch {
	case c.applicationName != "" && c.applicationVersion != "":
		c.userAgent = c.applicationName + "/" + c.applicationVersion + " " + UserAgent
	case c.applicationName != "" && c.applicationVersion == "":
		c.userAgent = c.applicationName + " " + UserAgent
	default:
		c.userAgent = UserAgent
	}
}
