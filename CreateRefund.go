package taxjar

import "encoding/json"

// CreateRefundParams should be passed to `CreateRefund` to create a refund․
type CreateRefundParams struct {
	TransactionID          string           `json:"transaction_id,omitempty"`
	TransactionReferenceID string           `json:"transaction_reference_id,omitempty"`
	TransactionDate        string           `json:"transaction_date,omitempty"`
	Provider               string           `json:"provider,omitempty"`
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
	Amount                 float64          `json:"amount"`
	Shipping               float64          `json:"shipping"`
	SalesTax               float64          `json:"sales_tax"`
	CustomerID             string           `json:"customer_id,omitempty"`
	ExemptionType          string           `json:"exemption_type,omitempty"`
	LineItems              []RefundLineItem `json:"line_items,omitempty"`
}

// CreateRefundResponse is the structure returned from `CreateRefund`․
//
// Access the created refund with `CreateRefundResponse.Refund`․
type CreateRefundResponse struct {
	Refund Refund `json:"refund"`
}

// CreateRefund creates a new refund in TaxJar․
//
// See https://developers.taxjar.com/api/reference/?go#post-create-a-refund-transaction for more details․
func (client *Config) CreateRefund(params CreateRefundParams) (*CreateRefundResponse, error) {
	res, err := client.post("transactions/refunds", params)
	if err != nil {
		return nil, err
	}

	refund := new(CreateRefundResponse)
	if err := json.Unmarshal(res.([]byte), &refund); err != nil {
		if typeError, ok := err.(*json.UnmarshalTypeError); ok {
			// Ignores JSON line_item.id type errors due to API's conversion of numeric strings to integers
			if !(typeError.Field == "refund.line_items.id" && typeError.Value == "number") {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return refund, nil
}
