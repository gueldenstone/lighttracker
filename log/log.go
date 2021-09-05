package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

type LogLevel uint8

const (
	NONE LogLevel = 0
	STAT LogLevel = 1 << iota
	DEBUG
	INFO
	WARN
	ERROR
)

var l logger

func init() {
	l.out = os.Stdout
	l.level = DEBUG
	l.stat = log.New(l.out, White+"[STAT]  "+Reset, log.LstdFlags)
	l.debug = log.New(l.out, Purple+"[DEBU]  "+Reset, log.LstdFlags)
	l.info = log.New(l.out, Blue+"[INFO]  "+Reset, log.LstdFlags)
	l.warn = log.New(l.out, Yellow+"[WARN]  "+Reset, log.LstdFlags)
	l.err = log.New(l.out, Red+"[ERRO]  "+Reset, log.LstdFlags)
}

type logger struct {
	level LogLevel
	stat  *log.Logger
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	err   *log.Logger
	out   io.Writer
}

func Println(v ...interface{}) {
	level := NONE
	if l, ok := v[0].(LogLevel); ok {
		level = l
	}
	if level == 0 {
		level |= STAT
	}
	if level&STAT == STAT {
		l.stat.Println(v...)
	}
	if level&DEBUG == DEBUG {
		l.debug.Println(v...)
	}
	if level&INFO == INFO {
		l.info.Println(v...)
	}
	if level&WARN == WARN {
		l.warn.Println(v...)
	}
	if level&ERROR == ERROR {
		l.err.Println(v...)
	}
}

func Stat(msg ...interface{}) {
	if l.level&STAT == STAT {
		l.stat.Println(msg...)
	}
}

func Debug(msg ...interface{}) {
	if l.level&DEBUG == DEBUG {
		_, file, no, ok := runtime.Caller(1)
		if !ok {
			l.debug.Println(msg...)
		}
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		l.debug.Printf("%s:%d: %s\n", file, no, fmt.Sprint(msg...))
	}
}
func Info(msg ...interface{}) {
	if l.level&INFO == INFO {
		l.info.Println(msg...)
	}
}
func Warn(msg ...interface{}) {
	if l.level&WARN == WARN {
		l.warn.Println(msg...)
	}
}
func Error(msg ...interface{}) {
	if l.level&ERROR == ERROR {
		l.err.Println(msg...)
	}
}

func Fatal(msg ...interface{}) {
	l.err.Println(msg...)
}

func Statf(format string, args ...interface{}) {
	if l.level&STAT == STAT {
		l.stat.Printf(format, args...)
	}
}

func Debugf(format string, args ...interface{}) {
	if l.level&DEBUG == DEBUG {
		_, file, no, ok := runtime.Caller(1)
		if !ok {
			l.debug.Printf(format, args...)
		}
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		l.debug.Printf(fmt.Sprintf("%s:%d: %s", file, no, format), args...)
	}
}
func Infof(format string, args ...interface{}) {
	if l.level&INFO == INFO {
		l.info.Printf(format, args...)
	}
}
func Warnf(format string, args ...interface{}) {
	if l.level&WARN == WARN {
		l.warn.Printf(format, args...)
	}
}
func Errorf(format string, args ...interface{}) {
	if l.level&ERROR == ERROR {
		l.err.Printf(format, args...)
	}
}

func Fatalf(format string, args ...interface{}) {
	l.err.Printf(format, args...)
	os.Exit(1)
}

//? ... Panic(), Fatal() etc. could also be implemented

//? possibly print to only level specified as parameter? Or as "highest" LogLevel?
func SetOutputLevel(level LogLevel) {
	switch level {
	case STAT:
		l.level = STAT | DEBUG | INFO | WARN | ERROR
	case DEBUG:
		l.level = DEBUG | INFO | WARN | ERROR
	case INFO:
		l.level = INFO | WARN | ERROR
	case WARN:
		l.level = WARN | ERROR
	case ERROR:
		l.level = ERROR
	}
}
