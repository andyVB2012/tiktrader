package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

var priceMocks = map[string]float64{
	"BTC":  35000,
	"ETH":  2700,
	"STFX": 1,
}

func (s *priceFetcher) FetchPrice(_ context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return 0.0, fmt.Errorf("price for ticker (%s) is not available", ticker)
	}

	return price, nil
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {

	time.Sleep(100 * time.Millisecond)
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("ticker not found )%s) ", ticker)
	}
	return price, nil
}

type loggingService struct {
	next PriceFetcher
}

func (s loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		reqID := ctx.Value("requestID")

		logrus.WithFields(logrus.Fields{
			"requestID": reqID,
			"took":      time.Since(begin),
			"err":       err,
			"price":     price,
			"ticker":    ticker,
		}).Info("FetchPrice")
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)
}
