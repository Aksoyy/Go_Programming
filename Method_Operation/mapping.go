package main

import "fmt"

func main() {

	monty_python_films_by_year := map[string]string{
		"1971": "Monty Python and the Holy Grail",
		"1979": "Life of Brian",
	}

	monty_python_films_by_year["1983"] = "The Meaning of Life"

	for key, value := range monty_python_films_by_year {
		fmt.Println(key, ":", value)
	}
	//fmt.Println(monty_python_films_by_year)
}

// 1979 : Life of Brian
// 1983 : The Meaning of Life
// 1971 : Monty Python and the Holy Grail
