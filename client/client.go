package client

import "context"

// type Client interface {
// 	FetchPrice(ctx context.Context, ticker string) (float64, error)
// }

type Client struct {
	endpoint string
}

func NewClient(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return 0, nil
}
