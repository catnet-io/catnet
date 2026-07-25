package cli

import (
	"bytes"
	"testing"

	"github.com/catnet-io/engine/pkg/targets"
)

func TestParseRangeAuto(t *testing.T) {
	ips, err := targets.ParseRange("auto")
	if err != nil {
		t.Fatalf("ParseRange('auto') returned unexpected error: %v", err)
	}
	if len(ips) == 0 {
		t.Errorf("Expected ParseRange('auto') to return at least one IP address")
	}
}

func TestScanCmdAutoTargetFeedback(t *testing.T) {
	t.Cleanup(func() {
		rootCmd.SetOut(nil)
		rootCmd.SetErr(nil)
		rootCmd.SetArgs(nil)
	})

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)
	rootCmd.SetArgs([]string{"scan", "auto", "--no-ports", "--ping-timeout", "10"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Expected scan auto to succeed, got %v", err)
	}

	if !bytes.Contains(buf.Bytes(), []byte("Auto-detected")) {
		t.Errorf("Expected CLI stderr to contain 'Auto-detected', got: %s", buf.String())
	}
}

func TestScanCmdDefaultsToAuto(t *testing.T) {
	t.Cleanup(func() {
		rootCmd.SetOut(nil)
		rootCmd.SetErr(nil)
		rootCmd.SetArgs(nil)
	})

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)
	rootCmd.SetArgs([]string{"scan", "--no-ports", "--ping-timeout", "10"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Expected catnet scan without arguments to default to auto and succeed, got %v", err)
	}
}
