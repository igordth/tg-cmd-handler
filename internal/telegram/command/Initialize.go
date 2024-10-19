package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/igordth/tg-cmd-handler/internal/telegram/messanger"
	"go.uber.org/zap"
)

type Message string

func (m Message) String() string {
	return string(m)
}

type Initialize interface {
	Init(up tgbotapi.Update) Runner
}

type cmd struct {
	api    *tgbotapi.BotAPI
	sender messanger.Sender
	up     tgbotapi.Update
	log    *zap.Logger
}

func New(api *tgbotapi.BotAPI, sender messanger.Sender, log *zap.Logger) Initialize {
	return &cmd{
		api:    api,
		sender: sender,
		log:    log,
	}
}

func (c *cmd) Init(up tgbotapi.Update) Runner {
	clone := *c
	clone.up = up
	return &clone
}
