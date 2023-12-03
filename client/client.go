package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andyVB2012/tiktrader/mytypes"
	"github.com/andyVB2012/tiktrader/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//	type Client interface {
//		FetchPrice(ctx context.Context, ticker string) (*mytypes.PriceResponse, error)
//	}

func NewGRPCClient(remoteAddr string) (proto.PriceFetcherClient, error) {
	conn, err := grpc.Dial(remoteAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := proto.NewPriceFetcherClient(conn)
	return client, nil
}

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*mytypes.PriceResponse, error) {
	endpoint := fmt.Sprintf("%s?ticker=%s", c.endpoint, ticker)
	fmt.Println(endpoint)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		httpErr := map[string]any{}
		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("unexpected error: %s", httpErr["error"])
	}

	priceResp := new(mytypes.PriceResponse)
	if err := json.NewDecoder(resp.Body).Decode(priceResp); err != nil {
		return nil, err
	}
	return priceResp, nil
}
