package router

import (
	"github.com/gin-gonic/gin"

	rulesHandler "cpds/cpds-analyzer/internal/handlers/rules"
)

func setRulesRouter(api *gin.RouterGroup, r *resource) {
	rulesApi := api.Group("rules")
	{
		rulesHandler := rulesHandler.New(r.logger, r.db)
		rulesApi.GET("", rulesHandler.Get())
		rulesApi.POST("/create", rulesHandler.Create())
		rulesApi.POST("/delete", rulesHandler.Delete())
		rulesApi.POST("/update", rulesHandler.Update())
	}
}
