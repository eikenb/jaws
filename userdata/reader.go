package userdata

import (
	"io"
	"strings"

	"github.com/eikenb/jaws/metadata"
)

var userdata = metadata.New("169.254.169.254", "latest", "user-data")

// Get the userdata file, optionally block until it is present.
// Blocking is here as the userdata file is often not available for a period
// after the instance comes up. This is to wait till it is.
func Reader() (io.Reader, error) {
	s, e := userdata.Lookup("")
	return strings.NewReader(s), e
}
