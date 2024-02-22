package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type loginService struct {
	next PriceFetcher
}

func NewLoginService(nex PriceFetcher) PriceFetcher {
	return &loginService{
		next: nex,
	}
}

func (s *loginService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("Fetching price", ticker)
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"ticker": ticker,
			"took":   time.Since(begin),
			"reqID":  ctx.Value("reqID"),
			"err":    err,
			"price":  price,
		}).Info("fetchPrice")

	}(time.Now())
	return s.next.FetchPrice(ctx, ticker)

}
