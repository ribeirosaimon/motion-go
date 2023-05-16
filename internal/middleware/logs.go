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
	Api          string    `json:"api"`
	HttpRequest  request   `json:"httpRequest"`
	HttpResponse response  `json:"httpResponse"`
	Timestamp    time.Time `json:"timestamp"`
}

type request struct {
	Proto         string `json:"proto"`
	RemoteIP      string `json:"remoteIP"`
	RequestMethod string `json:"requestMethod"`
	RequestPath   string `json:"requestPath"`
}

type response struct {
	Byte    int           `json:"byte"`
	Status  int           `json:"status"`
	Latency time.Duration `json:"latency"`
}

func NewLogger(api string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		var httpRequest = request{
			Proto:         c.Request.Proto,
			RemoteIP:      c.Request.RemoteAddr,
			RequestMethod: c.Request.Method,
			RequestPath:   c.Request.RequestURI,
		}

		end := time.Now()
		latency := end.Sub(start)

		var httpResponse = response{
			Byte:    c.Writer.Size(),
			Status:  c.Writer.Status(),
			Latency: latency,
		}

		var motionLog = motionLog{
			Api:          api,
			Timestamp:    start,
			HttpRequest:  httpRequest,
			HttpResponse: httpResponse,
		}

		jsonData, err := json.Marshal(motionLog)
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
