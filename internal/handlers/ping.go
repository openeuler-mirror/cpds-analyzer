package handlers

import (
	"cpds/cpds-analyzer/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	r := model.GetPingResult()
	c.JSON(http.StatusOK, r)
}
