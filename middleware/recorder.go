package middleware

import (
	"fmt"
	"time"

	"log"

	"github.com/gin-gonic/gin"
)

func logFormatter(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}

	if len(param.ErrorMessage) > 0 {
		log.Printf("%s,%s,%s\n", param.ClientIP, param.Method, param.ErrorMessage)
	}

	return fmt.Sprintf("[file-server] %v recorder.go:29\t[%s %3d %s| %13v | %15s |%s %-7s %s %#v]\n",
		param.TimeStamp.Format("2006-01-02 15:04:05.000000"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
	)
}

func Recorder() gin.HandlerFunc {
	return gin.LoggerWithFormatter(logFormatter)
}
