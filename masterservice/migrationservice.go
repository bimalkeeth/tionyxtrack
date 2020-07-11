package main

import "tionyxtrack/masterservice/mappers"
import "log"

func main() {
	mapp := mappers.New()
	err := mapp.GenerateSchema()
	if err != nil {
		log.Fatal("Error")
	}
}
