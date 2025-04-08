package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"

	"LutiLeech/pkg/logger/types"

	"go.uber.org/zap/zapcore"
)

// Данный логгер был написан мной в проекте cu-clubs-bot. Внесены некоторые изменения

// Logger represents a logger
type Logger struct {
	*zap.SugaredLogger
	LogsPath string
	Name     string
}

var (
	// Log is the global logger
	Log     *Logger
	logHook types.LogHook
)

// Config represents configuration options for logger initialization
type Config struct {
	Debug        bool           // Enable debug logging
	TimeLocation *time.Location // Set the time location
	LogToFile    bool           // Enable logging to a file
	LogsDir      string         // Set the directory for logs (default: current working directory)
}

// SetLogHook sets a hook function that will be called for each log entry
func SetLogHook(hook types.LogHook) {
	logHook = hook
	Log.Debug("Log hook set")
}

// Init is a function to initialize logger with extended configuration
func Init(config Config) (*Logger, error) {
	l := new(Logger)
	l.Name = "main"

	wd, err := os.Getwd()
	if err != nil {
		return l, err
	}

	// Set log directory, default to current working directory
	if config.LogsDir == "" {
		l.LogsPath = wd
	} else {
		l.LogsPath = filepath.Join(wd, config.LogsDir)
	}

	// Ensure log directory exists
	err = os.MkdirAll(l.LogsPath, os.ModePerm)
	if err != nil {
		return l, err
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "timestamp",
		NameKey:        "logger",
		CallerKey:      "caller",
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	if config.TimeLocation != nil {
		encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.In(config.TimeLocation).Format("2006-01-02 15:04:05"))
		}
	}

	var level zapcore.Level
	if config.Debug {
		level = zapcore.DebugLevel
	} else {
		level = zapcore.InfoLevel
	}

	// Console encoder with colors
	consoleEncoderConfig := encoderConfig
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	// File encoder without colors and with proper JSON formatting
	fileEncoderConfig := encoderConfig
	fileEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	fileEncoder := zapcore.NewJSONEncoder(fileEncoderConfig)

	var cores []zapcore.Core

	// Add console output
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), level)
	cores = append(cores, consoleCore)

	// Add file output if enabled
	if config.LogToFile {
		mainLogPath := filepath.Join(l.LogsPath, fmt.Sprintf("%s.log", time.Now().Format("2006-01-02 15:04")))
		fileWriter, errOpenFile := os.OpenFile(mainLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if errOpenFile != nil {
			return l, errOpenFile
		}

		fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(fileWriter), level)
		cores = append(cores, fileCore)
	}

	// Create combined core
	combinedCore := zapcore.NewTee(cores...)

	// Create logger with hook
	log := zap.New(combinedCore, zap.AddCaller(), zap.Hooks(func(entry zapcore.Entry) error {
		if logHook != nil {
			logHook(types.Log{
				Timestamp:  entry.Time,
				Caller:     entry.Caller.String(),
				LoggerName: entry.LoggerName,
				Level:      entry.Level,
				Message:    entry.Message,
			})
		}
		return nil
	}))

	l.SugaredLogger = log.Named(l.Name).Sugar()
	Log = l

	return l, nil
}

// Named returns a new logger with the specified name ("bot", "database", etc.)
func (l *Logger) Named(name string) *Logger {
	return &Logger{
		SugaredLogger: Log.SugaredLogger.Named(name),
		LogsPath:      Log.LogsPath,
		Name:          name,
	}
}

// customTimeEncoder formats time in GMT+0
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
