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
