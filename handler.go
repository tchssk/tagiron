package main

import (
	"fmt"
	"strconv"
	"strings"

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

func (s *SlackListener) handleQuestion(ev *slack.MessageEvent, m []string) error {
	switch len(m) {
	case 1:
		return s.handleQuestionList(ev)
	case 2:
		return s.handleQuestionAsk(ev, m)
	}
	return nil
}

func (s *SlackListener) handleQuestionList(ev *slack.MessageEvent) error {
	ss := []string{"質問:"}
	for i, question := range game.Questions {
		ss = append(ss, fmt.Sprintf("%d. %s", i+1, question))
	}
	if err := PostMessage(s.client, ev.Channel, strings.Join(ss, "\n")); err != nil {
		return fmt.Errorf("failed to post message: %v", err)
	}
	return nil
}

func (s *SlackListener) handleQuestionAsk(ev *slack.MessageEvent, m []string) error {
	index, err := strconv.Atoi(m[1])
	if err != nil {
		return err
	}
	player := game.Players.FindByID(ev.User)
	questions := game.Questions.Remove(index-1, index)
	player.Questions.Add(questions)
	game.Questions.Add(game.Stock.Remove(0, 1))
	ss := []string{fmt.Sprintf("%s さんが質問しました。", player.Name)}
	for _, question := range questions {
		ss = append(ss, question)
	}
	if err := PostMessage(s.client, ev.Channel, strings.Join(ss, "\n")); err != nil {
		return fmt.Errorf("failed to post message: %v", err)
	}
	return nil
}
