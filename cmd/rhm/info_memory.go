package rhm

import (
	"context"
	"fmt"

	"github.com/amanycodes/redis-handyman/internal/rredis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	mem := &cobra.Command{
		Use:   "memory",
		Short: "Show INFO memory (raw Redis output)",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithTimeout(cmd.Context(), viper.GetDuration("timeout"))
			defer cancel()

			client, closer, err := rredis.NewClientFromViper(ctx)
			if err != nil {
				return err
			}
			defer closer()

			res, err := client.Info(ctx, "memory").Result()
			if err != nil {
				return err
			}
			fmt.Print(res)
			return nil
		},
	}
	infoCmd.AddCommand(mem)
}
