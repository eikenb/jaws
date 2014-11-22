package userdata

import (
	"bufio"
	"io"
	"strings"
)

const ws = " \r\n\t"

type Userdata interface {
	error
	Set(string, string)
	Del(string)
	Get(string) string
	Reader() io.Reader
	Ok() bool
}

// Return all shell variable export lines (like 'export foo=bar')
func getLines(f io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, ws)
		if !ignore(line) {
			line = strings.TrimPrefix(line, "export")
			lines = append(lines, strings.Trim(line, ws))
		}
	}
	return lines
}

// Ignore empty lines and comments
func ignore(line string) bool {
	return len(line) == 0 ||
		strings.HasPrefix(line, "#") ||
		!strings.HasPrefix(line, "export") ||
		!strings.Contains(line, "=")
}
