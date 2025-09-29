package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

type JSONRPCRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
	ID     int         `json:"id"`
}

type JSONRPCResponse struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
	ID     int         `json:"id"`
}

type JSONRPCClient struct {
	client   *http.Client
	endpoint string
	session  *Session
	id       int
}

func NewJSONRPCClient(client *http.Client, session *Session) *JSONRPCClient {
	return &JSONRPCClient{
		client:  client,
		session: session,
		id:      1,
	}
}

func (c *JSONRPCClient) SetEndpoint(endpoint string) {
	c.endpoint = endpoint
}

func (c *JSONRPCClient) SetHTTPClient(client *http.Client) {
	c.client = client
}

func (c *JSONRPCClient) Call(ctx context.Context, method string, params interface{}) (map[string]interface{}, error) {
	request := JSONRPCRequest{
		Method: method,
		Params: params,
		ID:     c.id,
	}
	c.id++

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	log.Debug().
		Str("method", method).
		Str("endpoint", c.endpoint).
		RawJSON("request", requestBody).
		Msg("JSON-RPC request")

	req, err := http.NewRequestWithContext(ctx, "POST", c.endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "inwx-go/1.0.0")

	// Add session cookies
	for _, cookie := range c.session.GetCookies() {
		req.AddCookie(cookie)
		log.Debug().
			Str("cookie_name", cookie.Name).
			Str("cookie_value", cookie.Value).
			Msg("Adding session cookie")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("HTTP request failed")
		return nil, err
	}
	defer resp.Body.Close()

	// Store session cookies
	c.session.StoreCookies(resp.Cookies())
	for _, cookie := range resp.Cookies() {
		log.Debug().
			Str("cookie_name", cookie.Name).
			Str("cookie_value", cookie.Value).
			Msg("Received session cookie")
	}

	if resp.StatusCode != http.StatusOK {
		log.Error().
			Int("status_code", resp.StatusCode).
			Str("status", resp.Status).
			Msg("HTTP error response")
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read response body")
		return nil, err
	}

	log.Debug().
		Str("method", method).
		RawJSON("response", responseBody).
		Msg("JSON-RPC response")

	var response JSONRPCResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		log.Error().Err(err).Str("body", string(responseBody)).Msg("Failed to parse JSON response")
		return nil, err
	}

	if response.Error != nil {
		return nil, fmt.Errorf("JSON-RPC error: %v", response.Error)
	}

	// For INWX API, the response structure is different from standard JSON-RPC
	// The actual response data is directly in the response body, not in a "result" field

	// If this is a standard JSON-RPC response with a result field
	if response.Result != nil {
		// Try to convert result to map[string]interface{}
		switch result := response.Result.(type) {
		case map[string]interface{}:
			return result, nil
		case []interface{}:
			// If result is an array, wrap it in a map
			return map[string]interface{}{"result": result}, nil
		case string, float64, bool:
			// If result is a primitive type, wrap it in a map
			return map[string]interface{}{"result": result}, nil
		default:
			// Try to marshal and unmarshal to convert to map
			data, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("unexpected response format: %T", result)
			}

			var resultMap map[string]interface{}
			if err := json.Unmarshal(data, &resultMap); err != nil {
				// If that fails, wrap the original result
				return map[string]interface{}{"result": result}, nil
			}

			return resultMap, nil
		}
	}

	// For INWX API, parse the response body directly as the API response
	var apiResponse map[string]interface{}
	if err := json.Unmarshal(responseBody, &apiResponse); err != nil {
		log.Error().Err(err).Str("body", string(responseBody)).Msg("Failed to parse API response")
		return nil, err
	}

	return apiResponse, nil
}
