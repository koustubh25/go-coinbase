# Coinbase Advanced Trade API Go wrapper

This repository contains a Go wrapper around the official [Coinbase Advanced Trade API](https://docs.cloud.coinbase.com/advanced-trade-api/docs/welcome).

## Features

- Easy interaction with the Coinbase Advanced Trade API using Go.
- Handles JWT token authentication internally.

## Usage

### Prerequisite

1. Need to have the following two environment variables set
   1. `COINBASE_KEY_NAME` - name of the API key
   2. `COINBASE_KEY_SECRET` - private key
   
For obtaining these, refer the steps [here](
https://docs.cloud.coinbase.com/advanced-trade-api/docs/rest-api-auth#cloud-api-trading-keys)

First, import the package into your Go project:

```go
import "github.com/koustubh25/go-coinbase"
```

Then, create a new client and use it to interact with the API:
```go
import (
	"context"
	"fmt"

	"github.com/koustubh25/go-coinbase/pkg/advancedtradeapi"
	"github.com/koustubh25/go-coinbase/pkg/advancedtradeapi/rest"
)

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
```
Find more examples in the [examples](/pkg/examples/) directory