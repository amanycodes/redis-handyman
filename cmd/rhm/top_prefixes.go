package rhm

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "top-prefixes",
		Short: "Scan keys and summarize memory by key prefix (WIP)",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("top-prefixes: coming soon (Day 2+) 🚧")
		},
	}
	// Future flags (wired later):
	// cmd.Flags().String("pattern", "*", "Match pattern for SCAN")
	// cmd.Flags().String("delim", ":", "Key prefix delimiter")
	// cmd.Flags().Int("depth", 2, "Max prefix depth to display")
	// cmd.Flags().Int("limit", 50, "Number of rows to display")
	// cmd.Flags().String("output", "table", "table|json")
	// cmd.Flags().Int("batch", 100, "SCAN COUNT per iteration")
	// cmd.Flags().Int("rate", 0, "Max requests per second (0 = unlimited)")
	rootCmd.AddCommand(cmd)
}
