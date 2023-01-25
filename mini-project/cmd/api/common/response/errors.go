package response

import (
	"net/http"
	"time"

	"github.com/dewidyabagus/go-hexagonal/business"
	"github.com/dewidyabagus/go-hexagonal/utils/validator"

	"github.com/gin-gonic/gin"
)

func JSONErrorResponse(ctx *gin.Context, err error) {
	switch v := err.(type) {
	case *business.ErrDuplicateEntry:
		ctx.JSON(
			http.StatusBadRequest,
			BasicResponse{
				Status:    http.StatusBadRequest,
				Message:   v.Error(),
				Timestamp: time.Now().UnixNano(),
			},
		)

	case validator.ValidationErrors:
		ctx.JSON(
			http.StatusBadRequest,
			BasicResponse{
				Status:    http.StatusBadRequest,
				Message:   validator.TranslateError(v).Error(),
				Timestamp: time.Now().UnixNano(),
			},
		)

	case *business.ErrUnauthorized:
		ctx.JSON(
			http.StatusUnauthorized,
			BasicResponse{
				Status:    http.StatusUnauthorized,
				Message:   v.Error(),
				Timestamp: time.Now().UnixNano(),
			},
		)

	default:
		ctx.JSON(
			http.StatusInternalServerError,
			BasicResponse{
				Status:    http.StatusInternalServerError,
				Message:   v.Error(),
				Timestamp: time.Now().UnixNano(),
			},
		)
	}
}
