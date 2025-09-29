package api

import (
	"context"
	"math"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

type Transport struct {
	client    *http.Client
	endpoint  string
	userAgent string
	session   *Session
	jsonrpc   *JSONRPCClient
}

func NewTransport() (*Transport, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	session := NewSession()
	jsonrpc := NewJSONRPCClient(client, session)

	return &Transport{
		client:    client,
		userAgent: "inwx-go/1.0.0",
		session:   session,
		jsonrpc:   jsonrpc,
	}, nil
}

func (t *Transport) SetEndpoint(endpoint string) {
	t.endpoint = endpoint
	t.jsonrpc.SetEndpoint(endpoint)
}

func (t *Transport) SetTimeout(timeout time.Duration) {
	t.client.Timeout = timeout
}

func (t *Transport) SetHTTPClient(client *http.Client) {
	t.client = client
	t.jsonrpc.SetHTTPClient(client)
}

func (t *Transport) SetUserAgent(userAgent string) {
	t.userAgent = userAgent
}

func (t *Transport) Login(ctx context.Context, username, password string) error {
	log.Debug().
		Str("username", username).
		Str("endpoint", t.endpoint).
		Msg("Attempting login")

	params := map[string]interface{}{
		"user": username,
		"pass": password,
		"lang": "en",
	}

	response, err := t.jsonrpc.Call(ctx, "account.login", params)
	if err != nil {
		log.Error().Err(err).Msg("Login call failed")
		return err
	}

	log.Debug().Interface("response", response).Msg("Login response received")

	if code, ok := response["code"].(float64); ok && code != 1000 {
		msg := "Login failed"
		if message, ok := response["msg"].(string); ok {
			msg = message
		}
		log.Error().
			Float64("code", code).
			Str("message", msg).
			Msg("Login failed with API error")
		return NewAPIError(int(code), msg)
	}

	log.Debug().Msg("Login successful")
	return nil
}

func (t *Transport) Logout(ctx context.Context) error {
	_, err := t.jsonrpc.Call(ctx, "account.logout", map[string]interface{}{})
	return err
}

func (t *Transport) Call(ctx context.Context, method string, params map[string]interface{}) (map[string]interface{}, error) {
	return t.callWithRetry(ctx, method, params, 3)
}

func (t *Transport) callWithRetry(ctx context.Context, method string, params map[string]interface{}, maxRetries int) (map[string]interface{}, error) {
	log.Debug().
		Str("method", method).
		Interface("params", params).
		Msg("Making API call")

	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			// Exponential backoff
			backoff := time.Duration(math.Pow(2, float64(attempt-1))) * time.Second
			log.Debug().
				Int("attempt", attempt+1).
				Dur("backoff", backoff).
				Str("method", method).
				Msg("Retrying API call after backoff")

			select {
			case <-time.After(backoff):
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}

		response, err := t.jsonrpc.Call(ctx, method, params)
		if err != nil {
			lastErr = err

			// Check if this is a retryable error
			if t.isRetryableError(err) && attempt < maxRetries-1 {
				log.Warn().
					Err(err).
					Int("attempt", attempt+1).
					Str("method", method).
					Msg("Retryable error, will retry")
				continue
			}

			log.Error().
				Err(err).
				Str("method", method).
				Msg("API call failed")
			return nil, err
		}

		log.Debug().
			Str("method", method).
			Interface("response", response).
			Msg("API call response received")

		if code, ok := response["code"].(float64); ok && code != 1000 {
			msg := "API call failed"
			if message, ok := response["msg"].(string); ok {
				msg = message
			}

			reasonCode := ""
			if rc, ok := response["reasonCode"].(string); ok {
				reasonCode = rc
			}

			reason := ""
			if r, ok := response["reason"].(string); ok {
				reason = r
			}

			log.Error().
				Float64("code", code).
				Str("message", msg).
				Str("reasonCode", reasonCode).
				Str("reason", reason).
				Str("method", method).
				Msg("API call returned error")

			if reasonCode != "" || reason != "" {
				return nil, NewAPIErrorWithReason(int(code), msg, reasonCode, reason)
			}

			return nil, NewAPIError(int(code), msg)
		}

		log.Debug().
			Str("method", method).
			Interface("full_response", response).
			Msg("Returning successful response")

		return response, nil
	}

	return nil, lastErr
}

func (t *Transport) isRetryableError(err error) bool {
	// Check for network timeouts
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return true
	}

	// Check for temporary DNS errors
	if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
		return true
	}

	// Check for connection refused (temporary network issue)
	if opErr, ok := err.(*net.OpError); ok {
		if syscallErr, ok := opErr.Err.(*os.SyscallError); ok {
			if syscallErr.Err == syscall.ECONNREFUSED {
				return true
			}
		}
	}

	// Check for HTTP 429 (rate limiting) - this would need to be detected in the HTTP layer
	// For now, we'll handle this in the JSON-RPC client

	return false
}
