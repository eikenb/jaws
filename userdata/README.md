
A very simple little library to read a file of shell script style variables as
is exported by AWS ASGs (auto-scaling-groups), though it would work for any
file containing shell variables.

The user-data is obtained either by reading one of a couple different files [1]
or by grabbing the output from the meta data url [2].

[1] /etc/default/ec2-user-data  /etc/ec2-user-data.txt
[2] curl http://169.254.169.254/2008-02-01/user-data

