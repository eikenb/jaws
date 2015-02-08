package userdata

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

func Map(in io.Reader) *UdMap {
	lines := getLines(in)
	envmap := makeMap(lines)
	u := &UdMap{envmap, nil}
	return u
}

func (u *UdMap) Set(key, value string) { u.data[key] = value }

func (u *UdMap) Get(key string) string { return u.data[key] }

func (u *UdMap) Del(key string) { delete(u.data, key) }

func (u *UdMap) Ok() bool { return u.err == nil }

func (u *UdMap) Error() string {
	if !u.Ok() {
		return u.err.Error()
	}
	return ""
}

func (u *UdMap) Reader() io.Reader {
	js, err := json.Marshal(u.data)
	if err != nil {
		return strings.NewReader(err.Error())
	}
	return bytes.NewReader(js)
}

// --------------------------------------------------------------------

type UdMap struct {
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
