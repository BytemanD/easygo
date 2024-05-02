package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type LogLevel uint32

const (
	ERROR   LogLevel = 1
	WARNING LogLevel = 2
	INFO    LogLevel = 3
	DEBUG   LogLevel = 4
)

type Logger struct {
	Level          LogLevel
	MaxErrorStacks int
	Out            io.Writer
	callerSkip     int
	enableFileLine bool
}

func (logger *Logger) isLevelEnable(logLevel LogLevel) bool {
	return logger.Level >= logLevel
}

func (logger *Logger) SetLevel(logLevel LogLevel) {
	logger.Level = logLevel
}
func (logger *Logger) EnableFileLine() {
	logger.enableFileLine = true
}

func (logger *Logger) IncrementCallerSkip() {
	logger.callerSkip += 1
}

func (logger *Logger) SetMaxErrorStacks(num int) {
	logger.MaxErrorStacks = num
}

func (logger *Logger) prefix() string {
	prefix := fmt.Sprintf("%d", os.Getpid())
	if logger.enableFileLine {
		_, file, line, _ := runtime.Caller(logger.callerSkip + 1)
		shortFile := filepath.Base(file)
		prefix += fmt.Sprintf(" %s:%d", shortFile, line)
	}
	return prefix
}

func (logger *Logger) Debug(format string, args ...interface{}) {
	if !logger.isLevelEnable(DEBUG) {
		return
	}
	fmt.Print("\033[2K\r")
	log.Printf("%s DEBUG %s", logger.prefix(), fmt.Sprintf(format, args...))
}

func (logger *Logger) Info(format string, args ...interface{}) {
	if !logger.isLevelEnable(INFO) {
		return
	}
	fmt.Print("\033[2K\r")
	log.Printf("%s INFO %s", logger.prefix(), fmt.Sprintf(format, args...))
}
func (logger *Logger) Warning(format string, args ...interface{}) {
	if !logger.isLevelEnable(WARNING) {
		return
	}
	fmt.Print("\033[2K\r")
	log.Printf("%s WARNING %s", logger.prefix(), fmt.Sprintf(format, args...))
}
func (logger *Logger) Error(format string, args ...interface{}) {
	if !logger.isLevelEnable(ERROR) {
		return
	}
	fmt.Print("\033[2K\r")
	log.Printf("%s ERROR %s", logger.prefix(), fmt.Sprintf(format, args...))
}
func (logger *Logger) Panic(err error) {
	log.Panic(err)
}
func (logger *Logger) Exception(err error) {
	logger.Error("%s Exception: %+v", logger.prefix(), err)
	for i := 1; i <= logger.MaxErrorStacks; i++ {
		pc, file, line, _ := runtime.Caller(i)
		method := runtime.FuncForPC(pc)
		if file == "" {
			break
		}
		logger.Error("    %s(...)", method.Name())
		logger.Error("        %s:%d", file, line)
		if method.Name() == "main.main" {
			break
		}
	}
}
func SetOutput(file string) {
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
}
func NewLogger() *Logger {
	logger := &Logger{
		Level:          ERROR,
		MaxErrorStacks: MAX_ERROR_STACKS,
		callerSkip:     1,
	}
	return logger
}
