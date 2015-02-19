package jaws

import (
	"bytes"
	"net/http"
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

var content = []byte("foo")
var domain = "http://foo.com"

func TestGet(t *testing.T) {
	Aws.Cli = Reply{Status: 200, Body: content}
	r, e := Aws.Get(domain)
	assert.Nil(t, e)
	body, e := ReadBody(r)
	assert.Nil(t, e)
	assert.Equal(t, body, content)
}

func TestPost(t *testing.T) {
	assert := assert.New(t)
	Aws.Cli = Reply{Status: 200, Body: content,
		Reqtester: func(r *http.Request) {
			assert.Equal(r.URL.Host, "foo.com")
			assert.Equal(r.Header.Get("Content-Type"),
				"application/x-www-form-urlencoded")
		}}
	_, e := Aws.Post(domain, NewAwsParams("foo", "1"))
	assert.Nil(e)
}

func TestPut(t *testing.T) {
	assert := assert.New(t)
	Aws.Cli = Reply{Status: 200, Body: content,
		Reqtester: func(r *http.Request) {
			assert.Equal(r.URL.Host, "foo.com")
			assert.Equal(r.Header.Get("Content-Type"), "text/json")
		}}
	_, e := Aws.Put(domain, bytes.NewReader(content))
	assert.Nil(e)
}

func TestErr(t *testing.T) {
	Aws.Cli = Reply{Status: 500, Body: content}
	r, e := Aws.Get(domain)
	assert.Equal(t, r.StatusCode, 500)
	assert.Equal(t, e.Error(), "HTTP call failed bad status.")
	body, e := ReadBody(r)
	assert.Nil(t, e)
	assert.Equal(t, body, content)
}
