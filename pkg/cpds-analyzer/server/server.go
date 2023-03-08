package server

import (
	"cpds/cpds-analyzer/pkg/cpds-analyzer/config"
	"cpds/cpds-analyzer/pkg/cpds-analyzer/options"
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	s := options.NewServerRunOptions()

	cmd := &cobra.Command{
		Use:          "cpds-analyzer",
		Long:         `Analyze exceptions for Container Problem Detect System.`,
		Version:      "undefined",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if errs := s.Validate(); len(errs) != 0 {
				for _, v := range errs {
					panic(v)
				}
				// return utilerrors.NewAggregate(errs)
			}
			return Run(s)
		},
	}

	cobra.OnInitialize(func() {
		// Load configuration from file
		conf, err := config.TryLoadFromDisk(s.ConfigFile, s.DebugMode)
		if err == nil {
			s = &options.ServerRunOptions{
				Config:    conf,
				DebugMode: s.DebugMode,
			}
		} else {
			// TODO
			panic(fmt.Errorf("failed to load configuration from disk: %s", err))
		}
	})

	flags := cmd.Flags()
	flags.AddFlagSet(s.Flags())

	// usageFmt := "Usage:\n  %s\n"
	// cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	// cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
	// 	fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
	// 	cliflag.PrintSections(cmd.OutOrStdout(), namedFlagSets, cols)
	// })

	return cmd
}

func Run(s *options.ServerRunOptions) error {
	analyzer, err := s.NewAnalyzer()
	if err != nil {
		return err
	}

	err = analyzer.PrepareRun()
	if err != nil {
		return err
	}

	return analyzer.Run()
}
