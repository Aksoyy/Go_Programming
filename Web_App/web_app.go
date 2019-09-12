package main

import (
	http "net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r http.Request) {
		w.Write([]byte("Test App"))
	})
	http.ListenAndServe(":8000", nil)
}
