package utils

import (
	"duckdb-version-manager/models"
	"duckdb-version-manager/stacktrace"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func RemoveFileOrDie(file string) {
	if _, err := os.Stat(file); err == nil {
		err := os.RemoveAll(file)
		if err != nil {
			ExitWithError(stacktrace.Wrap(err))
		}
	}
}

// Renames instead of truncating, so readers never see a partial file and running executables avoid ETXTBSY.
func WriteFileAtomic(path string, data []byte, perm os.FileMode) stacktrace.Error {
	tmp, err := os.CreateTemp(filepath.Dir(path), filepath.Base(path)+".*.tmp")
	if err != nil {
		return stacktrace.Wrap(err)
	}
	defer os.Remove(tmp.Name())

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return stacktrace.Wrap(err)
	}
	if err := tmp.Chmod(perm); err != nil {
		tmp.Close()
		return stacktrace.Wrap(err)
	}
	if err := tmp.Close(); err != nil {
		return stacktrace.Wrap(err)
	}
	return stacktrace.Wrap(os.Rename(tmp.Name(), path))
}

// Windows cannot replace a running .exe but can rename it, so it is moved into oldDir first. Elsewhere that step
// is skipped because it would briefly leave path missing.
func ReplaceExecutable(path string, data []byte, oldDir string) stacktrace.Error {
	if GetDeviceInfo().Platform == models.PlatformWindows {
		if err := moveAside(path, oldDir); err != nil {
			return err
		}
	}
	if err := WriteFileAtomic(path, data, 0700); err != nil {
		return err
	}
	RemoveOldExecutables(oldDir)
	return nil
}

func moveAside(path string, oldDir string) stacktrace.Error {
	target := filepath.Join(oldDir, fmt.Sprintf("%s.%d", filepath.Base(path), time.Now().UnixNano()))
	err := os.Rename(path, target)
	if errors.Is(err, fs.ErrNotExist) {
		return nil
	}
	return stacktrace.Wrap(err)
}

// Errors are ignored: binaries still running on Windows cannot be deleted yet and are retried on the next run.
func RemoveOldExecutables(oldDir string) {
	entries, _ := os.ReadDir(oldDir)
	for _, entry := range entries {
		_ = os.Remove(filepath.Join(oldDir, entry.Name()))
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
