// Read fonksiyonu içerisinde dosyanın kapalı olup olmadığını
// kontrol edin. Dosya kapalı ise "file: file is closed"
// mesajı ile birlikte yeni bir hata döndürün.

package main

import (
	"errors"
	"fmt"
)

type File struct {
	closed bool
}

func Read(file File) error {

	if file.closed == true {
		fmt.Println("file: file is closed")
		return errors.New("error type")
	}
	return nil
}

func main() {
	f := File{closed: true}
	err := Read(f)
	if err != nil {
		fmt.Println(err)
	}
}

// file: file is closed
// error type
