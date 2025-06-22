
package accrual

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
)

type Good struct {
    Description string `json:"description"`
    Price       int    `json:"price"`
}

type OrderRequest struct {
    Order string `json:"order"`
    Goods []Good `json:"goods"`
}

type OrderResponse struct {
    Order   string  `json:"order"`
    Status  string  `json:"status"`
    Accrual float64 `json:"accrual,omitempty"`
}

type Client struct {
    baseURL string
    client  *http.Client
}

func New(baseURL string) *Client {
    return &Client{
        baseURL: baseURL,
        client: &http.Client{
            Timeout: 5 * time.Second,
        },
    }
}

func (c *Client) SendOrder(req OrderRequest) error {
    body, err := json.Marshal(req)
    if err != nil {
        return err
    }

    httpReq, err := http.NewRequest("POST", c.baseURL+"/api/orders", bytes.NewReader(body))
    if err != nil {
        return err
    }

    httpReq.Header.Set("Content-Type", "application/json")

    resp, err := c.client.Do(httpReq)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusAccepted && resp.StatusCode != http.StatusConflict {
        b, _ := io.ReadAll(resp.Body)
        return fmt.Errorf("accrual system error: %s", string(b))
    }

    return nil
}

func (c *Client) GetOrder(number string) (*OrderResponse, error) {
    url := fmt.Sprintf("%s/api/orders/%s", c.baseURL, number)
    resp, err := c.client.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusNoContent {
        return nil, nil
    }

    var res OrderResponse
    err = json.NewDecoder(resp.Body).Decode(&res)
    if err != nil {
        return nil, err
    }

    return &res, nil
}
