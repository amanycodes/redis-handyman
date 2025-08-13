package rhm

import (
	"context"
	"fmt"

	"github.com/amanycodes/redis-handyman/internal/rredis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	// Define the 'memory' subcommand for 'info'
	mem := &cobra.Command{
		Use:   "memory",                              // Command name
		Short: "Show INFO memory (raw Redis output)", // Short description
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get timeout duration from configuration and create a context with it
			ctx, cancel := context.WithTimeout(cmd.Context(), viper.GetDuration("timeout"))
			defer cancel()

			// Create a Redis client using configuration from Viper
			client, closer, err := rredis.NewClientFromViper(ctx)
			if err != nil {
				return err
			}
			// Ensure client resources are cleaned up
			defer closer()

			// Execute the Redis INFO memory command
			res, err := client.Info(ctx, "memory").Result()
			if err != nil {
				return err
			}
			// Print the raw INFO memory output
			fmt.Print(res)
			return nil
		},
	}
	// Add the 'memory' subcommand to the 'info' command
	infoCmd.AddCommand(mem)
}
