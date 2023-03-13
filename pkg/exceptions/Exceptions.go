package exceptions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type exceptions struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

func Unauthorized(c *gin.Context) {
	e := exceptions{
		Status:  http.StatusConflict,
		Message: "you not have permission",
		Date:    time.Now(),
	}
	c.JSON(http.StatusConflict, e)
	c.Abort()
}

func BodyError(c *gin.Context) {
	e := exceptions{
		Status:  http.StatusBadRequest,
		Message: "bad request",
		Date:    time.Now(),
	}
	c.JSON(http.StatusBadRequest, e)
	c.Abort()
}

func FieldError(c *gin.Context, field string) {
	e := exceptions{
		Status:  http.StatusBadRequest,
		Message: fmt.Sprintf("bad field error: %s", field),
		Date:    time.Now(),
	}
	c.JSON(http.StatusBadRequest, e)
	c.Abort()
}

func InternalServer(c *gin.Context, field string) {
	e := exceptions{
		Status:  http.StatusInternalServerError,
		Message: fmt.Sprintf("Internal server error: %s", field),
		Date:    time.Now(),
	}
	c.JSON(http.StatusInternalServerError, e)
	c.Abort()
}

func Forbidden(c *gin.Context) {
	e := exceptions{
		Status:  http.StatusForbidden,
		Message: fmt.Sprintf("Forbidden"),
		Date:    time.Now(),
	}
	c.JSON(http.StatusForbidden, e)
	c.Abort()
}
