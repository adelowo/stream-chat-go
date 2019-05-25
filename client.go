package stream

import (
	"net/http"
)

type (
	// Client interfaces with Stream's chat API
	Client struct {
		inner *http.Client
	}
)

// New creates a new Client object which has been configured for usage.
// If no options are provided, it uses sensible defaults
func New(_ ...Option) (*Client, error) {
	return &Client{}, nil
}
