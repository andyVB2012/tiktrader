package main

import (
	"context"
	"net"

	"github.com/andyVB2012/tiktrader/proto"
	"google.golang.org/grpc"
)

func makeGRPCServerAndRun(listenAddr string, svc PriceFetcher) error {
	// start grpc server
	grpcServer := NewGRPCPriceService(svc)

	ln, err := net.Listen("tcp", listenAddr)

	if err != nil {
		return err
	}
	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	proto.RegisterPriceFetcherServer(server, grpcServer)
	return server.Serve(ln)

}

type GRPCPriceFetcherService struct {
	svc PriceFetcher
	proto.UnimplementedPriceFetcherServer
}

func NewGRPCPriceService(svc PriceFetcher) *GRPCPriceFetcherService {
	return &GRPCPriceFetcherService{svc: svc}
}

func (s *GRPCPriceFetcherService) FetchPrice(ctx context.Context, req *proto.PriceRequest) (*proto.PriceResponse, error) {
	price, err := s.svc.FetchPrice(ctx, req.Ticker)
	if err != nil {
		return nil, err
	}
	resp := &proto.PriceResponse{
		Ticker: req.Ticker,
		Price:  float32(price),
	}
	return resp, nil
}
