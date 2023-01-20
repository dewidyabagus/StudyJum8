package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message     string `json:"message,omitempty"`
	Status      uint16 `json:"status,omitempty"`
	Description string `json:"description"`
}

// Format closure gin.RecoveryFunc
// func nameFunction(ctx *gin.Context, err any) {}
func errorHandler(ctx *gin.Context, err any) {
	val, ok := err.(error)
	if ok {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			Response{
				Message:     "Internal Server Error",
				Status:      http.StatusInternalServerError,
				Description: val.Error(),
			},
		)
	}
}

func main() {
	r := gin.Default()

	// Method yang digunakan untuk meregistrasikan closure proses recovery ketika terjadi error (panic)
	// r.Use(gin.CustomRecovery())
	// Controller itu dijalankan, nantinya ketika ada panic akan dikirimkan ke proses recovery
	r.Use(gin.CustomRecovery(errorHandler))

	r.GET("/divide", func(ctx *gin.Context) {
		var (
			y, _ = strconv.Atoi(ctx.Query("y"))
			x, _ = strconv.Atoi(ctx.Query("x"))
			z    = (x / y)
		)
		ctx.JSON(
			http.StatusOK,
			Response{
				Message:     "OK",
				Status:      http.StatusOK,
				Description: fmt.Sprintf("Divide: %d / %d = %d", x, y, z),
			})
	})

	r.Run(":8001")
}
