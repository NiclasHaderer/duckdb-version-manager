package manager

import (
	"duckdb-version-manager/api"
	"duckdb-version-manager/config"
	"duckdb-version-manager/models"
	"duckdb-version-manager/utils"
	"encoding/json"
	"errors"
	"os"
	"syscall"
	"time"
)

type VersionManager interface {
	InstallVersion(version string) error
	UninstallVersion(version string) error
	ListInstalledVersions() []models.LocalInstallationInfo
	GetDefaultVersion() *models.LocalInstallationInfo
	SetDefaultVersion(version *string) error
	Run(version string, args []string) error
	VersionIsInstalled(version string) bool
	saveLocalConfig() error
}

type VersionManagerImpl struct {
	client      api.Client
	localConfig models.LocalConfig
}

func (v VersionManagerImpl) saveLocalConfig() error {
	//TODO implement me
	panic("implement me")
}

func (v VersionManagerImpl) InstallVersion(version string) error {
	release, err := v.client.GetRelease(version)
	if err != nil {
		return err
	}

	downloadUrl, err := utils.GetDownloadUrlFrom(release)
	if err != nil {
		return err
	}

	githubAsset, err := utils.GetResponseBodyFrom(v.client.Get(), downloadUrl)
	duckDb, err := utils.ExtractDuckdbFile(githubAsset)
	if err != nil {
		return err
	}

	fileLocation := config.InstallDir + "/" + release.Version
	if err := os.WriteFile(fileLocation, duckDb, 0700); err != nil {
		return err
	}

	v.localConfig.LocalInstallations[release.Version] = models.LocalInstallationInfo{
		Version:          release.Version,
		Location:         fileLocation,
		InstallationDate: time.Now().String(),
	}

	return v.saveConfig()
}

func (v VersionManagerImpl) UninstallVersion(version string) error {
	if _, ok := v.localConfig.LocalInstallations[version]; !ok {
		return errors.New("version not installed")
	}

	// Check if the version is the default version
	if *v.localConfig.DefaultVersion == version {
		err := v.SetDefaultVersion(nil)
		if err != nil {
			return err
		}
	}

	if err := os.Remove(v.localConfig.LocalInstallations[version].Location); err != nil {
		return err
	}
	delete(v.localConfig.LocalInstallations, version)
	return v.saveConfig()
}

func (v VersionManagerImpl) ListInstalledVersions() []models.LocalInstallationInfo {
	return utils.Values(v.localConfig.LocalInstallations)
}

func (v VersionManagerImpl) GetDefaultVersion() *models.LocalInstallationInfo {
	if v.localConfig.DefaultVersion == nil {
		return nil
	}
	tmp := v.localConfig.LocalInstallations[*v.localConfig.DefaultVersion]
	return &tmp
}

func (v VersionManagerImpl) SetDefaultVersion(version *string) error {
	if _, err := os.Lstat(config.DefaultVersionFile); err != nil {
		err := os.Remove(config.DefaultVersionFile)
		if err != nil {
			return err
		}
	}
	if version == nil {
		v.localConfig.DefaultVersion = nil
		return v.saveConfig()
	}

	versionToInstall, ok := v.localConfig.LocalInstallations[*version]
	if !ok {
		err := v.InstallVersion(*version)
		if err != nil {
			return err
		}
	}

	return os.Symlink(versionToInstall.Location, config.DefaultVersionFile)
}

func (v VersionManagerImpl) saveConfig() error {
	configAsBytes, err := json.Marshal(v.localConfig)
	if err != nil {
		return err
	}

	return os.WriteFile(config.File, configAsBytes, 0700)
}

func (v VersionManagerImpl) Run(version string, args []string) error {
	if !v.VersionIsInstalled(version) {
		err := v.InstallVersion(version)
		if err != nil {
			return err
		}
	}

	versionPath := v.localConfig.LocalInstallations[version].Location
	args = utils.Prepend(args, versionPath)
	err := syscall.Exec(args[0], args, os.Environ())
	if err != nil {
		return err
	}
	return nil
}

func (v VersionManagerImpl) VersionIsInstalled(version string) bool {
	_, ok := v.localConfig.LocalInstallations[version]
	return ok
}
