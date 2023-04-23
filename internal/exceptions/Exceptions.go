package exceptions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

func (e Error) Error() string {
	return e.Message
}

func Unauthorized() *Error {
	return &Error{
		Status:  http.StatusConflict,
		Message: "you not have permission",
		Date:    time.Now(),
	}

}

func BodyError() *Error {
	return &Error{
		Status:  http.StatusBadRequest,
		Message: "bad request",
		Date:    time.Now(),
	}
}

func FieldError(field string) *Error {
	return &Error{
		Status:  http.StatusBadRequest,
		Message: fmt.Sprintf("bad field error: %s", field),
		Date:    time.Now(),
	}
}

func InternalServer(field string) *Error {
	return &Error{
		Status:  http.StatusInternalServerError,
		Message: fmt.Sprintf("Internal server error: %s", field),
		Date:    time.Now(),
	}
}

func NotFound() *Error {
	return &Error{
		Status:  http.StatusConflict,
		Message: fmt.Sprintf("Not Found"),
		Date:    time.Now(),
	}
}

func MotionError(err string) *Error {
	return &Error{
		Status:  http.StatusConflict,
		Message: fmt.Sprintf(err),
		Date:    time.Now(),
	}
}
func Forbidden() *Error {
	return &Error{
		Status:  http.StatusForbidden,
		Message: fmt.Sprintf("Forbidden"),
		Date:    time.Now(),
	}
}

func (e Error) Throw(c *gin.Context) {
	c.JSON(e.Status, e)
	c.Abort()
}
