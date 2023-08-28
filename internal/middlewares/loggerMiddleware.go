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

package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			logger.Info(
				"gin",
				zap.String("request_id", param.Request.Header.Get("X-Request-Id")),
				zap.String("client_ip", param.ClientIP),
				zap.Int("status_code", param.StatusCode),
				zap.String("method", param.Method),
				zap.String("path", param.Path),
				zap.String("error_message", param.ErrorMessage),
			)
			return ""
		},
	})
}

func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        method := c.Request.Method
        origin := c.Request.Header.Get("Origin") 
        if origin != "" {
            c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
            c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
            c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
            c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
            c.Header("Access-Control-Max-Age", "172800")
            c.Header("Access-Control-Allow-Credentials", "true")
        }

        if method == "OPTIONS" {
            c.Header("Access-Control-Allow-Origin", "*")
            c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization") //自定义 Header
            c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
            c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
            c.Header("Access-Control-Allow-Credentials", "true")
            c.AbortWithStatus(http.StatusNoContent)
        }
 
        defer func() {
            if err := recover(); err != nil {
                log.Printf("Panic info is: %v", err)
            }
        }()
 
        c.Next()
    }
}