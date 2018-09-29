package check

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Shows OK as the ping-pong result
// @Description Shows OK as the ping-pong result
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200 {string} plain "OK"
// @Router /ping [get]
func HealthCheck(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, message)
}
