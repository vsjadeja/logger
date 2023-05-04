package logger

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

// Fields provides an easy to use KV colletion for logging arbitary fields
type Fields map[string]interface{}

// LogLevel determines the severity/importance of the log statement
type LogLevel int

const (
	// Info informational message
	Info LogLevel = 0
	// Warning the system has encounted an unexpected state that can be recovered
	Warning LogLevel = 1
	// Error the system has encounted an unexpected state that cannot be recovered
	Error LogLevel = 2
)

// LoggerInterface presents an easy to use interface implemented by log.Logger
// provided here to support mocking and unit testing
type LoggerInterface interface {
	Log(level LogLevel, fields Fields)
}

// JSONLogger provides a simple implementation for a Logger
type JSONLogger struct {
	Logger *log.Logger
	Level  LogLevel
}

// Log writes the error
func (l *JSONLogger) Log(level LogLevel, fields Fields) {
	if l.Logger == nil {
		return
	}
	if !(level >= l.Level) {
		return
	}
	data, _ := json.Marshal(fields)
	var prefix string
	switch level {
	case Error:
		prefix = "{\"error\": "
	case Info:
		prefix = "{\"info\": "
	case Warning:
		prefix = "{\"warning\": "
	}
	l.Logger.Print(prefix, bytes.NewBuffer(data).String(), "}")
}

// NewJSONLogger returns a trivial implementation for a logger
func NewJSONLogger(level LogLevel) *JSONLogger {
	return &JSONLogger{Logger: log.New(os.Stderr, "", 0), Level: level}
}

// NewNilLogger returns a logger thats an NoOp
func NewNilLogger() *JSONLogger {
	return &JSONLogger{Logger: nil, Level: Info}
}
