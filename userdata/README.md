
A very simple little library to read a file of shell script style variables. It
takes input as an io.Reader, reads it in, strips spurious data and converts it
to a map[string]string. It only supports some minor changing and exporting to
Json.

I used this for AWS ASG user-data. It stores data in a couple different files
[1] or makes it accessible from the meta data url [2]. I wanted it as simple as
possible.

[1] /etc/default/ec2-user-data  /etc/ec2-user-data.txt
[2] curl http://169.254.169.254/2008-02-01/user-data


Example usage:

    package main

    import (
        "fmt"
        "io/ioutil"
        "os"

        "github.com/eikenb/jaws/userdata"
    )

    func main() {
        f, _ := os.Open("/etc/ec2-user-data.txt")
        ud := userdata.New(f)
        json_bytes, _ := ioutil.ReadAll(ud)
        fmt.Println(string(json_bytes))
    }
