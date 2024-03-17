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

type clientImpl struct {
	Host     string
	Client   *http.Client
	BasePath string
}

func (c clientImpl) GetAllReleases() (models.Releases, stacktrace.Error) {
	releaseDict, err := c.ListAllReleasesDict()
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}

	finalReleases := make(models.Releases)
	for version, versionPath := range releaseDict {
		release, err := c.GetReleaseWithLocation(versionPath)
		if err != nil {
			return nil, stacktrace.Wrap(err)
		}
		finalReleases[version] = *release
	}

	return finalReleases, nil
}

func (c clientImpl) ListAllReleasesDict() (models.VersionDict, stacktrace.Error) {
	url := c.Host + c.BasePath + "/versions.json"
	resp, err := c.Client.Get(url)
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

func (c clientImpl) ListAllReleases() (models.VersionList, stacktrace.Error) {
	result, err := c.ListAllReleasesDict()
	if err != nil {
		return nil, err
	}

	return toVersionList(result), nil
}

func (c clientImpl) GetRelease(version string) (*models.Release, stacktrace.Error) {
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

func (c clientImpl) GetReleaseWithLocation(versionPath string) (*models.Release, stacktrace.Error) {
	url := c.Host + c.BasePath + "/" + versionPath

	resp, err := c.Client.Get(url)
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

func (c clientImpl) LatestDuckVmRelease() (*models.Release, stacktrace.Error) {
	url := c.Host + c.BasePath + "/latest-vm.json"

	resp, err := c.Client.Get(url)
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

func (c clientImpl) Get() *http.Client {
	return c.Client
}
