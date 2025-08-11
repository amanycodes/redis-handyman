package rhm

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "ttl-report",
		Short: "TTL hygiene summary by prefix (WIP)",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("ttl-report: coming soon (Day 2+) 🚧")
		},
	}
	rootCmd.AddCommand(cmd)
}
