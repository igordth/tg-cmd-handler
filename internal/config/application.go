package config

import (
	"github.com/igordth/tg-cmd-handler/internal/logger"
	"github.com/igordth/tg-cmd-handler/internal/telegram"
)

type Application struct {
	LoggerMode       logger.Mode    `yaml:"logger_mode"`
	TelegramBotToken telegram.Token `yaml:"telegram_bot_token"`
}
