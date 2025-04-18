package models

// Request represents a generic request structure
type Request struct {
	Method  string            `json:"method"`
	Path    string            `json:"path"`
	Headers map[string]string `json:"headers"`
	Body    interface{}       `json:"body,omitempty"`
}

// NewRequest creates a new Request instance
func NewRequest() *Request {
	return &Request{
		Headers: make(map[string]string),
	}
}
