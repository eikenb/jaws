package jaws

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContentLengthable(t *testing.T) {
	rawdata := []byte{'f', 'o', 'o'}

	newdata := contentLengthable(Testbody(rawdata))
	var buf *bytes.Buffer
	assert.IsType(t, buf, newdata)

	rdr := bytes.NewReader(rawdata)
	newdata = contentLengthable(rdr)
	assert.IsType(t, rdr, newdata)
}
