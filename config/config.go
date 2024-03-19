package config

import (
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

func init() {
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
