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
	res, err := client.delete("transactions/refunds/"+transactionID, params)
	if err != nil {
		return nil, err
	}
	refund := new(DeleteRefundResponse)
	json.Unmarshal(res.([]byte), &refund)
	return refund, nil
}