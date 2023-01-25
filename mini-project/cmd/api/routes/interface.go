package routes

import "github.com/gin-gonic/gin"

type healthCheck interface {
	Ping(ctx *gin.Context)
}
