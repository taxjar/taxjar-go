package mocks

import (
	"encoding/json"

	"github.com/taxjar/taxjar-go"
)

// DeleteRefund - mock response
var DeleteRefund = new(taxjar.DeleteRefundResponse)
var _ = json.Unmarshal([]byte(DeleteRefundJSON), &DeleteRefund)

// DeleteRefundJSON - mock DeleteRefund JSON
var DeleteRefundJSON = `{
  "refund": {
    "transaction_id": "24-refund",
    "user_id": 121449,
    "provider": "api",
    "transaction_date": null,
    "transaction_reference_id": null,
    "customer_id": null,
    "exemption_type": null,
    "from_country": null,
    "from_zip": null,
    "from_state": null,
    "from_city": null,
    "from_street": null,
    "to_country": null,
    "to_zip": null,
    "to_state": null,
    "to_city": null,
    "to_street": null,
    "amount": null,
    "shipping": null,
    "sales_tax": null,
    "line_items": []
  }
}`
