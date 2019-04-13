package main

import (
	"github.com/nlopes/slack"
)

type SlackListener struct {
	client *slack.Client
	botID  string
}
