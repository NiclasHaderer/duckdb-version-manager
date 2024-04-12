package utils

import (
	"duckdb-version-manager/stacktrace"
	"io"
	"os"
)

func RemoveFileOrDie(file string) {
	if _, err := os.Stat(file); err == nil {
		err := os.RemoveAll(file)
		if err != nil {
			ExitWithError(stacktrace.Wrap(err))
		}
	}
}

func CopyFile(src, dst string) stacktrace.Error {
	srcFile, err := os.Open(src)
	if err != nil {
		return stacktrace.Wrap(err)
	}

	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return stacktrace.Wrap(err)
	}

	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return stacktrace.Wrap(err)
	}

	return nil
}
