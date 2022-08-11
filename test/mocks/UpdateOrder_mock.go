package mocks

import (
	"encoding/json"

	"github.com/taxjar/taxjar-go"
)

// UpdateOrder - mock response
var UpdateOrder = new(taxjar.UpdateOrderResponse)
var _ = json.Unmarshal([]byte(UpdateOrderJSON), &UpdateOrder)

// UpdateOrderJSON - mock UpdateOrder JSON
var UpdateOrderJSON = `{
  "order": {
    "transaction_id": "24",
    "user_id": 121449,
    "provider": "api",
    "transaction_date": "2019-08-26T00:00:00.000Z",
    "transaction_reference_id": null,
    "customer_id": "123",
    "exemption_type": "non_exempt",
    "from_country": "US",
    "from_zip": "94043",
    "from_state": "CA",
    "from_city": "MOUNTAIN VIEW",
    "from_street": "311 Moffet Blvd",
    "to_country": "US",
    "to_zip": "10019",
    "to_state": "NY",
    "to_city": "NEW YORK",
    "to_street": "1697 Broadway",
    "amount": "161.0",
    "shipping": "5.0",
    "sales_tax": "10.3",
    "line_items": [
      {
        "id": "0",
        "quantity": 1,
        "product_identifier": "10-12345-987",
        "product_tax_code": "20010",
        "description": "10-gallon Hat",
        "unit_price": "45.0",
        "discount": "0.0",
        "sales_tax": "0.0"
      },
      {
        "id": "1",
        "quantity": 1,
        "product_identifier": "78-95432-101",
        "product_tax_code": "20010",
        "description": "Extra-long Chaps",
        "unit_price": "111.0",
        "discount": "0.0",
        "sales_tax": "9.85"
      }
    ]
  }
}`
