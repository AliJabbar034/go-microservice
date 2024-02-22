package main

import (
	"context"
	"fmt"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct {
}

func (pf *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	fmt.Printf("Fetching price for %s\n", ticker)
	return MockPriceFetcher(ctx, ticker)

}

var priceMocks = map[string]float64{
	"AAPL": 100.0,
	"MSFT": 200.0,
	"AMZN": 300.0,
	"NFLX": 400.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	fmt.Println(price, ok, ticker)
	if !ok {
		return price, fmt.Errorf("the give ticker (%s) is not supported", ticker)
	}

	return price, nil
}
