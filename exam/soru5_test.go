// Stringify fonksiyonu çalıştırılınca oluşabilecek
// herhangi 2 davranışı test edin.
package main

import (
	"errors"
	"fmt"
	"testing"
)

// Stringify turns the supported value types into a string
// with a specified format
func Stringify(value interface{}) (string, error) {
	switch value.(type) {
	case string:
		return value.(string), nil
	case float32, float64:
		return fmt.Sprintf("%.2f", value), nil
	case int:
		return fmt.Sprintf("%d", value), nil
	default:
		return "", errors.New("Invalid type")
	}
}

func TestStringify(t *testing.T) {

	t.Run("string testing", func(t *testing.T) {
		got, errorName := Stringify("9.0")
		wantType := ""
		if got != wantType {
			t.Errorf("got %q want %q --> %s", got, wantType, errorName)
		}
	})

	// t.Run("float testing", func(t *testing.T) {
	// 	got, errorName := Stringify(9.0)
	// 	wantType, wantError := 9.0, false

	// 	if got != wantType {
	// 		t.Errorf("got %q want %q", got, wantType)
	// 	}
	// })
}

// soru5_test.go:32: got "9.0" want "" --> %!s(<nil>)
