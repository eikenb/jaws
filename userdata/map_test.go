package userdata

import (
	"bytes"
	"testing"
)

var expected_mapjson = []byte(`{"FOO":"bar","ZED":"die","zoe":"zoom"}`)

func TestUserdataMap(t *testing.T) {
	ud_reader := bytes.NewReader(test_userdata)
	ud := Map(ud_reader)
	commonTests(t, ud, expected_mapjson)
}
