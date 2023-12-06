package utils

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// function that generate a random string
func GenerateRandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func GenerateCode(prefix string, number int) string {
	UpperCasePrefix := strings.ToUpper(prefix)
	var suffix string

	if number < 10 {
		suffix = "000" + strconv.Itoa(number)
	} else if number < 100 {
		suffix = "00" + strconv.Itoa(number)
	} else if number < 1000 {
		suffix = "0" + strconv.Itoa(number)
	} else {
		suffix = strconv.Itoa(number)
	}
	return fmt.Sprintf("%s-%s", UpperCasePrefix, suffix)

}

func GetCurrentDate() string {
	currentDate := time.Now().Format("2006-01-02")
	return currentDate
}

func FormatDate(date string) string {
	// log.Printf("date: %s", date)
	parsedDate, _ := time.Parse(time.RFC3339, date)
	// log.Printf("parsedDate: %s", parsedDate)
	return parsedDate.Local().Format("2006-01-02")
}

// Logger is a global logger instance
var Logger = NewZeroLogger()

type ZeroLogger struct {
	logger zerolog.Logger
}

func NewZeroLogger() *ZeroLogger {
	output := zerolog.ConsoleWriter{Out: os.Stdout}
	output.TimeFormat = "2006-01-02 15:04:05"
	logger := zerolog.New(output).With().Timestamp().Logger()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	return &ZeroLogger{logger: logger}
}

// Info logs an informational message.
func (l *ZeroLogger) Info(message string) {
	l.logger.Info().Msg(message + "\n")
}

// Error logs an error message.
func (l *ZeroLogger) Error(message string) {
	l.logger.Error().Msg(message + "\n")
}

// Debug logs a debug message.
func (l *ZeroLogger) Debug(message string) {
	l.logger.Debug().Msg(message + "\n")
}

// Panic logs a panic message and panics.
func (l *ZeroLogger) Panic(message string) {
	l.logger.Panic().Msg(message + "\n")
}

// -------------------------------------------
