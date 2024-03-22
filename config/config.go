package config

import (
	"duckdb-version-manager/stacktrace"
	"duckdb-version-manager/utils"
	"log"
	"os"
)

var File string
var Dir string
var VersionDir string
var DefaultDuckdbFile string
var DuckmanBinaryFile string
var DuckDBName = "duckdb"

// Version Set using compile flags
var Version string

func init() {
	if Version == "" {
		utils.ExitWithError(stacktrace.New("Version not set using compile flags. Use -ldflags \"-X 'duckdb-version-manager/config.Version=1.0.0'\" to set the version."))
	}
	EnsureFoldersExist()
}

func EnsureFoldersExist() {
	homeDir, err := os.UserHomeDir()
	deviceInfo := utils.GetDeviceInfo()
	if err != nil {
		log.Fatalf("Error getting homeDir's home directory: %s", err)
	}
	Dir = homeDir + "/.config/duckman"
	File = Dir + "/config.json"
	VersionDir = Dir + "/versions"
	binaryDir := homeDir + "/.local/bin"
	DefaultDuckdbFile = binaryDir + "/" + DuckDBName
	DuckmanBinaryFile = binaryDir + "/duckman"

	if deviceInfo.Platform == "windows" {
		DefaultDuckdbFile += ".exe"
		DuckmanBinaryFile += ".exe"
	}

	// Ensure the directories exist
	err = os.MkdirAll(VersionDir, 0700)
	if err != nil {
		log.Fatalf("Error creating version directory: %s", err)
	}
	err = os.MkdirAll(binaryDir, 0700)
	if err != nil {
		log.Fatalf("Error creating install directory: %s", err)
	}
}
