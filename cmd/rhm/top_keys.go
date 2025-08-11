package rhm

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "top-keys",
		Short: "Show heaviest keys by MEMORY USAGE (WIP)",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("top-keys: coming soon (Day 2+) 🚧")
		},
	}
	rootCmd.AddCommand(cmd)
}
