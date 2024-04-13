package main

import (
	"context"
	"fmt"
	"log"

	"github.com/koustubh25/go-coinbase/pkg/advancedtradeapi"
	"github.com/koustubh25/go-coinbase/pkg/advancedtradeapi/websocket"
)

type Candle struct {
	Start     string `json:"start"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Open      string `json:"open"`
	Close     string `json:"close"`
	Volume    string `json:"volume"`
	ProductID string `json:"product_id"`
}

type Event struct {
	Type    string   `json:"type"`
	Candles []Candle `json:"candles"`
}

type CandlesMessage struct {
	Channel     string  `json:"channel"`
	ClientID    string  `json:"client_id"`
	Timestamp   string  `json:"timestamp"`
	SequenceNum int     `json:"sequence_num"`
	Events      []Event `json:"events"`
}

type Trade struct {
	TradeID   string `json:"trade_id"`
	ProductID string `json:"product_id"`
	Price     string `json:"price"`
	Size      string `json:"size"`
	Side      string `json:"side"`
	Time      string `json:"time"`
}

type TradeEvent struct {
	Type   string  `json:"type"`
	Trades []Trade `json:"trades"`
}

type MarketTradesMessage struct {
	Channel     string       `json:"channel"`
	ClientID    string       `json:"client_id"`
	Timestamp   string       `json:"timestamp"`
	SequenceNum int          `json:"sequence_num"`
	Events      []TradeEvent `json:"events"`
}

func subscribe() error {
	websocketClient, err := advancedtradeapi.NewWebSocketClient(advancedtradeapi.WithProduction())
	if err != nil {
		return err
	}
	ctx := context.Background()
	msgs, err := websocket.Subscribe[CandlesMessage](ctx, "candles", []string{"ETH-USD"}, websocketClient)
	if err != nil {
		log.Fatal(err)
	}
	for msg := range msgs {
		fmt.Println(msg)
	}
	return nil
}
