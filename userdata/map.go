package userdata

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

func NewMap(in io.Reader) *Map {
	lines := getLines(in)
	envmap := makeMap(lines)
	u := &Map{envmap, nil}
	return u
}

func (u *Map) Set(key, value string) {
	u.data[key] = value
}

func (u *Map) Get(key string) string {
	return u.data[key]
}

func (u *Map) Del(key string) { delete(u.data, key) }

func (u *Map) Ok() bool { return u.err == nil }

func (u *Map) Error() string {
	if !u.Ok() {
		return u.err.Error()
	}
	return ""
}

func (u *Map) Reader() io.Reader {
	js, err := json.Marshal(u.data)
	if err != nil {
		return strings.NewReader(err.Error())
	}
	return bytes.NewReader(js)
}

// --------------------------------------------------------------------

type Map struct {
	data map[string]string
	err  error
}

func makeMap(lines []string) map[string]string {
	em := make(map[string]string)
	for _, line := range lines {
		split := strings.SplitN(line, "=", 2)
		em[split[0]] = split[1]
	}
	return em
}
