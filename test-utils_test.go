package jaws

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Use these structures and methods to help in testing. They make it easy to
// mock out results and test input requests. See the other tests in this and
// the submodules for more examples.

func TestDoer(t *testing.T) {
	doer := (*doer)(nil)
	assert.Implements(t, doer, &http.Client{})
	assert.Implements(t, doer, Reply{})
	assert.Implements(t, doer, &replies{})
}

func TestReply(t *testing.T) {
	Aws.Cli = Reply{Status: 1, Body: []byte("foo"),
		Reqtester: func(r *http.Request) {
			assert.Equal(t, r.URL.Host, "foo.net")
			assert.Equal(t, r.URL.Scheme, "http")
			assert.Equal(t, r.URL.Path, "/bar")
		}}
	r, e := Aws.Get("http://foo.net/bar")
	assert.Equal(t, e, BadStatus)
	body, _ := ReadBody(r)
	assert.Equal(t, body, []byte("foo"))
}

func TestReplies(t *testing.T) {
	Aws.Cli = NewReplies(Reply{Status: 201}, Reply{Status: 202})
	r, _ := Aws.Get("http://foo.net/bar")
	assert.Equal(t, r.StatusCode, 201)
	r, _ = Aws.Get("http://foo.net/bar")
	assert.Equal(t, r.StatusCode, 202)
	assert.Panics(t, func() { Aws.Cli.Do(&http.Request{}) },
		"Should panic if more replies are needed than provided.")
}
