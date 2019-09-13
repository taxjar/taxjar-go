package taxjar

import "encoding/json"

// DeleteOrderParams should be passed to `DeleteOrder` to delete an order․
type DeleteOrderParams struct {
	Provider string `url:"provider,omitempty"`
}

// DeleteOrderResponse is the structure returned from `DeleteOrder`․
//
// Access the deleted order with `DeleteOrderResponse.Order`․
type DeleteOrderResponse struct {
	Order Order `json:"order"`
}

// DeleteOrder deletes an order in TaxJar․
//
// See https://developers.taxjar.com/api/reference/?go#delete-delete-an-order-transaction for more details․
func (client *Config) DeleteOrder(transactionID string, params ...DeleteOrderParams) (*DeleteOrderResponse, error) {
	res, err := client.delete("transactions/orders/"+transactionID, params)
	if err != nil {
		return nil, err
	}
	order := new(DeleteOrderResponse)
	json.Unmarshal(res.([]byte), &order)
	return order, nil
}
