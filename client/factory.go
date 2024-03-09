package client

import (
	"net/http"
	"time"
)

func New() Client {
	return &ApiClient{
		// Host: "https://raw.githubusercontent.com/NiclasHaderer/duckdb-version-manager/main/",
		Host: "http://127.0.0.1:8000",
		Client: &http.Client{
			Timeout: 5 * time.Minute,
		},
		BasePath: "/versions",
	}
}
