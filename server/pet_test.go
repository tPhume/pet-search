package server

import (
	"context"
	"github.com/tPhume/pet-search/model"
)

// Represent our pet for testing
// the setters does nothing here
type sushi struct{}

func (s sushi) SetId(string) {}

func (s sushi) GetId() string {
	return "1"
}

func (s sushi) SetName(string) {}

func (s sushi) GetName() string {
	return "sushi"
}

func (s sushi) SetDesc(string) {}

func (s sushi) GetDesc() string {
	return "Sushi is a good boy."
}

// Is used for testing the happy path
type petSearchHappy struct{}

func (p petSearchHappy) CheckStatus() error {
	return nil
}

func (p petSearchHappy) AddPet(ctx context.Context, pm model.PetModel) (string, error) {
	return "1", nil
}

func (p petSearchHappy) SearchPetById(ctx context.Context, id string) (model.PetModel, error) {
	return nil, nil
}

func (p petSearchHappy) UpdatePetAll(ctx context.Context, pm model.PetModel) error {
	return nil
}

func (p petSearchHappy) DeletePetById(ctx context.Context, id string) error {
	return nil
}

func (p petSearchHappy) ListPetByName(ctx context.Context, name string) ([]model.PetModel, error) {
	return nil, nil
}

func (p petSearchHappy) ListPetAll(ctx context.Context) ([]model.PetModel, error) {
	return nil, nil
}
