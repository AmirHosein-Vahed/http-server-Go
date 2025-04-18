package models

// Response represents a generic response structure
type Response struct {
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers"`
	Body       interface{}       `json:"body,omitempty"`
}

// NewResponse creates a new Response instance
func NewResponse() *Response {
	return &Response{
		Headers: make(map[string]string),
	}
}
