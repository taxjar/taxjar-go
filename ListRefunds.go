package taxjar

import "encoding/json"

// ListRefundsParams - TODO (document this)
type ListRefundsParams struct {
	TransactionDate     string `url:"transaction_date,omitempty"`
	FromTransactionDate string `url:"from_transaction_date,omitempty"`
	ToTransactionDate   string `url:"to_transaction_date,omitempty"`
	Provider            string `url:"provider,omitempty"`
}

// ListRefundsResponse TODO (document this)
type ListRefundsResponse struct {
	Refunds []string `json:"refunds"`
}

// ListRefunds - TODO (document this)
func (client *Config) ListRefunds(params ListRefundsParams) (*ListRefundsResponse, error) {
	res, err := client.get("transactions/refunds", params)
	if err != nil {
		return nil, err
	}
	refunds := new(ListRefundsResponse)
	json.Unmarshal(res.([]byte), &refunds)
	return refunds, nil
}
