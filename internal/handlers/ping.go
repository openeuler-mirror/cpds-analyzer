package handlers

import (
	"cpds/cpds-analyzer/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	r := models.GetPingResult()
	c.JSON(http.StatusOK, r)
}
