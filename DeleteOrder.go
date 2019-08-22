package taxjar

import "encoding/json"

// DeleteOrderParams - TODO (document this)
type DeleteOrderParams struct {
	Provider string `url:"provider,omitempty"`
}

// DeleteOrderResponse - TODO (document this)
type DeleteOrderResponse struct {
	Order struct {
		TransactionID string `json:"transaction_id"`
		UserID        int    `json:"user_id"`
		Provider      string `json:"provider"`
	} `json:"order"`
}

// DeleteOrder - TODO (document this)
func (client *Config) DeleteOrder(transactionID string, params ...DeleteOrderParams) (*DeleteOrderResponse, error) {
	res, err := client.delete("transactions/orders/"+transactionID, params)
	if err != nil {
		return nil, err
	}
	order := new(DeleteOrderResponse)
	json.Unmarshal(res.([]byte), &order)
	return order, nil
}
