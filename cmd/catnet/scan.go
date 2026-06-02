package catnet

import (
	"context"
	"fmt"
	"os"

	"github.com/mendsec/catnet-core/pkg/events"
	"github.com/mendsec/catnet-core/pkg/export"
	"github.com/mendsec/catnet-core/pkg/profile"
	"github.com/mendsec/catnet-core/pkg/results"
	"github.com/mendsec/catnet-core/pkg/scan"
	"github.com/mendsec/catnet-core/pkg/targets"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(scanCmd)
}

var scanCmd = &cobra.Command{
	Use:   "scan [targets]",
	Short: "Run a network scan against specified targets",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		targetInput := args[0]
		ips, err := targets.ParseRange(targetInput)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse targets: %v\n", err)
			os.Exit(1)
		}

		engine := scan.NewEngine()
		prof := profile.DefaultProfile()
		eventChan := make(chan events.Event)

		var scannedHosts []results.HostResult

		go func() {
			for ev := range eventChan {
				switch ev.Type {
				case events.ScanStarted:
					fmt.Printf("[*] Scan started on %d targets\n", ev.Data)
				case events.HostDiscovered:
					data, ok := ev.Data.(events.HostDiscoveredData)
					if ok {
						scannedHosts = append(scannedHosts, data.Host)
						if data.Host.Alive {
							fmt.Printf("[+] Host UP: %s (MAC: %s) - Ports: %v\n", data.Host.IP, data.Host.MAC, data.Host.OpenPorts)
						}
					}
				case events.ScanCompleted:
					fmt.Println("[*] Scan completed")
				}
			}
		}()

		err = engine.ScanStream(context.Background(), ips, prof, eventChan)
		close(eventChan)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan failed: %v\n", err)
			os.Exit(1)
		}

		// Export results to stdout as JSON
		jsonBytes, _ := export.ExportJSON(scannedHosts)
		fmt.Println("\n--- Scan Results ---")
		fmt.Println(string(jsonBytes))
	},
}
