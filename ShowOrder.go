package taxjar

import "encoding/json"

// ShowOrderParams - TODO (document this)
type ShowOrderParams struct {
	Provider string `url:"provider,omitempty"`
}

// ShowOrderResponse - TODO (document this)
type ShowOrderResponse struct {
	Order Order `json:"order"`
}

// OrderLineItem - TODO (document this)
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

// ShowOrder - TODO (document this)
func (client *Config) ShowOrder(transactionID string, params ...ShowOrderParams) (*ShowOrderResponse, error) {
	res, err := client.get("transactions/orders/"+transactionID, params)
	if err != nil {
		return nil, err
	}
	order := new(ShowOrderResponse)
	json.Unmarshal(res.([]byte), &order)
	return order, nil
}
