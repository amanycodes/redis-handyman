package config

import "time"

// Conn holds connection and runtime settings. Populated by flags/env (viper) in root.
type Conn struct {
	Addr          string        `mapstructure:"addr"`
	Username      string        `mapstructure:"user"`
	Password      string        `mapstructure:"password"`
	DB            int           `mapstructure:"db"`
	TLS           bool          `mapstructure:"tls"`
	TLSServerName string        `mapstructure:"tls-server-name"`
	TLSSkipVerify bool          `mapstructure:"tls-skip-verify"`
	DialTimeout   time.Duration `mapstructure:"dial-timeout"`
	CmdTimeout    time.Duration `mapstructure:"timeout"`
	Verbose       bool          `mapstructure:"verbose"`
}
