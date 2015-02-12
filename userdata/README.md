
A very simple little library to read a file of shell script style variables. It
takes input as an io.Reader, reads it in, strips spurious data and converts it
to a map[string]string. It only supports some minor changing and exporting to
Json.

I used this for AWS ASG user-data. It stores data in a couple different files
or makes it accessible from the meta data url [1]. I used one of the files as
default at right now as I wanted it as simple as possible. I should probably
change it to use the URL though as it is the only real standard source (the
files can vary per OS).

[1] curl http://169.254.169.254/2008-02-01/user-data

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
