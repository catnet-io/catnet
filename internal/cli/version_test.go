package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestVersionCmdOutput(t *testing.T) {
	t.Cleanup(func() {
		rootCmd.SetOut(nil)
		rootCmd.SetErr(nil)
		rootCmd.SetArgs(nil)
	})

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)
	rootCmd.SetArgs([]string{"version"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error executing version command: %v", err)
	}

	outStr := buf.String()
	if !strings.Contains(outStr, "engine/") {
		t.Errorf("Expected version output to contain 'engine/', got: %s", outStr)
	}
}
