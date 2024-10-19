package messanger

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Sender interface {
	Send(ctx context.Context, chatId int64, message string) error
}

type messenger struct {
	api *tgbotapi.BotAPI
	log *zap.Logger
}

func New(api *tgbotapi.BotAPI, log *zap.Logger) Sender {
	return &messenger{
		api: api,
		log: log.Named("telegram-messenger"),
	}
}

func (m *messenger) Send(_ context.Context, chatId int64, message string) (err error) {
	msg := tgbotapi.NewMessage(chatId, message)
	msg.ParseMode = tgbotapi.ModeMarkdownV2
	_, err = m.api.Send(msg)
	return
}
