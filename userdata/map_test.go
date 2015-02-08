package userdata

import (
	"testing"
)

var expected_mapjson = []byte(`{"FOO":"bar","ZED":"die","zoe":"zoom"}`)

func TestUserdataMap(t *testing.T) {
	ud_reader := testData()
	ud := Map(ud_reader)
	commonTests(t, ud, expected_mapjson)
}
