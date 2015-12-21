Warning
-------

I stopped developing this library as I had switched to use github.com/inconshreveable/log15
instead.

go-log
======

Simple log facility for all related projects.

Usage
-----

    import (
        "flag"
        "github.com/iwat/log"
    )
    
    var flagVerbose = false
    var flagVeryVerbose = false
    
    func init() {
        flag.BoolVar(&flagVerbose, "v", false, "Verbose")
        flag.BoolVar(&flagVeryVerbose, "vv", false, "Very verbose, implies -v")
    }
    
    func main() {
        flag.Parse()
        
        log.Init(flagVerbose, flagVeryVerbose)
        
        log.TRC.Println("Trace level")
        log.DBG.Println("Debug level")
        log.NFO.Println("Info level")
        log.WRN.Println("Warning level")
        log.ERR.Println("Error level")
        log.ERR.Fatalln("Error & exit")
    }

Legal
-----

This application is developed under MIT license, and can be used for open and
proprietary projects.

Copyright (c) 2015 Chaiwat Shuetrakoonpaiboon

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
