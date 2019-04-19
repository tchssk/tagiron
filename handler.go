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

func (s *SlackListener) handleJoin(ev *slack.MessageEvent, m []string) error {
	user, err := s.client.GetUserInfo(ev.User)
	if err != nil {
		return err
	}
	player := NewPlayer(user.Profile.RealName, user.ID)
	player.Tiles.Add(game.Tiles.Remove(0, 5))
	player.Tiles.Sort()
	game.Players.Add(player)

	if err := PostMessage(s.client, ev.Channel, fmt.Sprintf("%s さんが参加しました。", user.Profile.RealName)); err != nil {
		return fmt.Errorf("failed to post message: %v", err)
	}
	if err := PostEphemeral(s.client, ev.Channel, ev.User, fmt.Sprintf("あなたのタイル:\n%v", player.Tiles)); err != nil {
		return fmt.Errorf("failed to post message: %v", err)
	}
	return nil
}
