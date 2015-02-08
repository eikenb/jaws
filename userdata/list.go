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
	u := &UdList{envmap}
	return u
}

// Set a value (uniform api with map)
func (u *UdList) Set(key, value string) {
	u.Data = append(u.Data, pair{key, value})
}

// Get a value (uniform api with map)
func (u *UdList) Get(key string) string {
	for _, p := range u.Data {
		if p.Name == key {
			return p.Value
		}
	}
	return ""
}

// Delete a value (uniform api with map)
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

// Json reader for list
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
}

func makeList(lines []string) []pair {
	pairs := make([]pair, 0, len(lines))
	for _, line := range lines {
		split := strings.SplitN(line, "=", 2)
		pairs = append(pairs, pair{split[0], split[1]})
	}
	return pairs
}
