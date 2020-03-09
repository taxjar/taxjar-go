package taxjar

import "encoding/json"

// ListOrdersParams should be passed to `ListOrders` to list existing order IDs․
type ListOrdersParams struct {
	TransactionDate     string `url:"transaction_date,omitempty"`
	FromTransactionDate string `url:"from_transaction_date,omitempty"`
	ToTransactionDate   string `url:"to_transaction_date,omitempty"`
	Provider            string `url:"provider,omitempty"`
}

// ListOrdersResponse is the structure returned from `ListOrders`․
//
// Access the order list with `ListOrdersResponse.Orders`․
type ListOrdersResponse struct {
	Orders []string `json:"orders"`
}

// ListOrders lists existing order IDs in TaxJar․
//
// See https://developers.taxjar.com/api/reference/?go#get-list-order-transactions for more details․
func (client *Config) ListOrders(params ListOrdersParams) (*ListOrdersResponse, error) {
	res, err := client.get("transactions/orders", params)
	if err != nil {
		return nil, err
	}
	orders := new(ListOrdersResponse)
	if err := json.Unmarshal(res.([]byte), &orders); err != nil {
		return nil, err
	}
	return orders, nil
}
