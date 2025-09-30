package inwx

import (
	"context"
	"net/http"
	"time"

	"github.com/nmeilick/inwx-cli/internal/api"
)

const (
	// DefaultDNSTTL is the default TTL for DNS records in seconds
	DefaultDNSTTL = 3600
)

type Environment int

const (
	Production Environment = iota
	Testing
)

type Client struct {
	transport      *api.Transport
	username       string
	password       string
	env            Environment
	customEndpoint bool
}

type ClientOption func(*Client)

func WithCredentials(username, password string) ClientOption {
	return func(c *Client) {
		c.username = username
		c.password = password
	}
}

func WithEnvironment(env Environment) ClientOption {
	return func(c *Client) {
		c.env = env
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.transport.SetTimeout(timeout)
	}
}

func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.transport.SetHTTPClient(client)
	}
}

func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) {
		c.transport.SetUserAgent(userAgent)
	}
}

func WithEndpoint(endpoint string) ClientOption {
	return func(c *Client) {
		c.customEndpoint = true
		c.transport.SetEndpoint(endpoint)
	}
}

func NewClient(opts ...ClientOption) (*Client, error) {
	client := &Client{
		env: Production,
	}

	transport, err := api.NewTransport()
	if err != nil {
		return nil, err
	}
	client.transport = transport

	for _, opt := range opts {
		opt(client)
	}

	// Only set default endpoint if no custom endpoint was provided
	if !client.customEndpoint {
		var endpoint string
		switch client.env {
		case Testing:
			endpoint = "https://api.ote.domrobot.com/jsonrpc/"
		default:
			endpoint = "https://api.domrobot.com/jsonrpc/"
		}
		client.transport.SetEndpoint(endpoint)
	}

	return client, nil
}

// Login authenticates with the INWX API using the configured credentials
func (c *Client) Login(ctx context.Context) error {
	return c.transport.Login(ctx, c.username, c.password)
}

// Logout ends the current API session
func (c *Client) Logout(ctx context.Context) error {
	return c.transport.Logout(ctx)
}

// DNS creates a new DNS service instance with the specified options
func (c *Client) DNS(opts ...DNSOption) *DNSService {
	service := &DNSService{
		client:     c,
		defaultTTL: DefaultDNSTTL,
	}

	for _, opt := range opts {
		opt(service)
	}

	return service
}
