package rssapi

import (
	"net/http"
	"time"
)

// client
type Client struct {
	httpClient http.Client
}

// New client
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
