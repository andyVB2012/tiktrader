package main

import (
	"flag"
	"time"

	"github.com/andyVB2012/tiktrader/client"
)

func main() {
	var (
		jsonAddr = flag.String("port", ":3000", "server listen address of json transport")
		grpc     = flag.String("grpc", ":4000", "server listen address of grpc transport")
	)
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	grpcClient, err := client.NewGRPCClient(":4000")
	if err != nil {
		panic(err)
	}
	go func() {
		time.Sleep(3 * time.Second)
		grpcClient.makeGRPCServerAndRun(*grpc, svc)
	}()
	go makeGRPCServerAndRun(*grpc, svc)
	jsonServer := NewJSONAPIServer(*jsonAddr, svc)

	jsonServer.Run()

}
