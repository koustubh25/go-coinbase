package advancedtradeapi

import "github.com/koustubh25/go-coinbase/pkg/internal"

// NewRESTClient creates a new client with the provided options
// if different values that update the same field are provided, the last one will be used
// e.g. NewRESTClient(WithSandbox(), WithProduction()) will use the production endpoint
func NewRESTClient(opts ...Option) (*internal.RESTClient, error) {
	o := getOptionsWithDefaults(opts)
	c := &internal.RESTClient{
		Logger:                  o.Logger,
		AdvanceTradeAPIEndpoint: o.AdvanceTradeAPIEndpoint,
	}
	var err error
	c.HttpClient, err = o.NewHTTPClient()
	if err != nil {
		return nil, err
	}
	return c, nil
}
