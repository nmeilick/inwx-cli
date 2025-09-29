package api

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
)

type XMLRPCMethodCall struct {
	XMLName xml.Name     `xml:"methodCall"`
	Method  string       `xml:"methodName"`
	Params  XMLRPCParams `xml:"params"`
}

type XMLRPCParams struct {
	Param []XMLRPCParam `xml:"param"`
}

type XMLRPCParam struct {
	Value XMLRPCValue `xml:"value"`
}

type XMLRPCValue struct {
	Struct XMLRPCStruct `xml:"struct,omitempty"`
	String string       `xml:"string,omitempty"`
	Int    int          `xml:"int,omitempty"`
	Bool   bool         `xml:"boolean,omitempty"`
}

type XMLRPCStruct struct {
	Member []XMLRPCMember `xml:"member"`
}

type XMLRPCMember struct {
	Name  string      `xml:"name"`
	Value XMLRPCValue `xml:"value"`
}

type XMLRPCMethodResponse struct {
	XMLName xml.Name     `xml:"methodResponse"`
	Params  XMLRPCParams `xml:"params"`
}

type XMLRPCClient struct {
	client   *http.Client
	endpoint string
	session  *Session
}

func NewXMLRPCClient(client *http.Client, session *Session) *XMLRPCClient {
	return &XMLRPCClient{
		client:  client,
		session: session,
	}
}

func (c *XMLRPCClient) SetEndpoint(endpoint string) {
	c.endpoint = endpoint
}

func (c *XMLRPCClient) SetHTTPClient(client *http.Client) {
	c.client = client
}

func (c *XMLRPCClient) Call(ctx context.Context, method string, params map[string]interface{}) (map[string]interface{}, error) {
	request := XMLRPCMethodCall{
		Method: method,
		Params: c.buildParams(params),
	}

	requestBody, err := xml.Marshal(request)
	if err != nil {
		return nil, err
	}

	fullBody := []byte(`<?xml version="1.0" encoding="UTF-8"?>` + "\n" + string(requestBody))

	req, err := http.NewRequestWithContext(ctx, "POST", c.endpoint, bytes.NewBuffer(fullBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("User-Agent", "inwx-go/1.0.0")

	// Add session cookies
	for _, cookie := range c.session.GetCookies() {
		req.AddCookie(cookie)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Store session cookies
	c.session.StoreCookies(resp.Cookies())

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	var response XMLRPCMethodResponse
	if err := xml.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return c.parseResponse(response), nil
}

func (c *XMLRPCClient) buildParams(params map[string]interface{}) XMLRPCParams {
	var xmlParams XMLRPCParams

	if len(params) > 0 {
		param := XMLRPCParam{
			Value: XMLRPCValue{
				Struct: c.buildStruct(params),
			},
		}
		xmlParams.Param = []XMLRPCParam{param}
	}

	return xmlParams
}

func (c *XMLRPCClient) buildStruct(data map[string]interface{}) XMLRPCStruct {
	var members []XMLRPCMember

	for key, value := range data {
		member := XMLRPCMember{
			Name:  key,
			Value: c.buildValue(value),
		}
		members = append(members, member)
	}

	return XMLRPCStruct{Member: members}
}

func (c *XMLRPCClient) buildValue(value interface{}) XMLRPCValue {
	switch v := value.(type) {
	case string:
		return XMLRPCValue{String: v}
	case int:
		return XMLRPCValue{Int: v}
	case bool:
		return XMLRPCValue{Bool: v}
	case map[string]interface{}:
		return XMLRPCValue{Struct: c.buildStruct(v)}
	default:
		return XMLRPCValue{String: fmt.Sprintf("%v", v)}
	}
}

func (c *XMLRPCClient) parseResponse(response XMLRPCMethodResponse) map[string]interface{} {
	result := make(map[string]interface{})

	if len(response.Params.Param) > 0 {
		param := response.Params.Param[0]
		if param.Value.Struct.Member != nil {
			for _, member := range param.Value.Struct.Member {
				result[member.Name] = c.parseValue(member.Value)
			}
		}
	}

	return result
}

func (c *XMLRPCClient) parseValue(value XMLRPCValue) interface{} {
	if value.String != "" {
		return value.String
	}
	if value.Int != 0 {
		return value.Int
	}
	if value.Bool {
		return value.Bool
	}
	if value.Struct.Member != nil {
		result := make(map[string]interface{})
		for _, member := range value.Struct.Member {
			result[member.Name] = c.parseValue(member.Value)
		}
		return result
	}
	return nil
}
