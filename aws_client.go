package jaws

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/eikenb/jaws/log"
	"github.com/smartystreets/go-aws-auth"
)

//----------------------------------------------------------------------
// aws http authed client
type doer interface {
	Do(*http.Request) (*http.Response, error)
}

// custom client to allow replacing doer for testing
type client struct {
	Cli doer
}

var Aws *client

func init() { Aws = New() }

func New() *client { return &client{Cli: &http.Client{}} }

// client.Do with auth and timeout added
func (c client) do(req *http.Request) (r *http.Response, e error) {
	awsauth.Sign4(req)
	do := func() error { r, e = c.Cli.Do(req); return e }
	if e := Timeout(5).Retry(do, 3); e == nil {
		if r.StatusCode < 200 || r.StatusCode > 299 {
			e = errors.New("HTTP call failed with status: " + r.Status)
		}
		return r, e
	}
	if e == TimeoutError {
		e = errors.New("Timeout connecting to AWS")
	}
	return r, e
}

// get + better do
func (c client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

// post w/ content-type + better do
func (c client) Post(url string, data url.Values) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.do(req)
}

// put w/ content-type, date headers + better do
func (c client) Put(url string, data io.Reader) (*http.Response, error) {
	data = contentLengthable(data)
	req, err := http.NewRequest("PUT", url, data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "text/json")
	req.Header.Set("Date", time.Now().UTC().Format(http.TimeFormat))
	return c.do(req)
}

// make sure data passed in is of a type that built-in http client
// can calculate its content-length, if not make it one
func contentLengthable(data io.Reader) io.Reader {
	switch data.(type) {
	case *bytes.Buffer, *bytes.Reader, *strings.Reader:
	default: // convert to type that calculates content-length
		b := new(bytes.Buffer)
		b.ReadFrom(data)
		data = b
	}
	return data
}

//----------------------------------------------------------------------
// utility functions
func NewAwsParams(action, version string) url.Values {
	vals := make(url.Values)
	vals.Set("Action", action)
	vals.Set("Version", version)
	return vals
}

// util function to read the body of a response
func ReadBody(resp *http.Response) ([]byte, error) {
	if resp == nil {
		return nil, errors.New("nil response object")
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
