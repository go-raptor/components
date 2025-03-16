package components

import (
	"log/slog"
	"os"
	"strings"

	"github.com/go-raptor/config"
	"github.com/go-raptor/connector"
	"github.com/pwntr/tinter"
)

type Utils struct {
	Config *config.Config

	Log      *slog.Logger
	LogLevel *slog.LevelVar

	DB connector.DatabaseConnector
}

func NewUtils() *Utils {
	levelVar := &slog.LevelVar{}

	opts := &tinter.Options{
		Level: levelVar,
	}

	return &Utils{
		Log:      slog.New(tinter.NewHandler(os.Stderr, opts)),
		LogLevel: levelVar,
	}
}

func (u *Utils) SetDB(db connector.DatabaseConnector) {
	u.DB = db
}

func (u *Utils) SetConfig(config *config.Config) {
	u.Config = config
}

func (u *Utils) SetLogLevel(logLevel string) {
	var level slog.Level
	switch strings.ToUpper(logLevel) {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}
	u.LogLevel.Set(level)
}

func (u *Utils) SetHandler(handler slog.Handler) {
	u.Log = slog.New(handler)
}
