package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

var (
	game Game
)

func (s *SlackListener) handleNew(ev *slack.MessageEvent, m []string) error {
	game = NewGame()

	if err := PostMessage(s.client, ev.Channel, "新しいゲームを始めます。"); err != nil {
		return fmt.Errorf("failed to post message: %v", err)
	}
	return nil
}
