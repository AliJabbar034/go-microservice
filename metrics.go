package main

import (
	"context"
	"fmt"
)

type metricsService struct {
	next PriceFetcher
}

func NewMetricsService(next PriceFetcher) PriceFetcher {
	return &metricsService{
		next: next,
	}
}

func (m *metricsService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("index ", ticker)
	return m.next.FetchPrice(ctx, ticker)

}
