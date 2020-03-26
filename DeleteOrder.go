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
	var p interface{}
	if len(params) > 0 {
		p = params[0]
	}

	res, err := client.delete("transactions/orders/"+transactionID, p)
	if err != nil {
		return nil, err
	}

	order := new(DeleteOrderResponse)
	if err := json.Unmarshal(res.([]byte), &order); err != nil {
		return nil, err
	}

	return order, nil
}
