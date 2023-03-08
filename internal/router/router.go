package router

import (
	"cpds/cpds-analyzer/internal/handlers"
	"cpds/cpds-analyzer/internal/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRouter(debug bool, logger *zap.Logger) *gin.Engine {
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(middlewares.LoggerMiddleware(logger))

	// test route
	router.GET("/ping", handlers.GetPing)

	return router
}
