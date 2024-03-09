package utils

import (
	"archive/zip"
	"bytes"
	"duckdb-version-manager/models"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func getZipFileContent(file *zip.File) ([]byte, error) {
	fileReader, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fileReader.Close()

	return io.ReadAll(fileReader)
}

func saveZip(body []byte, dest string) error {
	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return err
	}

	for _, file := range zipReader.File {
		fileName := file.Name
		if fileName == "duckdb" {
			fileContent, err := getZipFileContent(file)
			if err != nil {
				return err
			}

			return os.WriteFile(dest, fileContent, 0700)
		}
		if strings.Contains(fileName, "duckdb_cli") && strings.HasSuffix(fileName, ".zip") {
			fileContent, err := getZipFileContent(file)
			if err != nil {
				return err
			}
			return saveZip(fileContent, dest)
		}
	}
	return errors.New("no duckdb binary found in zip")
}

func DownloadUrlTo(url string, dest string, isZip bool) error {
	httpClient := &http.Client{Timeout: 5 * time.Minute}
	resp, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if isZip {
		err = saveZip(body, dest)
	} else {
		err = os.WriteFile(dest, body, 0700)
	}

	if err != nil {
		return err
	}

	fmt.Printf("Downloaded %s to %s\n", url, dest)
	return nil
}

func GetDownloadUrlFrom(release *models.Release) (string, error) {
	sysInfo := GetDeviceInfo()

	platform, ok := release.Platforms[sysInfo.Platform]
	if !ok {
		log.Fatalf("Platform %s not supported. Supported architectures are: %s", sysInfo.Platform, strings.Join(Keys(release.Platforms), ", "))
	}

	arch, ok := platform[sysInfo.Architecture]
	if !ok {
		if res, ok := platform[models.ArchitectureUniversal]; ok {
			return res.DownloadUrl, nil
		}

		log.Fatalf("Architecture %s not supported. Supported architectures are: %s", sysInfo.Architecture, strings.Join(Keys(platform), ", "))
	}

	return arch.DownloadUrl, nil
}
