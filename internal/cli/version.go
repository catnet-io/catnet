package cli

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/spf13/cobra"
)

// Variables set by GoReleaser ldflags at build time.
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)
var shortVersion bool

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		out := cmd.OutOrStdout()
		if shortVersion {
			fmt.Fprintln(out, Version)
			return
		}

		fmt.Fprintf(out, "catnet/%s (%s/%s) %s\n", Version, runtime.GOOS, runtime.GOARCH, runtime.Version())
		fmt.Fprintf(out, "commit: %s  built: %s\n", Commit, Date)

		coreVersion := "unknown"
		if buildInfo, ok := debug.ReadBuildInfo(); ok {
			for _, dep := range buildInfo.Deps {
				if dep.Path == "github.com/catnet-io/engine" {
					coreVersion = dep.Version
					break
				}
			}
		}
		fmt.Fprintf(out, "engine/%s\n", coreVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolVar(&shortVersion, "short", false, "Print only the version number")
}
