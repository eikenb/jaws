package metadata

import (
	"testing"

	"github.com/eikenb/jaws"
	"github.com/stretchr/testify/assert"
)

func TestLookup(t *testing.T) {
	Latest.Mock(jaws.Reply{Status: 200, Body: []byte("us-west-2a")})
	val, err := Lookup("placement/availability-zone")
	assert.Nil(t, err)
	assert.Equal(t, val, "us-west-2a")
	// test cache
	Latest.Mock(jaws.Reply{Status: 200, Body: []byte("")})
	val, err = Lookup("placement/availability-zone")
	assert.Nil(t, err)
	assert.Equal(t, val, "us-west-2a")
}

func TestCache(t *testing.T) {
	Latest.Mock(jaws.Reply{Status: 200, Body: []byte("us-west-2a")})
	Lookup("placement/availability-zone") // prime cache
	// empty Get to fail if cache doesn't work
	Latest.Mock(jaws.Reply{Status: 200, Body: []byte("")})
	val, err := Lookup("placement/availability-zone")
	assert.Nil(t, err)
	assert.Equal(t, val, "us-west-2a")
}
