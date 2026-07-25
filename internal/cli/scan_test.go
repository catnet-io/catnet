package cli

import (
	"bytes"
	"context"
	"strings"
	"testing"
	"time"

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
		rootCmd.SetContext(nil)
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)
	rootCmd.SetContext(ctx)
	rootCmd.SetArgs([]string{"scan", "auto", "--no-ports", "--ping-timeout", "10"})

	_ = rootCmd.Execute()

	if !strings.Contains(buf.String(), "Auto-detected") {
		t.Errorf("Expected CLI stderr to contain 'Auto-detected', got: %s", buf.String())
	}
}

func TestScanCmdDefaultsToAuto(t *testing.T) {
	t.Cleanup(func() {
		rootCmd.SetOut(nil)
		rootCmd.SetErr(nil)
		rootCmd.SetArgs(nil)
		rootCmd.SetContext(nil)
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)
	rootCmd.SetContext(ctx)
	rootCmd.SetArgs([]string{"scan", "--no-ports", "--ping-timeout", "10"})

	_ = rootCmd.Execute()

	if !strings.Contains(buf.String(), "Auto-detected") {
		t.Errorf("Expected catnet scan without arguments to auto-detect targets, got: %s", buf.String())
	}
}
