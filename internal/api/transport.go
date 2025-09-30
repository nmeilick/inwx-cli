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

const (
	// DefaultTimeout is the default HTTP client timeout
	DefaultTimeout = 30 * time.Second
	// DefaultMaxRetries is the default number of retry attempts for API calls
	DefaultMaxRetries = 3
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
		Timeout: DefaultTimeout,
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
	return t.loginWithRetry(ctx, username, password, DefaultMaxRetries)
}

func (t *Transport) loginWithRetry(ctx context.Context, username, password string, maxRetries int) error {
	log.Debug().
		Str("username", username).
		Str("endpoint", t.endpoint).
		Msg("Attempting login")

	params := map[string]interface{}{
		"user": username,
		"pass": password,
		"lang": "en",
	}

	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			// Exponential backoff
			backoff := time.Duration(math.Pow(2, float64(attempt-1))) * time.Second
			log.Debug().
				Int("attempt", attempt+1).
				Dur("backoff", backoff).
				Msg("Retrying login after backoff")

			select {
			case <-time.After(backoff):
			case <-ctx.Done():
				return ctx.Err()
			}
		}

		response, err := t.jsonrpc.Call(ctx, "account.login", params)
		if err != nil {
			lastErr = err

			// Check if this is a retryable error
			if t.isRetryableError(err) && attempt < maxRetries-1 {
				log.Warn().
					Err(err).
					Int("attempt", attempt+1).
					Msg("Retryable login error, will retry")
				continue
			}

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

	return lastErr
}

func (t *Transport) Logout(ctx context.Context) error {
	return t.logoutWithRetry(ctx, DefaultMaxRetries)
}

func (t *Transport) logoutWithRetry(ctx context.Context, maxRetries int) error {
	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			// Exponential backoff
			backoff := time.Duration(math.Pow(2, float64(attempt-1))) * time.Second
			log.Debug().
				Int("attempt", attempt+1).
				Dur("backoff", backoff).
				Msg("Retrying logout after backoff")

			select {
			case <-time.After(backoff):
			case <-ctx.Done():
				return ctx.Err()
			}
		}

		_, err := t.jsonrpc.Call(ctx, "account.logout", map[string]interface{}{})
		if err != nil {
			lastErr = err

			// Check if this is a retryable error
			if t.isRetryableError(err) && attempt < maxRetries-1 {
				log.Warn().
					Err(err).
					Int("attempt", attempt+1).
					Msg("Retryable logout error, will retry")
				continue
			}

			return err
		}

		return nil
	}

	return lastErr
}

func (t *Transport) Call(ctx context.Context, method string, params map[string]interface{}) (map[string]interface{}, error) {
	return t.callWithRetry(ctx, method, params, DefaultMaxRetries)
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

	// Check for connection refused (temporary network issue)
	if opErr, ok := err.(*net.OpError); ok {
		if syscallErr, ok := opErr.Err.(*os.SyscallError); ok {
			if syscallErr.Err == syscall.ECONNREFUSED {
				return true
			}
		}
	}

	// Check for HTTP 429 (rate limiting)
	if httpErr, ok := err.(*HTTPError); ok && httpErr.IsRateLimitError() {
		log.Warn().Msg("Rate limit detected (HTTP 429), will retry with backoff")
		return true
	}

	return false
}
