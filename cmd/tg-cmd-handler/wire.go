//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/igordth/tg-cmd-handler/internal/app"
	"github.com/igordth/tg-cmd-handler/internal/config"
	"github.com/igordth/tg-cmd-handler/internal/logger"
	"github.com/igordth/tg-cmd-handler/internal/telegram"
	"github.com/igordth/tg-cmd-handler/internal/telegram/command"
	"github.com/igordth/tg-cmd-handler/internal/telegram/listener"
	"github.com/igordth/tg-cmd-handler/internal/telegram/messanger"
)

func InitializeApp() (*app.App, error) {
	panic(wire.Build(
		config.NewYaml,
		wire.FieldsOf(new(config.Application), "LoggerMode", "TelegramBotToken"),
		logger.NewZap,
		telegram.NewApi,
		messanger.New,
		command.New,
		listener.New,
		app.New,
	))
}
