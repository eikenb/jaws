// +build !release

package log

import (
	"log"
	"os"
)

// better logger for testing
var lg = log.New(os.Stderr, "", log.Ltime|log.Lshortfile)
var Print = lg.Print
var Fatal = lg.Fatal
