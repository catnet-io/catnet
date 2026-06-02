package catnet

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of CatNet",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CatNet CLI v0.1.0")
	},
}
