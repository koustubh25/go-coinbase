package main

import (
	"context"
	"fmt"

	"github.com/koustubh25/go-coinbase/pkg/advancedtradeapi"
	"github.com/koustubh25/go-coinbase/pkg/advancedtradeapi/rest"
)

func getAccounts() {
	client, err := advancedtradeapi.NewRESTClient(advancedtradeapi.WithProduction())
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	a, err := rest.ListAccounts[rest.AccountsEnvelope](ctx, client)
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
	b, err := rest.GetAccount[rest.AccountEnvelope](ctx, client, "431c504f-4fd5-5e50-9ed3-1e1e7e08a8ec")
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
}
