package jaws

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSucess(t *testing.T) {
	yesman := func() error { return nil }
	assert.Nil(t, Timeout(100).Run(yesman))
	assert.Nil(t, Timeout(100).Retry(yesman, 2))
}
func TestError(t *testing.T) {
	err := errors.New("foo")
	troublemaker := func() error { return err }
	assert.Equal(t, timeout(100).Run(troublemaker), err)
	multiplier = time.Microsecond
	assert.Equal(t, timeout(100).Retry(troublemaker, 2), err)
}
func TestTimeout(t *testing.T) {
	multiplier = time.Microsecond
	slow := func() error {
		time.Sleep(time.Millisecond)
		return nil
	}
	assert.Equal(t, timeout(1).Run(slow), TimeoutError)
	assert.Equal(t, timeout(1).Retry(slow, 2), TimeoutError)
	multiplier = time.Second
}
