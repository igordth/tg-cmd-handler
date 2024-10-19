package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Token string

func (tb Token) String() string {
	return string(tb)
}

func NewApi(t Token) (*tgbotapi.BotAPI, error) {
	return tgbotapi.NewBotAPI(t.String())
}
