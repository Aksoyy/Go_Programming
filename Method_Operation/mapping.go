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

/*
	Burada belirtilen map içindeki anahtarın değerine bakarken;
	monty_python_films_by_year["1983"] --> value + kontrol
	değeri geriye dönmektedir. Burada belirtilen anahtara
	karşılık düşen bir değerin olup olmadıgının kontrolu yapılabılır.

	film_name, registered := monty_python_films_by_year["2019"]

	if !registered {
		fmt.Println("Monty python did not relase a film at 2019")
	} else {
		fmt.Println(film_name)
	}

*/
