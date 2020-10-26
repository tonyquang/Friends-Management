package response

import "fmt"

// ResponseError http error response struct
type ResponseError struct {
	Code        int    `json:"-"`
	Description string `json:"error_description"`
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("%s", e.Description)
}
