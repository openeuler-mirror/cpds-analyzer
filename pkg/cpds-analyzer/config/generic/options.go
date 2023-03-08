package generic

import (
	"cpds/cpds-analyzer/pkg/utils/net"
	"fmt"

	"github.com/spf13/pflag"
)

type Options struct {
	BindAddress string `json:"bindAddress,omitempty" yaml:"bindAddress,omitempty"`
	Port        int    `json:"port,omitempty" yaml:"port,omitempty"`
}

func NewGenericOptions() *Options {
	return &Options{
		BindAddress: "0.0.0.0",
		Port:        19091,
	}
}

func (s *Options) Validate() []error {
	errs := []error{}

	if !net.IsValidIPAdress(s.BindAddress) {
		errs = append(errs, fmt.Errorf("wrong IP Address format: %s", s.BindAddress))
	}

	if !net.IsValidPort(s.Port) {
		errs = append(errs, fmt.Errorf("invalid port number: %d, should be 0 - 65535", s.Port))
	}

	return errs
}

func (s *Options) AddFlags(fs *pflag.FlagSet, c *Options) {
	fs.StringVar(&s.BindAddress, "bind-address", c.BindAddress, "server bind address")
	fs.IntVar(&s.Port, "port", c.Port, "insecure port number")
}
