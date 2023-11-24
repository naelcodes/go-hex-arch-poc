package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

// NewLogger creates a new instance of the logger.
func NewLogger() *Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout}
	output.TimeFormat = "2006-01-02 15:04:05"
	logger := zerolog.New(output).With().Timestamp().Logger()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	return &Logger{logger: logger}
}

// Info logs an informational message.
func (l *Logger) Info(message string) {
	l.logger.Info().Msg(message)
}

// Error logs an error message.
func (l *Logger) Error(message string) {
	l.logger.Error().Msg(message)
}

// Debug logs a debug message.
func (l *Logger) Debug(message string) {
	l.logger.Debug().Msg(message)
}

// Panic logs a panic message and panics.
func (l *Logger) Panic(message string) {
	l.logger.Panic().Msg(message)
}
