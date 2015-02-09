package jaws

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContentLengthable(t *testing.T) {
	data := testbody([]byte{'f', 'o', 'o'})
	newdata := contentLengthable(data)
	var buf *bytes.Buffer
	assert.IsType(t, buf, newdata)
}
