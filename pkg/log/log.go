package log

import (
	"tachyon/pkg/logger"
)

var l = logger.New("", logger.Reset)

func Debug(v ...interface{}) {
	l.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	l.Debugf(format, v...)
}

func Debugln(v ...interface{}) {
	l.Debugln(v...)
}

func Print(v ...interface{}) {
	l.Print(v...)
}

func Printf(format string, v ...interface{}) {
	l.Printf(format, v...)
}

func Println(v ...interface{}) {
	l.Println(v...)
}

func Error(v ...interface{}) {
	l.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	l.Errorf(format, v...)
}

func Errorln(v ...interface{}) {
	l.Errorln(v...)
}

func Trace() {
	l.Trace()
}
