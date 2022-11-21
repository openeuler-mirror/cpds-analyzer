package analyzer

import (
	"gitee.com/cpds/cpds-analyzer/config"
	"github.com/sirupsen/logrus"
)

func RunAnalyzer(opts *config.Options) error {
	logrus.Infof("Starting cpds-analyzer......")
	logrus.Infof("Using config: database address: %s, database port: %s", opts.DatabaseAddress, opts.DatabasePort)
	// TODO: complete this function
	return nil
}
