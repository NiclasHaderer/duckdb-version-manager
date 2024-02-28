package client

import (
	"duckdb-version-manager/models"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Client interface {
	GetAllReleases() (models.DuckdbReleases, error)
	GetRelease(version string) (*models.DuckdbRelease, error)
}

type ApiClient struct {
	Host        string
	Client      *http.Client
	VersionPath string
}

func (receiver ApiClient) GetAllReleases() (models.DuckdbReleases, error) {
	url := receiver.Host + receiver.VersionPath

	req, err := http.NewRequest("GET", url, nil)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(req.Body)

	if err != nil {
		return nil, err
	}

	var releases models.DuckdbReleases
	err = json.NewDecoder(req.Body).Decode(&releases)
	if err != nil {
		return nil, err
	}

	return releases, nil
}

func (receiver ApiClient) GetRelease(version string) (*models.DuckdbRelease, error) {

	allReleases, err := receiver.GetAllReleases()
	if err != nil {
		return nil, err
	}

	release, ok := allReleases[version]
	if !ok {
		return nil, errors.New("release not found")
	}

	return &release, nil
}
