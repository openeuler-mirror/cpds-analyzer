package analyzer

import (
	"context"
	"cpds/cpds-analyzer/internal/router"
	"cpds/cpds-analyzer/pkg/cpds-analyzer/config"
	"cpds/cpds-analyzer/pkg/logger"
	"cpds/cpds-analyzer/pkg/mariadb"
	timeutils "cpds/cpds-analyzer/pkg/utils/time"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Analyzer struct {
	Config *config.Config
	Logger *zap.Logger
	DB     *gorm.DB

	Debug bool
}

func (s *Analyzer) PrepareRun() error {
	var err error
	s.Logger, err = logger.NewLogger(
		logger.WithDisableConsole(),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileRotationP(
			s.Config.LoggerOptions.FileName,
			s.Config.LoggerOptions.MaxSize,
			s.Config.LoggerOptions.MaxBackups,
			s.Config.LoggerOptions.MaxAge,
		),
	)
	if err != nil {
		return err
	}

	dbLifeTime, err := timeutils.ParseDuration(s.Config.DatabaseOptions.MaxLifetime)
	if err != nil {
		return err
	}
	db := &mariadb.MariaDB{
		Host:        s.Config.DatabaseOptions.Host,
		Port:        s.Config.DatabaseOptions.Port,
		Username:    s.Config.DatabaseOptions.Username,
		Password:    s.Config.DatabaseOptions.Password,
		MaxOpenConn: s.Config.DatabaseOptions.MaxOpenConnections,
		MaxIdleConn: s.Config.DatabaseOptions.MaxIdleConnections,
		MaxLifetime: dbLifeTime,
	}
	s.DB, err = db.Connect()
	if err != nil {
		return err
	}

	return nil
}

func (s *Analyzer) Run() error {
	r := router.InitRouter(s.Debug, s.Logger, s.DB)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Config.GenericOptions.Port),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Logger.Info(fmt.Sprintf("Start listening on %d", s.Config.GenericOptions.Port))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	s.Logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		s.Logger.Fatal(fmt.Sprintf("Server Shutdown: %s", err))
	}
	log.Println()
	s.Logger.Info("Server exiting")

	return nil
}
