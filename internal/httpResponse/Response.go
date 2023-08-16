package httpResponse

import (
	"github.com/gin-gonic/gin"
)

func Entity(c *gin.Context, httpStatus int, body interface{}) {
	c.JSON(httpStatus, body)
}
