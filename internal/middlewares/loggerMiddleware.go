package middlewares

import (
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
