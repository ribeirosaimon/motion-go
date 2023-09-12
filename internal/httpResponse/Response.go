package httpResponse

import (
	"github.com/gin-gonic/gin"
)

func Entity(c *gin.Context, httpStatus int, body interface{}) {
	if body == nil {
		c.Status(httpStatus)
	} else {
		c.JSON(httpStatus, body)
	}
}
