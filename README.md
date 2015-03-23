Collection of small libraries for use with AWS. Added here to enable sharing
and easier reuse.

jaws client
-----------
Top level is an AWS client. Very simple net/http like set of functions that use
[go-aws-auth] (http://github.com/smartystreets/go-aws-auth) for AWS
authentication, handles timeouts, and provides a way to easily create mocked
responses.

userdata
--------
A lib to parse a shell formatted EC2 user-data string of exported environment
variables. It pulls the data from the instance metadata url (see below) and
converts it to either a map[string]string struct or an slice of name/value
structs (useful for json encoding).

metadata
--------
A lib to access the meta-data of an EC2 AWS instance. It hits the local 169
address and caches the results. It is designed to return single values, not a
struct of all the metadata.


