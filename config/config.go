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
	if deviceInfo.Platform == "windows" {
		DefaultDuckdbFile = binaryDir + "/duckdb" + ".exe"
	} else {
		DefaultDuckdbFile = binaryDir + "/duckdb"
	}
	DuckmanBinaryFile = binaryDir + "/duckman.exe"

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
