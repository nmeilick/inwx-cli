package inwx

import (
	"context"
)

type AccountService struct {
	client *Client
}

type AccountInfo struct {
	AccountID  int    `json:"accountId"`
	CustomerID int    `json:"customerId"`
	Username   string `json:"username"`
	Email      string `json:"email"`
}

func (c *Client) Account() *AccountService {
	return &AccountService{
		client: c,
	}
}

func (s *AccountService) Info(ctx context.Context) (*AccountInfo, error) {
	response, err := s.client.transport.Call(ctx, "account.info", map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	info := &AccountInfo{}
	if resData, ok := response["resData"].(map[string]interface{}); ok {
		if accountID, ok := resData["accountId"].(float64); ok {
			info.AccountID = int(accountID)
		}
		if customerID, ok := resData["customerId"].(float64); ok {
			info.CustomerID = int(customerID)
		}
		if username, ok := resData["username"].(string); ok {
			info.Username = username
		}
		if email, ok := resData["email"].(string); ok {
			info.Email = email
		}
	}

	return info, nil
}
