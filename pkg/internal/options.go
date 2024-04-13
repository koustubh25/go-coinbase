package internal

import (
	"log/slog"
	"net/http"

	"github.com/hashicorp/go-cleanhttp"
)

type AdvanceTradeAPIEndpoint struct {
	Production *Endpoint
	Sandbox    *Endpoint
}

type Endpoint struct {
	RESTBaseURL  string
	WebSocketURL string
}

type Options struct {
	AdvanceTradeAPIEndpoint *AdvanceTradeAPIEndpoint
	Logger                  *slog.Logger
	JWT                     string // it's almost useless to set JWT at this level, but keeping this for backward compatibility
}

func (o Options) NewTransport() (http.RoundTripper, error) {
	baseTransport := cleanhttp.DefaultPooledTransport()
	return &AuthTransport{
		Base: baseTransport,
		JWT:  o.JWT,
	}, nil
}

func (o Options) NewHTTPClient() (*http.Client, error) {
	transport, err := o.NewTransport()
	if err != nil {
		return nil, err
	}
	return &http.Client{Transport: transport}, nil
}

// AuthTransport is a transport that adds an Authorization header to the request.
type AuthTransport struct {
	Base http.RoundTripper
	JWT  string
}

var _ http.RoundTripper = (*AuthTransport)(nil)

func (a *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if a.JWT != "" {
		req.Header.Set("Authorization", "Bearer "+a.JWT)
	}
	return a.Base.RoundTrip(req)
}
