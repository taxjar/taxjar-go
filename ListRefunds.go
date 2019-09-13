package taxjar

import "encoding/json"

// ListRefundsParams should be passed to `ListRefunds` to list existing refund IDs․
type ListRefundsParams struct {
	TransactionDate     string `url:"transaction_date,omitempty"`
	FromTransactionDate string `url:"from_transaction_date,omitempty"`
	ToTransactionDate   string `url:"to_transaction_date,omitempty"`
	Provider            string `url:"provider,omitempty"`
}

// ListRefundsResponse is the structure returned from `ListRefunds`․
//
// Access the refund list with `ListRefundsResponse.Refunds`․
type ListRefundsResponse struct {
	Refunds []string `json:"refunds"`
}

// ListRefunds lists existing refund IDs in TaxJar․
//
// See https://developers.taxjar.com/api/reference/?go#get-list-refund-transactions for more details․
func (client *Config) ListRefunds(params ListRefundsParams) (*ListRefundsResponse, error) {
	res, err := client.get("transactions/refunds", params)
	if err != nil {
		return nil, err
	}
	refunds := new(ListRefundsResponse)
	json.Unmarshal(res.([]byte), &refunds)
	return refunds, nil
}
