package inwx

import (
	"context"
)

type DomainService struct {
	client *Client
}

type Domain struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func (c *Client) Domain() *DomainService {
	return &DomainService{
		client: c,
	}
}

func (s *DomainService) List(ctx context.Context) ([]Domain, error) {
	response, err := s.client.transport.Call(ctx, "domain.list", map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var domains []Domain
	if resData, ok := response["resData"].(map[string]interface{}); ok {
		if domainList, ok := resData["domain"].([]interface{}); ok {
			for _, d := range domainList {
				if domain, ok := d.(map[string]interface{}); ok {
					domainObj := Domain{}
					if name, ok := domain["domain"].(string); ok {
						domainObj.Name = name
					}
					if status, ok := domain["status"].(string); ok {
						domainObj.Status = status
					}
					domains = append(domains, domainObj)
				}
			}
		}
	}

	return domains, nil
}
