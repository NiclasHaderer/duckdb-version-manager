package api

import (
	"net/http"
	"time"
)

func New() Client {
	return &ApiClient{
		Host: "https://raw.githubusercontent.com/NiclasHaderer/duckdb-version-manager/main/",
		Client: &http.Client{
			Timeout: 5 * time.Minute,
		},
		BasePath: "/versions",
	}
}
