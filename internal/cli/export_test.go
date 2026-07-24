package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestExportCmdNoArguments(t *testing.T) {
	t.Cleanup(func() {
		rootCmd.SetOut(nil)
		rootCmd.SetErr(nil)
		rootCmd.SetArgs(nil)
	})

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)
	rootCmd.SetArgs([]string{"export"})

	err := rootCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error for missing export input argument, got nil")
	}

	if exitErr, ok := err.(*ExitError); ok {
		if exitErr.Code != ExitCodeInputError {
			t.Errorf("Expected ExitCodeInputError (%d), got %d", ExitCodeInputError, exitErr.Code)
		}
	} else {
		t.Errorf("Expected ExitError, got %T: %v", err, err)
	}

	if !strings.Contains(err.Error(), "requires exactly 1 input file argument") {
		t.Errorf("Expected error message to contain export guidance, got: %v", err)
	}
}
