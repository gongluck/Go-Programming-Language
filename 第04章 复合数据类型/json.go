package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movie = Movie{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}}

	data, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("JSON marshaling failed : ", err)
	} else {
		fmt.Printf("%s\n", data)
	}

	data, err = json.MarshalIndent(movie, "+", "\t")
	if err != nil {
		fmt.Println("JSON marshaling failed : ", err)
	} else {
		fmt.Printf("%s\n", data)
	}
}
