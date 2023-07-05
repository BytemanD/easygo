package logging

import (
	"fmt"
	"io"
	"log"
	"os"
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
}

func init() {
	std.SetLevel(ERROR)
	std.SetMaxErrorStacks(20)
}

func (logger *Logger) isLevelEnable(logLevel LogLevel) bool {
	return logger.Level >= logLevel
}

func (logger *Logger) SetLevel(logLevel LogLevel) {
	logger.Level = logLevel
}

func (logger *Logger) SetMaxErrorStacks(num int) {
	logger.MaxErrorStacks = num
}

func (logger *Logger) Debug(format string, args ...interface{}) {
	if !logger.isLevelEnable(DEBUG) {
		return
	}
	log.Printf(" %d DEBUG %s", os.Getpid(), fmt.Sprintf(format, args...))
}

func (logger *Logger) Info(format string, args ...interface{}) {
	if !logger.isLevelEnable(INFO) {
		return
	}
	log.Printf("%d INFO %s", os.Getpid(), fmt.Sprintf(format, args...))
}
func (logger *Logger) Warning(format string, args ...interface{}) {
	if !logger.isLevelEnable(WARNING) {
		return
	}
	log.Printf("%d WARNING %s", os.Getpid(), fmt.Sprintf(format, args...))
}
func (logger *Logger) Error(format string, args ...interface{}) {
	if !logger.isLevelEnable(ERROR) {
		return
	}
	log.Printf("%d ERROR %s", os.Getpid(), fmt.Sprintf(format, args...))
}
func (logger *Logger) Panic(err error) {
	log.Panic(err)
}
func (logger *Logger) Exception(err error) {
	logger.Error("%d Exception: %+v", os.Getpid(), err)
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
