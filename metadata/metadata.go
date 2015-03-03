package metadata

import (
	"strings"

	"github.com/eikenb/jaws"
)

// Access to EC2 metadata.
// http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html

type metadata struct {
	ip      string
	version string
	path    string
	client  *jaws.Client
	cache   chan map[string]string
}

func New(ip, version, path string) *metadata {
	cache := make(chan map[string]string, 1)
	cache <- make(map[string]string)
	client := jaws.New(false)
	return &metadata{ip, version, path, client, cache}
}

// Provide latest version as default.
var Latest = New("169.254.169.254", "latest", "meta-data")
var Lookup = Latest.Lookup

// Build the URL
func (md *metadata) Url(key string) string {
	return "http://" +
		strings.Join([]string{md.ip, md.version, md.path, key}, "/")
}

// Get one value from metadata (with key with meta-data/ as root)
func (md *metadata) Lookup(key string) (string, error) {
	c := <-md.cache
	defer func() { md.cache <- c }()
	if v, ok := c[key]; ok {
		return v, nil
	}

	r, err := md.client.Get(md.Url(key))
	var body []byte
	if err == nil {
		body, err = jaws.ReadBody(r)
	}
	if err != nil {
		return "", err
	} // no metadata == bad
	c[key] = string(body)
	return c[key], nil
}

func (md *metadata) Mock(replies ...jaws.Reply) {
	if len(replies) == 1 {
		md.client.Cli = replies[0]
	}
	md.client.Cli = jaws.NewReplies(replies...)
}
