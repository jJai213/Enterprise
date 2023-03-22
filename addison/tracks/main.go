package main

import (
	"log"
	"net/http"
	"tracks/repository"
	"tracks/resources"
)

func main() {
	repository.Init()   // initialise repository
	repository.Clear()  // clear any existing tables
	repository.Create() // create a new table

	log.Fatal(http.ListenAndServe(":3000", resources.Router()))
}
