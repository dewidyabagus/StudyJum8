package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/slack-go/slack"
)

// Invalid pewarisan
// type SlackClient2 slack.Client {
// 	Field1 string
// }

type SlackClient struct {
	channelID string
	*slack.Client
}

func NewSlackClient(channelID, token string) *SlackClient {
	return &SlackClient{
		channelID: channelID,
		Client:    slack.New(token, slack.OptionDebug(false)),
	}
}

func (s *SlackClient) ErrorReporting(ctx context.Context, errMsg, stacktrace string) error {
	ctxWT, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Send Message
	attachment := slack.Attachment{
		ServiceName: "Service API Calculation",
		AuthorName:  "Error Reporting",
		Color:       "#eb0729",
		Pretext:     ":boom: Panic Recovery From Service",
		Text:        "There is a problem with the message:\n`" + errMsg + "`\n\nStacktrace:\n```" + stacktrace + "```",
	}
	channel, timestamp, err := s.PostMessageContext(
		ctxWT,
		s.channelID,
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		return fmt.Errorf("post message slack: %w", err)
	}
	log.Printf("Message successfully sent to channel %s at %s\n", channel, timestamp)

	return nil
}
