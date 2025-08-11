package rhm

import "github.com/spf13/cobra"

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Redis INFO helpers",
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
