package main

import (
	"fmt"
	"net/http"
)

type myHandler struct {
	greeting string
}

func (my myHandler) ServeHTTP(w http.ResponseWriter, r http.Request) {
	w.Write([]byte(fmt.Sprintf("%v world", my.greeting)))
}

func main() {
	http.Handle("/", &myHandler{greeting: "Test"})
	http.ListenAndServe(":8000", nil)
}
