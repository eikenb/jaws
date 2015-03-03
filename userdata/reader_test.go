package userdata

import (
	"io/ioutil"
	"testing"

	"github.com/eikenb/jaws"
	"github.com/stretchr/testify/assert"
)

var test_data = []byte(`
export FOO=bar
export ZED=die
iam=ignored
me too
#export me=three
`)

func TestReader(t *testing.T) {
	userdata.Mock(jaws.Reply{Status: 200, Body: test_data})

	ud_reader, err := Reader()
	assert.Nil(t, err)
	bod, _ := ioutil.ReadAll(ud_reader)
	assert.Equal(t, string(test_data), string(bod))

	ud_reader, err = Reader()
	assert.Nil(t, err)
	bod, _ = ioutil.ReadAll(ud_reader)
	assert.Equal(t, string(test_data), string(bod))
}
