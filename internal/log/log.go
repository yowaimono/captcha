package internal

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

// Logger 是自定义的日志记录器
type Logger struct {
	logger *log.Logger
}

var (
	instanceLog *Logger
	olog        sync.Once
)

// GetLogger 返回单例模式的日志记录器
func GetLogger() *Logger {
	olog.Do(func() {
		instanceLog = &Logger{
			logger: log.New(os.Stdout, "", 0),
		}
	})
	return instanceLog
}

// logf 是日志记录的通用函数
func (l *Logger) logf(level, format string, v ...interface{}) {
	now := time.Now().Format("2006-01-02 15:04:05")
	_, file, line, _ := runtime.Caller(2)
	file = file[strings.LastIndex(file, "/")+1:]
	message := fmt.Sprintf(format, v...)

	// 添加颜色
	var levelColor string
	switch level {
	case "INFO":
		levelColor = "\033[32m" // 绿色
	case "ERROR":
		levelColor = "\033[31m" // 红色
	case "WARN":
		levelColor = "\033[33m" // 黄色
	case "DEBUG":
		levelColor = "\033[36m" // 青色
	case "FATAL":
		levelColor = "\033[35m" // 紫色
	default:
		levelColor = "\033[0m" // 默认颜色
	}

	l.logger.Printf("[\033[35mMIN\033[0m] [\033[34m%s\033[0m] [\033[36m%s:%d\033[0m] [\033[33m%s%s\033[0m] %s", now, file, line, levelColor, level, message)
}

// Info 记录信息级别的日志
func Info(format string, v ...interface{}) {
	GetLogger().logf("INFO", format, v...)
}

// Error 记录错误级别的日志
func Error(format string, v ...interface{}) {
	GetLogger().logf("ERROR", format, v...)
}

// Warn 记录警告级别的日志
func Warn(format string, v ...interface{}) {
	GetLogger().logf("WARN", format, v...)
}

// Debug 记录调试级别的日志
func Debug(format string, v ...interface{}) {
	GetLogger().logf("DEBUG", format, v...)
}

// Fatal 记录致命错误级别的日志并退出程序
func Fatal(format string, v ...interface{}) {
	GetLogger().logf("FATAL", format, v...)
	os.Exit(1)
}
