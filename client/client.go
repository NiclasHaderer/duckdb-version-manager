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
	GetReleaseWithLocation(versionPath string) (*models.Release, error)
	ListAllReleases() (models.VersionList, error)
	ListAllReleasesDict() (models.VersionDict, error)
}

type ApiClient struct {
	Host     string
	Client   *http.Client
	BasePath string
}

func (receiver ApiClient) GetAllReleases() (models.Releases, error) {
	releasesDict, err := receiver.ListAllReleasesDict()
	if err != nil {
		return nil, err
	}

	finalReleases := make(models.Releases)
	for version, versionPath := range releasesDict {
		release, err := receiver.GetReleaseWithLocation(versionPath)
		if err != nil {
			return nil, err
		}
		finalReleases[version] = *release
	}

	return finalReleases, nil
}

func (receiver ApiClient) ListAllReleasesDict() (models.VersionDict, error) {
	url := receiver.Host + receiver.BasePath + "/versions.json"
	resp, err := receiver.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var releases models.VersionDict
	err = json.NewDecoder(resp.Body).Decode(&releases)
	if err != nil {
		return nil, err
	}

	return releases, nil
}

func (receiver ApiClient) ListAllReleases() (models.VersionList, error) {
	result, err := receiver.ListAllReleasesDict()
	if err != nil {
		return nil, err
	}

	return toVersionList(result), nil
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

	return receiver.GetReleaseWithLocation(versionPath)
}

func (receiver ApiClient) GetReleaseWithLocation(versionPath string) (*models.Release, error) {
	url := receiver.Host + receiver.BasePath + "/" + versionPath

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
