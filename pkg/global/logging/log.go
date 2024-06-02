package logging

import (
	"os"
)

var LOGGER = New()
var MAX_ERROR_STACKS = 20

func New() *Logger {
	logger := NewLogger()
	logger.IncrementCallerSkip()
	return logger
}

type LogConfig struct {
	Level          LogLevel
	MaxErrorStacks int
	EnableFileLine bool
	Output         string
	EnableColor    bool
}

func BasicConfig(config LogConfig) {
	if config.Level != 0 {
		LOGGER.SetLevel(config.Level)
	}
	if config.MaxErrorStacks > 0 {
		LOGGER.SetMaxErrorStacks(config.MaxErrorStacks)
	}
	if config.EnableFileLine {
		LOGGER.EnableFileLine()
	}
	if config.Output != "" {
		SetOutput(config.Output)
	}
	if config.EnableColor {
		LOGGER.EnableColor()
	}
}

func Debug(format string, args ...interface{}) {
	LOGGER.Debug(format, args...)
}
func Info(format string, args ...interface{}) {
	LOGGER.Info(format, args...)
}

func Warning(format string, args ...interface{}) {
	LOGGER.Warning(format, args...)
}

func Error(format string, args ...interface{}) {
	LOGGER.Error(format, args...)
}
func Fatal(format string, args ...interface{}) {
	LOGGER.Error(format, args...)
	os.Exit(1)
}
func Success(format string, args ...interface{}) {
	LOGGER.Success(format, args...)
}

func Exception(err error) {
	LOGGER.Exception(err)
}

func Panic(err error) {
	LOGGER.Panic(err)
}

func init() {
	LOGGER.SetLevel(ERROR)
	LOGGER.SetMaxErrorStacks(20)
}
