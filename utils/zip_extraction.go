package utils

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"strings"
)

func getZipFileContent(file *zip.File) ([]byte, error) {
	fileReader, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fileReader.Close()

	return io.ReadAll(fileReader)
}

func ExtractDuckdbFile(assetBytes []byte) ([]byte, error) {
	zipReader, err := zip.NewReader(bytes.NewReader(assetBytes), int64(len(assetBytes)))
	if err != nil {
		return nil, err
	}
	for _, file := range zipReader.File {
		fileName := file.Name

		// Normal release -> just one file with the name "duckdb"
		if fileName == "duckdb" {
			fileContent, err := getZipFileContent(file)
			if err != nil {
				return nil, err
			}

			return fileContent, nil
		} else if strings.Contains(fileName, "duckdb_cli") && strings.HasSuffix(fileName, ".zip") {
			// Nightly release -> multiple zip files within the zip file
			fileContent, err := getZipFileContent(file)
			if err != nil {
				return nil, err
			}
			return ExtractDuckdbFile(fileContent)
		}
	}

	return nil, errors.New("no duckdb binary found in zip")
}
