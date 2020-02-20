package main

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/tPhume/pet-search/search"
	"io/ioutil"
	"log"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	failOnError(err, "could not connect to elasticsearch")

	sushi := search.NewPetModel("sushi", "He is a good boy.")
	petClient := search.NewPetClient(es)

	// add pet
	res, err := petClient.AddPet(context.Background(), sushi)
	failOnError(err, "")

	resByte, err := ioutil.ReadAll(res.Body)
	failOnError(err, "failed to read response")

	log.Println(string(resByte))
	defer res.Body.Close()

	// list pet = sushi
	res, err = petClient.ListPetByName(context.Background(), "sushi")
	failOnError(err, "")

	resByte, err = ioutil.ReadAll(res.Body)
	failOnError(err, "failed to read response")

	log.Println(string(resByte))
	defer res.Body.Close()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
