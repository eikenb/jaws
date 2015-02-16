package metadata

import (
	"github.com/eikenb/jaws"
)

// Access to EC2 metadata.
// http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html

type metadata struct {
	ip      string
	version string
	cache   chan map[string]string
}

func New(ip, version string) *metadata {
	cache := make(chan map[string]string, 1)
	cache <- make(map[string]string)
	return &metadata{ip: ip, version: version, cache: cache}
}

// Provide latest version as default.
var Latest = New("169.254.169.254", "latest")
var Lookup = Latest.Lookup

// http client
var client = jaws.New(false)

// Build the URL
func (md metadata) Url(key string) string {
	return "http://" + md.ip + "/" + md.version + "/meta-data/" + key
}

// Get one value from metadata (with key with meta-data/ as root)
func (md metadata) Lookup(key string) (string, error) {
	c := <-md.cache
	defer func() { md.cache <- c }()
	if v, ok := c[key]; ok {
		return v, nil
	}

	r, err := client.Get(md.Url(key))
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
