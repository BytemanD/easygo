package logging

import (
	"log"
	"os"
)

var std = New()
var MAX_ERROR_STACKS = 20

func New() *Logger {
	return &Logger{
		Level:          ERROR,
		MaxErrorStacks: MAX_ERROR_STACKS,
	}
}

type LogConfig struct {
	Level          LogLevel
	MaxErrorStacks int
}

func BasicConfig(config LogConfig) {
	if config.Level != 0 {
		std.SetLevel(config.Level)
	}
	if config.MaxErrorStacks > 0 {
		std.SetMaxErrorStacks(config.MaxErrorStacks)
	}
}

func SetOutput(file string) {
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
}

func Debug(format string, args ...interface{}) {
	std.Debug(format, args...)
}
func Info(format string, args ...interface{}) {
	std.Info(format, args...)
}

func Warning(format string, args ...interface{}) {
	std.Warning(format, args...)
}

func Error(format string, args ...interface{}) {
	std.Error(format, args...)
}
func Fatal(format string, args ...interface{}) {
	std.Error(format, args...)
	os.Exit(1)
}
func Exception(err error) {
	std.Exception(err)
}

func Panic(err error) {
	std.Panic(err)
}
