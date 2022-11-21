package config

import (
	"github.com/spf13/pflag"
)

type Options struct {
	Debug           bool
	LogLevel        string
	DatabaseAddress string
	DatabasePort    string
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) InstallFlags(flags *pflag.FlagSet) {
	flags.BoolVarP(&o.Debug, "debug", "D", false, "Enable debug mode")
	flags.StringVarP(&o.LogLevel, "log-level", "l", "info", `Set the logging level ("debug"|"info"|"warn"|"error"|"fatal")`)

	flags.StringVar(&o.DatabaseAddress, "db_address", "localhost", "specify database address")
	flags.StringVar(&o.DatabasePort, "db_port", "3306", "specify database port")
}
