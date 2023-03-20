package motionHttp

import "github.com/gin-gonic/gin"

func Ok(c *gin.Context, body interface{}) {
	c.Status(200)
	c.JSON(200, body)
}

func Created(c *gin.Context, body interface{}) {
	c.Status(201)
	c.JSON(201, body)
}
