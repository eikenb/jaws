package client

import (
	"errors"
	"time"
)

type timeout int

// variable so it can be changed for testing
var multiplier = time.Second

// set time-out (default: in seconds)
func Timeout(t int) timeout { return timeout(t) }

var TimeoutError = errors.New("Timeout")

// run passed in func()error with timeout
// return error from func(), timeout error or nil
func (t timeout) Run(f func() error) (e error) {
	done := make(chan error, 1)
	go func() { done <- f() }()
	select {
	case e = <-done:
		return e
	case <-time.After(time.Duration(t) * multiplier):
		return TimeoutError
	}
}

// Try to Run() the function N times before returning the error
// Sleep 'number of try' * multiplier each time through.
func (t timeout) Retry(f func() error, N int) (e error) {
	for i := 0; i < N; i++ {
		if e = t.Run(f); e == nil {
			return
		}
		time.Sleep(multiplier * time.Duration(i+1))
	}
	return e
}
