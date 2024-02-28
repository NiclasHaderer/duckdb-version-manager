package client

import (
	"net/http"
	"time"
)

func New() Client {
	return &ApiClient{
		Host: "http://localhost:8080",
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
		VersionPath: "/versions/index.json",
	}
}
