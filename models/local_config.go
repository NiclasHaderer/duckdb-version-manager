package models

type LocalConfig struct {
	LocalInstallations LocalInstallations `json:"localInstallations"`
	DefaultVersion     *string            `json:"defaultVersion"`
	Version            string             `json:"version"`
}

type LocalInstallationId = string
type LocalInstallations = map[LocalInstallationId]LocalInstallationInfo

type LocalInstallationInfo struct {
	Version          string `json:"version"`
	Location         string `json:"location"`
	InstallationDate string `json:"installationDate"`
}
