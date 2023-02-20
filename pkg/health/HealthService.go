package health

import (
	"github.com/gin-gonic/gin"
)

func getHealthService(c *gin.Context) {
	c.String(200, "Olá, mundo!")
}
