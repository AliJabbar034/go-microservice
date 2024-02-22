package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alijabbar034/go-microservice/types"
)

type client struct {
	endpoint string
}

func NewClient(endpoint string) *client {
	return &client{
		endpoint: endpoint,
	}
}

func (c *client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	endpoint := fmt.Sprintf("%s?ticker=%s", c.endpoint, ticker)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Server response status: %d", res.StatusCode)
	}
	defer res.Body.Close()
	priceres := new(types.PriceResponse)
	if err := json.NewDecoder(res.Body).Decode(priceres); err != nil {
		return nil, err
	}
	fmt.Println("Price", priceres)
	return priceres, nil
}
