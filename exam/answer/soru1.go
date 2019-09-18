// Read fonksiyonu içerisinde dosyanın kapalı olup olmadığını
// kontrol edin. Dosya kapalı ise "file: file is closed"
// mesajı ile birlikte yeni bir hata döndürün.

package main

import "fmt"

type File struct {
	closed bool
}

type ErrorFileClosed struct{}

func (e ErrorFileClosed) Error() string {
	return "file: file is closed"
}
func Read(file File) error {
	if file.closed {
		return ErrorFileClosed{}
	}
	return nil
}

func main() {
	f := File{closed: true}
	err := Read(f)
	fmt.Println(err)
}
