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
A lib to parse a shell formatted ec2-user-data file written by an AWS ASG on
the instance. It can read the default ec2-user-data file off the system or be
provided an alternative. It only understands exported shell variables as might
be included by shell scripts as this is a common format as it has multiple
uses.

metadata
--------
A lib to access the meta-data of an EC2 AWS instance. It hits the local 169
address and caches the results.


