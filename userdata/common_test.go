package userdata

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_userdata = []byte(`
export FOO=bar
export ZED=die
iam=ignored
me too
#export me=three
`)

func commonTests(t *testing.T, ud Userdata, expected_json []byte) {
	ud.Set("zoe", "zoom")
	out, _ := ioutil.ReadAll(ud.Reader())
	assert.Equal(t, string(out), string(expected_json))
	assert.Equal(t, "die", string(ud.Get("ZED")))
	ud.Del("zoe")
	assert.Equal(t, "", string(ud.Get("zoe")))
}
