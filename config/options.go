package config

import "github.com/spf13/pflag"

type Options struct {
	DatabaseAddress string
	DatabasePort    string
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) InstallFlags(flags *pflag.FlagSet) {
	flags.StringVar(&o.DatabaseAddress, "db_address", "localhost", "specify database address")
	flags.StringVar(&o.DatabasePort, "db_port", "3306", "specify database port")
}
