package metadata

import (
	"net/http"
	"testing"

	"github.com/eikenb/jaws/client"
	"github.com/stretchr/testify/assert"
)

func TestLookup(t *testing.T) {
	Get = func(_ string) (*http.Response, error) {
		resp := &http.Response{
			StatusCode: 200,
			Body:       client.Testbody([]byte("us-west-2a")),
		}
		return resp, nil
	}
	val, err := Lookup("placement/availability-zone")
	assert.Nil(t, err)
	assert.Equal(t, val, "us-west-2a")
	// test cache
	Get = func(_ string) (*http.Response, error) { return nil, nil }
	val, err = Lookup("placement/availability-zone")
	assert.Nil(t, err)
	assert.Equal(t, val, "us-west-2a")
}

func TestCache(t *testing.T) {
	Get = func(_ string) (*http.Response, error) {
		resp := &http.Response{
			StatusCode: 200,
			Body:       client.Testbody([]byte("us-west-2a")),
		}
		return resp, nil
	}
	Lookup("placement/availability-zone") // prime cache
	// empty Get to fail if cache doesn't work
	Get = func(_ string) (*http.Response, error) { return nil, nil }
	val, err := Lookup("placement/availability-zone")
	assert.Nil(t, err)
	assert.Equal(t, val, "us-west-2a")
}
