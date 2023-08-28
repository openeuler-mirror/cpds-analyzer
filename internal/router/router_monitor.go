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
