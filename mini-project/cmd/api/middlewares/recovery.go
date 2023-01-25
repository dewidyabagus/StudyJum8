package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(ctx *gin.Context, err any) {
	if val, ok := err.(error); ok {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			ErrResponse{
				Message:     "Internal Server Error",
				Status:      http.StatusInternalServerError,
				Description: val.Error(),
			},
		)
	}
}
