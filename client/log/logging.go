package log

import (
	"log"
	"log/syslog"
	"os"
)

// by default, log to stderr
var lg = log.New(os.Stderr, "", log.Ltime|log.Lshortfile)
var Print = lg.Print
var Fatal = lg.Fatal

// log to syslog in production
// set in 'release' file specified with tagging
func Syslog() {
	var priority = syslog.LOG_WARNING | syslog.LOG_USER
	var flags = log.LstdFlags | log.Lshortfile

	var lg, err = syslog.NewLogger(priority, flags)
	if err != nil {
		log.Fatal(err)
	}

	Print = lg.Print
	Fatal = lg.Fatal
}
