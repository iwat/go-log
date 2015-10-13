package log

import (
	"fmt"
	"io/ioutil"
	internal "log"
	"os"
)

var TRC = internal.New(ioutil.Discard, "TRC ", internal.LstdFlags|internal.Lshortfile)
var DBG = internal.New(ioutil.Discard, "DBG ", internal.LstdFlags|internal.Lshortfile)
var NFO = internal.New(os.Stdout, "NFO ", internal.LstdFlags|internal.Lshortfile)
var WRN = internal.New(os.Stderr, "WRN ", internal.LstdFlags|internal.Lshortfile)
var ERR = internal.New(os.Stderr, "ERR ", internal.LstdFlags|internal.Lshortfile)

func init() {
	internal.SetOutput(os.Stdout)
	internal.SetPrefix("NFO ")
	internal.SetFlags(internal.LstdFlags | internal.Lshortfile)
}

func Init(verbose, veryVerbose bool) {
	if veryVerbose {
		TRC = internal.New(os.Stdout, "TRC ", internal.LstdFlags|internal.Lshortfile)
		verbose = true
	}

	if verbose {
		DBG = internal.New(os.Stdout, "DBG ", internal.LstdFlags|internal.Lshortfile)
	}
}

func Trace(v ...interface{}) {
	TRC.Output(3, fmt.Sprint(v...))
}

func Tracef(format string, v ...interface{}) {
	TRC.Output(3, fmt.Sprintf(format, v...))
}

func Traceln(v ...interface{}) {
	TRC.Output(3, fmt.Sprintln(v...))
}

func Debug(v ...interface{}) {
	DBG.Output(3, fmt.Sprint(v...))
}

func Debugf(format string, v ...interface{}) {
	DBG.Output(3, fmt.Sprintf(format, v...))
}

func Debugln(v ...interface{}) {
	DBG.Output(3, fmt.Sprintln(v...))
}

func Info(v ...interface{}) {
	NFO.Output(3, fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
	NFO.Output(3, fmt.Sprintf(format, v...))
}

func Infoln(v ...interface{}) {
	NFO.Output(3, fmt.Sprintln(v...))
}

func Warn(v ...interface{}) {
	WRN.Output(3, fmt.Sprint(v...))
}

func Warnf(format string, v ...interface{}) {
	WRN.Output(3, fmt.Sprintf(format, v...))
}

func Warnln(v ...interface{}) {
	WRN.Output(3, fmt.Sprintln(v...))
}

func Error(v ...interface{}) {
	ERR.Output(3, fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	ERR.Output(3, fmt.Sprintf(format, v...))
}

func Errorln(v ...interface{}) {
	ERR.Output(3, fmt.Sprintln(v...))
}

func Fatal(v ...interface{}) {
	ERR.Output(3, fmt.Sprint(v...))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	ERR.Output(3, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Fatalln(v ...interface{}) {
	ERR.Output(3, fmt.Sprintln(v...))
	os.Exit(1)
}

func Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	ERR.Output(3, s)
	panic(s)
}

func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	ERR.Output(3, s)
	panic(s)
}

func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	ERR.Output(3, s)
	panic(s)
}
