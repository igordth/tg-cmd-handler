// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/igordth/tg-cmd-handler/internal/app"
	"github.com/igordth/tg-cmd-handler/internal/config"
	"github.com/igordth/tg-cmd-handler/internal/logger"
	"github.com/igordth/tg-cmd-handler/internal/telegram"
	"github.com/igordth/tg-cmd-handler/internal/telegram/command"
	"github.com/igordth/tg-cmd-handler/internal/telegram/listener"
	"github.com/igordth/tg-cmd-handler/internal/telegram/messanger"
)

// Injectors from wire.go:

func InitializeApp() (*app.App, error) {
	application, err := config.NewYaml()
	if err != nil {
		return nil, err
	}
	token := application.TelegramBotToken
	botAPI, err := telegram.NewApi(token)
	if err != nil {
		return nil, err
	}
	mode := application.LoggerMode
	zapLogger, err := logger.NewZap(mode)
	if err != nil {
		return nil, err
	}
	sender := messanger.New(botAPI, zapLogger)
	initialize := command.New(botAPI, sender, zapLogger)
	runner := listener.New(botAPI, initialize, zapLogger)
	appApp := app.New(runner)
	return appApp, nil
}
