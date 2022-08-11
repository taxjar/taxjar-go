package taxjar

import "encoding/json"

// ShowOrderParams should be passed to `ShowOrder` to show an order․
type ShowOrderParams struct {
	Provider string `url:"provider,omitempty"`
}

// ShowOrderResponse is the structure returned from `ShowOrder`․
//
// Access the order with `ShowOrderResponse.Order`․
type ShowOrderResponse struct {
	Order Order `json:"order"`
}

// OrderLineItem is the structure for a line item passed within `CreateOrderParams.LineItems` and `UpdateOrderParams.LineItems`․
//
// OrderLineItem is also the structure for a line item returned within `CreateOrderResponse.Order.LineItems`, `UpdateOrderResponse.Order.LineItems`, `ShowOrderResponse.Order.LineItems`, and `DeleteOrderResponse.Order.LineItems`․
type OrderLineItem struct {
	ID                string  `json:"id,omitempty"`
	Quantity          int     `json:"quantity,omitempty"`
	ProductIdentifier string  `json:"product_identifier,omitempty"`
	Description       string  `json:"description,omitempty"`
	ProductTaxCode    string  `json:"product_tax_code,omitempty"`
	UnitPrice         float64 `json:"unit_price,omitempty,string"`
	Discount          float64 `json:"discount,omitempty,string"`
	SalesTax          float64 `json:"sales_tax,omitempty,string"`
}

// ShowOrder shows an existing order in TaxJar․
//
// See https://developers.taxjar.com/api/reference/?go#get-show-an-order-transaction for more details․
func (client *Config) ShowOrder(transactionID string, params ...ShowOrderParams) (*ShowOrderResponse, error) {
	var p interface{}
	if len(params) > 0 {
		p = params[0]
	}

	res, err := client.get("transactions/orders/"+transactionID, p)
	if err != nil {
		return nil, err
	}

	order := new(ShowOrderResponse)
	if err := json.Unmarshal(res.([]byte), &order); err != nil {
		return nil, err
	}

	return order, nil
}
