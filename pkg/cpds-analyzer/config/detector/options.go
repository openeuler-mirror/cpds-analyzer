package detector

import (
	"cpds/cpds-analyzer/pkg/utils/net"
	"fmt"

	"github.com/spf13/pflag"
)

type Options struct {
	Host string `json:"host,omitempty" yaml:"host,omitempty"`
	Port int    `json:"port,omitempty" yaml:"port,omitempty"`
}

func NewDetectorOptions() *Options {
	return &Options{
		Host: "127.0.0.1",
		Port: 19092,
	}
}

func (s *Options) Validate() []error {
	errs := []error{}

	if !net.IsValidIPAdress(s.Host) {
		errs = append(errs, fmt.Errorf("wrong IP Address format: %s", s.Host))
	}

	if !net.IsValidPort(s.Port) {
		errs = append(errs, fmt.Errorf("invalid port number range: %d, should be 0 - 65535", s.Port))
	}

	return errs
}

func (s *Options) AddFlags(fs *pflag.FlagSet, c *Options) {
	fs.StringVar(&s.Host, "detector-host", c.Host, "detector host IP address")
	fs.IntVar(&s.Port, "detector-port", c.Port, "detector port number")
}
