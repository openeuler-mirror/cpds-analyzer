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
