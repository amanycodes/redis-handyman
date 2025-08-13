package rhm

import (
	"context"
	"fmt"
	"time"

	"github.com/amanycodes/redis-handyman/internal/rredis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// init registers the "ping" command to the root command.
// The "ping" command pings the Redis server and prints the response and elapsed time.
func init() {
	cmd := &cobra.Command{
		Use:   "ping",                  // Command name
		Short: "Ping the Redis server", // Short description
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get timeout duration from configuration
			timeout := viper.GetDuration("timeout")
			// Create a context with timeout for the operation
			ctx, cancel := context.WithTimeout(cmd.Context(), timeout)
			defer cancel()

			// Create a Redis client using configuration from Viper
			client, closer, err := rredis.NewClientFromViper(ctx)
			if err != nil {
				return err
			}
			defer closer() // Ensure client resources are cleaned up

			// Measure the time taken to ping Redis
			start := time.Now()
			res, err := client.Ping(ctx).Result()
			elapsed := time.Since(start)
			if err != nil {
				return err
			}
			// Print the ping response and elapsed time
			fmt.Printf("PING => %s (%s)\n", res, elapsed)
			return nil
		},
	}
	// Add the ping command to the root command
	rootCmd.AddCommand(cmd)
}
