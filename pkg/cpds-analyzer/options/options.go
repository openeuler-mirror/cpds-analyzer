package options

import (
	analyzer "cpds/cpds-analyzer/pkg/cpds-analyzer"
	"cpds/cpds-analyzer/pkg/cpds-analyzer/config"

	"github.com/spf13/pflag"
)

type ServerRunOptions struct {
	ConfigFile string
	*config.Config

	DebugMode bool
}

func NewServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{
		Config: config.New(),
	}
}

func (s *ServerRunOptions) Flags() *pflag.FlagSet {
	fs := pflag.NewFlagSet("generic", pflag.ExitOnError)
	fs.StringVar(&s.ConfigFile, "config", config.DefaultConfigurationPath, "Directory where configuration files are stored")
	fs.BoolVar(&s.DebugMode, "debug", false, "Don't enable this if you don't know what it means.")

	return fs
}

func (s *ServerRunOptions) NewAnalyzer() (*analyzer.Analyzer, error) {
	analyzer := &analyzer.Analyzer{
		Config: s.Config,
		Debug:  s.DebugMode,
	}
	return analyzer, nil
}
