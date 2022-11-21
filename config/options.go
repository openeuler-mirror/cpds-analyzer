package config

import (
	"github.com/spf13/pflag"
)

type Options struct {
	Debug            bool
	LogLevel         string
	DatabaseAddress  string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	BindAddress      string
	Port             string
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) InstallFlags(flags *pflag.FlagSet) {
	flags.BoolVarP(&o.Debug, "debug", "D", false, "Enable debug mode")
	flags.StringVarP(&o.LogLevel, "log-level", "l", "info", `Set the logging level ("debug"|"info"|"warn"|"error"|"fatal")`)

	flags.StringVar(&o.DatabaseAddress, "db-address", "localhost", "Specify database address")
	flags.StringVar(&o.DatabasePort, "db-port", "3306", "Specify database port")
	flags.StringVar(&o.DatabaseUser, "db-user", "admin", "Database username")
	flags.StringVar(&o.DatabasePassword, "db-password", "", "Database password")

	flags.StringVar(&o.BindAddress, "bind-address", "0.0.0.0", "Server bind address")
	flags.StringVarP(&o.Port, "port", "p", "9090", "Port number to listen")
}
