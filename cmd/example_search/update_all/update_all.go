package main

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/tPhume/pet-search/model"
	"github.com/tPhume/pet-search/search"
	"log"
	"os"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	failOnError(err, "could not connect to elasticsearch")

	// create pet and client
	sushi, _ := model.NewPetInstanceWithId(os.Args[1], os.Args[2], os.Args[3])
	petClient := search.NewPetClient(es)

	// add pet
	log.Println("---- UPDATING PET ----")
	err = petClient.UpdatePetAll(context.Background(), sushi)
	failOnError(err, "")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
