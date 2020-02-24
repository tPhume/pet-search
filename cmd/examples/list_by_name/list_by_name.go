package main

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/tPhume/pet-search/search"
	"log"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	failOnError(err, "could not connect to elasticsearch")

	petClient := search.NewPetClient(es)

	// list pet = sushi
	log.Println("---- LIST PET BY NAME ----")
	resList, err := petClient.ListPetByName(context.Background(), "sushi")
	failOnError(err, "")

	for _, r := range resList {
		log.Printf("\n%s\n", r)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
