package rest

import (
	"fmt"

	"github.com/koustubh25/go-coinbase/pkg/internal"
)

const (
	requestAPIPath = "/api/v3/brokerage"
)

// getBaseURL returns the correct base URL for the API endpoint
// this api is based on the fact that production endpoints are nil when sandbox is not nil and vice versa
func getBaseURL(a *internal.AdvanceTradeAPIEndpoint) (string, error) {
	var baseUrl string
	if a.Sandbox != nil {
		baseUrl = a.Sandbox.RESTBaseURL
	} else if a.Production != nil {
		baseUrl = a.Production.RESTBaseURL
	} else {
		return "", fmt.Errorf(`both production and sandbox endpoints are nil`)
	}
	return baseUrl, nil
}
