package config

import (
	"github.com/spf13/pflag"
)

func (c *Config) installFlags() {
	pflag.StringVar(&c.ConfigPath, "config-path", defaultConfigPath, "Configuration file Path")
	pflag.BoolVarP(&c.Debug, "debug", "D", false, "Enable debug mode")
	pflag.StringVarP(&c.LogLevel, "log-level", "l", "info", `Set the logging level ("debug"|"info"|"warn"|"error"|"fatal")`)

	pflag.StringVar(&c.DatabaseAddress, "db-address", "localhost", "Specify database address")
	pflag.StringVar(&c.DatabasePort, "db-port", "3306", "Specify database port")
	pflag.StringVar(&c.DatabaseUser, "db-user", "admin", "Database username")
	pflag.StringVar(&c.DatabasePassword, "db-password", "", "Database password")

	pflag.StringVar(&c.BindAddress, "bind-address", "0.0.0.0", "Server bind address")
	pflag.StringVarP(&c.Port, "port", "p", "9090", "Port number to listen")
}
