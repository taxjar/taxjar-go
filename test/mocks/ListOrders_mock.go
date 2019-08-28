package mocks

import (
	"encoding/json"

	"github.com/taxjar/taxjar-go"
)

// ListOrders - mock response
var ListOrders = new(taxjar.ListOrdersResponse)
var _ = json.Unmarshal([]byte(ListOrdersJSON), &ListOrders)

// ListOrdersJSON - mock ListOrders JSON
var ListOrdersJSON = `{
	"orders": [
		"123",
		"246",
		"359"
	]
}`
