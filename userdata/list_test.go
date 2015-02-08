package userdata

import (
	"bytes"
	"testing"
)

var expected_listjson = []byte(
	`{"data":[{"name":"FOO","value":"bar"},{"name":"ZED","value":"die"},{"name":"zoe","value":"zoom"}]}`)

func TestUserdataList(t *testing.T) {
	ud_reader := bytes.NewReader(test_userdata)
	ud := List(ud_reader)
	commonTests(t, ud, expected_listjson)
}
