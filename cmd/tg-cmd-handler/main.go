package main

import (
	"context"
)

func main() {
	app, err := InitializeApp()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	app.TgListener.Run(ctx)
}
