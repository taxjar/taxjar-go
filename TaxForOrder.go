package taxjar

import "encoding/json"

// TaxForOrder - TODO (document this)
func (client *Config) TaxForOrder(params TaxForOrderParams) (*TaxForOrderResponse, error) {
	res, err := client.post("taxes", params)
	if err != nil {
		return nil, err
	}
	tax := new(TaxForOrderResponse)
	json.Unmarshal(res.([]byte), &tax)
	return tax, nil
}
