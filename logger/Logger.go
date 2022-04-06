package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Level int

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO  "
	case ERROR:
		return "ERROR "
	case FATAL:
		return "FATAL"
	case PANIC:
		return "PANIC"
	case WARN:
		return "WARN"
	}
	return ""
}

const (
	DEBUG Level = iota + 1
	INFO
	WARN
	ERROR
	FATAL
	PANIC
)

type Logger struct {
	*log.Logger
	level  Level
	prefix string
}

var Log = NewLogger().WithLongFile()

func SetOut(out io.Writer) {
	Log.SetOut(out)
}

func SetLevel(level Level) {
	Log.SetLevel(level)
}

func Debug(args ...interface{}) {
	Log.Output(3, DEBUG, "[DEBUG]", fmt.Sprint(args...))
}

func Debugf(format string, args ...interface{}) {
	Log.Output(3, DEBUG, "[DEBUG]", fmt.Sprintf(format, args...))
}

func Info(args ...interface{}) {
	Log.Output(3, INFO, "[INFO ]", fmt.Sprint(args...))
}

func Infof(format string, args ...interface{}) {
	Log.Output(3, INFO, "[INFO ]", fmt.Sprintf(format, args...))
}

func Warn(args ...interface{}) {
	Log.Output(3, WARN, "[WARN ]", fmt.Sprint(args...))
}

func Warnf(format string, args ...interface{}) {
	Log.Output(3, WARN, "[WARN ]", fmt.Sprintf(format, args...))
}

func Error(args ...interface{}) {
	Log.Output(3, ERROR, "[ERROR]", fmt.Sprint(args...))
}

func Errorf(format string, args ...interface{}) {
	Log.Output(3, ERROR, "[ERROR]", fmt.Sprintf(format, args...))
}

func Fatal(args ...interface{}) {
	if Log.level <= FATAL {
		Log.Output(3, FATAL, "[FATAL]", fmt.Sprint(args...))
		os.Exit(1)
	}
}

func Fatalf(format string, args ...interface{}) {
	if Log.level <= FATAL {
		Log.Output(3, FATAL, "[FATAL]", fmt.Sprintf(format, args...))
		os.Exit(1)
	}
}

func Panic(args ...interface{}) {
	if Log.level <= PANIC {
		Log.Output(3, PANIC, "[PANIC]", fmt.Sprint(args...))
		panic("")
	}
}

func Panicf(format string, args ...interface{}) {
	if Log.level <= PANIC {
		Log.Output(3, PANIC, "[PANIC]", fmt.Sprintf(format, args...))
		panic("")
	}
}

func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stderr, "", log.LstdFlags|log.Lmsgprefix),
		level:  INFO,
	}
}

func (l *Logger) SetLevel(level Level) {
	l.level = level
}

func (l *Logger) SetOutFile(file string) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	l.SetOut(f)
}

func (l *Logger) SetOut(out io.Writer) {
	l.Logger.SetOutput(out)
}

func (l *Logger) SetPrefix(prefix string) {
	l.prefix = prefix
	l.Logger.SetPrefix("")
}

func (l *Logger) WithShortFile() *Logger {
	l.SetFlags(l.Flags() | log.Lshortfile)
	return l
}

func (l *Logger) WithLongFile() *Logger {
	l.SetFlags(l.Flags() | log.Llongfile)
	return l
}

func (l *Logger) Output(calldepth int, le Level, prefix string, log string) {
	if l.level <= le {
		if l.prefix != "" {
			prefix = fmt.Sprintf("[%s] ", l.prefix) + prefix
			l.Logger.SetPrefix("")
		}
		_ = l.Logger.Output(calldepth, fmt.Sprintf(" %s %s", prefix, log))
	}
}

func (l *Logger) Debug(args ...interface{}) {
	l.Output(3, DEBUG, "[DEBUG]", fmt.Sprint(args...))
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Output(3, DEBUG, "[DEBUG]", fmt.Sprintf(format, args...))
}

func (l *Logger) Info(args ...interface{}) {
	l.Output(3, INFO, "[INFO ]", fmt.Sprint(args...))
}

func (l *Logger) INFO(format string, args ...interface{}) {
	l.Output(3, INFO, "[INFO ]", fmt.Sprintf(format, args...))
}

func (l *Logger) Warn(args ...interface{}) {
	l.Output(3, WARN, "[WARN ]", fmt.Sprint(args...))
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Output(3, WARN, "[WARN ]", fmt.Sprintf(format, args...))
}

func (l *Logger) Error(args ...interface{}) {
	l.Output(3, ERROR, "[ERROR]", fmt.Sprint(args...))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Output(3, ERROR, "[ERROR]", fmt.Sprintf(format, args...))
}

func (l *Logger) Fatal(args ...interface{}) {
	if l.level <= FATAL {
		l.Output(3, FATAL, "[FATAL]", fmt.Sprint(args...))
		os.Exit(1)
	}
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	if l.level <= FATAL {
		l.Output(3, FATAL, "[FATAL]", fmt.Sprintf(format, args...))
		os.Exit(1)
	}
}

func (l *Logger) Panic(args ...interface{}) {
	if l.level <= PANIC {
		s := fmt.Sprint(args...)
		l.Output(3, PANIC, "[PANIC]", s)
		panic(s)
	}
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	if l.level <= PANIC {
		s := fmt.Sprintf(format, args...)
		l.Output(3, PANIC, "[PANIC]", s)
		panic(s)
	}
}

func (l *Logger) Print(v ...interface{}) { //wrap gorm logger
	if len(v) == 0 {
		return
	}

	if v[0] == "log" {
		l.Output(5, ERROR, "[ERROR]", fmt.Sprintln(v[1:]...))
	} else if v[0] == "sql" {
		l.Output(5, DEBUG, "[DEBUG]", fmt.Sprintln(v[1:]...))
	} else if v[0] == "error" {
		l.Output(3, ERROR, "[ERROR]", fmt.Sprintln(v[1:]...))
	} else {
		l.Output(3, INFO, "[INFO ]", fmt.Sprintln(v...))
	}
}

func (l *Logger) Println(v ...interface{}) {
	if len(v) == 0 {
		return
	}
	l.Output(3, INFO, "[INFO ]", fmt.Sprintln(v...))
}
