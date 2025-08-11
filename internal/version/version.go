package rhm

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rhm",
	Short: "redis-handyman: fast, safe Redis CLI",
	Long:  "redis-handyman (rhm): a developer/sysadmin-focused CLI to inspect Redis memory, TTL hygiene, and more.",
}

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print version info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("rhm %s (commit %s, built %s)\n", version, commit, date)
		},
	})
}
