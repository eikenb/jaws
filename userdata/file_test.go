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

func TestReader(t *testing.T) {
	Ec2UserdataPath = "./testdata/ec2-user-data"
	ud_reader, err := Reader(false)
	assert.Nil(t, err)
	bod, _ := ioutil.ReadAll(ud_reader)
	assert.Equal(t, string(test_userdata), string(bod))

	ud_reader, err = Reader(true)
	assert.Nil(t, err)
	bod, _ = ioutil.ReadAll(ud_reader)
	assert.Equal(t, string(test_userdata), string(bod))
}
