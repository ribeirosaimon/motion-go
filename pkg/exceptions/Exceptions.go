package exceptions

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	notFound            int = 401
	badRequest          int = 400
	internalServererror int = 500
)

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

func BodyError(c *gin.Context) {
	e := exceptions{
		Status:  badRequest,
		Message: "bad request",
		Date:    time.Now(),
	}
	c.JSON(badRequest, e)
	c.Abort()
}

func FieldError(c *gin.Context, field string) {
	e := exceptions{
		Status:  badRequest,
		Message: fmt.Sprintf("bad field error: %s", field),
		Date:    time.Now(),
	}
	c.JSON(badRequest, e)
	c.Abort()
}

func InternalServer(c *gin.Context, field string) {
	e := exceptions{
		Status:  internalServererror,
		Message: fmt.Sprintf("Internal server error: %s", field),
		Date:    time.Now(),
	}
	c.JSON(badRequest, e)
	c.Abort()
}
