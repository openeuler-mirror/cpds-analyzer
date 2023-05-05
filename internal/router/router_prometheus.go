package router

import (
	"github.com/gin-gonic/gin"

	prometheusHandler "cpds/cpds-analyzer/internal/handlers/prometheus"
)

func setPrometheusRouter(api *gin.RouterGroup, r *resource) {
	rulesApi := api.Group("prometheus")
	{
		prometheusHandler := prometheusHandler.New(r.logger, r.config)
		rulesApi.GET("/query", prometheusHandler.Query())
		rulesApi.GET("/query_range", prometheusHandler.QueryRange())
		rulesApi.GET("/query_validate", prometheusHandler.QueryValidate())
	}
}
