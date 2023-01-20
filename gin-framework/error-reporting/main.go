package main

import (
	"github.com/gin-gonic/gin"
)

// Slack Channel ID dan Token
const (
	channelID  = "<<Input Channel ID"
	slackToken = "<<Input Token>>"
)

func main() {
	r := gin.Default()

	// slack := NewSlackClient()
	recovery := &Recovery{
		Slack: NewSlackClient(channelID, slackToken),
	}
	r.Use(gin.CustomRecovery(recovery.ErrorHandle))
	r.GET("/divide", divideController)

	r.Run(":8001")
}
