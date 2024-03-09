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
	InstallDir = homeDir + "/.local/bin"

	// Ensure the directories exist
	err = os.MkdirAll(VersionDir, 0700)
	if err != nil {
		log.Fatalf("Error creating version directory: %s", err)
	}
	err = os.MkdirAll(InstallDir, 0700)
	if err != nil {
		log.Fatalf("Error creating install directory: %s", err)
	}
}
