package inwx

import (
	"context"
	"net/http"
	"time"

	"github.com/nmeilick/inwx-cli/internal/api"
)

type Environment int

const (
	Production Environment = iota
	Testing
)

type Client struct {
	transport *api.Transport
	username  string
	password  string
	env       Environment
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

	var endpoint string
	switch client.env {
	case Testing:
		endpoint = "https://api.ote.domrobot.com/jsonrpc/"
	default:
		endpoint = "https://api.domrobot.com/jsonrpc/"
	}
	client.transport.SetEndpoint(endpoint)

	return client, nil
}

func (c *Client) Login(ctx context.Context) error {
	return c.transport.Login(ctx, c.username, c.password)
}

func (c *Client) Logout(ctx context.Context) error {
	return c.transport.Logout(ctx)
}

func (c *Client) DNS(opts ...DNSOption) *DNSService {
	service := &DNSService{
		client:     c,
		defaultTTL: 3600,
	}

	for _, opt := range opts {
		opt(service)
	}

	return service
}
