package taxjar

import "encoding/json"

// UpdateOrderParams should be passed to `UpdateOrder` to update an existing order․
type UpdateOrderParams struct {
	TransactionID   string          `json:"transaction_id,omitempty"`
	TransactionDate string          `json:"transaction_date,omitempty"`
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
	Amount          *float64        `json:"amount,omitempty"`
	Shipping        *float64        `json:"shipping,omitempty"`
	SalesTax        *float64        `json:"sales_tax,omitempty"`
	CustomerID      string          `json:"customer_id,omitempty"`
	ExemptionType   string          `json:"exemption_type,omitempty"`
	LineItems       []OrderLineItem `json:"line_items,omitempty"`
}

// UpdateOrderResponse is the structure returned from `UpdateOrder`․
//
// Access the updated order with `UpdateOrderResponse.Order`․
type UpdateOrderResponse struct {
	Order Order `json:"order"`
}

// UpdateOrder updates an existing order in TaxJar․
//
// See https://developers.taxjar.com/api/reference/?go#put-update-an-order-transaction for more details․
func (client *Config) UpdateOrder(params UpdateOrderParams) (*UpdateOrderResponse, error) {
	res, err := client.put("transactions/orders/"+params.TransactionID, params)
	if err != nil {
		return nil, err
	}

	order := new(UpdateOrderResponse)
	if err := json.Unmarshal(res.([]byte), &order); err != nil {
		return nil, err
	}

	return order, nil
}

// Float64 returns a pointer to the float64 value passed in.
func Float64(v float64) *float64 {
	return &v
}

// Float64Value returns the value of the float64 pointer passed in or
// 0 if the pointer is nil.
func Float64Value(v *float64) float64 {
	if v != nil {
		return *v
	}
	return 0
}
