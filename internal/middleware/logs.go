package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type motionLog struct {
	Api          string      `json:"api"`
	HttpRequest  request     `json:"httpRequest"`
	HttpResponse response    `json:"httpResponse"`
	LoggedUser   *LoggedUser `json:"loggedUser,omitempty"`
	Timestamp    time.Time   `json:"timestamp"`
}

type request struct {
	Proto         string `json:"proto"`
	RemoteIP      string `json:"remoteIP"`
	RequestMethod string `json:"requestMethod"`
}

type response struct {
	Byte    int           `json:"byte"`
	Status  int           `json:"status"`
	Latency time.Duration `json:"latency"`
}

func NewLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		var httpRequest = request{
			Proto:         c.Request.Proto,
			RemoteIP:      c.Request.RemoteAddr,
			RequestMethod: c.Request.Method,
		}

		end := time.Now()
		latency := end.Sub(start)

		var httpResponse = response{
			Byte:    c.Writer.Size(),
			Status:  c.Writer.Status(),
			Latency: latency,
		}

		user, err := GetLoggedUser(c)

		var motionLogger = motionLog{
			Api:          c.Request.RequestURI,
			Timestamp:    start,
			HttpRequest:  httpRequest,
			HttpResponse: httpResponse,
		}

		if err == nil {
			motionLogger.LoggedUser = &user
		}

		jsonData, err := json.Marshal(motionLogger)
		if err != nil {
			fmt.Println("Erro ao converter para JSON:", err)
			return
		}

		var msg string
		switch {
		case strconv.Itoa(c.Writer.Status())[0] == '5':
			msg = "31mError"
		case strconv.Itoa(c.Writer.Status())[0] == '4':
			msg = "33mForbidden"
		default:
			msg = "32mSuccess"
		}

		log.Printf(fmt.Sprintf("\033[%s:\033[0m %s.\"", msg, jsonData))
	}
}
