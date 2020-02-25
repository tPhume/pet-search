package main

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	"github.com/tPhume/pet-search/search"
	"github.com/tPhume/pet-search/server"
	"log"
)

func main() {
	router := gin.Default()
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	server.RegisterPetRoutes(router, search.NewPetClient(es))
	log.Fatal(router.Run("0.0.0.0:8080"))
}
