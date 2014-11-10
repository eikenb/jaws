package userdata

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

func New(in io.Reader) *Userdata {
	lines := getLines(in)
	envmap := getMap(lines)
	u := &Userdata{envmap, nil, nil}
	u.setJson()
	return u
}

func (u *Userdata) Set(key, value string) {
	u.data[key] = value
	u.setJson()
}

func (u *Userdata) Get(key string) string {
	return u.data[key]
}

func (u *Userdata) Del(key string) {
	delete(u.data, key)
	u.setJson()
}

func (u *Userdata) Ok() bool { return u.err == nil }

func (u *Userdata) Error() string {
	if !u.Ok() {
		return u.err.Error()
	}
	return ""
}

func (u *Userdata) Read(p []byte) (int, error) {
	return u.reader.Read(p)
}

// --------------------------------------------------------------------

type Userdata struct {
	data   map[string]string
	err    error
	reader *bytes.Reader
}

func (u *Userdata) setJson() {
	u.err, u.reader = nil, nil // reset json state
	result, err := json.MarshalIndent(u.data, "", "  ")
	if err != nil {
		u.err = err
	} else {
		u.reader = bytes.NewReader(result)
	}
}

func getLines(f io.Reader) []string {
	var lines []string
	var line string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line = strings.TrimPrefix(scanner.Text(), "export")
		line = strings.Trim(line, " \r\n\t")
		if !ignore(line) {
			lines = append(lines, line)
		}
	}
	return lines
}

func getMap(lines []string) map[string]string {
	em := make(map[string]string)
	for _, line := range lines {
		split := strings.SplitN(line, "=", 2)
		em[split[0]] = split[1]
	}
	return em
}

// Ignore empty lines and comments
func ignore(line string) bool {
	return len(line) == 0 ||
		strings.HasPrefix(line, "#") ||
		!strings.Contains(line, "=")
}
