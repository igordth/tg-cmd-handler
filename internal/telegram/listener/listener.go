package listener

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/igordth/tg-cmd-handler/internal/telegram/command"
	"go.uber.org/zap"
)

type Runner interface {
	Run(ctx context.Context)
}

type listener struct {
	api *tgbotapi.BotAPI
	cmd command.Initialize
	log *zap.Logger
}

func New(api *tgbotapi.BotAPI, cmd command.Initialize, log *zap.Logger) Runner {
	return &listener{
		api: api,
		cmd: cmd,
		log: log.Named("telegram"),
	}
}

func (b *listener) Run(ctx context.Context) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.api.GetUpdatesChan(u)
	for update := range updates {
		b.handlingUpdate(ctx, update)
	}
}

func (b *listener) handlingUpdate(ctx context.Context, up tgbotapi.Update) {
	if up.Message == nil {
		return
	}
	if up.Message.IsCommand() {
		b.handlingCommand(ctx, up)
		return
	}
}

func (b *listener) handlingCommand(ctx context.Context, up tgbotapi.Update) {
	err := b.cmd.Init(up).Run(ctx)
	if err != nil {
		b.log.Error("telegram command failed", zap.Error(err))
	}
}
