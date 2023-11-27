package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andyVB2012/tiktrader/mytypes"
)

// type Client interface {
// 	FetchPrice(ctx context.Context, ticker string) (float64, error)
// }

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
	req, err := http.NewRequest("get", endpoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	priceResp := new(mytypes.PriceResponse)
	if err := json.NewDecoder(resp.Body).Decode(priceResp); err != nil {
		return nil, err
	}
	return priceResp, nil
}
