package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Digunakan untuk menulis ulang header request ketika
// suatu kunci optional yang dibutuhkan kosong.

// https://http.dev/x-request-id

const requestID = "X-Request-ID"

func Headers() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader(requestID)

		if strings.TrimSpace(traceID) == "" {
			c.Request.Header.Set(requestID, uuid.New().String())
		}

		c.Next()
	}
}
