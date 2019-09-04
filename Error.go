package taxjar

import "fmt"

// Error - TODO (document this)
type Error struct {
	Err    string `json:"error"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("taxjar: %v %v - %v", err.Status, err.Err, err.Detail)
}
