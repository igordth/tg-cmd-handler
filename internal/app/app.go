package app

import "github.com/igordth/tg-cmd-handler/internal/telegram/listener"

type App struct {
	TgListener listener.Runner
}

func New(tgListener listener.Runner) *App {
	return &App{
		TgListener: tgListener,
	}
}
