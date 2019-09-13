package taxjar

import "encoding/json"

// DeleteRefundParams - TODO (document this)
type DeleteRefundParams struct {
	Provider string `url:"provider,omitempty"`
}

// DeleteRefundResponse - TODO (document this)
type DeleteRefundResponse struct {
	Refund Refund `json:"refund"`
}

// DeleteRefund - TODO (document this)
func (client *Config) DeleteRefund(transactionID string, params ...DeleteRefundParams) (*DeleteRefundResponse, error) {
	res, err := client.delete("transactions/refunds/"+transactionID, params)
	if err != nil {
		return nil, err
	}
	refund := new(DeleteRefundResponse)
	json.Unmarshal(res.([]byte), &refund)
	return refund, nil
}
