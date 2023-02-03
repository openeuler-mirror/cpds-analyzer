package analyzer

import (
	"fmt"
	"net/http"
	"time"

	"gitee.com/cpds/cpds-analyzer/analyzer/debug"
	"gitee.com/cpds/cpds-analyzer/config"
	commonv1 "gitee.com/cpds/cpds-analyzer/pkgs/apis/common/v1"
	rulesv1 "gitee.com/cpds/cpds-analyzer/pkgs/apis/rules/v1"
	"gitee.com/cpds/cpds-analyzer/pkgs/rules"
	restful "github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
)

var (
	serverTimeout = 5000 * time.Millisecond
)

type Analyzer struct {
	*config.Config
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

func (d *Analyzer) Run(conf *config.Config) error {
	if err := conf.CheckConfig(); err != nil {
		return err
	}

	if err := configureLogLevel(conf); err != nil {
		return err
	}

	if conf.Debug {
		debug.Enable()
		logrus.Debugf("enable debug mode")
	}

	logrus.Infof("starting cpds-analyzer......")
	logrus.Infof("using config: database address: %s, database port: %s", conf.DatabaseAddress, conf.DatabasePort)
	logrus.Infof("using config: bind address: %s, listening port: %s", conf.BindAddress, conf.Port)

	wsContainer := restful.NewContainer()
	logrus.Debug("creating new container")
	installAPIs(wsContainer)
	setRestfulConf(wsContainer)
	conf.RegisterSwagger(wsContainer)

	tlsconf := config.GetTlsConf()
	server := &http.Server{
		Addr:        ":" + conf.Port,
		Handler:     wsContainer,
		TLSConfig:   tlsconf,
		ReadTimeout: serverTimeout,
	}
	if err := server.ListenAndServeTLS(conf.CertFile, conf.KeyFile); err != nil {
		logrus.Infof("failed to listen https://%s:%s: %w", conf.BindAddress, conf.Port, err)
	}
	defer server.Close()

	return nil
}

// configureLogLevel "debug"|"info"|"warn"|"error"|"fatal", default: "info"
func configureLogLevel(conf *config.Config) error {
	logrus.Infof("configure Log Level: %s", conf.LogLevel)
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
	logrus.Debug("installing APIs")
	r := rules.New()
	commonv1.AddToContainer(c)
	rulesv1.AddToContainer(c, r)
}

func setRestfulConf(c *restful.Container) {
	logrus.Debug("setting restful configuration")
	// Add cross origin filter
	cors := config.GetCors(c)
	c.Filter(cors.Filter)

	// Add container filter to respond to OPTIONS
	c.Filter(c.OPTIONSFilter)
}
