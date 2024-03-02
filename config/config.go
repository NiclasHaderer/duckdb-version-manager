package config

import (
	"log"
	"os"
)

var ConfigPath string

var VersionDir string

var InstallDir string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting homeDir's home directory: %s", err)
	}
	ConfigPath = homeDir + "/.config/duck-vm"
	VersionDir = ConfigPath + "/versions"
	InstallDir = homeDir + ".local/bin"
}
