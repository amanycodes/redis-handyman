package rhm

import (
	"context"
	"fmt"
	"time"

	"github.com/amanycodes/redis-handyman/internal/rredis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cmd := &cobra.Command{
		Use:   "ping",
		Short: "Ping the Redis server",
		RunE: func(cmd *cobra.Command, args []string) error {
			timeout := viper.GetDuration("timeout")
			ctx, cancel := context.WithTimeout(cmd.Context(), timeout)
			defer cancel()

			client, closer, err := rredis.NewClientFromViper(ctx)
			if err != nil {
				return err
			}
			defer closer()

			start := time.Now()
			res, err := client.Ping(ctx).Result()
			elapsed := time.Since(start)
			if err != nil {
				return err
			}
			fmt.Printf("PING => %s (%s)\n", res, elapsed)
			return nil
		},
	}
	rootCmd.AddCommand(cmd)
}
