package userdata

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

// Make the userdata as a dict-like map.
func Map(in io.Reader) *UdMap {
	lines := getLines(in)
	envmap := makeMap(lines)
	u := &UdMap{envmap}
	return u
}

// Set a value (uniform api with list)
func (u *UdMap) Set(key, value string) { u.Data[key] = value }

// Get a value (uniform api with list)
func (u *UdMap) Get(key string) string { return u.Data[key] }

// Delete a value (uniform api with list)
func (u *UdMap) Del(key string) { delete(u.Data, key) }

// Json reader for list
func (u *UdMap) Reader() io.Reader {
	js, err := json.Marshal(u.Data)
	if err != nil {
		return strings.NewReader(err.Error())
	}
	return bytes.NewReader(js)
}

// --------------------------------------------------------------------

type UdMap struct {
	Data map[string]string
}

func makeMap(lines []string) map[string]string {
	em := make(map[string]string)
	for _, line := range lines {
		split := strings.SplitN(line, "=", 2)
		em[split[0]] = split[1]
	}
	return em
}
