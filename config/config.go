package config

import (
	"log"
	"os"
)

var path string
var File string
var VersionDir string
var InstallDir string
var DefaultVersionFile string
var DuckVMName = "duck-vm"

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting homeDir's home directory: %s", err)
	}
	path = homeDir + "/.config/duck-vm"
	File = path + "/config.json"
	VersionDir = path + "/versions"
	InstallDir = homeDir + "/.local/bin"
	DefaultVersionFile = InstallDir + "/duckdb"

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
