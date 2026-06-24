package integration_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var binaryPath string

func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "hello-world-test")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	binaryPath = filepath.Join(dir, "hello-world")
	cmd := exec.Command("go", "build", "-o", binaryPath, "./cmd/hello-world")
	cmd.Dir = findProjectRoot()
	if out, err := cmd.CombinedOutput(); err != nil {
		panic("failed to build binary: " + string(out))
	}

	os.Exit(m.Run())
}

func findProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			panic("could not find project root")
		}
		dir = parent
	}
}

func TestCLI_HappyPath(t *testing.T) {
	cmd := exec.Command(binaryPath, "World")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := strings.TrimRight(string(out), "\n")
	want := "hello World"
	if got != want {
		t.Errorf("output = %q, want %q", got, want)
	}
}

func TestCLI_MultiWordName(t *testing.T) {
	cmd := exec.Command(binaryPath, "Jane Doe")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := strings.TrimRight(string(out), "\n")
	want := "hello Jane Doe"
	if got != want {
		t.Errorf("output = %q, want %q", got, want)
	}
}

func TestCLI_MissingArgument(t *testing.T) {
	cmd := exec.Command(binaryPath)
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatal("expected non-zero exit code, got 0")
	}
	exitErr, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatalf("expected ExitError, got %T: %v", err, err)
	}
	if exitErr.ExitCode() != 1 {
		t.Errorf("exit code = %d, want 1", exitErr.ExitCode())
	}
	got := strings.TrimRight(string(out), "\n")
	want := "Usage: hello-world <name>"
	if got != want {
		t.Errorf("stderr = %q, want %q", got, want)
	}
}

func TestCLI_EmptyStringArgument(t *testing.T) {
	cmd := exec.Command(binaryPath, "")
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatal("expected non-zero exit code, got 0")
	}
	exitErr, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatalf("expected ExitError, got %T: %v", err, err)
	}
	if exitErr.ExitCode() != 1 {
		t.Errorf("exit code = %d, want 1", exitErr.ExitCode())
	}
	got := strings.TrimRight(string(out), "\n")
	want := "Usage: hello-world <name>"
	if got != want {
		t.Errorf("stderr = %q, want %q", got, want)
	}
}

func TestCLI_ExtraArguments(t *testing.T) {
	cmd := exec.Command(binaryPath, "Alice", "Bob")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := strings.TrimRight(string(out), "\n")
	want := "hello Alice"
	if got != want {
		t.Errorf("output = %q, want %q", got, want)
	}
}
