package api

import (
	"duckdb-version-manager/models"
	"duckdb-version-manager/stacktrace"
	"duckdb-version-manager/utils"
	"encoding/json"
	"net/http"
	"time"
)

type Client interface {
	GetAllReleases() (map[models.VersionStr]models.VersionInformation, stacktrace.Error)
	GetRelease(version string) (*models.VersionInformation, stacktrace.Error)
	GetReleaseWithLocation(versionPath string) (*models.VersionInformation, stacktrace.Error)
	ListAllReleases() (models.RemoteVersions, stacktrace.Error)
	ListAllReleasesDict() (models.RemoteVersionDict, stacktrace.Error)
	LatestDuckVmRelease(timeout time.Duration) (*models.VersionInformation, stacktrace.Error)
	Get() *http.Client
}

type clientImpl struct {
	Host     string
	Client   *http.Client
	BasePath string
}

func (c clientImpl) GetAllReleases() (map[models.VersionStr]models.VersionInformation, stacktrace.Error) {
	releaseDict, err := c.ListAllReleasesDict()
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}

	finalReleases := make(map[models.VersionStr]models.VersionInformation)
	for version, versionPath := range releaseDict {
		release, err := c.GetReleaseWithLocation(versionPath)
		if err != nil {
			return nil, stacktrace.Wrap(err)
		}
		finalReleases[version] = *release
	}

	return finalReleases, nil
}

func (c clientImpl) ListAllReleasesDict() (models.RemoteVersionDict, stacktrace.Error) {
	url := c.Host + c.BasePath + "/versions.json"
	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}
	defer resp.Body.Close()

	var releases models.RemoteVersionDict
	err = json.NewDecoder(resp.Body).Decode(&releases)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}

	return releases, nil
}

func (c clientImpl) ListAllReleases() (models.RemoteVersions, stacktrace.Error) {
	result, err := c.ListAllReleasesDict()
	if err != nil {
		return nil, err
	}

	return utils.ToVersionList(&result), nil
}

func (c clientImpl) GetRelease(version string) (*models.VersionInformation, stacktrace.Error) {
	versions, err := c.ListAllReleasesDict()
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

	return c.GetReleaseWithLocation(versionPath)
}

func (c clientImpl) GetReleaseWithLocation(versionPath string) (*models.VersionInformation, stacktrace.Error) {
	url := c.Host + c.BasePath + "/" + versionPath

	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}
	defer resp.Body.Close()

	var release models.VersionInformation
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}

	return &release, nil
}

func (c clientImpl) LatestDuckVmRelease(timeout time.Duration) (*models.VersionInformation, stacktrace.Error) {
	url := c.Host + c.BasePath + "/latest-vm.json"

	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}
	defer resp.Body.Close()

	var release models.VersionInformation
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}

	return &release, nil
}

func (c clientImpl) Get() *http.Client {
	return c.Client
}
