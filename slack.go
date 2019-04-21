package main

import (
	"fmt"

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

func PostEphemeral(c *slack.Client, channelID, userID, message string) error {
	_, err := c.PostEphemeral(
		channelID,
		userID,
		slack.MsgOptionText(message, false),
		slack.MsgOptionPostMessageParameters(slack.NewPostMessageParameters()),
	)
	return err
}

func unknown(command string) string {
	return fmt.Sprintf("Unknown command \"%s\".", command)
}
