package exceptions

import (
	"time"

	"github.com/gin-gonic/gin"
)

const notFound int = 401

type exceptions struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

func Unauthorized(c *gin.Context) {
	e := exceptions{
		Status:  notFound,
		Message: "not found",
		Date:    time.Now(),
	}
	c.JSON(notFound, e)
	c.Abort()
}
