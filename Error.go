package taxjar

import "fmt"

// Error is the custom error type returned in the second return value of each TaxJar API method (e.g., `TaxForOrder`)․
//
// See here for example error handling that extracts `Err`, `Detail`, and `Status` fields and displays a stack trace: https://github.com/taxjar/taxjar-go/blob/master/README.md#error-handling
type Error struct {
	Err    string `json:"error"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("taxjar: %v %v - %v", err.Status, err.Err, err.Detail)
}
