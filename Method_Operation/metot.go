package main

import (
	"fmt"
	"strconv"
)

type Connection struct {
	host     string
	port     int
	timeOut  int
	isClosed bool
}

func (conn *Connection) Close() {
	conn.isClosed = true
	fmt.Println("Connection", conn.host+":"+strconv.Itoa(conn.port), "closed.")
}

func main() {
	c := Connection{
		host:     "0.0.0.0",
		port:     8080,
		timeOut:  1,
		isClosed: false,
	}
	c.Close() // Connection 0.0.0.0:8080 closed.
}
