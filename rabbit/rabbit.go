package rabbit

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/tPhume/pet-search/model"
)

// operations that async must support
type Pet interface {
	AddPet(context.Context, model.PetInstance) error
}

// Concrete implementation
type PetClient struct {
	Channel *amqp.Channel
}

func (pc *PetClient) AddPet(ctx context.Context, petModel model.PetInstance) error {
	body, err := json.Marshal(petModel)
	if err != nil {
		return err
	}

	if err = pc.Channel.Publish(
		"pet",
		"pet.add",
		false,
		false,
		amqp.Publishing{
			Type: "json",
			Body: body,
		},
	); err != nil {
		return err
	}

	return nil
}
