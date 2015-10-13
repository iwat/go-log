package log

import (
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
	TRC.Print(v...)
}

func Tracef(format string, v ...interface{}) {
	TRC.Printf(format, v...)
}

func Traceln(v ...interface{}) {
	TRC.Println(v...)
}

func Debug(v ...interface{}) {
	DBG.Print(v...)
}

func Debugf(format string, v ...interface{}) {
	DBG.Printf(format, v...)
}

func Debugln(v ...interface{}) {
	DBG.Println(v...)
}

func Info(v ...interface{}) {
	NFO.Print(v...)
}

func Infof(format string, v ...interface{}) {
	NFO.Printf(format, v...)
}

func Infoln(v ...interface{}) {
	NFO.Println(v...)
}

func Warn(v ...interface{}) {
	WRN.Print(v...)
}

func Warnf(format string, v ...interface{}) {
	WRN.Printf(format, v...)
}

func Warnln(v ...interface{}) {
	WRN.Println(v...)
}

func Error(v ...interface{}) {
	ERR.Print(v...)
}

func Errorf(format string, v ...interface{}) {
	ERR.Printf(format, v...)
}

func Errorln(v ...interface{}) {
	ERR.Println(v...)
}

func Fatal(v ...interface{}) {
	ERR.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	ERR.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	ERR.Fatalln(v...)
}

func Panic(v ...interface{}) {
	ERR.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	ERR.Panicf(format, v...)
}

func Panicln(v ...interface{}) {
	ERR.Panicln(v...)
}
