// +build release

package log

import (
	"log"
	"log/syslog"
)

var priority = syslog.LOG_WARNING | syslog.LOG_USER
var flags = log.LstdFlags | log.Lshortfile

// Log to syslog in production
var lg, err = syslog.NewLogger(priority, flags)

var Print = lg.Print
var Fatal = lg.Fatal

func init() {
	if err != nil {
		log.Fatal(err)
	}
}
