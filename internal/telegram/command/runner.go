package command

import (
	"context"
	"errors"
	"fmt"
)

const (
	msgStart Message = "start"
	msgHelp  Message = "help"

	msgStartAnswer = "Hello, %s"
	msgHelpAnswer  = `List of all commands:
/start
/help

this is all`
)

var ErrNotFound = errors.New("command not found")

type Runner interface {
	Run(ctx context.Context) error
}

func (c *cmd) Run(ctx context.Context) error {
	switch c.up.Message.Command() {
	case msgStart.String():
		return c.runStart(ctx)
	case msgHelp.String():
		return c.runHelp(ctx)
	default:
		return ErrNotFound
	}
}

func (c *cmd) runStart(ctx context.Context) (err error) {
	err = c.sender.Send(ctx, c.up.Message.Chat.ID, fmt.Sprintf(msgStartAnswer, c.up.SentFrom().String()))
	return
}

func (c *cmd) runHelp(ctx context.Context) (err error) {
	err = c.sender.Send(ctx, c.up.Message.Chat.ID, msgHelpAnswer)
	return
}
