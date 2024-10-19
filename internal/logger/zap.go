package logger

import (
	"errors"
	"go.uber.org/zap"
)

type Mode string

const (
	ModeDevelopment Mode = "development"
	ModeProduction  Mode = "production"
)

var ErrInvalidMode = errors.New("invalid Mode")

func NewZap(mode Mode) (l *zap.Logger, err error) {
	switch mode {
	case ModeDevelopment:
		l, err = zap.NewDevelopment()
		if l != nil {
			l.Debug("zap logger init")
		}
	case ModeProduction:
		l, err = zap.NewProduction()
	default:
		return nil, ErrInvalidMode
	}
	return
}
