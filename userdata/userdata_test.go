package userdata

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_userdata = []byte(`
FOO=bar
ZED=die
hi
`)

var expected_json = []byte(`{
  "FOO": "bar",
  "ZED": "die",
  "zoe": "zoom"
}`)

func TestUserdata(t *testing.T) {
	ud_reader := bytes.NewReader(test_userdata)
	ud := NewUserdata(ud_reader)
	ud.Set("zoe", "zoom")
	out, _ := ioutil.ReadAll(ud)
	assert.Equal(t, out, expected_json)
	assert.Equal(t, "die", ud.Get("ZED"))
}
