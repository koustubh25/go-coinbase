package websocket

import (
	"fmt"

	"github.com/koustubh25/go-coinbase/pkg/internal"
)

// getBaseURL returns the correct base URL for the WebSocket API endpoint
func getBaseURL(a *internal.AdvanceTradeAPIEndpoint) (string, error) {
	var baseUrl string
	if a.Sandbox != nil && a.Sandbox.WebSocketURL != "" {
		baseUrl = a.Sandbox.WebSocketURL
	} else if a.Production != nil && a.Production.WebSocketURL != "" {
		baseUrl = a.Production.WebSocketURL
	} else {
		return "", fmt.Errorf(`both production and sandbox websocket api endpoints are nil`)
	}
	return baseUrl, nil
}
