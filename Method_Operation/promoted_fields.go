package main

import (
	"fmt"
)

type Content struct {
	body string
	size int
}

type Page struct {
	nextUrl     string
	previousURL string
	Content
}

func main() {

	content := Content{
		body: "test11",
		size: 11,
	}

	page := Page{
		nextUrl:     "fener",
		previousURL: "bahce",
		Content:     content,
	}

	fmt.Println(page)         // {fener bahce {test11 11}}
	fmt.Println(page.Content) // {test11 11}

}

// https://play.golang.org/p/Co-j-ybYBcA
