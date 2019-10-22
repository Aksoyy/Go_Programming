package main

import (
	"encoding/xml"
	"fmt"
)

type XMLData struct {
	XMLName  xml.Name `xml:"xml'inismi"`
	Element1 string   `xml:"Group1>Element1,omitempty"`
	Element2 string   `xml:"Group1>Element2,omitempty"`
	Element3 string   `xml:"Group2>Example3,omitempty"`
}

func main() {
	foo := &XMLData{}
	foo.Element1 = "Value1"
	foo.Element2 = "Value2"
	foo.Element3 = "Value3"

	fooOut, _ := xml.Marshal(foo)
	fmt.Println(string(fooOut))

	bar := &XMLData{}
	//bar.Element3 = "Value3"
	barOut, _ := xml.Marshal(bar)
	fmt.Println(string(barOut))
}
