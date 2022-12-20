package app

import (
	"gitee.com/cpds/cpds-analyzer/analyzer"
	"gitee.com/cpds/cpds-analyzer/config"
	"github.com/spf13/cobra"
)

func NewAnalyzer() (*cobra.Command, error) {
	conf := config.New()
	cmd := &cobra.Command{
		Use:                   "cpds-analyzer [OPTIONS]",
		Short:                 "Analyze exceptions for Container Problem Detect System",
		Version:               "undefined",
		SilenceUsage:          true,
		SilenceErrors:         true,
		DisableFlagsInUseLine: true,
		Args:                  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(conf)
		},
	}

	flags := cmd.Flags()
	conf.LoadConfig(flags)

	return cmd, nil
}

func Run(conf *config.Config) error {
	analyzer := analyzer.NewAnalyzer()
	return analyzer.Run(conf)
}
