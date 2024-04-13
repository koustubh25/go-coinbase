package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"

	urlpkg "net/url"

	"github.com/koustubh25/go-coinbase/pkg/auth"
)

// RESTClient represents a client for making HTTP requests.
type RESTClient struct {
	HttpClient              *http.Client
	Logger                  *slog.Logger
	AdvanceTradeAPIEndpoint *AdvanceTradeAPIEndpoint
}

const (
	httpProtocolScheme = "https"
)

// DoRequest sends an HTTP request and processes the response.
// it is pretty dumb in the sense, it doesn't automatically verify the url endpoints e.g. websocket or REST
func (c *RESTClient) DoRequest(ctx context.Context, method, url string, query urlpkg.Values, result, reqBody interface{}) error {
	url = httpProtocolScheme + "://" + url
	if len(query) > 0 {
		url += "?" + query.Encode()
	}
	var reqBodyReader io.Reader
	if reqBody != nil {
		reqBodyBytes, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}
		reqBodyReader = bytes.NewReader(reqBodyBytes)
	}
	req, err := http.NewRequestWithContext(ctx, method, url, reqBodyReader)
	if err != nil {
		return err
	}
	if result != nil {
		req.Header.Set("Accept", "application/json")
	}
	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	c.Logger.Debug("setting auth token if not set already...", "method", method, "url", url)
	err = auth.SetToken(req, method+" "+strings.TrimPrefix(url, httpProtocolScheme+"://"))
	if err != nil {
		return err
	}
	c.Logger.Debug("setting auth token if not set already...done", "method", method, "url", url)
	c.Logger.Debug("doing request", "method", method, "url", url)
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error doing request %s %s: %w", req.Method, req.URL, err)
	}
	defer resp.Body.Close()
	logger := c.Logger.With("method", resp.Request.Method, "url", resp.Request.URL, "status", resp.StatusCode)
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading body of %d-response of request %s %s: %w", resp.StatusCode, resp.Request.Method, resp.Request.URL, err)
	}
	logger.Debug("checking response code...")
	if resp.StatusCode/100 != 2 {
		return fmt.Errorf(`received %d response code for request %s %s: %s`, resp.StatusCode, resp.Request.Method, resp.Request.URL, string(respBodyBytes))
	}
	logger.Debug("response code 2xx found in response of request")
	logger.Debug("unmarshalling response of request", "body", string(respBodyBytes))
	if err := json.Unmarshal(respBodyBytes, &result); err != nil {
		return fmt.Errorf("error unmarshalling body of %d-response of request %s %s: %s", resp.StatusCode, resp.Request.Method, resp.Request.URL, string(respBodyBytes))
	}
	logger.Debug("successfully unmarshalled response of request", "result", result)
	return nil
}
