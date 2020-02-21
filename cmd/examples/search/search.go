package main

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/tPhume/pet-search/model"
	"github.com/tPhume/pet-search/search"
	"log"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	failOnError(err, "could not connect to elasticsearch")

	// create pet and client
	sushi := model.NewPetInstance("sushi", "He is a good boy.")
	petClient := search.NewPetClient(es)

	// add pet
	res, err := petClient.AddPet(context.Background(), sushi)
	failOnError(err, "")
	log.Printf("\n%s\n", res)

	// list pet = sushi
	resList, err := petClient.ListPetByName(context.Background(), "sushi")
	failOnError(err, "")

	for _, r := range resList {
		log.Printf("\n%s\n", r)
	}

	// list all pet
	resList, err = petClient.ListAllPet(context.Background())
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
