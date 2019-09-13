package mocks

import (
	"encoding/json"

	"github.com/taxjar/taxjar-go"
)

// Error - mock
var Error = taxjar.Error{
	Status: 401,
	Err:    "Unauthorized",
	Detail: "Not authorized for route 'POST /v2/taxes'",
}

// ErrorJSON - mock Error JSON
var ErrorJSON, _ = json.Marshal(Error)
