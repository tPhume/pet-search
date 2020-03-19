package main

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"github.com/tPhume/pet-search/rabbit"
	"github.com/tPhume/pet-search/search"
	"github.com/tPhume/pet-search/server"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError("fail to connect to rabbitmq", err)

	ch, err := conn.Channel()
	failOnError("fail to create a channel", err)

	err = ch.ExchangeDeclare(
		"pet",
		amqp.ExchangeDirect,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError("could not declare exchange", err)

	queue, err := ch.QueueDeclare(
		"Add pet",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError("could not declare queue", err)

	err = ch.QueueBind(queue.Name, "pet.add", "pet", false, nil)
	failOnError("fail to bind queue", err)

	router := gin.Default()
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	server.RegisterPetRoutes(router, search.NewPetClient(es), &rabbit.PetClient{ch})
	log.Fatal(router.Run("0.0.0.0:8080"))
}

func failOnError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
