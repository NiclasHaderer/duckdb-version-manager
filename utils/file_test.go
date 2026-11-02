package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestWriteFileAtomicReplacesRunningExecutable(t *testing.T) {
	sleep, err := exec.LookPath("sleep")
	if err != nil {
		t.Skip("sleep binary not available")
	}
	original, err := os.ReadFile(sleep)
	if err != nil {
		t.Fatal(err)
	}

	binary := filepath.Join(t.TempDir(), "binary")
	if err := os.WriteFile(binary, original, 0700); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(binary, "10")
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	defer cmd.Process.Kill()

	if err := WriteFileAtomic(binary, []byte("new"), 0700); err != nil {
		t.Fatal(err)
	}

	got, err := os.ReadFile(binary)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "new" {
		t.Fatalf("got %q, want %q", got, "new")
	}
	info, err := os.Stat(binary)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0700 {
		t.Fatalf("got mode %v, want 0700", info.Mode().Perm())
	}
}
