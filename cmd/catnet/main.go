package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/catnet-io/catnet/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		var exitErr *cli.ExitError
		if errors.As(err, &exitErr) {
			os.Exit(exitErr.Code)
		}
		os.Exit(1)
	}
}
