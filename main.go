package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/andyVB2012/tiktrader/client"
	"github.com/andyVB2012/tiktrader/proto"
)

func main() {
	var (
		jsonAddr = flag.String("port", ":3000", "server listen address of json transport")
		grpcAddr = flag.String("grpc", ":4000", "server listen address of grpc transport")
		svc      = loggingService{&priceFetcher{}}
		ctx      = context.Background()
	)
	flag.Parse()

	grpcClient, err := client.NewGRPCClient(":4000")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		time.Sleep(3 * time.Second)
		resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "BTC"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", resp)
	}()
	go makeGRPCServerAndRun(*grpcAddr, svc)
	jsonServer := NewJSONAPIServer(*jsonAddr, svc)

	jsonServer.Run()

}
