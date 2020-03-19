package main

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/tPhume/pet-search/consumer"
	"github.com/tPhume/pet-search/search"
	"log"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	consumer.ConsumePet(search.NewPetClient(es))
}
