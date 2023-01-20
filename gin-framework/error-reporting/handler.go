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

func divideController(ctx *gin.Context) {
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
}

// func uploadController() {
// 	// Proses notif
// }
