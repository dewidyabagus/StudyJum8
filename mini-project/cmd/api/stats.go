package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type dependencies struct {
	db *gorm.DB
}

func (d *dependencies) Ping(ctx *gin.Context) {
	var response = map[string]string{
		"database": "down",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, 10*time.Second)
	defer cancel1()

	db, _ := d.db.DB()
	if err := db.PingContext(ctx1); err == nil {
		response["database"] = "up"
	}

	var statusCode = http.StatusOK

	for _, status := range response {
		if status == "down" {
			statusCode = http.StatusInternalServerError
			break
		}
	}

	ctx.JSON(statusCode, response)
}
