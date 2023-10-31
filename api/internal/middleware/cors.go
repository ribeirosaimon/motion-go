package middleware

import "github.com/gin-gonic/gin"

func CorsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, MotionRole")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}

	c.Next()
}
