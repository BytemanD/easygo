package logging

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

func Exception(err error) {
	std.Exception(err)
}

func Panic(err error) {
	std.Panic(err)
}
