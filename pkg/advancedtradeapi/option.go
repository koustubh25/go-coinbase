package advancedtradeapi

import (
	"log/slog"

	"github.com/koustubh25/go-coinbase/pkg/internal"
)

type Option func(*internal.Options)

const (
	ProductionRESTBaseURL  = "api.coinbase.com"
	ProductionWebSocketURL = "advanced-trade-ws.coinbase.com"
	SandboxRESTBaseURL     = "api-n5e1.coinbase.com"
	SandboxWebSocketURL    = "ws-md.n5e2.coinbase.com"
)

func WithJWT(jwt string) Option {
	return func(o *internal.Options) {
		o.JWT = jwt
	}
}

func WithProduction() Option {
	return func(o *internal.Options) {
		advanceTradeAPIEndpoint := &internal.AdvanceTradeAPIEndpoint{Production: &internal.Endpoint{RESTBaseURL: ProductionRESTBaseURL, WebSocketURL: ProductionWebSocketURL}, Sandbox: nil}
		o.AdvanceTradeAPIEndpoint = advanceTradeAPIEndpoint
	}
}

func WithSandbox() Option {
	return func(o *internal.Options) {
		advanceTradeAPIEndpoint := &internal.AdvanceTradeAPIEndpoint{Sandbox: &internal.Endpoint{RESTBaseURL: SandboxRESTBaseURL, WebSocketURL: SandboxWebSocketURL}, Production: nil}
		o.AdvanceTradeAPIEndpoint = advanceTradeAPIEndpoint
	}
}

// getOptionsWithDefaults returns a new Options with default values set. If options are passed as paramters they will override the defautl values
func getOptionsWithDefaults(opts []Option) internal.Options {
	var options internal.Options
	options.Logger = slog.Default()
	// hit sandbox by default
	options.AdvanceTradeAPIEndpoint = &internal.AdvanceTradeAPIEndpoint{Sandbox: &internal.Endpoint{RESTBaseURL: SandboxRESTBaseURL, WebSocketURL: SandboxWebSocketURL}}
	for _, opt := range opts {
		opt(&options)
	}
	return options
}
