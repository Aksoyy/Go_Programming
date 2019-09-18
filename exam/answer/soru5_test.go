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
	t.Run("converts int", func(t *testing.T) {
		want := "1"
		got, _ := Stringify(1)
		if got != want {
			t.Errorf("Expected %s got %s", want, got)
		}
	})
	t.Run("converts float", func(t *testing.T) {
		want := "1.12"
		got, _ := Stringify(1.12312312)
		if got != want {
			t.Errorf("Expected %s got %s", want, got)
		}
	})
	t.Run("converts string", func(t *testing.T) {
		want := "test"
		got, _ := Stringify("test")
		if got != want {
			t.Errorf("Expected %s got %s", want, got)
		}
	})
	t.Run("handles unknown type", func(t *testing.T) {
		_, err := Stringify([]int{})
		if err == nil {
			t.Errorf("Expected an error")
		}
	})
}
