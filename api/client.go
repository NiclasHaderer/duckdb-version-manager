package api

import (
	"duckdb-version-manager/models"
	"duckdb-version-manager/stacktrace"
	"encoding/json"
	"net/http"
)

type Client interface {
	GetAllReleases() (models.Releases, stacktrace.Error)
	GetRelease(version string) (*models.Release, stacktrace.Error)
	GetReleaseWithLocation(versionPath string) (*models.Release, stacktrace.Error)
	ListAllReleases() (models.VersionList, stacktrace.Error)
	ListAllReleasesDict() (models.VersionDict, stacktrace.Error)
	LatestDuckVmRelease() (*models.Release, stacktrace.Error)
	Get() *http.Client
}

type ApiClient struct {
	Host     string
	Client   *http.Client
	BasePath string
}

func (receiver ApiClient) GetAllReleases() (models.Releases, stacktrace.Error) {
	releasesDict, err := receiver.ListAllReleasesDict()
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}

	finalReleases := make(models.Releases)
	for version, versionPath := range releasesDict {
		release, err := receiver.GetReleaseWithLocation(versionPath)
		if err != nil {
			return nil, stacktrace.Wrap(err)
		}
		finalReleases[version] = *release
	}

	return finalReleases, nil
}

func (receiver ApiClient) ListAllReleasesDict() (models.VersionDict, stacktrace.Error) {
	url := receiver.Host + receiver.BasePath + "/versions.json"
	resp, err := receiver.Client.Get(url)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}
	defer resp.Body.Close()

	var releases models.VersionDict
	err = json.NewDecoder(resp.Body).Decode(&releases)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}

	return releases, nil
}

func (receiver ApiClient) ListAllReleases() (models.VersionList, stacktrace.Error) {
	result, err := receiver.ListAllReleasesDict()
	if err != nil {
		return nil, err
	}

	return toVersionList(result), nil
}

func (receiver ApiClient) GetRelease(version string) (*models.Release, stacktrace.Error) {
	versions, err := receiver.ListAllReleasesDict()
	if err != nil {
		return nil, err
	}

	versionPath, ok := versions[version]
	if !ok {
		versionPath, ok = versions["v"+version]
		if !ok {
			return nil, stacktrace.NewF("Version '%s' not found in remote versions", version)
		}
	}

	return receiver.GetReleaseWithLocation(versionPath)
}

func (receiver ApiClient) GetReleaseWithLocation(versionPath string) (*models.Release, stacktrace.Error) {
	url := receiver.Host + receiver.BasePath + "/" + versionPath

	resp, err := receiver.Client.Get(url)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}
	defer resp.Body.Close()

	var release models.Release
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}

	return &release, nil
}

func (receiver ApiClient) LatestDuckVmRelease() (*models.Release, stacktrace.Error) {
	url := receiver.Host + receiver.BasePath + "/latest-vm.json"

	resp, err := receiver.Client.Get(url)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}
	defer resp.Body.Close()

	var release models.Release
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}

	return &release, nil
}

func (receiver ApiClient) Get() *http.Client {
	return receiver.Client
}
