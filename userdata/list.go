package userdata

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

func List(in io.Reader) *UdList {
	lines := getLines(in)
	envmap := makeList(lines)
	u := &UdList{envmap, nil}
	return u
}

func (u *UdList) Set(key, value string) {
	u.Data = append(u.Data, pair{key, value})
}

func (u *UdList) Get(key string) string {
	for _, p := range u.Data {
		if p.Name == key {
			return p.Value
		}
	}
	return ""
}

func (u *UdList) Del(key string) {
	idx := len(u.Data)
	for i, p := range u.Data {
		if p.Name == key {
			idx = i
			break
		}
	}
	u.Data[idx] = u.Data[len(u.Data)-1]
	u.Data = u.Data[:len(u.Data)-1]
}

func (u *UdList) Ok() bool { return u.err == nil }

func (u *UdList) Error() string {
	if !u.Ok() {
		return u.err.Error()
	}
	return ""
}

func (u *UdList) Reader() io.Reader {
	js, err := json.Marshal(u)
	if err != nil {
		return strings.NewReader(err.Error())
	}
	return bytes.NewReader(js)
}

// --------------------------------------------------------------------

type pair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type UdList struct {
	Data []pair `json:"data"`
	err  error
}

func makeList(lines []string) []pair {
	pairs := make([]pair, 0, len(lines))
	for _, line := range lines {
		split := strings.SplitN(line, "=", 2)
		pairs = append(pairs, pair{split[0], split[1]})
	}
	return pairs
}
