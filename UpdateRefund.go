package taxjar

import "encoding/json"

// UpdateRefundParams - TODO (document this)
type UpdateRefundParams struct {
	TransactionID          string           `json:"transaction_id,omitempty"`
	TransactionReferenceID string           `json:"transaction_reference_id,omitempty"`
	TransactionDate        string           `json:"transaction_date,omitempty"`
	FromCountry            string           `json:"from_country,omitempty"`
	FromZip                string           `json:"from_zip,omitempty"`
	FromState              string           `json:"from_state,omitempty"`
	FromCity               string           `json:"from_city,omitempty"`
	FromStreet             string           `json:"from_street,omitempty"`
	ToCountry              string           `json:"to_country,omitempty"`
	ToZip                  string           `json:"to_zip,omitempty"`
	ToState                string           `json:"to_state,omitempty"`
	ToCity                 string           `json:"to_city,omitempty"`
	ToStreet               string           `json:"to_street,omitempty"`
	Amount                 float64          `json:"amount,omitempty"`
	Shipping               float64          `json:"shipping,omitempty"`
	SalesTax               float64          `json:"sales_tax,omitempty"`
	CustomerID             string           `json:"customer_id,omitempty"`
	ExemptionType          string           `json:"exemption_type,omitempty"`
	LineItems              []RefundLineItem `json:"line_items,omitempty"`
}

// UpdateRefundResponse - TODO (document this)
type UpdateRefundResponse struct {
	Refund Order `json:"refund"`
}

// UpdateRefund - TODO (document this)
func (client *Config) UpdateRefund(params UpdateRefundParams) (*UpdateRefundResponse, error) {
	res, err := client.put("transactions/refunds/"+params.TransactionID, params)
	if err != nil {
		return nil, err
	}
	refund := new(UpdateRefundResponse)
	json.Unmarshal(res.([]byte), &refund)
	return refund, nil
}
