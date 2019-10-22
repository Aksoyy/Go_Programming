package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
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

func main() {

	xmlFile, err := os.Open("users.xml")
	if err != nil {
		fmt.Println("test")
	} else {
		fmt.Println("Successfully Opened users.xml")
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	//fmt.Println("gelenveri: ", string(byteValue))

	var users Users
	xml.Unmarshal(byteValue, &users)

	// EXAMPLE
	// fmt.Println(users.Users[0].Name) --> Elliot
	// fmt.Println(users.Users[1].Name) --> Fraser
	// fmt.Println(users.Users[2].Name) ---> index out of range

	fmt.Println(len(users.Users))
	fmt.Println(len(users.Users[1].Name))
	for i := 0; i < len(users.Users); i++ {
		// fmt.Println("User Type: " + users.Users[i].Type)
		// fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}
}
