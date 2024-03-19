package manager

import (
	"duckdb-version-manager/api"
	"duckdb-version-manager/config"
	"duckdb-version-manager/models"
	"duckdb-version-manager/stacktrace"
	"duckdb-version-manager/utils"
	"encoding/json"
	"github.com/spf13/cobra"
	"os"
	"syscall"
	"time"
)

type VersionManager interface {
	InstallVersion(version string) stacktrace.Error
	UninstallVersion(version string) stacktrace.Error
	ListInstalledVersions() []models.LocalInstallationInfo
	GetDefaultVersion() *models.LocalInstallationInfo
	SetDefaultVersion(version *string) stacktrace.Error
	Run(version string, args []string) stacktrace.Error
	VersionIsInstalled(version string) bool
	GetLocalReleaseInfo(version string) (*models.LocalInstallationInfo, stacktrace.Error)
	LocalVersionList(cmd *cobra.Command, args []string, complete string) ([]string, cobra.ShellCompDirective)
	RemoteVersionList(cmd *cobra.Command, args []string, complete string) ([]string, cobra.ShellCompDirective)
	saveLocalConfig() stacktrace.Error
}

type versionManagerImpl struct {
	client      api.Client
	localConfig models.LocalConfig
}

func (v versionManagerImpl) saveLocalConfig() stacktrace.Error {
	//TODO implement me
	panic("implement me")
}

func (v versionManagerImpl) InstallVersion(version string) stacktrace.Error {
	release, err := v.client.GetRelease(version)
	if err != nil {
		return err
	}

	downloadUrl, err := utils.GetDownloadUrlFrom(release)
	if err != nil {
		return err
	}

	githubAsset, err := utils.GetResponseBodyFrom(v.client.Get(), *downloadUrl)
	duckDb, err := utils.ExtractDuckdbFile(githubAsset)
	if err != nil {
		return err
	}

	fileLocation := config.VersionDir + "/" + release.Version
	if err := os.WriteFile(fileLocation, duckDb, 0700); err != nil {
		return stacktrace.Wrap(err)
	}

	installTime, _ := time.Now().MarshalText()
	v.localConfig.LocalInstallations[release.Version] = models.LocalInstallationInfo{
		Version:          release.Version,
		Location:         fileLocation,
		InstallationDate: string(installTime),
	}

	return v.saveConfig()
}

func (v versionManagerImpl) UninstallVersion(unreliableVersion string) stacktrace.Error {
	if !v.VersionIsInstalled(unreliableVersion) {
		return stacktrace.NewF("Version '%s' not installed", unreliableVersion)
	}

	release, _ := v.GetLocalReleaseInfo(unreliableVersion)

	// Check if the version is the default version
	if v.localConfig.DefaultVersion == &release.Version {
		err := v.SetDefaultVersion(nil)
		if err != nil {
			return err
		}
	}

	if err := os.Remove(v.localConfig.LocalInstallations[release.Version].Location); err != nil {
		return stacktrace.Wrap(err)
	}
	delete(v.localConfig.LocalInstallations, release.Version)
	return v.saveConfig()
}

func (v versionManagerImpl) ListInstalledVersions() []models.LocalInstallationInfo {
	return utils.Values(v.localConfig.LocalInstallations)
}

func (v versionManagerImpl) GetDefaultVersion() *models.LocalInstallationInfo {
	if v.localConfig.DefaultVersion == nil {
		return nil
	}
	tmp := v.localConfig.LocalInstallations[*v.localConfig.DefaultVersion]
	return &tmp
}

func (v versionManagerImpl) SetDefaultVersion(version *string) stacktrace.Error {
	if _, err := os.Lstat(config.DefaultDuckdbFile); err == nil {
		err := os.Remove(config.DefaultDuckdbFile)
		if err != nil {
			return stacktrace.Wrap(err)
		}
	}
	if version == nil {
		v.localConfig.DefaultVersion = nil
		return v.saveConfig()
	}

	if !v.VersionIsInstalled(*version) {
		err := v.InstallVersion(*version)
		if err != nil {
			return err
		}
	}

	versionToInstall, _ := v.GetLocalReleaseInfo(*version)
	err := os.Symlink(versionToInstall.Location, config.DefaultDuckdbFile)
	if err != nil {
		return stacktrace.Wrap(err)
	}

	v.localConfig.DefaultVersion = &versionToInstall.Version
	return v.saveConfig()
}

func (v versionManagerImpl) saveConfig() stacktrace.Error {
	configAsBytes, err := json.MarshalIndent(v.localConfig, "", "  ")
	if err != nil {
		return stacktrace.Wrap(err)
	}

	err = os.WriteFile(config.File, configAsBytes, 0700)
	if err != nil {
		return stacktrace.Wrap(err)
	}
	return nil
}

func (v versionManagerImpl) Run(version string, args []string) stacktrace.Error {
	if !v.VersionIsInstalled(version) {
		err := v.InstallVersion(version)
		if err != nil {
			return err
		}
	}

	release, _ := v.GetLocalReleaseInfo(version)
	installationTime, _ := time.Parse(time.RFC3339, release.InstallationDate)
	isOlderThanOneDay := time.Now().Sub(installationTime) > 24*time.Hour
	if release.Version == "nightly" && isOlderThanOneDay {
		err := v.InstallVersion(version)
		if err != nil {
			return err
		}
	}

	args = utils.Prepend(args, release.Location)
	err := syscall.Exec(args[0], args, os.Environ())
	if err != nil {
		return stacktrace.Wrap(err)
	}
	return nil
}

func (v versionManagerImpl) VersionIsInstalled(version string) bool {
	_, ok := v.localConfig.LocalInstallations[version]

	if !ok {
		_, ok = v.localConfig.LocalInstallations["v"+version]
		return ok
	}

	return ok
}

func (v versionManagerImpl) GetLocalReleaseInfo(version string) (*models.LocalInstallationInfo, stacktrace.Error) {
	li, ok := v.localConfig.LocalInstallations[version]
	if !ok {
		li, ok = v.localConfig.LocalInstallations["v"+version]
		if !ok {
			return nil, stacktrace.NewF("Version '%s' not found in local versions", version)
		}
	}
	return &li, nil
}

func (v versionManagerImpl) LocalVersionList(cmd *cobra.Command, args []string, complete string) ([]string, cobra.ShellCompDirective) {
	installedVersions := v.ListInstalledVersions()
	versionList := utils.Map(installedVersions, func(li models.LocalInstallationInfo) string {
		return li.Version
	})

	utils.SortVersions(versionList, func(version string) string {
		return version
	})
	return versionList, cobra.ShellCompDirectiveKeepOrder
}

func (v versionManagerImpl) RemoteVersionList(cmd *cobra.Command, args []string, complete string) ([]string, cobra.ShellCompDirective) {
	remoteVersions, err := v.client.ListAllReleases()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	return toVersionList(remoteVersions), cobra.ShellCompDirectiveKeepOrder
}

func toVersionList(versions []models.VersionInfo) []string {
	versionList := utils.Map(versions, func(release models.VersionInfo) string {
		return release.Version
	})

	utils.SortVersions(versionList, func(version string) string {
		return version
	})

	return versionList
}
