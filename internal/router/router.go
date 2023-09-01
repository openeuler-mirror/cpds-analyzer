/*
 *  Copyright 2023 CPDS Author
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       https://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package router

import (
	"cpds/cpds-analyzer/internal/handlers"
	"cpds/cpds-analyzer/internal/middlewares"
	dbinitiator "cpds/cpds-analyzer/internal/pkg/database"
	"cpds/cpds-analyzer/pkg/cpds-analyzer/config"

	gormlogger "gorm.io/gorm/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type resource struct {
	config *config.Config
	logger *zap.Logger
	db     *gorm.DB
}

func InitRouter(debug bool, config *config.Config, logger *zap.Logger, db *gorm.DB) *gin.Engine {
	r := &resource{
		config: config,
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
	router.Use(middlewares.Cors(),middlewares.LoggerMiddleware(logger))

	// test route
	router.GET("/ping", handlers.GetPing)

	apiv1 := router.Group("/api/v1")
	{
		setRulesRouter(apiv1, r)
		setAnalysisRouter(apiv1, r)
		setMonitorRouter(apiv1, r)
		setPrometheusRouter(apiv1, r)
	}

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
