package taxjar

import "encoding/json"

// ShowRefundParams - TODO (document this)
type ShowRefundParams struct {
	Provider string `url:"provider,omitempty"`
}

// ShowRefundResponse - TODO (document this)
type ShowRefundResponse struct {
	Refund Order `json:"refund"`
}

// RefundLineItem - TODO (document this)
type RefundLineItem struct {
	ID                string  `json:"id,omitempty"`
	Quantity          int     `json:"quantity,omitempty"`
	ProductIdentifier string  `json:"product_identifier,omitempty"`
	Description       string  `json:"description,omitempty"`
	ProductTaxCode    string  `json:"product_tax_code,omitempty"`
	UnitPrice         float64 `json:"unit_price,omitempty"`
	Discount          float64 `json:"discount,omitempty"`
	SalesTax          float64 `json:"sales_tax,omitempty"`
}

// ShowRefund - TODO (document this)
func (client *Config) ShowRefund(transactionID string, params ...ShowRefundParams) (*ShowRefundResponse, error) {
	res, err := client.get("transactions/refunds/"+transactionID, params)
	if err != nil {
		return nil, err
	}
	refund := new(ShowRefundResponse)
	json.Unmarshal(res.([]byte), &refund)
	return refund, nil
}
