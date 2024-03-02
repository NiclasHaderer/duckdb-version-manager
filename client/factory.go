package client

import (
	"net/http"
	"time"
)

func New() Client {
	return &ApiClient{
		Host: "http://localhost:8000",
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
		BasePath: "/versions",
	}
}
