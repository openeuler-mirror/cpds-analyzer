package router

import (
	"cpds/cpds-analyzer/internal/handlers"
	"cpds/cpds-analyzer/internal/middlewares"
	dbinitiator "cpds/cpds-analyzer/internal/pkg/database"

	gormlogger "gorm.io/gorm/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type resource struct {
	logger *zap.Logger
	db     *gorm.DB
}

func InitRouter(debug bool, logger *zap.Logger, db *gorm.DB) *gin.Engine {
	r := &resource{
		logger: logger,
		db:     db,
	}

	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		db.Logger.LogMode(gormlogger.Silent)
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(middlewares.LoggerMiddleware(logger))

	// test route
	router.GET("/ping", handlers.GetPing)

	apiv1 := router.Group("/api/v1")
	setRulesRouter(apiv1, r)
	initDatabaseTable(db)

	return router
}

func initDatabaseTable(db *gorm.DB) error {
	d := dbinitiator.New(db)
	if err := d.Init(); err != nil {
		return err
	}
	return nil
}
