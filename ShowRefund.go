package taxjar

import "encoding/json"

// ShowRefundParams should be passed to `ShowRefund` to show a refund․
type ShowRefundParams struct {
	Provider string `url:"provider,omitempty"`
}

// Refund is the structure for a refund returned within `CreateRefundResponse`, `ShowRefundResponse`, `UpdateRefundResponse`, and `DeleteRefundResponse`․
type Refund struct {
	Order
}

// ShowRefundResponse is the structure returned from `ShowRefund`․
//
// Access the refund with `ShowRefundResponse.Refund`․
type ShowRefundResponse struct {
	Refund Refund `json:"refund"`
}

// RefundLineItem is the structure for a line item passed within `CreateRefundParams.LineItems` and `UpdateRefundParams.LineItems`․
//
// RefundLineItem is also the structure for a line item returned within `CreateRefundResponse.Refund.LineItems`, `UpdateRefundResponse.Refund.LineItems`, `ShowRefundResponse.Refund.LineItems`, and `DeleteRefundResponse.Refund.LineItems`․
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

// ShowRefund shows an existing refund in TaxJar․
//
// See https://developers.taxjar.com/api/reference/?go#get-show-a-refund-transaction for more details․
func (client *Config) ShowRefund(transactionID string, params ...ShowRefundParams) (*ShowRefundResponse, error) {
	var p interface{}
	if len(params) > 0 {
		p = params[0]
	}

	res, err := client.get("transactions/refunds/"+transactionID, p)
	if err != nil {
		return nil, err
	}

	refund := new(ShowRefundResponse)
	if err := json.Unmarshal(res.([]byte), &refund); err != nil {
		return nil, err
	}

	return refund, nil
}
