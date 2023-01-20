package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/go-errors/errors"
)

// Library stacktrace: github.com/go-errors/errors

type Recovery struct {
	Slack *SlackClient
}

func (r *Recovery) ErrorHandle(ctx *gin.Context, err any) {
	wrapErr := errs.Wrap(err, 2)

	if val, ok := err.(error); ok {
		r.Slack.ErrorReporting(ctx, wrapErr.Error(), wrapErr.ErrorStack())

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
