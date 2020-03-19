package rabbit

import (
	"context"
	"github.com/tPhume/pet-search/model"
)

// operations that async must support
type Pet interface {
	AddPet(context.Context, model.PetModel) error
	UpdatePet(context.Context, model.PetModel) error
	DeletePet(context.Context, string) error
}
