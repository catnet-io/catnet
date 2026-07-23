package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExportCmdPathHandling(t *testing.T) {
	tmpDir := t.TempDir()

	// Prepare valid input file
	jsonPath := filepath.Join(tmpDir, "test_scan.json")
	validJSON := `{"schemaVersion":"2.0.0","devices":[{"ip":"127.0.0.1","status":"Alive"}]}`
	if err := os.WriteFile(jsonPath, []byte(validJSON), 0600); err != nil {
		t.Fatalf("Failed to write test input: %v", err)
	}

	t.Run("valid path with dotdot normalization", func(t *testing.T) {
		subDir := filepath.Join(tmpDir, "sub")
		if err := os.MkdirAll(subDir, 0700); err != nil {
			t.Fatalf("Failed to create subdir: %v", err)
		}
		dotDotPath := filepath.Join(subDir, "..", "test_scan.json")

		var buf bytes.Buffer
		rootCmd.SetOut(&buf)
		rootCmd.SetErr(&buf)
		rootCmd.SetArgs([]string{"export", dotDotPath, "--format", "json"})

		err := rootCmd.Execute()
		if err != nil {
			t.Fatalf("Expected success for normalized path with '..', got: %v", err)
		}

		if !strings.Contains(buf.String(), "schemaVersion") {
			t.Errorf("Expected output to contain 'schemaVersion', got: %s", buf.String())
		}
	})

	t.Run("valid absolute path export csv", func(t *testing.T) {
		absPath, err := filepath.Abs(jsonPath)
		if err != nil {
			t.Fatalf("Failed to get abs path: %v", err)
		}

		var buf bytes.Buffer
		rootCmd.SetOut(&buf)
		rootCmd.SetErr(&buf)
		rootCmd.SetArgs([]string{"export", absPath, "--format", "csv"})

		err = rootCmd.Execute()
		if err != nil {
			t.Fatalf("Expected success for absolute path export, got: %v", err)
		}

		if !strings.Contains(buf.String(), "127.0.0.1") {
			t.Errorf("Expected CSV output to contain IP, got: %s", buf.String())
		}
	})

	t.Run("non-existent file path returns exit code input error", func(t *testing.T) {
		var buf bytes.Buffer
		rootCmd.SetOut(&buf)
		rootCmd.SetErr(&buf)
		rootCmd.SetArgs([]string{"export", filepath.Join(tmpDir, "nonexistent.json"), "--format", "json"})

		err := rootCmd.Execute()
		if err == nil {
			t.Fatalf("Expected error for non-existent file, got nil")
		}
		if exitErr, ok := err.(*ExitError); ok {
			if exitErr.Code != ExitCodeInputError {
				t.Errorf("Expected ExitCodeInputError (%d), got %d", ExitCodeInputError, exitErr.Code)
			}
		} else {
			t.Errorf("Expected ExitError, got %T: %v", err, err)
		}
	})
}
