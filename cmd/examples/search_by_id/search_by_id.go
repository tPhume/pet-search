package main

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/tPhume/pet-search/search"
	"log"
	"os"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	failOnError(err, "could not connect to elasticsearch")

	pc := search.NewPetClient(es)
	log.Println("---- MATCHING PET BY ID ----")
	res, err := pc.SearchPetByID(context.Background(), os.Args[1])
	failOnError(err, "")
	log.Printf("\n%s\n", res)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
