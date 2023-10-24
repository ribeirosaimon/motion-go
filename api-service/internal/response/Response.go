package response

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/internal/middleware"
)

func Entity(c *gin.Context, httpStatus int, body interface{}) {
	modify := middleware.GetCache().NextModify
	if modify != nil {
		c.Writer.Header().Set("Next-Modify", strconv.FormatInt(*modify, 10))
	}

	if body == nil {
		c.Status(httpStatus)
	} else {
		c.JSON(httpStatus, body)
	}
}
