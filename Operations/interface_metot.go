package main

import "fmt"

type IOInterface interface {
	Read() []byte
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

type Logger struct {
}

func (log *Logger) Log(io IOInterface) {
	for _, data := range io.Read() {
		fmt.Print(string(data))
	}
	fmt.Println()
}

func main() {
	file := &File{
		name: "test",
	}
	sock := &Socket{
		ip:   "0.0.0.0",
		port: 21,
	}
	logger := Logger{}

	data := []byte("data")

	file.Write(data)
	sock.Write(file.Read())

	logger.Log(file) // data
	logger.Log(sock) // data
}
