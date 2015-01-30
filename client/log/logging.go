package log

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
)

// by default, log to stderr
var lg = log.New(os.Stderr, "", log.Ltime|log.Lshortfile)

// Our logging calls, needed to override default calldepth
func Print(v ...interface{}) { lg.Output(3, fmt.Sprint(v...)) }
func Fatal(v ...interface{}) {
	lg.Output(3, fmt.Sprint(v...))
	os.Exit(1)
}

// log to syslog in production
// set in 'release' file specified with tagging
func Syslog() {
	var priority = syslog.LOG_WARNING | syslog.LOG_USER
	var flags = log.LstdFlags | log.Lshortfile
	var err error
	lg, err = syslog.NewLogger(priority, flags)
	if err != nil {
		log.Fatal(err)
	}

}
