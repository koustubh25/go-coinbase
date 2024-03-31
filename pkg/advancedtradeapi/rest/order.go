package rest

import (
	"context"

	"github.com/koustubh25/go-coinbase/pkg/internal"
)

// CreateOrder creates an order with a specified product_id (asset-pair), side (buy/sell), etc.
// https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_postorder
func CreateOrder[T any](ctx context.Context, c *internal.RESTClient, opts *OrderOptions) (any, error) {
	var a T
	baseURL, err := getBaseURL(c.AdvanceTradeAPIEndpoint)
	if err != nil {
		return nil, err
	}
	if err := c.DoRequest(ctx, "POST", baseURL+requestAPIPath+"/orders", nil, &a, opts); err != nil {
		return nil, err
	}
	return a, nil
}
