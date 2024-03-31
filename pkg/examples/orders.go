package main

import (
	"context"
	"fmt"

	"github.com/koustubh25/go-coinbase/pkg/advancedtradeapi"
	"github.com/koustubh25/go-coinbase/pkg/advancedtradeapi/rest"
)

func createOrder() {
	client, err := advancedtradeapi.NewRESTClient(advancedtradeapi.WithProduction())
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	a, err := rest.CreateOrder[rest.OrderResponse](ctx, client, &rest.OrderOptions{
		ClientOrderID: "random",
		ProductID:     "BTC-USD",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}
