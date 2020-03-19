package consumer

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/tPhume/pet-search/model"
	"github.com/tPhume/pet-search/search"
	"log"
)

// simple consumer that will consume and run AddPet
func ConsumePet(search search.Pet) {
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

	msgs, err := ch.Consume(
		queue.Name,
		"consume add pet",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError("could not start consumer", err)

	var id string

	for d := range msgs {
		err = nil

		body := &model.PetInstance{}

		err = json.Unmarshal(d.Body, body)
		if err != nil {
			_ = d.Nack(false, true)
		} else {
			id, err = search.AddPet(context.Background(), body)
			if err != nil {
				_ = d.Nack(false, true)
			}
		}

		_ = d.Ack(false)
		log.Printf("Pet: %s added with ID: %s\n", body.GetName(), id)
	}
}

func failOnError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
