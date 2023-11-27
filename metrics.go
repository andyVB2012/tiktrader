package main

import (
	"context"
	"fmt"
)

type MetricService struct {
	next PriceFetcher
}

func NewMetricService(svc PriceFetcher) PriceFetcher {
	return &MetricService{next: svc}
}

func (s *MetricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("metrics service")
	// metrics storage
	return s.next.FetchPrice(ctx, ticker)
}
