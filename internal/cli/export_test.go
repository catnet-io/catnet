package cli

import (
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
		outPath := filepath.Join(tmpDir, "out_dotdot.json")

		rootCmd.SetArgs([]string{"export", dotDotPath, "--format", "json", "--output", outPath})

		err := rootCmd.Execute()
		if err != nil {
			t.Fatalf("Expected success for normalized path with '..', got: %v", err)
		}

		cleanOutPath := filepath.Clean(outPath)
		// nosemgrep: go.lang.security.audit.path-traversal
		outBytes, err := os.ReadFile(cleanOutPath) // #nosec G304
		if err != nil {
			t.Fatalf("Failed to read output json: %v", err)
		}
		if !strings.Contains(string(outBytes), "schemaVersion") {
			t.Errorf("Expected output to contain 'schemaVersion', got: %s", string(outBytes))
		}
	})

	t.Run("valid absolute path export csv", func(t *testing.T) {
		absPath, err := filepath.Abs(jsonPath)
		if err != nil {
			t.Fatalf("Failed to get abs path: %v", err)
		}

		outPath := filepath.Join(tmpDir, "output.csv")
		rootCmd.SetArgs([]string{"export", absPath, "--format", "csv", "--output", outPath})

		err = rootCmd.Execute()
		if err != nil {
			t.Fatalf("Expected success for absolute path export, got: %v", err)
		}

		cleanOutCSV := filepath.Clean(outPath)
		// nosemgrep: go.lang.security.audit.path-traversal
		outBytes, err := os.ReadFile(cleanOutCSV) // #nosec G304
		if err != nil {
			t.Fatalf("Failed to read output csv: %v", err)
		}
		if !strings.Contains(string(outBytes), "127.0.0.1") {
			t.Errorf("Expected CSV output to contain IP, got: %s", string(outBytes))
		}
	})

	t.Run("non-existent file path returns exit code input error", func(t *testing.T) {
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
