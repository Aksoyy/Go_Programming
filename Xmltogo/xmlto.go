import (
  	...
    "encoding/xml"
    ...
)

type Users struct {
    XMLName xml.Name `xml:"users"`
    Users   []User   `xml:"user"`
}

type User struct {
    XMLName xml.Name `xml:"user"`
    Type    string   `xml:"type,attr"`
    Name    string   `xml:"name"`
    Social  Social   `xml:"social"`
}

type Social struct {
    XMLName  xml.Name `xml:"social"`
    Facebook string   `xml:"facebook"`
    Twitter  string   `xml:"twitter"`
    Youtube  string   `xml:"youtube"`
}