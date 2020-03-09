package taxjar

import "encoding/json"

// TaxForOrder shows the sales tax that should be collected for a given order․
//
// See https://developers.taxjar.com/api/reference/?go#post-calculate-sales-tax-for-an-order for more details․
func (client *Config) TaxForOrder(params TaxForOrderParams) (*TaxForOrderResponse, error) {
	res, err := client.post("taxes", params)
	if err != nil {
		return nil, err
	}
	tax := new(TaxForOrderResponse)
	if err := json.Unmarshal(res.([]byte), &tax); err != nil {
		return nil, err
	}
	return tax, nil
}
