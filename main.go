package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/andyVB2012/tiktrader/client"
)

func main() {
	client := client.New("http://localhost:3000")

	price, err := client.FetchPrice(context.Background(), "ETH")
	if err != nil {
		fmt.Println("error fetching price")
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", price)
	return
	listenAddr := flag.String("listen-addr", ":3000", "server listen address the service is running")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)

	server.Run()

}
