package internal

import (
	"log/slog"
	"net/url"

	"github.com/gorilla/websocket"
)

// WebSocketClient represents a client for making HTTP requests.
type WebSocketClient struct {
	Connection              *websocket.Conn
	Logger                  *slog.Logger
	AdvanceTradeAPIEndpoint *AdvanceTradeAPIEndpoint
}

const (
	webSocketProtocolScheme = "wss"
)

func (w *WebSocketClient) SetConnection(baseUrl string) error {
	// create a new websocket connection
	u := url.URL{Scheme: webSocketProtocolScheme, Host: baseUrl, Path: "/"}
	var err error
	w.Connection, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	return nil
}
