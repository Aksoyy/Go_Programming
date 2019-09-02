package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {

	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "tet"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b", "abc", "ab"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 2)) // "1" olaydı: f0o
	p("Split:     ", s.Split("a-b-c-d-e", "-"))     // in list
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	p("Len: ", len("hello"))
	p("Char:", "Aello"[0])
}

// Contains:   true
// Count:      2
// HasPrefix:  false
// HasSuffix:  true
// Index:      1
// Join:       a-b-abc-ab
// Repeat:     aaaaa
// Replace:    f00
// Replace:    f00
// Split:      [a b c d e]
// ToLower:    test
// ToUpper:    TEST

// Len:  5
// Char: 65
