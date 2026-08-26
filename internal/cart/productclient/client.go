package productclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	SKU   int64  `json:"sku_id"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func New(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetProduct(sku int64) (Product, bool, error) {
	url := fmt.Sprintf("%s/products/%d", c.baseURL, sku)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return Product{}, false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return Product{}, false, nil
	}

	var p Product
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return Product{}, false, err
	}

	return p, true, nil
}
