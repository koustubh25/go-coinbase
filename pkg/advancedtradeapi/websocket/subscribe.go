package websocket

import (
	"context"
	"log"

	"github.com/koustubh25/go-coinbase/pkg/auth"
	"github.com/koustubh25/go-coinbase/pkg/internal"
)

func Subscribe[T any](ctx context.Context, channel string, products []string, c *internal.WebSocketClient) (chan T, error) {
	baseURL, err := getBaseURL(c.AdvanceTradeAPIEndpoint)
	if err != nil {
		return nil, err
	}
	c.SetConnection(baseURL)

	// build a new jwt token everytime as the token is valid for 2 minutes
	jwt, err := auth.BuildJWTForWebSocketAPI()
	if err != nil {
		return nil, err
	}

	// subscribe to the specified channel
	err = c.Connection.WriteJSON(map[string]interface{}{
		"type":        "subscribe",
		"channel":     channel,
		"product_ids": products,
		"jwt":         jwt,
	})
	if err != nil {
		return nil, err
	}

	// create a channel to receive the messages
	msgs := make(chan T)

	// start a goroutine to read messages from the websocket connection
	go func() {
		defer close(msgs)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				var msg T
				if err := c.Connection.ReadJSON(&msg); err != nil {
					log.Println("read:", err)
					return
				}
				msgs <- msg
			}
		}
	}()
	return msgs, nil
}
