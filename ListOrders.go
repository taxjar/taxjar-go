package taxjar

import "encoding/json"

// ListOrdersParams - TODO (document this)
type ListOrdersParams struct {
	TransactionDate     string `url:"transaction_date,omitempty"`
	FromTransactionDate string `url:"from_transaction_date,omitempty"`
	ToTransactionDate   string `url:"to_transaction_date,omitempty"`
	Provider            string `url:"provider,omitempty"`
}

// ListOrdersResponse - TODO (document this)
type ListOrdersResponse struct {
	Orders []string `json:"orders"`
}

// ListOrders - TODO (document this)
func (client *Config) ListOrders(params ListOrdersParams) (*ListOrdersResponse, error) {
	res, err := client.get("transactions/orders", params)
	if err != nil {
		return nil, err
	}
	orders := new(ListOrdersResponse)
	json.Unmarshal(res.([]byte), &orders)
	return orders, nil
}
