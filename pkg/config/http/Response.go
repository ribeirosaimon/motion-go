package http

import "github.com/gin-gonic/gin"

func Ok(c *gin.Context, body interface{}) {
	c.Status(200)
	c.JSON(200, body)
}

func Created(c *gin.Context, body interface{}) {
	c.Status(201)
	c.JSON(201, body)
}

func Forbidden(c *gin.Context) {
	c.Status(403)
	c.JSON(403, "You do not have permission")
}
