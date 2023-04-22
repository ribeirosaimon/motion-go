package httpresponse

import "github.com/gin-gonic/gin"

func Ok(c *gin.Context, body interface{}) {
	c.Status(200)
	if body != nil {
		c.JSON(200, body)
	}
}

func Created(c *gin.Context, body interface{}) {
	c.Status(201)
	if body != nil {
		c.JSON(201, body)
	}
}
