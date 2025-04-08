package types

import (
	"time"

	"go.uber.org/zap/zapcore"
)

// Log represents a log entry
type Log struct {
	Timestamp  time.Time
	Caller     string
	LoggerName string
	Level      zapcore.Level
	Message    string
}

// LogHook is a function that will be called for each log entry
type LogHook func(log Log)
