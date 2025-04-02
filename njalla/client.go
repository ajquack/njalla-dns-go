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
func Application(name, version string) ClientOption {
	return func(client *Client) {
		client.applicationName = name
		client.applicationVersion = version
	}
}

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
