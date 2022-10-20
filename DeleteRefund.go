package taxjar

import "encoding/json"

// DeleteRefundParams should be passed to `DeleteRefund` to delete a refund․
type DeleteRefundParams struct {
	Provider string `url:"provider,omitempty"`
}

// DeleteRefundResponse is the structure returned from `DeleteRefund`․
//
// Access the deleted refund with `DeleteRefundResponse.Refund`․
type DeleteRefundResponse struct {
	Refund Refund `json:"refund"`
}

// DeleteRefund deletes a refund in TaxJar․
//
// See https://developers.taxjar.com/api/reference/?go#delete-delete-a-refund-transaction for more details․
func (client *Config) DeleteRefund(transactionID string, params ...DeleteRefundParams) (*DeleteRefundResponse, error) {
	var p interface{}
	if len(params) > 0 {
		p = params[0]
	}

	res, err := client.delete("transactions/refunds/"+transactionID, p)
	if err != nil {
		return nil, err
	}

	refund := new(DeleteRefundResponse)
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
