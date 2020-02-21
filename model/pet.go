package model

import "errors"

var (
	errorEmptyId   = errors.New("id is empty")
	errorEmptyName = errors.New("name is empty")
)

type PetModel interface {
	GetId() string
	GetName() string
	GetDesc() string
}

type PetInstance struct {
	id   string
	name string
	desc string
}

func (p *PetInstance) GetId() string {
	return p.id
}

func (p *PetInstance) GetName() string {
	return p.name
}

func (p *PetInstance) GetDesc() string {
	return p.desc
}
