package taxjar

import "encoding/json"

// CreateOrderParams should be passed to `CreateOrder` to create an order․
type CreateOrderParams struct {
	TransactionID   string          `json:"transaction_id,omitempty"`
	TransactionDate string          `json:"transaction_date,omitempty"`
	Provider        string          `json:"provider,omitempty"`
	FromCountry     string          `json:"from_country,omitempty"`
	FromZip         string          `json:"from_zip,omitempty"`
	FromState       string          `json:"from_state,omitempty"`
	FromCity        string          `json:"from_city,omitempty"`
	FromStreet      string          `json:"from_street,omitempty"`
	ToCountry       string          `json:"to_country,omitempty"`
	ToZip           string          `json:"to_zip,omitempty"`
	ToState         string          `json:"to_state,omitempty"`
	ToCity          string          `json:"to_city,omitempty"`
	ToStreet        string          `json:"to_street,omitempty"`
	Amount          float64         `json:"amount"`
	Shipping        float64         `json:"shipping"`
	SalesTax        float64         `json:"sales_tax"`
	CustomerID      string          `json:"customer_id,omitempty"`
	ExemptionType   string          `json:"exemption_type,omitempty"`
	LineItems       []OrderLineItem `json:"line_items,omitempty"`
}

// Order is the structure for an order returned within `CreateOrderResponse`, `ShowOrderResponse`, `UpdateOrderResponse`, and `DeleteOrderResponse`․
type Order struct {
	TransactionID          string          `json:"transaction_id"`
	UserID                 int             `json:"user_id"`
	TransactionDate        string          `json:"transaction_date"`
	TransactionReferenceID string          `json:"transaction_reference_id"`
	Provider               string          `json:"provider"`
	ExemptionType          string          `json:"exemption_type,omitempty"`
	FromCountry            string          `json:"from_country"`
	FromZip                string          `json:"from_zip"`
	FromState              string          `json:"from_state"`
	FromCity               string          `json:"from_city"`
	FromStreet             string          `json:"from_street"`
	ToCountry              string          `json:"to_country"`
	ToZip                  string          `json:"to_zip"`
	ToState                string          `json:"to_state"`
	ToCity                 string          `json:"to_city"`
	ToStreet               string          `json:"to_street"`
	Amount                 float64         `json:"amount,string"`
	Shipping               float64         `json:"shipping,string"`
	SalesTax               float64         `json:"sales_tax,string"`
	LineItems              []OrderLineItem `json:"line_items"`
}

// CreateOrderResponse is the structure returned from `CreateOrder`․
//
// Access the created order with `CreateOrderResponse.Order`․
type CreateOrderResponse struct {
	Order Order `json:"order"`
}

// CreateOrder creates a new order in TaxJar․
//
// See https://developers.taxjar.com/api/reference/?go#post-create-an-order-transaction for more details․
func (client *Config) CreateOrder(params CreateOrderParams) (*CreateOrderResponse, error) {
	res, err := client.post("transactions/orders", params)
	if err != nil {
		return nil, err
	}

	order := new(CreateOrderResponse)
	if err := json.Unmarshal(res.([]byte), &order); err != nil {
		return nil, err
	}

	return order, nil
}
