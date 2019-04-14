package main

import (
	"github.com/nlopes/slack"
)

type SlackListener struct {
	client *slack.Client
	botID  string
}

func PostMessage(c *slack.Client, channelID, message string) error {
	_, _, err := c.PostMessage(
		channelID,
		slack.MsgOptionText(message, false),
		slack.MsgOptionPostMessageParameters(slack.NewPostMessageParameters()),
	)
	return err
}
