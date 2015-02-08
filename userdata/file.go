package userdata

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"time"
)

// Location of ec2-user-data file.
// change this for testing or when otherwise desired
var Ec2UserdataPath = "/etc/default/ec2-user-data"

// Limit retrying to something so we will eventually exit with an error
var Retry = 60
var Sleep = time.Second

// Eventual error given if reading fails
var Timeout = errors.New("Timeout trying to read " + Ec2UserdataPath)

// Get the userdata file, optionally block until it is present.
// Blocking is here as the userdata file is often not available for a period
// after the instance comes up. This is to wait till it is.
func Reader(block bool) (io.Reader, error) {
	if !block {
		return readdata(Ec2UserdataPath)
	}
	for i := 0; i < Retry; i++ {
		buf, err := readdata(Ec2UserdataPath)
		if err == nil && buf.Len() > 0 {
			return buf, err
		}
		log.Println("No userdata... Retrying.")
		time.Sleep(Sleep)
	}
	return nil, Timeout
}

func readdata(path string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	fp, err := os.Open(path)
	defer fp.Close()
	if err == nil {
		buf.ReadFrom(fp)
	}
	return buf, err

}
