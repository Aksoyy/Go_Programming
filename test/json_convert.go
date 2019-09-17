package main

import (
	"fmt"
	"reflect"
)

func main() {

	type Film struct {
		Isim        string  `myjson:"film_name"`
		Yazar       string  `myjson:"director_name"`
		Imdb        float32 `myjson:"imdb_score"`
		CikisTarihi int     `myjson:"year_of_release"`
	}

	myFilm := Film{}
	myReflect := reflect.TypeOf(myFilm)

	for _, value := range []string{"Isim", "Yazar", "Imdb", "CikisTarihi"} {
		field, found := myReflect.FieldByName(value)
		if !found {
			continue // pass
		}
		fmt.Println(field.Tag.Get("myjson"))
	}

}

// film_name
// director_name
// imdb_score
// year_of_release
