package logger

import "github.com/theartofdevel/logging"

type Logger struct {
	Log *logging.Logger
}

func NewLogger() (*Logger, error) {
	return &Logger{
		Log: logging.NewLogger(
			logging.WithIsJSON(true),
			logging.WithAddSource(true),
		),
	}, nil
}
