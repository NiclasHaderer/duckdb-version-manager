package client

import (
	"duckdb-version-manager/models"
	"encoding/json"
	"errors"
	"net/http"
)

type Client interface {
	GetAllReleases() (models.Releases, error)
	GetRelease(version string) (*models.Release, error)
	ListAllReleases() (models.VersionList, error)
	ListAllReleasesDict() (models.VersionDict, error)
}

type ApiClient struct {
	Host            string
	Client          *http.Client
	VersionsPath    string
	VersionBasePath string
}

func (receiver ApiClient) GetAllReleases() (models.Releases, error) {
	url := receiver.Host + receiver.VersionsPath
	resp, err := receiver.Client.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var releases models.VersionDict
	err = json.NewDecoder(resp.Body).Decode(&releases)
	if err != nil {
		return nil, err
	}

	// TODO return not nil
	return nil, nil
}

func (receiver ApiClient) ListAllReleasesDict() (models.VersionDict, error) {
	url := receiver.Host + receiver.VersionsPath
	resp, err := receiver.Client.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var releases models.VersionDict
	err = json.NewDecoder(resp.Body).Decode(&releases)
	if err != nil {
		return nil, err
	}

	return releases, nil
}

func (receiver ApiClient) ListAllReleases() (models.VersionList, error) {
	_, err := receiver.ListAllReleasesDict()
	if err != nil {
		return nil, err
	}

	// TODO return not nil
	//  https://github.com/hashicorp/go-version?tab=readme-ov-file for sorting the versions
	return nil, nil
}

func (receiver ApiClient) GetRelease(version string) (*models.Release, error) {
	versions, err := receiver.ListAllReleasesDict()
	if err != nil {
		return nil, err
	}

	versionPath, ok := versions[version]
	if !ok {
		return nil, errors.New("version not found")
	}

	url := receiver.Host + receiver.VersionBasePath + "/" + versionPath

	resp, err := receiver.Client.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var release models.Release
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return nil, err
	}

	return &release, nil
}
