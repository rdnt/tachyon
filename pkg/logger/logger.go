package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"

	"github.com/rdnt/tachyon/pkg/syncwriter"

	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
)

// Color is an alias to aurora's color
type Color aurora.Color

const (
	Reset = Color(0)

	BlackFg   = Color(aurora.BlackFg)
	RedFg     = Color(aurora.RedFg)
	GreenFg   = Color(aurora.GreenFg)
	YellowFg  = Color(aurora.YellowFg)
	BlueFg    = Color(aurora.BlueFg)
	MagentaFg = Color(aurora.MagentaFg)
	CyanFg    = Color(aurora.CyanFg)
	WhiteFg   = Color(aurora.WhiteFg)

	BlackBg   = Color(aurora.BlackBg)
	RedBg     = Color(aurora.RedBg)
	GreenBg   = Color(aurora.GreenBg)
	YellowBg  = Color(aurora.YellowBg)
	BlueBg    = Color(aurora.BlueBg)
	MagentaBg = Color(aurora.MagentaBg)
	CyanBg    = Color(aurora.CyanBg)
	WhiteBg   = Color(aurora.WhiteBg)
)

var (
	stdoutWriter *syncwriter.Writer
	stderrWriter *syncwriter.Writer
)

type Logger struct {
	sync.Mutex
	name   string
	stdout *log.Logger
	stderr *log.Logger
}

func init() {
	stdoutWriter = syncwriter.New(colorable.NewColorable(os.Stdout))
	stderrWriter = syncwriter.New(colorable.NewColorable(os.Stderr))
}

func colorize(s string, c Color) string {
	return aurora.Colorize(s, aurora.Color(c)).String()
}

func New(name string, color Color) *Logger {
	if name != "" {
		name = fmt.Sprintf("[%s] ", name)
		name = colorize(name, color)
	}

	return &Logger{
		name:   name,
		stdout: log.New(stdoutWriter, "", 0),
		stderr: log.New(stderrWriter, "", 0),
	}
}

var cwd *string

func init() {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	wd = filepath.ToSlash(wd + "/")
	cwd = &wd
}

func (l *Logger) prefix(caller bool) string {
	prefix := l.name
	if caller {
		_, file, line, ok := runtime.Caller(3)
		if !ok {
			file = "???"
			line = 0
		}
		if cwd != nil {
			file = strings.Replace(file, *cwd, "", 1)
		}
		prefix = fmt.Sprintf("%s%s:%d ", prefix, file, line)
	}
	return prefix
}

func (l *Logger) print(s string) {
	l.Lock()
	defer l.Unlock()

	s = strings.TrimRight(s, "\n")
	_ = l.stdout.Output(3, l.prefix(false)+s)
}

func (l *Logger) debugPrint(s string) {
	l.Lock()
	defer l.Unlock()

	s = strings.TrimRight(s, "\n")
	_ = l.stdout.Output(3, l.prefix(false)+s)
}

func (l *Logger) errorPrint(s string) {
	l.Lock()
	defer l.Unlock()

	s = strings.TrimRight(s, "\n")
	s = colorize(s, RedFg)
	_ = l.stderr.Output(3, l.prefix(true)+s)
}

func (l *Logger) tracePrint() {
	l.Lock()
	defer l.Unlock()

	stack := debug.Stack()
	s := colorize(string(stack), RedFg)
	_ = l.stderr.Output(3, l.prefix(true)+s)
}

func (l *Logger) Debug(v ...interface{}) {
	l.print(fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.print(fmt.Sprintf(format, v...))
}

func (l *Logger) Debugln(v ...interface{}) {
	l.print(fmt.Sprintln(v...))
}

func (l *Logger) Print(v ...interface{}) {
	l.debugPrint(fmt.Sprint(v...))
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.debugPrint(fmt.Sprintf(format, v...))
}

func (l *Logger) Println(v ...interface{}) {
	l.debugPrint(fmt.Sprintln(v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.errorPrint(fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.errorPrint(fmt.Sprintf(format, v...))
}

func (l *Logger) Errorln(v ...interface{}) {
	l.errorPrint(fmt.Sprintln(v...))
}

func (l *Logger) Trace() {
	l.tracePrint()
}
