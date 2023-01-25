package middlewares

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type logger struct {
	Level    string `json:"lvl"`
	Time     string `json:"time"`
	Msg      string `json:"msg"`
	Method   string `json:"method"`
	Path     string `json:"path"`
	Status   int    `json:"status"`
	Latency  string `json:"latency"`
	ClientIP string `json:"client_ip"`
	TraceID  string `json:"trace_id"`
	ErrorMsg string `json:"error,omitempty"`
}

func CustomLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		byteLogger, _ := json.Marshal(
			logger{
				Level:    "info",
				Time:     time.Now().Format("2006-01-02T15:04:05.999+0700"),
				Msg:      "request completed",
				Method:   param.Method,
				Path:     param.Path,
				Status:   param.StatusCode,
				Latency:  param.Latency.String(),
				ClientIP: param.ClientIP,
				TraceID:  param.Request.Header.Get("X-Request-ID"),
				ErrorMsg: param.ErrorMessage,
			},
		)
		return fmt.Sprintf("%s\n", string(byteLogger))
	})
}
