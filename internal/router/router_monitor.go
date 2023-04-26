package router

import (
	monitorHandler "cpds/cpds-analyzer/internal/handlers/monitor"

	"github.com/gin-gonic/gin"
)

func setMonitorRouter(api *gin.RouterGroup, r *resource) {
	monitorApi := api.Group("monitor")
	{
		handler := monitorHandler.New(r.config, r.logger)
		monitorApi.GET("/targets", handler.GetMonitorTargets())
		monitorApi.GET("/node_info", handler.GetNodeInfo())
		monitorApi.GET("/node_status", handler.GetNodeStatus())
		monitorApi.GET("/node_container_status", handler.GetNodeContainerStatus())
		monitorApi.GET("/node_resources", handler.GetNodeResource())
		monitorApi.GET("/cluster_resources", handler.GetClusterResource())
		monitorApi.GET("/cluster_container_status", handler.GetClusterContainerStatus())
	}
}
