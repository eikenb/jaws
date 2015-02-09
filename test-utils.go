package jaws

import (
	"bytes"
	"io"
	"net/http"
)

//----------------------------------------------------------------------
// test helper code

// Reply implements the doer interface, so it can be swapped in for the http
// client in and return mock data for testing.
type Reply struct {
	Status    int
	Body      []byte
	Reqtester func(*http.Request)
}

// Mimics the behaviour of http.Client's Do() method.
func (r Reply) Do(req *http.Request) (*http.Response, error) {
	if r.Reqtester != nil {
		r.Reqtester(req)
	}
	resp := &http.Response{StatusCode: r.Status, Body: Testbody(r.Body)}
	return resp, nil
}

//
type replies []Reply

func (rs *replies) Do(req *http.Request) (*http.Response, error) {
	if len(*rs) == 0 {
		panic("out of replies")
	}
	rsv := *rs
	r := rsv[0]
	*rs = rsv[1:]
	return r.Do(req)
}

func NewReplies(rs ...Reply) *replies {
	rpls := replies{}
	rpls = append(rpls, rs...)
	return &rpls
}

type body struct{ buf *bytes.Buffer }

func Testbody(b []byte) io.ReadCloser      { return &body{bytes.NewBuffer(b)} }
func (b body) Read(bs []byte) (int, error) { return b.buf.Read(bs) }
func (b body) Close() error                { return nil }
