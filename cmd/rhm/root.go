package rhm

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "rhm",
	Short: "redis-handyman: fast, safe Redis CLI",
	Long:  "redis-handyman (rhm): a developer/sysadmin-focused CLI to inspect Redis memory, TTL hygiene, and more.",
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags (env override with RHM_*).
	rootCmd.PersistentFlags().String("addr", "127.0.0.1:6379", "Redis address (host:port)")
	rootCmd.PersistentFlags().String("user", "", "Redis ACL username")
	rootCmd.PersistentFlags().String("password", "", "Redis password")
	rootCmd.PersistentFlags().Int("db", 0, "Redis logical DB")
	rootCmd.PersistentFlags().Bool("tls", false, "Enable TLS")
	rootCmd.PersistentFlags().Bool("tls-skip-verify", false, "Skip TLS cert verification (insecure)")
	rootCmd.PersistentFlags().Duration("timeout", 5*time.Second, "Per-command timeout")
	rootCmd.PersistentFlags().Bool("verbose", false, "Verbose logging")

	// Bind to viper with RHM_ prefix
	_ = viper.BindPFlag("addr", rootCmd.PersistentFlags().Lookup("addr"))
	_ = viper.BindPFlag("user", rootCmd.PersistentFlags().Lookup("user"))
	_ = viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	_ = viper.BindPFlag("db", rootCmd.PersistentFlags().Lookup("db"))
	_ = viper.BindPFlag("tls", rootCmd.PersistentFlags().Lookup("tls"))
	_ = viper.BindPFlag("tls_skip_verify", rootCmd.PersistentFlags().Lookup("tls-skip-verify"))
	_ = viper.BindPFlag("timeout", rootCmd.PersistentFlags().Lookup("timeout"))
	_ = viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

func initConfig() {
	viper.SetEnvPrefix("RHM")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	if viper.GetBool("verbose") {
		fmt.Fprintf(os.Stderr, "[rhm] addr=%s db=%d tls=%v timeout=%s\n",
			viper.GetString("addr"), viper.GetInt("db"), viper.GetBool("tls"), viper.GetDuration("timeout"))
	}
}

func Execute() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return rootCmd.ExecuteContext(ctx)
}
