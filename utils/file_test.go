package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

const sleepEnv = "DUCKMAN_TEST_SLEEP"

// Lets the test binary serve as a running executable on every platform.
func TestMain(m *testing.M) {
	if os.Getenv(sleepEnv) == "1" {
		time.Sleep(time.Minute)
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func TestReplaceExecutableWhileRunning(t *testing.T) {
	self, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	original, err := os.ReadFile(self)
	if err != nil {
		t.Fatal(err)
	}

	dir := t.TempDir()
	binary := filepath.Join(dir, "binary.exe")
	oldDir := filepath.Join(dir, "old")
	if err := os.Mkdir(oldDir, 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(binary, original, 0700); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command(binary)
	cmd.Env = append(os.Environ(), sleepEnv+"=1")
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	stopped := false
	stop := func() {
		if !stopped {
			stopped = true
			_ = cmd.Process.Kill()
			_ = cmd.Wait()
		}
	}
	defer stop()

	if err := ReplaceExecutable(binary, []byte("new"), oldDir); err != nil {
		t.Fatal(err)
	}

	got, err := os.ReadFile(binary)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "new" {
		t.Fatalf("got %q, want %q", got, "new")
	}
	if runtime.GOOS != "windows" {
		info, err := os.Stat(binary)
		if err != nil {
			t.Fatal(err)
		}
		if info.Mode().Perm() != 0700 {
			t.Fatalf("got mode %v, want 0700", info.Mode().Perm())
		}
	}

	wantOld := 0
	if runtime.GOOS == "windows" {
		wantOld = 1
	}
	assertEntries(t, oldDir, wantOld)

	RemoveOldExecutables(oldDir)
	assertEntries(t, oldDir, wantOld)

	stop()
	RemoveOldExecutables(oldDir)
	assertEntries(t, oldDir, 0)
}

func assertEntries(t *testing.T, dir string, want int) {
	t.Helper()
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != want {
		t.Fatalf("got %d entries in %s, want %d", len(entries), dir, want)
	}
}
