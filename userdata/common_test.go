package userdata

import (
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testData() io.Reader {
	Ec2UserdataPath = "./testdata/ec2-user-data"
	var ud_reader, _ = Reader(false)
	return ud_reader
}

func commonTests(t *testing.T, ud Userdata, expected_json []byte) {
	ud.Set("zoe", "zoom")
	out, _ := ioutil.ReadAll(ud.Reader())
	assert.Equal(t, string(out), string(expected_json))
	assert.Equal(t, "die", string(ud.Get("ZED")))
	ud.Del("zoe")
	assert.Equal(t, "", string(ud.Get("zoe")))
}
