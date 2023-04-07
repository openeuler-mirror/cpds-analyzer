package router

import (
	"github.com/gin-gonic/gin"

	analysisHandler "cpds/cpds-analyzer/internal/handlers/analysis"
)

func setAnalysisRouter(api *gin.RouterGroup, r *resource) {
	rulesApi := api.Group("analysis")
	{
		analysisHandler := analysisHandler.New(r.logger, r.db, r.config)
		rulesApi.GET("/result", analysisHandler.GetResult())
		rulesApi.POST("/result/delete", analysisHandler.DeleteResult())
		rulesApi.GET("/result/raw_data", analysisHandler.GetRawData())
	}
}
