package main

import "fmt"

type Reader interface {
	Read() []byte
}

type Writer interface {
	Write([]byte) int
}
type File struct {
	name    string
	content []byte
}

func (file *File) Read() []byte {
	return file.content
}

func (file *File) Write(content []byte) int {
	file.content = append(file.content, content...)
	return len(content)
}

type Socket struct {
	ip     string
	port   int
	buffer []byte
}

func (sock *Socket) Read() []byte {
	return sock.buffer
}

func (sock *Socket) Write(stream []byte) int {
	sock.buffer = append(sock.buffer, stream...)
	return len(stream)
}

func main() {
	var file Reader = &File{
		name:    "test",
		content: []byte("data"),
	}
	var sock Writer = &Socket{
		ip:   "0.0.0.0",
		port: 21,
	}

	sock.Write(file.Read())

	fmt.Println(sock.buffer) // sock.buffer undefined (type Writer has no field or method buffer)
}

// # command-line-arguments
// ./multi_interface.go:53:18: sock.buffer undefined (type Writer has no field or method buffer)
