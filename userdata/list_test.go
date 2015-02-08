package userdata

import (
	"testing"
)

var expected_listjson = []byte(
	`{"data":[{"name":"FOO","value":"bar"},{"name":"ZED","value":"die"},{"name":"zoe","value":"zoom"}]}`)

func TestUserdataList(t *testing.T) {
	ud_reader := testData()
	ud := List(ud_reader)
	commonTests(t, ud, expected_listjson)
}
