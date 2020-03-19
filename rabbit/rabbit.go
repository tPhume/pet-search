package rabbit

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/streadway/amqp"
	"github.com/tPhume/pet-search/model"
)

// operations that async must support
type Pet interface {
	AddPet(context.Context, model.PetModel) error
	UpdatePet(context.Context, model.PetModel) error
}

// Concrete implementation
type PetClient struct {
	channel *amqp.Channel
}

func (pc *PetClient) AddPet(ctx context.Context, petModel model.PetModel) error {
	var body bytes.Buffer
	enc := gob.NewEncoder(&body)

	if err := enc.Encode(petModel); err != nil {
		return err
	}

	if err := pc.channel.Publish(
		"pet",
		"pet.add",
		false,
		false,
		amqp.Publishing{
			Type: "gob",
			Body: body.Bytes(),
		},
	); err != nil {
		return err
	}

	return nil
}

func (pc *PetClient) UpdatePet(ctx context.Context, petModel model.PetModel) error {
	var body bytes.Buffer
	enc := gob.NewEncoder(&body)

	if err := enc.Encode(petModel); err != nil {
		return err
	}

	if err := pc.channel.Publish(
		"pet",
		"pet.update",
		false,
		false,
		amqp.Publishing{
			Type: "gob",
			Body: body.Bytes(),
		},
	); err != nil {
		return err
	}

	return nil
}
