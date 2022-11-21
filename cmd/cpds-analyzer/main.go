package main

import (
	"os"

	"gitee.com/cpds/cpds-analyzer/analyzer"
	"gitee.com/cpds/cpds-analyzer/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newAnalyzer() (*cobra.Command, error) {
	opts := config.NewOptions()
	cmd := &cobra.Command{
		Use:                   "cpds-analyzer [OPTIONS]",
		Short:                 "Analyze exceptions for Container Problem Detect System",
		Version:               "undefined",
		SilenceUsage:          true,
		SilenceErrors:         true,
		DisableFlagsInUseLine: true,
		Args:                  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return analyzer.RunAnalyzer(opts)
		},
	}
	flags := cmd.Flags()
	flags.BoolP("version", "v", false, "Print version information and quit")
	opts.InstallFlags(flags)

	return cmd, nil
}

func initLogging() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	logrus.SetOutput(os.Stdout)
}

func main() {
	initLogging()

	cmd, err := newAnalyzer()
	if err != nil {
		logrus.Error(err)
		// if cannot create new Analyzer, just exit
		os.Exit(1)
	}
	if err := cmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
