package rest

import (
	"context"

	"github.com/koustubh25/go-coinbase/pkg/internal"
)

const accountsPath = "/accounts"

// ListAccounts Get a list of authenticated accounts for the current user.
// https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getaccounts
func ListAccounts[T any](ctx context.Context, c *internal.RESTClient) (any, error) {
	var a T
	baseURL, err := getBaseURL(c.AdvanceTradeAPIEndpoint)
	if err != nil {
		return nil, err
	}
	if err := c.DoRequest(ctx, "GET", baseURL+requestAPIPath+accountsPath, nil, &a, nil); err != nil {
		return nil, err
	}
	return a, nil
}

// GetAccount Get a list of information about an account, given an account UUID.
// https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getaccount
func GetAccount[T any](ctx context.Context, c *internal.RESTClient, accountId string) (any, error) {
	var a T
	baseURL, err := getBaseURL(c.AdvanceTradeAPIEndpoint)
	if err != nil {
		return nil, err
	}
	if err := c.DoRequest(ctx, "GET", baseURL+requestAPIPath+accountsPath+"/"+accountId, nil, &a, nil); err != nil {
		return nil, err
	}
	return a, nil
}
