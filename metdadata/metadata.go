package metadata

import (
	"net/http"

	"github.com/eikenb/jaws/client"
)

// Access to EC2 metadata.
// http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html

// make http.Get() a var to allow for swapping out in tests
var Get = func(key string) (*http.Response, error) {
	return http.Get("http://169.254.169.254/latest/meta-data/" + key)
}

// Get one value from metadata (with key with meta-data/ as root)
func Lookup(key string) (string, error) {
	c := <-cache
	defer func() { cache <- c }()
	if v, ok := c[key]; ok {
		return v, nil
	}

	var r *http.Response
	var body []byte
	var err error
	get := func() error { r, err = Get(key); return err }
	err = client.Timeout(1).Retry(get, 3)
	if err == nil {
		body, err = client.ReadBody(r)
	}
	if err != nil {
		return "", err
	} // no metadata == very bad
	c[key] = string(body)
	return c[key], nil
}

// cache in a channel because I hate locks
var cache = make(chan map[string]string, 1)

func init() { cache <- make(map[string]string) }
