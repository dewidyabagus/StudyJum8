package response

import (
	"time"

	"github.com/gin-gonic/gin"
)

type BasicResponse struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

type BasicData struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

func JSONBasicResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(
		code,
		BasicResponse{
			Status:    code,
			Message:   message,
			Timestamp: time.Now().UnixNano(),
		},
	)
}

func JSONBasicData(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(
		code,
		BasicData{
			Status:    code,
			Message:   message,
			Timestamp: time.Now().UnixNano(),
			Data:      data,
		},
	)
}
