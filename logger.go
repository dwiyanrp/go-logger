package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

var (
	instance *Logger
	once     sync.Once
)

type Logger struct {
	config *viper.Viper
	level  LogLevel
}

func Init(v *viper.Viper) {
	once.Do(func() {
		instance = &Logger{config: v}
		instance.setLogLevel()
	})
}

func (l *Logger) setLogLevel() {
	levelStr := strings.ToLower(l.config.GetString("log.level"))
	switch levelStr {
	case "debug":
		l.level = DebugLevel
	case "info":
		l.level = InfoLevel
	case "warn":
		l.level = WarnLevel
	case "error":
		l.level = ErrorLevel
	case "fatal":
		l.level = FatalLevel
	default:
		l.level = InfoLevel // Default to Info if not specified
	}
}

func (l *Logger) log(level LogLevel, message string) {
	if level >= l.level {
		levelStr := []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}[level]
		fmt.Printf("[%s] %s\n", levelStr, message)
		if level == FatalLevel {
			os.Exit(1)
		}
	}
}

func Debug(category string, message string) {
	if instance.level <= DebugLevel && instance.config.GetBool(fmt.Sprintf("debug.%s", category)) {
		fmt.Printf("[DEBUG:%s] %s\n", category, message)
	}
}

func Info(message string) {
	instance.log(InfoLevel, message)
}

func Warn(message string) {
	instance.log(WarnLevel, message)
}

func Error(message string) {
	instance.log(ErrorLevel, message)
}

func Fatal(message string) {
	instance.log(FatalLevel, message)
}
