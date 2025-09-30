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

// Account creates a new account service instance for managing account information
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

	// Safety check: ensure response is not nil
	if response == nil {
		return info, nil
	}

	if resData, ok := response["resData"].(map[string]interface{}); ok && resData != nil {
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
