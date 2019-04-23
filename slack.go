package main

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/nlopes/slack"
)

const (
	commandPing     = "ping"
	commandHelp     = "help"
	commandNew      = "new"
	commandJoin     = "join"
	commandQuestion = "question"
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

func dispatchHandler(s *SlackListener, ev *slack.MessageEvent, m []string) error {
	switch m[0] {
	case commandPing:
		if err := PostMessage(s.client, ev.Channel, "pong"); err != nil {
			return fmt.Errorf("failed to post message: %v", err)
		}
	case commandHelp:
		if err := PostMessage(s.client, ev.Channel, help()); err != nil {
			return fmt.Errorf("failed to post message: %v", err)
		}
	case commandNew:
		return s.handleNew(ev, m)
	case commandJoin:
		return s.handleJoin(ev, m)
	case commandQuestion:
		return s.handleQuestion(ev, m)
	default:
		if err := PostMessage(s.client, ev.Channel, unknown(m[0])); err != nil {
			return fmt.Errorf("failed to post message: %v", err)
		}
		return fmt.Errorf("invalid message: %v", ev.Msg.Text)
	}
	return nil
}

func unknown(command string) string {
	return fmt.Sprintf("Unknown command \"%s\".", command)
}

func help() string {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	for _, command := range []struct {
		name        string
		description string
	}{
		{commandPing, ""},
		{commandHelp, ""},
		{commandNew, "新しいゲームを始める"},
		{commandJoin, "ゲームに参加する"},
		{commandQuestion, "場の質問を一覧する"},
		{fmt.Sprintf("%s <number>", commandQuestion), "質問する"},
	} {
		fmt.Fprintf(w, "%s\t%s\n", command.name, command.description)
	}
	w.Flush()
	return "```" + buf.String() + "```"
}
