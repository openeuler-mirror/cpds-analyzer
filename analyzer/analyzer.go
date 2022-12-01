package analyzer

import (
	"fmt"
	"net/http"
	"os"

	"gitee.com/cpds/cpds-analyzer/config"
	rulesv1 "gitee.com/cpds/cpds-analyzer/pkgs/apis/rules/v1"
	"gitee.com/cpds/cpds-analyzer/pkgs/rules"
	restful "github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
)

func RunAnalyzer(conf *config.Config) error {
	if err := configureLogLevel(conf); err != nil {
		return err
	}

	if conf.Debug {
		enableDebug()
		logrus.Debugf("Enable debug mode")
	}

	logrus.Infof("Starting cpds-analyzer......")
	logrus.Infof("Using config: database address: %s, database port: %s", conf.DatabaseAddress, conf.DatabasePort)
	logrus.Infof("Using config: bind address: %s, listening port: %s", conf.BindAddress, conf.Port)

	wsContainer := restful.NewContainer()
	installAPIs(wsContainer)

	// Add container filter to respond to OPTIONS
	wsContainer.Filter(wsContainer.OPTIONSFilter)

	server := &http.Server{
		Addr:    ":" + conf.Port,
		Handler: wsContainer,
	}
	if err := server.ListenAndServeTLS(conf.CertFile, conf.KeyFile); err != nil {
		logrus.Infof("Failed to listen https://%s:%s: %w", conf.BindAddress, conf.Port, err)
	}
	defer server.Close()

	return nil
}

// enableDebug sets the DEBUG env var to true
// and makes the logger to log at debug level.
func enableDebug() {
	os.Setenv("DEBUG", "1")
	logrus.SetLevel(logrus.DebugLevel)
}

// disableDebug sets the DEBUG env var to false
// and makes the logger to log at info level.
func disableDebug() {
	os.Setenv("DEBUG", "")
	logrus.SetLevel(logrus.InfoLevel)
}

// isDebugEnabled checks whether the debug flag is set or not.
func isDebugEnabled() bool {
	return os.Getenv("DEBUG") != ""
}

// configureLogLevel "debug"|"info"|"warn"|"error"|"fatal", default: "info"
func configureLogLevel(conf *config.Config) error {
	if conf.LogLevel != "" {
		lvl, err := logrus.ParseLevel(conf.LogLevel)
		if err != nil {
			return fmt.Errorf("unable to parse logging level: %s", conf.LogLevel)
		}
		logrus.SetLevel(lvl)
	} else {
		// Set InfoLevel as default logLevel
		// Only log the info severity or above.
		logrus.SetLevel(logrus.InfoLevel)
	}
	return nil
}

func installAPIs(c *restful.Container) {
	r := rules.New()
	rulesv1.AddToContainer(c, r)
}
