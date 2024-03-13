package models

type LocalConfig struct {
	LocalInstallations LocalInstallations `json:"localInstallations"`
	DefaultVersion     *string            `json:"defaultVersion"`
}

type LocalInstallationId = string
type LocalInstallations = map[LocalInstallationId]LocalInstallationInfo

type LocalInstallationInfo struct {
	Version          string
	Location         string
	InstallationDate string
}
